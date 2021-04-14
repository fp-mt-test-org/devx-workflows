package main

import (
	"os"

	"github.flexport.io/flexport/devx-workflow-scripts/flex/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		// Ensure that an error results in a failed exit status
		os.Exit(1)
	}
}
