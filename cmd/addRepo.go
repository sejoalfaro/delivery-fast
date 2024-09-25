package cmd

import (
	"database/sql"
	"delivery/internal/repository"
	"delivery/internal/usecase"
	"fmt"

	"github.com/spf13/cobra"
)

// Definición del comando `add-repo`.
var addRepoCmd = &cobra.Command{
	Use:   "add-repo [url] [branch]",
	Short: "Agrega un nuevo repositorio a la lista de monitorización",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		db, err := sql.Open("sqlite", repository.DBFileName)
		if err != nil {
			fmt.Println("Error al inicializar la base de datos:", err)
			return
		}
		defer db.Close()

		// Crear el repositorio y el caso de uso.
		repoRepository := repository.NewSQLiteRepo(db)
		repoUseCase := usecase.NewRepoUseCase(repoRepository)

		url := args[0]
		branch := args[1]

		err = repoUseCase.AddRepository(url, branch)
		if err != nil {
			fmt.Println("Error al agregar el repositorio:", err)
		} else {
			fmt.Println("Repositorio agregado exitosamente:", url)
		}
	},
}
