package database

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/kisalto/Feel-The-Night-Swagger/internal/models"
)

var DB *gorm.DB

func Connect() error {
	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("erro ao carregar o arquivo .env: %w", err)
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Asia/Shanghai",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"),
	)

	DB, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		return fmt.Errorf("falha ao conectar no banco de dados: %w", err)
	}

	err = DB.AutoMigrate(
		&models.User{},
		&models.Character{},
		&models.Event{},
		&models.Guide{},
		&models.LastEvent{},
	)
	if err != nil {
		return fmt.Errorf("falha ao executar o AutoMigrate: %w", err)
	}

	fmt.Println("Conexão com o banco de dados e migrações concluídas com sucesso!")
	return nil
}
