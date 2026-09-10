package dto

type CreateCharacterInput struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	ImageURL    string `json:"imageURL"`
	Type        string `json:"type" binding:"required"`
}

type CharacterIDInput struct {
	id uint `uri:"id" binding:"required"`
}

type UpdateCharacter struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	ImageURL    string `json:"imageURL"`
	Type        string `json:"type"`
}

type CharErrorResponse struct {
	Error string `json:"error" example:"Personagem não encontrado"`
}

type CharacterResponse struct {
	CharacterID uint   `json:"characterID"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ImageURL    string `json:"imageURL"`
	Type        string `json:"type"`
}
