package main

import (
	"bufio"
	"flag"
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
	lFlag := flag.Bool("l", false, "Print line count")
	wFlag := flag.Bool("w", false, "Print word count")
	cFlag := flag.Bool("c", false, "Print byte count")

	flag.Parse()
	files := flag.Args()

	if len(files) == 0 {
		fmt.Println("Usage: wc [-l] [-w] [-c] <file>")
		os.Exit(1)
	}

	for _, filename := range files {
		lines, words, bytes, err := countFile(filename)
		if err != nil {
			fmt.Printf("wc: %s: %v\n", filename, err)
			continue
		}

		output := []string{}
		if *lFlag {
			output = append(output, fmt.Sprintf("%d", lines))
		}
		if *wFlag {
			output = append(output, fmt.Sprintf("%d", words))
		}
		if *cFlag {
			output = append(output, fmt.Sprintf("%d", bytes))
		}

		if len(output) == 0 {
			output = append(output, fmt.Sprintf("%d %d %d", lines, words, bytes))
		}

		fmt.Printf("%s %s\n", strings.Join(output, " "), filename)
	}
}
