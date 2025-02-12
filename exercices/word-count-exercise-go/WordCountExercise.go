package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

// countStats holds the count of lines, words, and bytes in the input.
type countStats struct {
	lines int
	words int
	bytes int
}

// count processes the input reader and calculates line, word, and byte counts.
func count(r io.Reader) countStats {
	var stats countStats
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		line := scanner.Text()
		stats.lines++
		stats.words += len(strings.Fields(line))
		stats.bytes += len(line) + 1 // Assume newline exists
	}

	// If input is non-empty, remove the last extra newline count
	if stats.bytes > 0 {
		stats.bytes--
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
	}

	return stats
}

// processFile opens and processes a given file, returning its count statistics.
func processFile(filename string) (countStats, error) {
	file, err := os.Open(filename)
	if err != nil {
		return countStats{}, err // Return an error if the file cannot be opened
	}
	defer file.Close() // Ensure the file is closed after processing
	return count(file), nil
}

// formatOutput formats the count statistics for display.
func formatOutput(stats countStats, label string, countLines, countWords, countBytes bool) string {
	var output []string

	// Append only the requested counts; if no flag is set, default to all counts.
	if countLines || countWords || countBytes {
		if countLines {
			output = append(output, fmt.Sprintf("%6d", stats.lines))
		}
		if countWords {
			output = append(output, fmt.Sprintf("%6d", stats.words))
		}
		if countBytes {
			output = append(output, fmt.Sprintf("%6d", stats.bytes))
		}
	} else {
		output = append(output, fmt.Sprintf("%6d %6d %6d", stats.lines, stats.words, stats.bytes))
	}

	return fmt.Sprintf("%s %s", strings.Join(output, " "), label)
}

// main is the entry point of the program, handling CLI arguments and executing the word count logic.
func main() {
	// Define CLI flags
	lineFlag := flag.Bool("l", false, "Count lines")
	wordFlag := flag.Bool("w", false, "Count words")
	byteFlag := flag.Bool("c", false, "Count bytes")
	flag.Parse() // Parse command-line arguments

	files := flag.Args() // Get file names from arguments

	if len(files) == 0 {
		// No files provided; read from STDIN
		stats := count(os.Stdin)
		fmt.Println(formatOutput(stats, "(stdin)", *lineFlag, *wordFlag, *byteFlag))
		return
	}

	total := countStats{} // Accumulator for multiple file processing

	for _, file := range files {
		stats, err := processFile(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "wc: cannot open '%s' for reading: %v\n", file, err)
			continue // Continue processing other files even if one fails
		}

		// Update the total count
		total.lines += stats.lines
		total.words += stats.words
		total.bytes += stats.bytes

		// Print per-file statistics
		fmt.Println(formatOutput(stats, file, *lineFlag, *wordFlag, *byteFlag))
	}

	// Print the total summary if multiple files were processed
	if len(files) > 1 {
		fmt.Println(formatOutput(total, "total", *lineFlag, *wordFlag, *byteFlag))
	}
}
