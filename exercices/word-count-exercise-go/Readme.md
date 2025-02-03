# Word Count Utility (Go Implementation)

## Overview
This is a Go implementation of a word count utility similar to the Unix `wc` command. It reads text from files or standard input and counts lines, words, and bytes, providing an output similar to `wc`.

## Features
- Count the number of **lines**, **words**, and **bytes** in a file or standard input.
- Supports multiple file inputs, displaying individual counts and a total summary.
- Flags to selectively count lines, words, or bytes.
- Handles errors gracefully, continuing execution even if some files cannot be read.

## Usage
### Command-line Options
The program supports the following flags:

- `-l` : Count **lines**
- `-w` : Count **words**
- `-c` : Count **bytes**

If no flags are specified, the program defaults to counting all three metrics.

### Running the Program
#### Counting a Single File
```sh
$ go run main.go example.txt
```
Example output:
```
  10   50  250 example.txt
```
(10 lines, 50 words, 250 bytes in `example.txt`)

#### Counting Multiple Files
```sh
$ go run main.go file1.txt file2.txt
```
Example output:
```
  10   50  250 file1.txt
   5   30  180 file2.txt
  15   80  430 total
```

#### Counting from Standard Input
```sh
$ echo "Hello, World!" | go run main.go
```
Example output:
```
   1    2   14 (stdin)
```

#### Using Flags
To count only lines:
```sh
$ go run main.go -l example.txt
```
Output:
```
  10 example.txt
```

To count only words:
```sh
$ go run main.go -w example.txt
```
Output:
```
  50 example.txt
```

To count only bytes:
```sh
$ go run main.go -c example.txt
```
Output:
```
 250 example.txt
```

## Implementation Details
- **`countStats` Struct**: Stores line, word, and byte counts.
- **`count(r io.Reader)` Function**: Reads input and calculates counts using a `bufio.Scanner`.
- **`processFile(filename string)` Function**: Reads and processes a file, returning its count stats.
- **`formatOutput(stats countStats, label string, countLines, countWords, countBytes bool)` Function**: Formats the output correctly.
- **`main()` Function**: Parses flags, processes files, handles standard input, and prints results.

## Error Handling
- If a file cannot be opened, an error message is displayed, but other files are still processed.
- If standard input is unreadable, an error message is displayed and execution stops.

## Exercise
Modify the program to:
1. **Add character counting (`-m` flag)**: Modify `countStats` to include a `chars` field and update `count()` to track it.
2. **Handle empty files gracefully**: Ensure the output format remains consistent even for empty files.
3. **Improve formatting**: Align output in columns for better readability.
4. **Add unit tests**: Implement tests for `count()` and `processFile()` using Go's `testing` package.



