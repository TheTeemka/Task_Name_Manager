package main

import (
	"io"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/TheTeemka/hhChat/internal/config"
	"github.com/TheTeemka/hhChat/internal/database"
	"github.com/TheTeemka/hhChat/internal/repo"
	"github.com/TheTeemka/hhChat/internal/server"
	"github.com/TheTeemka/hhChat/internal/service"
)

func main() {
	slog.SetDefault(makeLog(os.Stdout, slog.LevelDebug))
	cfg := config.MustLoad()

	db := database.OpenPostgres(cfg.DBString)

	personRepo := repo.NewPersonRepository(db)
	personService := service.NewPersonService(personRepo)

	srv := server.NewServer(cfg.ServerPort, personService)
	srv.Serve()
}

func makeLog(w io.Writer, level slog.Level) *slog.Logger {
	return slog.New(slog.NewTextHandler(w, &slog.HandlerOptions{
		AddSource: true,
		Level:     &level,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.SourceKey {
				source, _ := a.Value.Any().(*slog.Source)
				if source != nil {
					source.File = filepath.Base(source.File)
				}
			}
			return a
		},
	}))
}
