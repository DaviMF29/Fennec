package routes

import (
	"github.com/go-chi/chi"
	"github.com/DaviMF29/fennec/handlers"
)

func RegisterUserRoutes() *chi.Mux {

	r := chi.NewRouter()

	r.Get("/{id}", handlers.GetUserHandler)    
	r.Post("/", handlers.InsertUserHandler)
	r.Delete("/{id}", handlers.DeleteUserHandler)
	r.Patch("/{id}", handlers.UpdateUserHandler)    

	return r
}