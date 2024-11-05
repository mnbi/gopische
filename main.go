package main

import (
	"flag"
	"fmt"
	"os"
)

var printVersion bool
var printUsage bool

func main() {
	flag.BoolVar(&printVersion, "v", false, "print version")
	flag.BoolVar(&printUsage, "h", false, "print usage")
	flag.Parse()

	if printVersion {
		version()
		os.Exit(1)
	}

	if printUsage {
		flag.Usage()
		os.Exit(1)
	}

	fmt.Println("Hello, world!")
}

func version() {
	fmt.Println("gopische version 0.1.0")
}
