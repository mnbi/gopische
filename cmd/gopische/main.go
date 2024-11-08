package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/mnbi/gopische"
)

var (
	versionFlag = flag.Bool("v", false, "show version")
	usageFlag   = flag.Bool("h", false, "show usage")
)

func main() {
	flag.Usage = gopische.Usage
	flag.Parse()

	if *versionFlag {
		gopische.ShowVersion()
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
