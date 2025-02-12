# gonix

**gonix** is a collection of simplified implementations of common **Unix
commands**, written in **Go**. The goal is to create lightweight, fast, and
minimal versions of these tools while learning **systems programming** in Go.

## Commands

| **Command**       | **Description**                                                                      |
| ----------------- | ------------------------------------------------------------------------------------ |
| [`ls`](/ls)       | Lists files and directories with support for hidden files and detailed views.        |
| [`mkdir`](/mkdir) | Creates new directories.                                                             |
| [`touch`](/touch) | Creates empty files or updates timestamps of existing files.                         |
| [`rm`](/rm)       | Removes files and directories (supports `-r` and `-f` flags).                        |
| [`cat`](/cat)     | Prints file contents with optional line numbering (`-n`).                            |
| [`wc`](/wc)       | Counts lines, words, and bytes in a file or from `stdin`.                            |
| [`grep`](/grep)   | Searches for a pattern in files or `stdin`, supports case-insensitive search (`-i`). |

## Installation

- **Clone the repository:**

```bash
git clone https://github.com/ashish0kumar/gonix.git
```

- **Navigate to a specific command's directory and build it:**

```bash
cd gonix/<command> 
go build <command>.go
```

Replace `<command>` with any of the available commands, like `ls`, `rm`, `grep`,
etc.

## Usage

Each command follows a similar usage pattern. For example:

- List files: `./ls`
- Remove a file: `./rm file.txt`
- Create a directory: `./mkdir newdir`
- Search for text: `./grep "pattern" file.txt`

Refer to each command’s subdirectory for detailed usage instructions.

## Why this project?

- Learn **systems programming** with Go.
- Implement essential **Unix utilities** in a lightweight manner.
- Improve understanding of **file handling, text processing, and CLI tools**.

## Contributions

Feel free to contribute, suggest improvements, or report issues!
