package cmd

import (
	"github.com/spf13/cobra"
	exec "github.flexport.io/flexport/devx-workflow-scripts/pkg/exec"
)

func init() {
	rootCmd.AddCommand(buildAndTestCmd)
}

var buildAndTestCmd = &cobra.Command{
	Use:   "build-test",
	Short: "Build and test your application",
	Long:  `Reads config from service_config.yml to build and test the application`,
	RunE: func(cmd *cobra.Command, args []string) error {
		execObj := new(exec.Obj)
		return buildAndTest(execObj)
	},
}

func buildAndTest(execObj exec.E) error {
	build(execObj)
	test(execObj)

	return nil
}
