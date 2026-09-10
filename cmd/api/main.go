package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	// Importe o pacote docs sem o "_" para poder acessar a variável SwaggerInfo
	"github.com/kisalto/Feel-The-Night-Swagger/docs"
	"github.com/kisalto/Feel-The-Night-Swagger/internal/database"
	"github.com/kisalto/Feel-The-Night-Swagger/internal/handler"
)

// @title           Feel The Night API
// @version         1.0
// @description     API para gerenciamento de guias, eventos e personagens.
// @BasePath        /
func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("WARN: Arquivo .env não encontrado, lendo variáveis de ambiente do sistema.")
	}

	// 1. Tratamento da Porta
	port := os.Getenv("PORT")
	if port == "" {
		port = os.Getenv("APP_PORT")
	}
	if port == "" {
		port = "8080"
	}

	// 2. Configuração Dinâmica do Swagger via docs.SwaggerInfo
	renderHost := os.Getenv("RENDER_EXTERNAL_HOSTNAME")

	if renderHost != "" {
		// Em produção no Render
		docs.SwaggerInfo.Host = renderHost
		docs.SwaggerInfo.Schemes = []string{"https"}
	} else {
		// Em desenvolvimento local
		docs.SwaggerInfo.Host = fmt.Sprintf("localhost:%s", port)
		docs.SwaggerInfo.Schemes = []string{"http"}
	}

	// 3. Conexão com o Banco de Dados
	log.Println("Iniciando conexão com o banco de dados...")
	if err := database.Connect(); err != nil {
		log.Fatalf("Erro crítico ao inicializar o banco: %v", err)
	}

	// 4. Carrega Rotas
	routes := handler.SetupRoutes()

	// 5. Inicia o Servidor
	serverPort := fmt.Sprintf(":%s", port)
	log.Printf("Servidor rodando na porta %s\n", serverPort)

	if err := http.ListenAndServe(serverPort, routes); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
