package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
)

func grep(pattern string, reader *bufio.Scanner, ignoreCase bool) {
	if ignoreCase {
		pattern = "(?i)" + pattern
	}
	re, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Fprintf(os.Stderr, "grep: invalid pattern: %v\n", err)
		os.Exit(1)
	}

	for reader.Scan() {
		line := reader.Text()
		if re.MatchString(line) {
			fmt.Println(line)
		}
	}

	if err := reader.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "grep: error reading input: %v\n", err)
		os.Exit(1)
	}
}

func main() {
	ignoreCase := flag.Bool("i", false, "Perform case-insensitive matching")
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("Usage: grep [-i] <pattern> [file...]")
		os.Exit(1)
	}

	pattern := args[0]
	files := args[1:]

	if len(files) == 0 {
		grep(pattern, bufio.NewScanner(os.Stdin), *ignoreCase)
	} else {
		for _, file := range files {
			f, err := os.Open(file)
			if err != nil {
				fmt.Fprintf(os.Stderr, "grep: cannot open '%s': %v\n", file, err)
				continue
			}
			grep(pattern, bufio.NewScanner(f), *ignoreCase)
			f.Close()
		}
	}
}
