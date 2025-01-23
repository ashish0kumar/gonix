package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	showHidden := flag.Bool("a", false, "Include hidden files in the listing")
	showDetails := flag.Bool("l", false, "Display detailed file information")
	flag.Parse()

	dir := "./"
	args := flag.Args()
	if len(args) > 1 {
		fmt.Println("Usage: ls [-a] [-l] [directory]")
		return
	}
	if len(args) == 1 {
		dir = args[0]
	}

	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Printf("Error reading directory: %v\n", err)
		return
	}

	dirs := []os.DirEntry{}
	normalFiles := []os.DirEntry{}
	for _, file := range files {
		if !*showHidden && file.Name()[0] == '.' {
			continue
		}
		if file.IsDir() {
			dirs = append(dirs, file)
		} else {
			normalFiles = append(normalFiles, file)
		}
	}

	listFiles(dirs, *showDetails)
	listFiles(normalFiles, *showDetails)
	fmt.Println()
}

func listFiles(files []os.DirEntry, showDetails bool) {
	for _, file := range files {
		if showDetails {
			info, err := file.Info()
			if err != nil {
				fmt.Printf("Error retrieving info for %s: %v\n", file.Name(), err)
				continue
			}

			name := file.Name()
			if file.IsDir() {
				name = fmt.Sprintf("\033[34m%s/\033[0m", file.Name())
			}

			fmt.Printf("%-10s %-8d %-20s %s\n",
				getPerms(info.Mode()),
				info.Size(),
				info.ModTime().Format(time.RFC822),
				name,
			)
		} else {
			if file.IsDir() {
				fmt.Printf("\033[34m%s/\033[0m ", file.Name())
			} else {
				fmt.Print(file.Name(), " ")
			}
		}
	}
}

func getPerms(mode os.FileMode) string {
	perms := ""
	if mode.IsDir() {
		perms += "d"
	} else {
		perms += "-"
	}

	perms += rwx(mode & 0700 >> 6)
	perms += rwx(mode & 0070 >> 3)
	perms += rwx(mode & 0007)
	return perms
}

func rwx(mode os.FileMode) string {
	r := "-"
	w := "-"
	x := "-"

	if mode&4 != 0 {
		r = "r"
	}
	if mode&2 != 0 {
		w = "w"
	}
	if mode&1 != 0 {
		x = "x"
	}
	return r + w + x
}
