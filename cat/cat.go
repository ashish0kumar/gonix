package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	files := os.Args[1:]

	if len(files) == 0 {
		_, err := io.Copy(os.Stdout, os.Stdin)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cat: error reading stdin: %v\n", err)
			os.Exit(1)
		}
		return
	}

	for _, file := range files {
		err := printFile(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cat: %s: %v\n", file, err)
			os.Exit(1)
		}
	}
}

func printFile(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = io.Copy(os.Stdout, f)
	return err
}
