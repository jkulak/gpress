package main

import (
	"fmt"
	"os"
)

// ErrorCheck checks for errors and acts upon them
func ErrorCheck(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
		// panic(err)
	}
}
