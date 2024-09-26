package cmd

import (
	"database/sql"
	"delivery/internal/repository"
	"delivery/internal/usecase"
	"fmt"

	"github.com/spf13/cobra"
)

// Definición del comando `list-repos`.
var listReposCmd = &cobra.Command{
	Use:   "list-repos",
	Short: "List all monitored repositories",
	Run: func(cmd *cobra.Command, args []string) {
		db, err := sql.Open("sqlite", repository.DBFileName)
		if err != nil {
			fmt.Println("Error to init the database", err)
			return
		}
		defer db.Close()

		repoRepo := repository.NewSQLiteRepo(db)
		repoUseCase := usecase.NewApplicationUseCase(repoRepo)

		repos, err := repoUseCase.ListApplications()
		if err != nil {
			fmt.Println("Error getting the repository`s list:", err)
			return
		}

		if len(repos) == 0 {
			fmt.Println("No repositories are being monitored.")
		} else {
			fmt.Println("List of monitored repositories:")
			for _, repo := range repos {
				fmt.Printf("-ID: %s, Name: %s, URL: %s, Branch: %s\n", repo.ID, repo.Name, repo.URL, repo.Branch)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(listReposCmd)
}
