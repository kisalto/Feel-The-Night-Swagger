package handler

import (
	"net/http"
)

// SetupRoutes registra todas as rotas da sua API
func SetupRoutes() http.Handler {
	mux := http.NewServeMux()

	// Exemplo de rotas
	mux.HandleFunc("GET /health", HealthCheckHandler)

	return mux
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status": "ok"}`))
}
