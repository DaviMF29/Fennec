package routes

import (
	"net/http"

	"github.com/go-chi/chi"
)

func RegisterRoutes() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("api wombat"))
	})



	r.Mount("/api/user", RegisterUserRoutes())
	r.Mount("/api/login", RegisterAuthRoutes())

	return r
}

