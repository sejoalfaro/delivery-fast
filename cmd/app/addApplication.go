package cmd

import (
	"database/sql"
	"delivery/internal/repository"
	"delivery/internal/usecase"
	"fmt"

	"github.com/spf13/cobra"
)

var branch string

var addApplicationCmd = &cobra.Command{
	Use:   "add [name] [url] -b [branch]",
	Short: "Add a new application",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		if branch == "" {
			fmt.Println("Error: Flag branch is required")
			cmd.Usage()
			return
		}

		db, err := sql.Open("sqlite", repository.DBFileName)
		if err != nil {
			fmt.Println("Error to init the database", err)
			return
		}
		defer db.Close()

		appRepository := repository.NewApplicationSQLRepository()
		appUseCase := usecase.NewApplicationUseCase(appRepository)

		name := args[0]
		url := args[1]

		err = appUseCase.AddApplication(url, name, branch)
		if err != nil {
			fmt.Printf("Error adding a new application: %s\n", err)
		} else {
			fmt.Println("New application has been created.")
		}
	},
}

func init() {
	addApplicationCmd.Flags().StringVarP(&branch, "branch", "b", "", "App branch (required)")
	addApplicationCmd.MarkFlagRequired("branch")
	AppCmd.AddCommand(addApplicationCmd)
}
