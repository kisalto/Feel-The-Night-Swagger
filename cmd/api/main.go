package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/kisalto/Feel-The-Night-Swagger/internal/database"
	"github.com/kisalto/Feel-The-Night-Swagger/internal/handler"
)

// @title           Feel The Night API
// @version         1.0
// @description     API para gerenciamento de guias, eventos e personagens.
// @host            localhost:8080
// @BasePath        /
func main() {
	err := godotenv.Load()
	if err != nil {
		slog.Warn("erro ao carregar o arquivo .env", "error", err)
	}

	// 1. Conecta no banco de dados e executa o AutoMigrate
	log.Println("Iniciando conexão com o banco de dados...")
	if err := database.Connect(); err != nil {
		log.Fatalf("Erro crítico ao inicializar o banco: %v", err)
	}

	// 2. Configura e carrega as rotas
	routes := handler.SetupRoutes()

	// 3. Inicia o servidor HTTP
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}
	serverPort := fmt.Sprintf(":%s", port)
	log.Printf("Servidor rodando na porta %s\n", serverPort)

	if err := http.ListenAndServe(serverPort, routes); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
