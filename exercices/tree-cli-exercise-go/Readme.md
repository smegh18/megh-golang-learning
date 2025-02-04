# Tree CLI Exercise (Go)

## Overview
This is a Go-based command-line utility that replicates the functionality of the Unix `tree` command. It recursively prints the directory structure in a tree-like format with various filtering and formatting options.

## Features
- Prints the directory tree structure with indentation.
- Supports various options:
  - `-f` : Print full relative file paths.
  - `-d` : Print only directories, ignoring files.
  - `-L <n>` : Limit the tree depth to `n` levels.
  - `-p` : Show file permissions.
  - `-t` : Sort by last modification time instead of alphabetically.
  - `-X` : Output in XML format.
  - `-J` : Output in JSON format.
  - `-if` : Print without indentation lines.
- Proper error handling for inaccessible directories.
- Displays summary (total directories and files count).

## Installation
### Prerequisites
- Go 1.18+ installed

### Build
```sh
go build -o tree
```

### Run
```sh
./tree <directory_path>
```

## Usage Examples
### Basic Usage
```sh
./tree .
```

### Using Flags
```sh
./tree -d .          # Show only directories
./tree -f .          # Show full paths
./tree -L 3 .        # Limit depth to 3 levels
./tree -p .          # Show file permissions
./tree -t .          # Sort by last modified time
./tree -X .          # Output in XML format
./tree -J .          # Output in JSON format
./tree -if .         # Print without indentation lines
```

## Running Tests
```sh
go test ./...
```

## Assumptions
- The program mimics the Unix `tree` command behavior.
- Output is printed to `STDOUT`.
- Proper permissions are required to read directories.


