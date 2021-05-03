package main

import (
	"flag"
	"fmt"
	"os"

	"devx-workflows/flex/cmd"
	"devx-workflows/pkg/version"
)

func main() {
	versionFlag := flag.Bool("version", false, "Prints the version of this tool.")
	flag.Parse()

	if *versionFlag {
		fmt.Println(version.FlexVersion)
		os.Exit(0)
	}

	err := cmd.Execute()
	if err != nil {
		// Ensure that an error results in a failed exit status
		os.Exit(1)
	}
}
