package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	outputFile := flag.String("o", "", "Write output to file")
	followRedirects := flag.Bool("L", false, "Follow redirects")

	flag.Parse()
	args := flag.Args()

	if len(args) == 0 {
		fmt.Println("Usage: curl [-L] [-o filename] <URL>")
		os.Exit(1)
	}

	url := args[0]

	client := &http.Client{}
	if *followRedirects {
		client.CheckRedirect = nil
	} else {
		client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		}
	}

	resp, err := client.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "curl: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		fmt.Fprintf(os.Stderr, "curl: HTTP error %d\n", resp.StatusCode)
		os.Exit(1)
	}

	var out io.Writer = os.Stdout
	if *outputFile != "" {
		file, err := os.Create(*outputFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "curl: cannot write to '%s': %v\n", *outputFile, err)
			os.Exit(1)
		}
		defer file.Close()
		out = file
	}

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "curl: error reading response: %v\n", err)
		os.Exit(1)
	}
}
