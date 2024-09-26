package cmd

import (
	"database/sql"
	"delivery/internal/repository"
	"delivery/internal/usecase"
	"fmt"

	"github.com/spf13/cobra"
)

var removeRepoCmd = &cobra.Command{
	Use:   "remove-repo [name]",
	Short: "Remove a repository by its Name",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		db, err := sql.Open("sqlite", repository.DBFileName)
		if err != nil {
			fmt.Println("Error to init db:", err)
			return
		}
		defer db.Close()

		database := repository.NewSQLiteRepo(db)
		repoUseCase := usecase.NewApplicationUseCase(database)

		name := args[0]

		err = repoUseCase.RemoveApplication(name)
		if err != nil {
			fmt.Println("Error removing the repository:", err)
		} else {
			fmt.Println("The repository has been removed:", name)
		}
	},
}

func init() {
	rootCmd.AddCommand(removeRepoCmd)
}
