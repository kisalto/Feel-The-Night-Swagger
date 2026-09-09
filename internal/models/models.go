package models

import (
	"time"
)

// User table
type User struct {
	UserID           uint      `gorm:"primaryKey;column:user_id"`
	Nickname         string    `gorm:"size:30;not null"`
	Email            string    `gorm:"size:50;not null"`
	DiscordID        string    `gorm:"size:25;column:dc_id"`
	Password         string    `gorm:"size:30;not null"`
	RegistrationDate time.Time `gorm:"type:date;default:CURRENT_DATE"`
	EventCount       int       `gorm:"default:0"`
	GuideCount       int       `gorm:"default:0"`
	IsModerator      bool      `gorm:"default:false"`
	IsVeteran        bool      `gorm:"default:false"`
}

func (User) TableName() string { return "User" }

// Character table
type Personagem struct {
	CharacterID uint   `gorm:"primaryKey;column:personagem_id"`
	Name        string `gorm:"size:20;not null"`
	Description string `gorm:"size:255;not null"`
	ImageURL    string `gorm:"size:255"`
	Type        string `gorm:"size:15;not null"`
}

func (Personagem) TableName() string { return "personagem" }

// Guide table
type Guia struct {
	GuideID      uint        `gorm:"primaryKey;column:guia_id"`
	Title        string      `gorm:"size:50;not null"`
	BannerURL    string      `gorm:"size:255"`
	Type         string      `gorm:"size:30"`
	Description  string      `gorm:"size:50;not null"`
	Link         string      `gorm:"size:2083;not null"`
	CreationDate time.Time   `gorm:"type:date;default:CURRENT_DATE"`
	Likes        int         `gorm:"default:0"`
	Dislikes     int         `gorm:"default:0"`
	UserID       uint        `gorm:"column:fk_user_id;not null"`
	CharacterID  *uint       `gorm:"column:fk_personagem_id"`
	User         User        `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Character    *Personagem `gorm:"foreignKey:CharacterID;constraint:OnDelete:SET NULL"`
}

func (Guia) TableName() string { return "guia" }

// Events table
type Eventos struct {
	EventID     uint      `gorm:"primaryKey;column:eventos_id"`
	Title       string    `gorm:"size:75;not null"`
	Description string    `gorm:"size:255;not null"`
	BannerURL   string    `gorm:"size:255"`
	Day         time.Time `gorm:"type:date"`
	UserID      uint      `gorm:"column:fk_user_id;not null"`
	User        User      `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}

func (Eventos) TableName() string { return "eventos" }

// Last event table
type UltimoEvento struct {
	LastEventID uint    `gorm:"primaryKey;column:ue_id"`
	Title       string  `gorm:"size:75;not null"`
	EventID     uint    `gorm:"column:fk_eventos_id;not null"`
	Event       Eventos `gorm:"foreignKey:EventID;constraint:OnDelete:CASCADE"`
}

func (UltimoEvento) TableName() string { return "ultimo_evento" }
