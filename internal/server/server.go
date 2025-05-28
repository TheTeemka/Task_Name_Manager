package server

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/TheTeemka/TaskNameManager/internal/service"
)

type Server struct {
	Port          string
	PersonHandler *PersonHandler
}

func NewServer(Port string, personService *service.PersonService) *Server {
	return &Server{
		Port:          Port,
		PersonHandler: NewPersonHandler(personService),
	}
}

func (s *Server) Serve() {
	srv := http.Server{
		Addr:    s.Port,
		Handler: s.Router(),
	}

	shutdown := make(chan error)
	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, os.Interrupt, os.Kill)

		slog.Info("Server is shutting down", "cause:", <-quit)
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		err := srv.Shutdown(ctx)
		shutdown <- err
	}()

	slog.Info("Server is starting", "port", s.Port)
	err := srv.ListenAndServe()
	if err != http.ErrServerClosed {
		log.Fatal(err)
	}

	err = <-shutdown
	if err != nil {
		log.Fatal(err)
	}

	slog.Info("Server is successfully closed")
}
