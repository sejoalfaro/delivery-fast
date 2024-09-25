package cmd

import (
	"database/sql"
	"delivery/internal/repository"
	"delivery/internal/usecase"
	"fmt"

	"github.com/spf13/cobra"
)

// Definición del comando `remove-repo`.
var removeRepoCmd = &cobra.Command{
	Use:   "remove-repo [url]",
	Short: "Elimina un repositorio de la lista de monitorización",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		db, err := sql.Open("sqlite3", "./repos.db")
		if err != nil {
			fmt.Println("Error al inicializar la base de datos:", err)
			return
		}
		defer db.Close()

		database := repository.NewSQLiteRepo(db)
		repoUseCase := usecase.NewRepoUseCase(database)

		url := args[0]

		err = repoUseCase.RemoveRepository(url)
		if err != nil {
			fmt.Println("Error al eliminar el repositorio:", err)
		} else {
			fmt.Println("Repositorio eliminado exitosamente:", url)
		}
	},
}
