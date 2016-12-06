package main

import "fmt"

// PrintVersion displays the application version and additional information
func PrintVersion() {
	fmt.Println(`gpress version ` + version + `🐣   GCOMPRESSION lib version 0.0.1/2

Copyright (C) 2016 by Jakub Kułak.
gpress comes with ABSOLUTELY NO WARRANTY.  This is free software, and you are welcome to redistribute it under certain conditions.  See the GNU General Public Licence for details.`)
}

// PrintHelp displays help, explaining how to use the application
func PrintHelp() {

	fmt.Println(`usage: gpress [--decompress] [--verbose] --sourcefile <file>

--decompress            decompress file
--help                  display help
--sourcefile            file to be compressed or decompressed
--verbose               add debug output
--version               display program version`)
}
