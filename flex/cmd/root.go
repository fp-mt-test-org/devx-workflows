package cmd

import (
	"fmt"

	exec "devx-workflows/pkg/exec"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Workflow struct {
	Command string
	Env     []string
}

const workflowKey = "flex.workflows"

var rootCmd = &cobra.Command{
	Use:   "flex",
	Short: "Flex for all of your CI/CD needs",
	Long:  `Execute custom workflows for your application with Flex`,
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		execObj := new(exec.Obj)
		return workflowExec(execObj, args[0])
	},
}

// Execute is a cobra requirement to execute our flex root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	viper.SetConfigName("service_config")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("Config file not found, use `flex init` to initialize")
		} else {
			panic(fmt.Errorf("fatal error config file: %s", err))
		}
	}
}

func workflowExec(execObj exec.E, workflowName string) error {
	workflowDefList, err := getWorkflowDefList()
	if err != nil {
		return err
	}
	workflowDef, exists := workflowDefList[workflowName]
	if !exists {
		return fmt.Errorf("could not find workflow definition for %s; run `flex list` for a list of available workflows", workflowName)
	}

	cmd := workflowDef.Command
	if len(cmd) > 0 {
		return execObj.ExecFn(cmd, workflowDef.Env...)
	}
	return fmt.Errorf("command for `%s` is empty in service_config.yml", workflowName)
}
