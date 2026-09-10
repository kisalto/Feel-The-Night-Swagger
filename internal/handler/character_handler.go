package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kisalto/Feel-The-Night-Swagger/internal/dto"
	"github.com/kisalto/Feel-The-Night-Swagger/internal/models"
	"github.com/kisalto/Feel-The-Night-Swagger/internal/services"
)

type CharacterHandler struct {
	characterService *services.CharacterService
}

func NewCharacterHandler(characterService *services.CharacterService) *CharacterHandler {
	return &CharacterHandler{characterService: characterService}
}

// CreateCharacter godoc
// @Summary Criar um novo personagem
// @Description Cria um novo personagem com os dados informados do body
// @Tags Character
// @Accept json
// @Produce json
// @Param character body dto.CreateCharacterInput true "Dados do Personagem"
// @Success 201 {object} models.Character
// @Router /char [post]
func (h *CharacterHandler) CreateCharacter(c *gin.Context) {
	var input dto.CreateCharacterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	character := models.Character{
		Name:        input.Name,
		Description: input.Description,
		ImageURL:    input.ImageURL,
		Type:        input.Type,
	}

	if err := h.characterService.CreateCharacter(&character); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, character)
}
