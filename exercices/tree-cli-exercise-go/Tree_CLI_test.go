package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func setupTestDir(base string) error {
	paths := []string{
		"dir1",
		"dir1/file1.txt",
		"dir1/file2.txt",
		"dir2",
		"dir2/file3.txt",
		"file4.txt",
	}
	for _, path := range paths {
		fullPath := filepath.Join(base, path)
		if filepath.Ext(path) == "" {
			if err := os.MkdirAll(fullPath, 0755); err != nil {
				return err
			}
		} else {
			f, err := os.Create(fullPath)
			if err != nil {
				return err
			}
			f.Close()
		}
	}
	return nil
}

func TestListDir(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "tree_test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	if err := setupTestDir(tempDir); err != nil {
		t.Fatal(err)
	}

	_, dirCount, fileCount := listDir(tempDir, 0)

	if dirCount != 2 {
		t.Errorf("Expected 2 directories, got %d", dirCount)
	}
	if fileCount != 4 {
		t.Errorf("Expected 4 files, got %d", fileCount)
	}
}

func TestJSONOutput(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "tree_test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	if err := setupTestDir(tempDir); err != nil {
		t.Fatal(err)
	}

	root, dirCount, fileCount := listDir(tempDir, 0)
	tree := Tree{Root: root, Directories: dirCount, Files: fileCount}
	_, err = json.MarshalIndent(tree, "", "  ")
	if err != nil {
		t.Errorf("Failed to generate JSON output: %v", err)
	}
}

func TestOnlyDirsFlag(t *testing.T) {
	oldVal := *onlyDirs
	defer func() { *onlyDirs = oldVal }()
	*onlyDirs = true

	tempDir, err := os.MkdirTemp("", "tree_test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	if err := setupTestDir(tempDir); err != nil {
		t.Fatal(err)
	}

	_, dirCount, fileCount := listDir(tempDir, 0)

	if fileCount != 0 {
		t.Errorf("Expected 0 files, got %d", fileCount)
	}
	if dirCount != 2 {
		t.Errorf("Expected 2 directories, got %d", dirCount)
	}
}
