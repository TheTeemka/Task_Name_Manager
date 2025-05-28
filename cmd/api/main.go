package main

import (
	"log/slog"
	"os"

	_ "github.com/TheTeemka/TaskNameManager/cmd/api/docs"
	"github.com/TheTeemka/TaskNameManager/internal/config"
	"github.com/TheTeemka/TaskNameManager/internal/database"
	"github.com/TheTeemka/TaskNameManager/internal/repo"
	"github.com/TheTeemka/TaskNameManager/internal/server"
	"github.com/TheTeemka/TaskNameManager/internal/service"
	"github.com/TheTeemka/TaskNameManager/pkg/utils"
)

func main() {
	slog.SetDefault(utils.Mylog(os.Stdout, slog.LevelDebug))
	cfg := config.MustLoad()

	db := database.OpenPostgres(cfg.DBString)

	personRepo := repo.NewPersonRepository(db)
	personService := service.NewPersonService(personRepo)

	srv := server.NewServer(cfg.ServerPort, personService)
	srv.Serve()
}
