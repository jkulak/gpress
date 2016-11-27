package main

import (
    "fmt"
    "flag"
    "io/ioutil"
    "gcompressor"
)

// Generic error handling
func check(e error) {
    if e != nil {
        fmt.Println("Problem reading the file!")
        panic(e)
    }
}

func main() {

    sourceFile := flag.String("sourcefile", "", "a filename")

    flag.Parse()

    if *sourceFile != "" {
        dat, err := ioutil.ReadFile(*sourceFile)
        check(err)
        fmt.Print(string(dat))

        // datCompressed := compress(string(dat))
        datCompressed := string(dat) + "compression?"
        datComp := []byte(datCompressed)

        error := ioutil.WriteFile(string(*sourceFile) + ".gpressed", datComp, 0644)
        check(error)
    }

    // Debug part
    fmt.Println("sourceFile:", *sourceFile)
    fmt.Println("tail:", flag.Args())
}
