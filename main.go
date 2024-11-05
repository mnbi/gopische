package main

import (
	"flag"
	"fmt"
	"os"
)

const ProgName = "gopische"
const ProgDesc = "A tiny scheme implmentation written in Go"
const ProgVersion = "0.1.0"
const ProgRelease = "2024-11-06"

func usage() {
	fmt.Fprintf(os.Stderr, "%s\n", ProgDesc)
	fmt.Fprintf(os.Stderr, "usage: %s [options] [file]\n", ProgName)
	flag.PrintDefaults()
	os.Exit(2)
}

func version() {
	fmt.Fprintf(os.Stderr, "%s version %s (%s)\n", ProgName, ProgVersion, ProgRelease)
	os.Exit(2)
}

var (
	versionFlag = flag.Bool("v", false, "show version")
	usageFlag   = flag.Bool("h", false, "show usage")
)

func main() {
	flag.Usage = usage
	flag.Parse()

	if *versionFlag {
		version()
	}

	if *usageFlag {
		flag.Usage()
	}

	args := flag.Args()

	if len(args) > 0 {
		fmt.Fprintf(os.Stderr, "ARGS:")
		for _, arg := range args {
			fmt.Fprintf(os.Stderr, " %s", arg)
		}
		fmt.Fprintf(os.Stderr, "\n")
	} else {
		fmt.Printf("Hi!\n")
	}
}
