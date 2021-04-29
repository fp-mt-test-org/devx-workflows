package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	exec "devx-workflows/pkg/exec"
)

func init() {
	rootCmd.AddCommand(buildCmd)
}

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build your application",
	Long:  `Reads config from service_config.yml to build`,
	RunE: func(cmd *cobra.Command, args []string) error {
		execObj := new(exec.Obj)
		return build(execObj)
	},
}

func build(execObj exec.E) error {
	buildCommand := viper.GetString("build.command")
	if len(buildCommand) > 0 {
		return execObj.ExecFn(buildCommand, viper.GetStringSlice("build.env")...)
	}
	return fmt.Errorf("Command not specified in service_config.yml; `flex init` before running this command")
}
