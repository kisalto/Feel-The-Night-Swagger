package models

import (
	"time"
)

// User table
type User struct {
	UserID           uint      `gorm:"primaryKey;column:user_id" json:"user_id"`
	Nickname         string    `gorm:"size:30;not null;column:nickname" json:"nickname"`
	Email            string    `gorm:"size:50;not null;column:email" json:"email"`
	DiscordID        string    `gorm:"size:25;column:dc_id" json:"discord_id"`
	Password         string    `gorm:"size:30;not null;column:password" json:"password"`
	RegistrationDate time.Time `gorm:"type:date;default:CURRENT_DATE;column:registration_date" json:"registration_date"`
	EventCount       int       `gorm:"default:0;column:event_count" json:"event_count"`
	GuideCount       int       `gorm:"default:0;column:guide_count" json:"guide_count"`
	IsModerator      bool      `gorm:"default:false;column:is_moderator" json:"is_moderator"`
	IsVeteran        bool      `gorm:"default:false;column:is_veteran" json:"is_veteran"`

	Events []Event `gorm:"foreignKey:UserID" json:"events,omitempty"`
	Guides []Guide `gorm:"foreignKey:UserID" json:"guides,omitempty"`
}

func (User) TableName() string { return "User" }

// Character table
type Character struct {
	CharacterID uint   `gorm:"primaryKey;column:character_id" json:"character_id"`
	Name        string `gorm:"size:20;not null;column:name" json:"name"`
	Description string `gorm:"size:255;not null;column:description" json:"description"`
	ImageURL    string `gorm:"size:255;column:image_url" json:"image_url"`
	Type        string `gorm:"size:15;not null;column:type" json:"type"`

	Guides []Guide `gorm:"foreignKey:CharacterID" json:"guides,omitempty"`
}

func (Character) TableName() string { return "character" }

// Guide table
type Guide struct {
	GuideID      uint      `gorm:"primaryKey;column:guide_id" json:"guide_id"`
	Title        string    `gorm:"size:50;not null;column:title" json:"title"`
	BannerURL    string    `gorm:"size:255;column:banner_url" json:"banner_url"`
	Type         string    `gorm:"size:30;column:type" json:"type"`
	Description  string    `gorm:"size:50;not null;column:description" json:"description"`
	Link         string    `gorm:"size:2083;not null;column:link" json:"link"`
	CreationDate time.Time `gorm:"type:date;default:CURRENT_DATE;column:creation_date" json:"creation_date"`
	Likes        int       `gorm:"default:0;column:likes" json:"likes"`
	Dislikes     int       `gorm:"default:0;column:dislikes" json:"dislikes"`

	CharacterID *uint `gorm:"column:fk_character_id" json:"character_id,omitempty"`
	UserID      uint  `gorm:"column:fk_user_id;not null" json:"user_id"`

	User      *User      `gorm:"foreignKey:UserID;references:UserID;constraint:OnDelete:CASCADE" json:"user,omitempty"`
	Character *Character `gorm:"foreignKey:CharacterID;references:CharacterID;constraint:OnDelete:SET NULL" json:"character,omitempty"`
}

func (Guide) TableName() string { return "guide" }

// Event table
type Event struct {
	EventID     uint      `gorm:"primaryKey;column:event_id" json:"event_id"`
	Title       string    `gorm:"size:75;not null;column:title" json:"title"`
	Description string    `gorm:"size:255;not null;column:description" json:"description"`
	BannerURL   string    `gorm:"size:255;column:banner_url" json:"banner_url"`
	Day         time.Time `gorm:"type:date;column:day" json:"day"`

	UserID uint  `gorm:"column:fk_user_id;not null" json:"user_id"`
	User   *User `gorm:"foreignKey:UserID;references:UserID;constraint:OnDelete:CASCADE" json:"user,omitempty"`
}

func (Event) TableName() string { return "event" }

// LastEvent table
type LastEvent struct {
	LastEventID uint   `gorm:"primaryKey;column:last_event_id" json:"last_event_id"`
	Title       string `gorm:"size:75;not null;column:title" json:"title"`

	EventID uint   `gorm:"column:fk_event_id;not null" json:"event_id"`
	Event   *Event `gorm:"foreignKey:EventID;references:EventID;constraint:OnDelete:CASCADE" json:"event,omitempty"`
}

func (LastEvent) TableName() string { return "last_event" }
