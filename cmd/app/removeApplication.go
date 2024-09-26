package cmd

import (
	"database/sql"
	"delivery/internal/repository"
	"delivery/internal/usecase"
	"fmt"

	"github.com/spf13/cobra"
)

var removeRepoCmd = &cobra.Command{
	Use:   "remove [name]",
	Short: "Remove an application by its name",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		db, err := sql.Open("sqlite", repository.DBFileName)
		if err != nil {
			fmt.Println("Error to init db:", err)
			return
		}
		defer db.Close()

		database := repository.NewSQLiteApplication(db)
		repoUseCase := usecase.NewApplicationUseCase(database)

		name := args[0]

		err = repoUseCase.RemoveApplication(name)
		if err != nil {
			fmt.Println("Error removing the application:", err)
		} else {
			fmt.Println("The application has been removed:", name)
		}
	},
}

func init() {
	AppCmd.AddCommand(removeRepoCmd)
}
