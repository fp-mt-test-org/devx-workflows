package main

import (
	"os"

	"devx-workflows/flex/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		// Ensure that an error results in a failed exit status
		os.Exit(1)
	}
}
