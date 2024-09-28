package cmd

import (
	"delivery/internal/repository"
	"delivery/internal/usecase"
	"fmt"

	"github.com/spf13/cobra"
)

var environment string
var traefikRule string
var version string

var addDeployCmd = &cobra.Command{
	Use:   "add [name] [application name] [domain] -e [environment] -t [traefikRule] -v [version]",
	Short: "Commnad to add a new deploy",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		// Add a new deploy
		name := args[0]
		applicationName := args[1]
		domainName := args[2]

		if name == "" || applicationName == "" || domainName == "" {
			cmd.Usage()
			fmt.Println("Error: name, application name and domain are required")
			return
		}

		appRepository := repository.NewApplicationSQLRepository()
		appUseCase := usecase.NewApplicationUseCase(appRepository)
		app, err := appUseCase.FindApplication(applicationName)

		if err != nil || app == nil {
			fmt.Println("Error: Application not found")
			return
		}

		deployRepository := repository.NewDeploySQLiteRepository()
		depoyUseCase := usecase.NewDeployUseCase(deployRepository)

		if environment == "" {
			environment = "na"
		}

		if traefikRule == "" {
			traefikRule = "na"
		}

		if version == "" {
			version = "1.0.0"
		}

		if err := depoyUseCase.AddDeploy(name, app.ID, environment, domainName, traefikRule, version); err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Deploy added successfully")
	},
}

func init() {
	addDeployCmd.Flags().StringVarP(&environment, "environmnet", "e", "", "Deploy environment variables (optional)")
	addDeployCmd.Flags().StringVarP(&traefikRule, "traefik-rule", "t", "", "Deploy traefik rule (optional)")
	addDeployCmd.Flags().StringVarP(&version, "version", "v", "", "Deploy version (optional)")
	DeployCmd.AddCommand(addDeployCmd)
}
