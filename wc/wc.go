package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func countFile(filename string) (lines, words, bytes int, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, 0, 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines++
		line := scanner.Text()
		words += countWords(line)
		bytes += len(line) + 1
	}

	if err := scanner.Err(); err != nil {
		return 0, 0, 0, err
	}

	return lines, words, bytes, nil
}

func countWords(line string) int {
	scanner := bufio.NewScanner(strings.NewReader(line))
	scanner.Split(bufio.ScanWords)

	wordCount := 0
	for scanner.Scan() {
		wordCount++
	}
	return wordCount
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: wc <file>")
		os.Exit(1)
	}

	filename := os.Args[1]
	lines, words, bytes, err := countFile(filename)
	if err != nil {
		fmt.Printf("wc: %s: %v\n", filename, err)
		os.Exit(1)
	}

	fmt.Printf("%d %d %d %s\n", lines, words, bytes, filename)
}
