package server

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func (s *Server) Router() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Route("/people", func(r chi.Router) {
		r.Post("/", s.PersonHandler.CreatePerson)
		r.Get("/", s.PersonHandler.GetByFilters)
		r.Get("/{id}", s.PersonHandler.GetByID)
		r.Delete("/{id}", s.PersonHandler.DeleteByID)
		r.Patch("/{id}", s.PersonHandler.UpdateByID)
	})
	return r
}
