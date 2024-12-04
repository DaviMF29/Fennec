package routes

import (
	"github.com/DaviMF29/fennec/handlers"
	"github.com/go-chi/chi"
)

func RegisterAuthRoutes() *chi.Mux {
	r := chi.NewRouter()

	r.Post("/", handlers.LoginHandler)
	return r
}
