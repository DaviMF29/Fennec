package main

import (
	"log"
	"net/http"
	"github.com/DaviMF29/wombat/config"
	"github.com/DaviMF29/wombat/routes"
)

func main() {
	err := config.Load()
	if err != nil {
		log.Fatalf("Erro ao carregar configuração: %v", err)
	}

	router := routes.RegisterRoutes()

	port := config.GetServerPort().Port
	log.Printf("Servidor iniciado na porta %s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
