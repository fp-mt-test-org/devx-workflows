package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List workflows specified by user",
	Long:  `Reads config from service_config.yml to list user defined workflows`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return list()
	},
}

func list() error {
	cmdDefList, err := getWorkflowDefList()
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

func getWorkflowDefList() (map[string]Workflow, error) {
	var workflowDefList map[string]Workflow
	if err := viper.UnmarshalKey(workflowKey, &workflowDefList); err != nil {
		return nil, fmt.Errorf("error unmarshalling cmd definition in service_config.yml: %s", err)
	}
	return workflowDefList, nil
}
