package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 2 {
		fmt.Println("Usage: ls [directory]")
		return
	}

	dir := "./"
	if len(os.Args) == 2 {
		dir = os.Args[1]
	}

	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Printf("Error reading directory: %v\n", err)
		return
	}

	for _, file := range files {
		if file.IsDir() {
			fmt.Printf("\033[34m%s/\033[0m ", file.Name())
		} else {
			fmt.Print(file.Name(), " ")
		}
	}
	fmt.Println()
}
