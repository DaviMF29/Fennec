package main

import (
	"log"
	"net/http"
	"os"
	"github.com/DaviMF29/fennec/config"
	_ "github.com/DaviMF29/fennec/docs"
	"github.com/DaviMF29/fennec/routes"
	"github.com/gorilla/handlers" // Pacote para configurar o CORS
	"github.com/joho/godotenv"
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
	loadEnv()
	initSecretKey()
	router := routes.RegisterRoutes()

	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
		handlers.AllowCredentials(),
	)

	port := config.GetServerPort().Port
	log.Printf("Servidor iniciado na porta %s", port)
	log.Fatal(http.ListenAndServe(":"+port, corsHandler(router)))
}

var SECRET_KEY []byte

func loadEnv() {
	env := os.Getenv("ENV")

	if env == "production" {
		log.Println("Running in production mode. Skipping .env loading.")
		return
	}

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	log.Println(".env file loaded successfully.")
}

func initSecretKey() {
	SECRET_KEY = []byte(os.Getenv("SECRET_KEY"))
	if len(SECRET_KEY) == 0 {
		log.Fatal("SECRET_KEY is not defined in the environment variables")
	}
}
