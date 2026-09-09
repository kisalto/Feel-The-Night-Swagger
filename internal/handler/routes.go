package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/kisalto/Feel-The-Night-Swagger/docs"
	"github.com/kisalto/Feel-The-Night-Swagger/internal/database"
	"github.com/kisalto/Feel-The-Night-Swagger/internal/services"
)

// SetupRoutes registra todas as rotas da sua API
func SetupRoutes() *gin.Engine {
	router := gin.Default()

	// Inicialização de serviços e handlers
	userService := services.NewUserService(database.DB) // Assumindo que seu pacote database expõe a variável 'DB' do GORM
	userHandler := NewUserHandler(userService)

	// Rota do Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Rotas gerais
	router.GET("/health", HealthCheckHandler)

	// Rotas de Usuários
	router.POST("/users", userHandler.CreateUser)
	router.GET("/users/:id", userHandler.GetUserById)

	return router
}

func HealthCheckHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
