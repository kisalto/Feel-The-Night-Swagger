package dto

// CreateUserInput representa apenas os dados permitidos no POST
type CreateUserInput struct {
	Nickname  string `json:"nickname" binding:"required"`
	Email     string `json:"email" binding:"required"`
	DiscordID string `json:"discordID"`
	Password  string `json:"password" binding:"required"`
}
