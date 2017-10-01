// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 10.
//!+

// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string][]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, files := range counts {
		if len(files) > 1 {
			fmt.Printf("%d\t%s\t%s\n", len(files), line, strings.Join(dedupeFiles(files), ", "))
		}
	}
}

func countLines(f *os.File, counts map[string][]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()] = append(counts[input.Text()], f.Name())
	}
	// NOTE: ignoring potential errors from input.Err()
}

func dedupeFiles(files []string) []string {
	encountered := map[string]bool{}

	for v := range files {
		encountered[files[v]] = true
	}

	result := []string{}
	for key, _ := range encountered {
		result = append(result, key)
	}
	return result
}

//!-
