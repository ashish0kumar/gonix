package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	rFlag := false
	fFlag := false
	processedArgs := []string{}

	for _, arg := range os.Args[1:] {
		if arg == "-rf" || arg == "-fr" {
			rFlag, fFlag = true, true
		} else if arg == "-r" {
			rFlag = true
		} else if arg == "-f" {
			fFlag = true
		} else if strings.HasPrefix(arg, "-") {
			fmt.Printf("rm: invalid option '%s'\n", arg)
			os.Exit(1)
		} else {
			processedArgs = append(processedArgs, arg)
		}
	}

	if len(processedArgs) == 0 {
		fmt.Println("Usage: rm [-r] [-f] <file>...")
		os.Exit(1)
	}

	for _, file := range processedArgs {
		var err error
		if rFlag {
			err = os.RemoveAll(file)
		} else {
			err = os.Remove(file)
		}

		if err != nil && !fFlag {
			fmt.Printf("rm: cannot remove '%s': %v\n", file, err)
			os.Exit(1)
		}
	}
}
