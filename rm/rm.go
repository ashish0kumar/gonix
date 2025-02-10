package main

import (
	"fmt"
	"os"
)

func main() {
	files := os.Args[1:]
	if len(files) == 0 {
		fmt.Println("Usage: rm <file>...")
		os.Exit(1)
	}

	for _, file := range files {
		err := os.Remove(file)
		if err != nil {
			fmt.Printf("rm: cannot remove '%s': %v\n", file, err)
			os.Exit(1)
		}
	}
}
