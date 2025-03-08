package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func countReader(reader *bufio.Scanner) (lines, words, bytes int) {
	for reader.Scan() {
		lines++
		line := reader.Text()
		words += countWords(line)
		bytes += len(line) + 1
	}
	return
}

func countFile(filename string) (lines, words, bytes int, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, 0, 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines, words, bytes = countReader(scanner)

	if err := scanner.Err(); err != nil {
		return 0, 0, 0, err
	}

	return
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
		scanner := bufio.NewScanner(os.Stdin)
		lines, words, bytes := countReader(scanner)
		fmt.Println(formatOutput(lines, words, bytes, *lFlag, *wFlag, *cFlag))
		return
	}

	for _, filename := range files {
		lines, words, bytes, err := countFile(filename)
		if err != nil {
			fmt.Printf("wc: %s: %v\n", filename, err)
			continue
		}
		fmt.Printf("%s %s\n", formatOutput(lines, words, bytes, *lFlag, *wFlag, *cFlag), filename)
	}
}

func formatOutput(lines, words, bytes int, lFlag, wFlag, cFlag bool) string {
	output := []string{}
	if lFlag {
		output = append(output, fmt.Sprintf("%d", lines))
	}
	if wFlag {
		output = append(output, fmt.Sprintf("%d", words))
	}
	if cFlag {
		output = append(output, fmt.Sprintf("%d", bytes))
	}

	if len(output) == 0 {
		output = append(output, fmt.Sprintf("%d %d %d", lines, words, bytes))
	}

	return strings.Join(output, " ")
}
