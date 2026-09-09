package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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
// @Description  Retorna os dados de um usuário pelo ID
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        user  body      dto.CreateUserInput  true  "Dados do usuário"
// @Success      200  {object}  models.User
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Router       /users/{id} [get]
func (h *UserHandler) GetUserById(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	user, err := h.userService.GetUserById(uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

// func GetAllUsers(c *gin.Context) {
// 	c.IndentedJSON(http.StatusOK, users)
// }
