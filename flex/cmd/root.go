package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "flex",
	Short: "Flex for all of your CI/CD needs",
	Long:  `Build, push, pull, and deploy your application with Flex`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Make sure to add `alias flex=\"./devx-workflow-scripts/exec/flex\" to your ~/.zshrc and/or ~/.bashrc\n" +
			"`flex help` for more info")
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
