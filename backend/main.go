package main

import (
	"log"
	"net/http"
	"github.com/DaviMF29/fennec/config"
	"github.com/DaviMF29/fennec/routes"
	_"github.com/DaviMF29/fennec/docs"
)


// @title API Fennec
// @version 1.0
// @description Fennec is a social platform designed for developers to connect, share knowledge, and showcase their projects. Built with Golang on the backend and React on the frontend, Fennec delivers high performance, scalability, and an engaging developer-centric experience.
// @host https://wombat-production-e2c6.up.railway.app/
// @BasePath /api


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
