// Package main provides full functionality of the
// gpress compressing application
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/jkulak/gcompressor"
)

// gpress version
const version string = "0.2"

func main() {

	// Read options passed to execuatable
	paramSourceFile := flag.String("sourcefile", "", "a filename")
	paramVerbose := flag.Bool("verbose", false, "be verbose")
	paramDecompress := flag.Bool("decompress", false, "decompress")
	paramHelp := flag.Bool("help", false, "display help")
	paramVersion := flag.Bool("version", false, "display program version")
	flag.Parse()

	if *paramVersion {
		PrintVersion()
		os.Exit(0)
	}

	if *paramHelp {
		PrintHelp()
		os.Exit(0)
	}

	if *paramSourceFile != "" {
		input, err := ioutil.ReadFile(*paramSourceFile)
		ErrorCheck(err)
		log(string(input), *paramVerbose)

		var output string

		if *paramDecompress {
			output = gcompressor.Compress(string(input)) + "\n"
			fmt.Printf("Compressing... Done\n")
		} else {
			output = gcompressor.Decompress(string(input)) + "\n"
			fmt.Printf("Decompressing... Done\n")
		}

		datComp := []byte(output)

		error := ioutil.WriteFile(string(*paramSourceFile)+".gpressed", datComp, 0644)
		ErrorCheck(error)
	} else {
		fmt.Printf("You need to specify the sourcefile")
	}

	// Debug part
	log("sourceFile: "+*paramSourceFile, *paramVerbose)
}
