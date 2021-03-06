// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 8.

// Echo3 prints its command-line arguments.
// TODO: Exercise 1.3
package main

import (
	"fmt"
	"os"
	"strings"
)

//!+
func main() {
	fmt.Println(strings.Join(os.Args, " "))
	for i, arg := range os.Args[1:] {
		fmt.Printf("%d: %s\n", i+1, arg)
	}
}

//!-
