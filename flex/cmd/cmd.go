package cmd

import (
	"fmt"

	exec "devx-workflows/pkg/exec"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Cmd struct {
	Command string
	Env     []string
}

const cmdKey = "flex.cmd"

func init() {
	rootCmd.AddCommand(command)
	command.AddCommand(listCmd)
}

var command = &cobra.Command{
	Use:   "cmd",
	Short: "Execute user-defined commands from service_config.yml",
	Long:  `Reads config from service_config.yml to list and execute user defined commands`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		execObj := new(exec.Obj)
		return cmdExec(execObj, args[0])
	},
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List commands specified by user",
	Long:  `Reads config from service_config.yml to list user defined commands`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return list()
	},
}

func cmdExec(execObj exec.E, cmdName string) error {
	cmdDefList, err := getCmdDefList()
	if err != nil {
		return err
	}
	cmdDef, exists := cmdDefList[cmdName]
	if !exists {
		return fmt.Errorf("could not find command definition for %s; run `flex cmd list` for a list of available commands", cmdName)
	}

	cmd := cmdDef.Command
	if len(cmd) > 0 {
		return execObj.ExecFn(cmd, cmdDef.Env...)
	}
	return fmt.Errorf("command for `%s` is empty in service_config.yml", cmdName)
}

func list() error {
	cmdDefList, err := getCmdDefList()
	if err != nil {
		return err
	}
	if len(cmdDefList) == 0 {
		return fmt.Errorf("no commands specified in service_config.yml; `flex init` before running this command")
	}

	fmt.Println("List of commands:")
	for key, el := range cmdDefList {
		fmt.Printf("  %s: %s\n", key, el.Command)
	}
	return nil
}

func getCmdDefList() (map[string]Cmd, error) {
	var cmdDefList map[string]Cmd
	if err := viper.UnmarshalKey(cmdKey, &cmdDefList); err != nil {
		return nil, fmt.Errorf("error unmarshalling cmd definition in service_config.yml: %s", err)
	}
	return cmdDefList, nil
}
