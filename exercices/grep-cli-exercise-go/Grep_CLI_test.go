package main

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
)

func TestPrintMatches(t *testing.T) {
	cases := []struct {
		name       string
		input      []string
		searchTerm string
		options    Options
		expected   string
	}{
		{
			"Case-sensitive match",
			[]string{"hello world", "goodbye world"},
			"hello",
			Options{},
			"STDIN:hello world\n",
		},
		{
			"Case-insensitive match",
			[]string{"Hello world", "goodbye world"},
			"hello",
			Options{ignoreCase: true},
			"STDIN:Hello world\n",
		},
		{
			"Count only",
			[]string{"hello world", "hello again", "no match here"},
			"hello",
			Options{countOnly: true},
			"STDIN:2\n",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var output bytes.Buffer
			printMatches(tc.input, tc.searchTerm, "STDIN", tc.options, &output)
			actual := output.String()
			if actual != tc.expected {
				t.Errorf("expected %q but got %q", tc.expected, actual)
			}
		})
	}
}

func TestProcessFile(t *testing.T) {
	tmpFile, err := os.CreateTemp("", "testfile")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpFile.Name())

	content := "hello world\nthis is a test\ngoodbye world"
	if _, err := tmpFile.WriteString(content); err != nil {
		t.Fatal(err)
	}
	tmpFile.Close()

	var output bytes.Buffer
	processFile(tmpFile.Name(), "hello", Options{}, &output)
	expected := tmpFile.Name() + ":hello world\n"
	actual := output.String()

	if actual != expected {
		t.Errorf("expected %q but got %q", expected, actual)
	}
}

func TestRecursiveSearch(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "testdir")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	files := []struct {
		name    string
		content string
	}{
		{"file1.txt", "hello world\n"},
		{"file2.txt", "hello again\n"},
	}

	for _, file := range files {
		path := filepath.Join(tmpDir, file.name)
		if err := os.WriteFile(path, []byte(file.content), 0644); err != nil {
			t.Fatal(err)
		}
	}

	var output bytes.Buffer
	recursiveSearch(tmpDir, "hello", Options{recursive: true}, &output)

	expected := filepath.ToSlash(filepath.Join(tmpDir, "file1.txt")) + ":hello world\n" +
		filepath.ToSlash(filepath.Join(tmpDir, "file2.txt")) + ":hello again\n"

	actual := filepath.ToSlash(output.String())

	if actual != expected {
		t.Errorf("expected\n%q\nbut got\n%q", expected, actual)
	}
}
