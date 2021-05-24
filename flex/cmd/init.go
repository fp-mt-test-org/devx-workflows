package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"devx-workflows/pkg/version"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize config",
	Long:  `Sets up service_config.yml to build application`,
	Run: func(cmd *cobra.Command, args []string) {
		viper.Set("flex.version", version.FlexVersion)

		fmt.Println("Service Name (ie. helloworld-service):")
		viper.Set("service.name", promptUser())

		fmt.Println("--- Configure flex commands ---")
		for input := ""; input != "n"; input = promptUser() {
			fmt.Println("Name of your command (eg. build): ")
			cmdName := promptUser()
			fmt.Println("Executable for this command (eg. docker build .): ")
			cmdCommand := promptUser()
			viper.Set(fmt.Sprintf("%s.%s.command", cmdKey, cmdName), cmdCommand)
			fmt.Println("Configure another command? [y/n]")
		}

		viper.WriteConfigAs("./service_config.yml")
	},
}

func promptUser() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSuffix(input, "\n")
}
