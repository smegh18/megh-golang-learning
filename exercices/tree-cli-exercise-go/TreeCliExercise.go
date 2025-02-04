package main

import (
	"encoding/json"
	"encoding/xml"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sort"
)

type FileInfo struct {
	Name        string     `json:"name,omitempty" xml:"name,attr"`
	Permissions string     `json:"permissions,omitempty" xml:"permissions,attr"`
	Children    []FileInfo `json:"children,omitempty" xml:"children>directory,omitempty"`
	IsDir       bool       `json:"-" xml:"-"`
}

type Tree struct {
	Root        FileInfo `json:"tree" xml:"tree"`
	Directories int      `json:"directories" xml:"report>directories"`
	Files       int      `json:"files" xml:"report>files"`
}

var (
	showFullPath  = flag.Bool("f", false, "Print relative file paths")
	onlyDirs      = flag.Bool("d", false, "Show directories only")
	maxDepth      = flag.Int("L", -1, "Limit depth of recursion (-1 for unlimited)")
	showPerms     = flag.Bool("p", false, "Print file permissions")
	sortByModTime = flag.Bool("t", false, "Sort files by modification time")
	outputXML     = flag.Bool("X", false, "Print output in XML format")
	outputJSON    = flag.Bool("J", false, "Print output in JSON format")
	noIndent      = flag.Bool("if", false, "Print without indentation")
)

func getFileInfo(path string, file os.DirEntry) (FileInfo, error) {
	info, err := file.Info()
	if err != nil {
		return FileInfo{}, err
	}

	permissions := ""
	if *showPerms {
		permissions = fmt.Sprintf("[%s]", info.Mode().Perm().String())
	}

	return FileInfo{
		Name:        path,
		Permissions: permissions,
		IsDir:       file.IsDir(),
	}, nil
}

func listDir(path string, depth int) (FileInfo, int, int) {
	if *maxDepth != -1 && depth > *maxDepth {
		return FileInfo{}, 0, 0
	}

	files, err := os.ReadDir(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading directory %s: %v\n", path, err)
		os.Exit(1)
	}

	if *sortByModTime {
		sort.Slice(files, func(i, j int) bool {
			iInfo, _ := files[i].Info()
			jInfo, _ := files[j].Info()
			return iInfo.ModTime().After(jInfo.ModTime())
		})
	}

	var root FileInfo
	if *showFullPath {
		root.Name = path
	} else {
		root.Name = filepath.Base(path)
	}

	var dirs, filesCount int
	for _, file := range files {
		filePath := filepath.Join(path, file.Name())

		fileInfo, err := getFileInfo(filePath, file)
		if err != nil {
			continue
		}

		if file.IsDir() {
			dirTree, dCount, fCount := listDir(filePath, depth+1)
			fileInfo.Children = dirTree.Children
			dirs += dCount + 1
			filesCount += fCount
		} else if *onlyDirs {
			continue
		} else {
			filesCount++
		}

		root.Children = append(root.Children, fileInfo)
	}

	return root, dirs, filesCount
}

func printTree(root FileInfo, prefix string, isLast bool) {
	if *noIndent {
		fmt.Println(root.Name)
	} else {
		connector := "├── "
		if isLast {
			connector = "└── "
		}

		fmt.Println(prefix + connector + root.Name)

		newPrefix := prefix + "│   "
		if isLast {
			newPrefix = prefix + "    "
		}

		for i, child := range root.Children {
			printTree(child, newPrefix, i == len(root.Children)-1)
		}
	}
}

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) == 0 {
		fmt.Println("Usage: tree [options] <directory>")
		flag.PrintDefaults()
		os.Exit(1)
	}

	dir := args[0]
	root, dirCount, fileCount := listDir(dir, 0)

	if *outputXML {
		xmlData, _ := xml.MarshalIndent(Tree{Root: root, Directories: dirCount, Files: fileCount}, "", "  ")
		fmt.Println(string(xmlData))
	} else if *outputJSON {
		jsonData, _ := json.MarshalIndent(Tree{Root: root, Directories: dirCount, Files: fileCount}, "", "  ")
		fmt.Println(string(jsonData))
	} else {
		printTree(root, "", true)
		fmt.Printf("\n%d directories, %d files\n", dirCount, fileCount)
	}
}
