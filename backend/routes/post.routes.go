package routes

import (
	"github.com/DaviMF29/fennec/handlers"
	"github.com/go-chi/chi"
)

func RegisterPostRoutes() *chi.Mux {
	r := chi.NewRouter()

	r.Post("/", handlers.InsertPostHandler)
	r.Get("/{id}", handlers.GetPostByIdHandler)
	r.Delete("/{id}", handlers.DeletePostByIdHandler)
	return r
}
