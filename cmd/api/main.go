package main

import (
	"log/slog"
	"os"

	_ "github.com/TheTeemka/hhChat/cmd/api/docs"
	"github.com/TheTeemka/hhChat/internal/config"
	"github.com/TheTeemka/hhChat/internal/database"
	"github.com/TheTeemka/hhChat/internal/repo"
	"github.com/TheTeemka/hhChat/internal/server"
	"github.com/TheTeemka/hhChat/internal/service"
	"github.com/TheTeemka/hhChat/pkg/utils"
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
