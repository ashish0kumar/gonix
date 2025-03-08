package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	files := os.Args[1:]
	if len(files) == 0 {
		fmt.Println("Usage: touch <file>...")
		os.Exit(1)
	}

	for _, file := range files {
		f, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Printf("touch: cannot create file '%s': %v\n", file, err)
			os.Exit(1)
		}
		f.Close()

		now := time.Now()
		err = os.Chtimes(file, now, now)
		if err != nil {
			fmt.Printf("touch: cannot update timestamp for '%s': %v\n", file, err)
			os.Exit(1)
		}
	}
}
