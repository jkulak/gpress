package main

import (
    "fmt"
    "flag"
    "io/ioutil"
    "os"
    "github.com/jkulak/gcompressor"
)

// Generic error handling
func check(err error) {
    if err != nil {
        fmt.Fprintf(os.Stderr, "error: %v\n", err)
        os.Exit(1)
        // panic(err)
    }
}

func log(s string, verbose bool) {
    if verbose {
        fmt.Printf(s)
    }
}

func main() {

    paramSourceFile := flag.String("sourcefile", "", "a filename")
    paramVerbose := flag.Bool("verbose", false, "be verbose")
    paramDecompress := flag.Bool("decompress", false, "decompress")
    paramHelp := flag.Bool("help", false, "display help")
    paramVersion := flag.Bool("version", false, "display program version")

    flag.Parse()

    if *paramVersion {
        fmt.Fprintf("gpress version 0.9   GCOMPRESSION lib version 0.0.1/2\nCopyright (C) 2016 by Jakub Kułak.\ngpress comes with ABSOLUTELY NO WARRANTY.  This is free software, and you are welcome to redistribute it under certain conditions.  See the GNU General Public Licence for details.");
        os.Exit(0)
    }

    if *paramSourceFile != "" {
        input, err := ioutil.ReadFile(*paramSourceFile)
        check(err)
        log(string(input), *paramVerbose)

        if *paramDecompress {
            output := gcompressor.Compress(string(input)) + "\n"
            fmt.Printf("Compressing... Done\n")
        } else {
            output := gcompressor.Decompress(string(input)) + "\n"
            fmt.Printf("Compressing... Done\n")
        }

        log(output, *paramVerbose)
        datComp := []byte(output)

        error := ioutil.WriteFile(string(*paramSourceFile) + ".gpressed", datComp, 0644)
        check(error)
    } else {
        fmt.Printf("You need to specify the sourcefile")

    }

    // Debug part
    log("sourceFile: " + *paramSourceFile, *paramVerbose)
    // log("tail:" + string(flag.Args()), *paramVerbose)
}
