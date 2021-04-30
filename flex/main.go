package main

import (
	"flag"
	"fmt"
	"os"

	"devx-workflows/flex/cmd"
)

// FlexVersion contains the semantic version of this tool and should be set via ldflags at build-time.
var FlexVersion = "unspecified"

func main() {
	version := flag.Bool("version", false, "Prints the version of this tool.")
	flag.Parse()
	if *version {
		fmt.Println(FlexVersion)
		os.Exit(0)
	}

	err := cmd.Execute()
	if err != nil {
		// Ensure that an error results in a failed exit status
		os.Exit(1)
	}
}
