package cmd

import (
	"database/sql"
	"delivery/internal/repository"
	"delivery/internal/usecase"
	"fmt"

	"github.com/spf13/cobra"
)

var name, branch string

var addRepoCmd = &cobra.Command{
	Use:   "add-repo [url] -n \"[name]\" -b \"[branch]\"",
	Short: "Add a new repository with a branch to track changes",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if name == "" || branch == "" {
			fmt.Println("Error: Flag name and branch are required")
			cmd.Usage()
			return
		}

		db, err := sql.Open("sqlite", repository.DBFileName)
		if err != nil {
			fmt.Println("Error to init the database", err)
			return
		}
		defer db.Close()

		repoRepository := repository.NewSQLiteRepo(db)
		repoUseCase := usecase.NewApplicationUseCase(repoRepository)

		url := args[0]

		err = repoUseCase.AddApplication(url, name, branch)
		if err != nil {
			fmt.Printf("Error adding a new repository: %s\n", err)
		} else {
			fmt.Println("New repository have been created.")
		}
	},
}

func init() {
	addRepoCmd.Flags().StringVarP(&name, "name", "n", "", "Name of the repository (required)")
	addRepoCmd.Flags().StringVarP(&branch, "branch", "b", "", "Branch of the repository (required)")
	addRepoCmd.MarkFlagRequired("name")
	addRepoCmd.MarkFlagRequired("branch")
	rootCmd.AddCommand(addRepoCmd)
}
