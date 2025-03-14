# gonix

**gonix** is a collection of simplified implementations of common **Unix
commands**, written in **Go**. The goal is to create lightweight, fast, and
minimal versions of these tools while learning **systems programming** in Go.

## Commands

| **Command**       | **Description**                                                                          |
| ----------------- | ---------------------------------------------------------------------------------------- |
| [`ls`](/cmd/ls)       | Lists files and directories with support for hidden files and detailed views.            |
| [`mkdir`](/cmd/mkdir) | Creates new directories.                                                                 |
| [`touch`](/cmd/touch) | Creates empty files or updates timestamps of existing files.                             |
| [`rm`](/cmd/rm)       | Removes files and directories (supports `-r` and `-f` flags).                            |
| [`cat`](/cmd/cat)     | Prints file contents with optional line numbering (`-n`).                                |
| [`wc`](/cmd/wc)       | Counts lines, words, and bytes in a file or from `stdin`.                                |
| [`grep`](/cmd/grep)   | Searches for a pattern in files or `stdin`, supports case-insensitive search (`-i`).     |
| [`curl`](/cmd/curl)   | Fetches content from a URL, follows redirects (`-L`), and saves output to a file (`-o`). |

## Installation

- **Clone the repository:**

```bash
git clone https://github.com/ashish0kumar/gonix.git
```

- **Navigate to a specific command's directory and build it:**

```bash
cd gonix/cmd/<command> 
go build <command>.go
```

Replace `<command>` with any of the available commands.

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

## License

[MIT](LICENSE)