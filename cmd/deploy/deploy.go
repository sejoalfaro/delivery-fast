package cmd

import "github.com/spf13/cobra"

var DeployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy a new application",
}
