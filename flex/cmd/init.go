package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Artifact represents a build artifact, exposed in CI, can be push/pulled if right type
type Artifact struct {
	RepoType    string
	Source      string
	Destination string
	Version     string
}

// Service contains all relevant info regarding a service for initialization of files
type Service struct {
	Name string
}

func init() {
	rootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize config",
	Long:  `Sets up service_config.yml to build application`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("What is the command to build your service (eg. docker build .): ")
		viper.Set("build.command", promptUser())

		fmt.Println("Specify build-time environment variables, separate by newline (eg. SOME_ENV_VAR=blah): ")
		fmt.Println("Type \"done\" when done")
		var buildEnv []string
		for input := promptUser(); input != "done"; input = promptUser() {
			buildEnv = append(buildEnv, input)
		}
		if len(buildEnv) == 0 {
			viper.Set("build.env", append(buildEnv, "SOME_ENV_VAR=blah"))
		} else {
			viper.Set("build.env", buildEnv)
		}

		viper.WriteConfigAs("./service_config.yml")

		var service Service
		fmt.Println("What is your service's name? (ie. fpos-services)")
		service.Name = promptUser()
	},
}

func promptUser() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSuffix(input, "\n")
}
