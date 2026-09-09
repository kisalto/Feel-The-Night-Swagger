package services

import (
	"errors"

	"github.com/kisalto/Feel-The-Night-Swagger/internal/models"
	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{db: db}
}

// Change it so userID, registrationDate, isModerator, isVeteran, eventCount and guideCount
// cant be passed througth the createUser post.
// userID will be added automaticaly  by the DB
// registrationDate will get date.now
// isModerator will always be False
// isVeteran will always be false
// eventCount and guideCount will be 0 when creating account
func (s *UserService) CreateUser(user *models.User) error {
	var existingUser models.User

	if err := s.db.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
		return errors.New("Email ja cadastrado")
	}

	return s.db.Create(user).Error
}

func (s *UserService) GetUserById(id uint) (*models.User, error) {
	var user models.User
	if err := s.db.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("Usuário nao econtrado")
		}
		return nil, err
	}
	return &user, nil
}

//func GetAllUsers()
