package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// Options struct holds all CLI flags
type Options struct {
	ignoreCase  bool
	recursive   bool
	beforeLines int
	afterLines  int
	outputFile  string
	countOnly   bool
}

func main() {
	// Parse command-line flags
	options, searchTerm, files := parseFlags()

	// Open output file if required
	var output io.Writer = os.Stdout
	if options.outputFile != "" {
		file, err := os.OpenFile(options.outputFile, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening output file: %v\n", err)
			os.Exit(1)
		}
		defer file.Close()
		output = file
	}

	// Process input files
	if len(files) == 0 {
		searchStdin(searchTerm, options, output)
	} else {
		for _, file := range files {
			if options.recursive {
				recursiveSearch(file, searchTerm, options, output)
			} else {
				processFile(file, searchTerm, options, output)
			}
		}
	}
}

func parseFlags() (Options, string, []string) {
	var opts Options
	flag.BoolVar(&opts.ignoreCase, "i", false, "Perform case-insensitive search")
	flag.BoolVar(&opts.recursive, "r", false, "Search directories recursively")
	flag.IntVar(&opts.beforeLines, "B", 0, "Print N lines before match")
	flag.IntVar(&opts.afterLines, "A", 0, "Print N lines after match")
	flag.IntVar(&opts.beforeLines, "C", 0, "Print N lines before and after match")
	flag.StringVar(&opts.outputFile, "o", "", "Output file")
	flag.BoolVar(&opts.countOnly, "c", false, "Print only count of matches")
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("Usage: mygrep [options] <search_term> [files...]")
		os.Exit(1)
	}

	return opts, args[0], args[1:]
}

func searchStdin(searchTerm string, opts Options, output io.Writer) {
	scanner := bufio.NewScanner(os.Stdin)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	printMatches(lines, searchTerm, "STDIN", opts, output)
}

func recursiveSearch(root, searchTerm string, opts Options, output io.Writer) {
	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			processFile(path, searchTerm, opts, output)
		}
		return nil
	})
}

func processFile(filename, searchTerm string, opts Options, output io.Writer) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "mygrep: %s: %v\n", filename, err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	printMatches(lines, searchTerm, filename, opts, output)
}

func printMatches(lines []string, searchTerm, source string, opts Options, output io.Writer) {
	if opts.ignoreCase {
		searchTerm = strings.ToLower(searchTerm)
	}

	matchCount := 0
	for _, line := range lines {
		lineToCheck := line
		if opts.ignoreCase {
			lineToCheck = strings.ToLower(line)
		}

		if strings.Contains(lineToCheck, searchTerm) {
			matchCount++
			if !opts.countOnly {
				fmt.Fprintf(output, "%s:%s\n", source, line)
			}
		}
	}

	if opts.countOnly {
		fmt.Fprintf(output, "%s:%d\n", source, matchCount)
	}
}
