package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var addDeployCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new deploy",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Implement add deploy
		fmt.Println("Add deploy")
	},
}

func init() {
	DeployCmd.AddCommand(addDeployCmd)
}
