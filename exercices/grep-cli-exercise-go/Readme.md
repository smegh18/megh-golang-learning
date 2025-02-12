# MyGrep: A Command-Line Grep Utility in Go

## Overview
MyGrep is a Unix-like `grep` command-line utility implemented in Go. It allows searching for a string in files, directories, and standard input, with support for various options like case-insensitive search, recursive search, context lines, and output redirection.

## Features
- ✅ **Basic text search** in files and directories
- ✅ **Recursive search (`-r`)** to traverse directories
- ✅ **Case-insensitive search (`-i`)**
- ✅ **Context control**: show lines before/after matches (`-A`, `-B`, `-C`)
- ✅ **Count matches only (`-c`)**
- ✅ **Read from STDIN** for interactive search
- ✅ **Redirect output to a file (`-o`)**

## Installation
### Prerequisites
Ensure you have Go installed on your system. You can check by running:
```sh
$ go version
```
Compile the program:
```sh
$ go build -o mygrep main.go
```
This creates an executable named `mygrep`.

## Usage

### 1. Basic Search in a File
```sh
$ ./mygrep "search_string" filename.txt
```

### 2. Recursive Search in a Directory
```sh
$ ./mygrep -r "test" my_folder/
```

### 3. Case-Insensitive Search
```sh
$ ./mygrep -i "Test" filename.txt
```

### 4. Search from Standard Input (Interactive Mode)
```sh
$ echo "this is a test" | ./mygrep "test"
```

### 5. Write Output to a File
```sh
$ ./mygrep "error" logfile.txt -o errors.txt
```

### 6. Count Only Matches
```sh
$ ./mygrep -c "test" filename.txt
```

### 7. Show N Lines Before/After Match
```sh
$ ./mygrep -B 2 "error" logfile.txt   # Show 2 lines before match
$ ./mygrep -A 3 "error" logfile.txt   # Show 3 lines after match
$ ./mygrep -C 2 "error" logfile.txt   # Show 2 lines before & after match
```

## Error Handling
- If the specified file does not exist:
  ```sh
  $ ./mygrep "test" nonexistent.txt
  mygrep: nonexistent.txt: No such file or directory
  ```
- If a directory is provided without `-r`:
  ```sh
  $ ./mygrep "test" my_folder
  mygrep: my_folder: Is a directory
  ```
- If output file already exists:
  ```sh
  $ ./mygrep "error" logfile.txt -o output.txt
  Error opening output file: file already exists
  ```


