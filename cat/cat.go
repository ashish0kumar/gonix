package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	nFlag := flag.Bool("n", false, "Number all output lines")
	flag.Parse()

	files := flag.Args()

	if len(files) == 0 {
		printStream(os.Stdin, *nFlag)
		return
	}

	for i, file := range files {
		if err := printFile(file, *nFlag); err != nil {
			fmt.Fprintf(os.Stderr, "cat: %s: %v\n", file, err)
			os.Exit(1)
		}

		if i < len(files)-1 {
			fmt.Println()
		}
	}
}

func printFile(filename string, numberLines bool) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	printStream(f, numberLines)
	return nil
}

func printStream(r io.Reader, numberLines bool) {
	if numberLines {
		scanner := bufio.NewScanner(r)
		lineNum := 1
		for scanner.Scan() {
			fmt.Printf("%6d  %s\n", lineNum, scanner.Text())
			lineNum++
		}
	} else {
		_, _ = io.Copy(os.Stdout, r)
		fmt.Println()
	}
}
