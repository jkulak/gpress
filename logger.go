package main

import "fmt"

// log logs data to the output
// It supports bool verbose parameter to log extra information
// Useful while searching for issues
func log(s string, verbose bool) {
	if verbose {
		fmt.Printf(s)
	}
}
