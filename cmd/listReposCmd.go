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
	Short: "Lista todos los repositorios monitorizados",
	Run: func(cmd *cobra.Command, args []string) {
		db, err := sql.Open("sqlite", repository.DBFileName)
		if err != nil {
			fmt.Println("Error al inicializar la base de datos:", err)
			return
		}
		defer db.Close()

		repoRepo := repository.NewSQLiteRepo(db)
		repoUseCase := usecase.NewRepoUseCase(repoRepo)

		repos, err := repoUseCase.ListRepositories()
		if err != nil {
			fmt.Println("Error al listar los repositorios:", err)
			return
		}

		if len(repos) == 0 {
			fmt.Println("No hay repositorios monitorizados.")
		} else {
			fmt.Println("Repositorios monitorizados:")
			for _, repo := range repos {
				fmt.Printf("-ID: %s, URL: %s, Branch: %s\n", repo.ID, repo.URL, repo.Branch)
			}
		}
	},
}
