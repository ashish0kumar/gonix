package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	pFlag := flag.Bool("p", false, "Create parent directories if needed")
	vFlag := flag.Bool("v", false, "Print message for each created directory")
	flag.Parse()

	dirs := flag.Args()
	if len(dirs) == 0 {
		fmt.Println("Usage: mkdir [-p] [-v] <directory>...")
		os.Exit(1)
	}

	for _, dir := range dirs {
		var err error
		if *pFlag {
			err = os.MkdirAll(dir, 0755)
		} else {
			err = os.Mkdir(dir, 0755)
		}

		if err != nil {
			if os.IsExist(err) {
				fmt.Printf("mkdir: cannot create directory '%s': File exists\n", dir)
			} else if os.IsPermission(err) {
				fmt.Printf("mkdir: cannot create directory '%s': Permission denied\n", dir)
			} else {
				fmt.Printf("mkdir: error creating '%s': %v\n", dir, err)
			}
			os.Exit(1)
		}

		if *vFlag {
			fmt.Printf("mkdir: created directory '%s'\n", dir)
		}
	}
}
