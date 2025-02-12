package main

import (
	"strings"
	"testing"
)

// TestCount checks if the count function correctly counts lines, words, and bytes.
func TestCount(t *testing.T) {
	tests := []struct {
		input    string
		expected countStats
	}{
		{"hello world", countStats{lines: 1, words: 2, bytes: 11}},
		{"hello\nworld", countStats{lines: 2, words: 2, bytes: 11}},
		{"one two three\nfour five six", countStats{lines: 2, words: 6, bytes: 27}},
		{"", countStats{lines: 0, words: 0, bytes: 0}},
	}

	for _, test := range tests {
		reader := strings.NewReader(test.input)
		result := count(reader)
		if result != test.expected {
			t.Errorf("For input %q, expected %+v but got %+v", test.input, test.expected, result)
		}
	}
}

// TestFormatOutput checks the output formatting.
func TestFormatOutput(t *testing.T) {
	stats := countStats{lines: 3, words: 5, bytes: 20}

	cases := []struct {
		countLines bool
		countWords bool
		countBytes bool
		expected   string
	}{
		{true, false, false, "     3 test.txt"},
		{false, true, false, "     5 test.txt"},
		{false, false, true, "    20 test.txt"},
		{true, true, true, "     3      5     20 test.txt"},
	}

	for _, c := range cases {
		output := formatOutput(stats, "test.txt", c.countLines, c.countWords, c.countBytes)
		if output != c.expected {
			t.Errorf("Expected %q but got %q", c.expected, output)
		}
	}
}
