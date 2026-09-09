package handler

import (
	"net/http"

	_ "github.com/kisalto/Feel-The-Night-Swagger/docs"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

// SetupRoutes registra todas as rotas da sua API
func SetupRoutes() http.Handler {
	mux := http.NewServeMux()

	// Rota do Swager
	mux.HandleFunc("GET /swagger/", httpSwagger.WrapHandler)

	// Exemplo de rotas
	mux.HandleFunc("GET /health", HealthCheckHandler)

	return mux
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status": "ok"}`))
}
