package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	dto "github.com/kisalto/Feel-The-Night-Swagger/internal/dto"
	"github.com/kisalto/Feel-The-Night-Swagger/internal/models"
	"github.com/kisalto/Feel-The-Night-Swagger/internal/services"
)

type UserHandler struct {
	userService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

// CreateUser godoc
// @Summary      Criar um novo usuário
// @Description  Cria um usuário com os dados informados no body.
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        user  body      dto.CreateUserInput  true  "Dados do usuário"
// @Success      201   {object}  models.User
// @Failure      400   {object}  map[string]string
// @Router       /users [post]
func (h *UserHandler) CreateUser(c *gin.Context) {
	var input dto.CreateUserInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		Nickname:  input.Nickname,
		Email:     input.Email,
		DiscordID: input.DiscordID,
		Password:  input.Password,
	}

	if err := h.userService.CreateUser(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}

// GetUserById godoc
// @Summary      Buscar usuário por ID
// @Tags         User
// @Produce      json
// @Param        id   path      int  true  "ID do usuário" minimum(1)
// @Success      200  {object}  dto.UserResponse
// @Failure      404  {object}  dto.ErrorResponse
// @Router       /users/{id} [get]
func (h *UserHandler) GetUserById(c *gin.Context) {
	var input dto.UserIDInput

	if err := c.ShouldBindUri(&input); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "id inválido"})
		return
	}

	user, err := h.userService.GetUserById(input.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, dto.ErrorResponse{Error: "usuário não encontrado"})
			return
		}
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	response := dto.UserResponse{
		UserID:           user.UserID,
		Nickname:         user.Nickname,
		Email:            user.Email,
		DiscordID:        user.DiscordID,
		RegistrationDate: user.RegistrationDate,
		EventCount:       user.EventCount,
		GuideCount:       user.GuideCount,
		IsModerator:      user.IsModerator,
		IsVeteran:        user.IsVeteran,
	}

	c.JSON(http.StatusOK, response)
}

// DeleteUserById godoc
// @Summary      Deletar usuário por ID
// @Tags         User
// @Produce      json
// @Param        id   path      int  true  "ID do usuário" minimum(1)
// @Success      200  {object}  map[string]string
// @Failure      404  {object}  dto.ErrorResponse
// @Router       /users/{id} [delete]
func (h *UserHandler) DeleteUserById(c *gin.Context) {
	var input dto.UserIDInput

	if err := c.ShouldBindUri(&input); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "id inválido"})
		return
	}

	if err := h.userService.DeleteUserById(input.ID); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, dto.ErrorResponse{Error: "usuário não encontrado"})
			return
		}
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	// Retorna 200 OK com uma mensagem em formato JSON
	c.JSON(http.StatusOK, gin.H{"message": "usuário deletado com sucesso"})
}
