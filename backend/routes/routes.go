package routes

import (
	"net/http"

	"github.com/go-chi/chi"
	httpSwagger "github.com/swaggo/http-swagger"
)

func RegisterRoutes() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("api fennec"))
	})

	r.Mount("/api/user", RegisterUserRoutes())
	r.Mount("/api/login", RegisterAuthRoutes())
	r.Mount("/api/post", RegisterPostRoutes())


	r.Get("/swagger/*", httpSwagger.WrapHandler)

	return r
}

