package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	exec "github.flexport.io/flexport/devx-workflow-scripts/pkg/exec"
)

func init() {
	rootCmd.AddCommand(testCmd)
}

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Test your application",
	Long:  `Reads config from service_config.yml to test application`,
	RunE: func(cmd *cobra.Command, args []string) error {
		execObj := new(exec.Obj)
		return test(execObj)
	},
}

func test(execObj exec.E) error {
	testCommand := viper.GetString("test.command")
	if len(testCommand) > 0 {
		return execObj.ExecFn(testCommand, viper.GetStringSlice("test.env")...)
	}

	return fmt.Errorf("Command not specified in service_config.yml; `flex init` before running this command")
}
