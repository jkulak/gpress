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
    paramDecompress := flag.Bool("d", false, "decompress")

    flag.Parse()

    if *paramSourceFile != "" {
        input, err := ioutil.ReadFile(*paramSourceFile)
        check(err)
        log(string(input), *paramVerbose)

        // output := ""

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
