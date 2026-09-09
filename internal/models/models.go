package models

import (
	"time"
)

// Tabela de usuário
type Usuario struct {
	UserID     uint      `gorm:"primaryKey;column:user_id"`
	Apelido    string    `gorm:"size:30;not null"`
	Email      string    `gorm:"size:50;not null"`
	DcID       string    `gorm:"size:25;column:dc_id"`
	Senha      string    `gorm:"size:30;not null"`
	DataReg    time.Time `gorm:"type:date;default:CURRENT_DATE"`
	EventosQnt int       `gorm:"default:0"`
	GuiasQnt   int       `gorm:"default:0"`
	IsMod      bool      `gorm:"default:false"`
	IsVet      bool      `gorm:"default:false"`
}

func (Usuario) TableName() string { return "usuario" }

// Tabela de personagem
type Personagem struct {
	PersonagemID uint   `gorm:"primaryKey;column:personagem_id"`
	Nome         string `gorm:"size:20;not null"`
	Descricao    string `gorm:"size:255;not null"`
	ImgURL       string `gorm:"size:255"`
	Tipo         string `gorm:"size:15;not null"`
}

func (Personagem) TableName() string { return "personagem" }

// Tabela de guia
type Guia struct {
	GuiaID       uint        `gorm:"primaryKey;column:guia_id"`
	Titulo       string      `gorm:"size:50;not null"`
	BannerURL    string      `gorm:"size:255"`
	Tipo         string      `gorm:"size:30"`
	Descricao    string      `gorm:"size:50;not null"`
	Link         string      `gorm:"size:2083;not null"`
	DataCr       time.Time   `gorm:"type:date;default:CURRENT_DATE"`
	Likes        int         `gorm:"default:0"`
	Dislikes     int         `gorm:"default:0"`
	FkUserID     uint        `gorm:"column:fk_user_id;not null"`
	FkPersonagem *uint       `gorm:"column:fk_personagem_id"`
	Usuario      Usuario     `gorm:"foreignKey:FkUserID;constraint:OnDelete:CASCADE"`
	Personagem   *Personagem `gorm:"foreignKey:FkPersonagem;constraint:OnDelete:SET NULL"`
}

func (Guia) TableName() string { return "guia" }

// Tabela de eventos
type Eventos struct {
	EventosID uint      `gorm:"primaryKey;column:eventos_id"`
	Titulo    string    `gorm:"size:75;not null"`
	Descricao string    `gorm:"size:255;not null"`
	BannerURL string    `gorm:"size:255"`
	Dia       time.Time `gorm:"type:date"`
	FkUserID  uint      `gorm:"column:fk_user_id;not null"`
	Usuario   Usuario   `gorm:"foreignKey:FkUserID;constraint:OnDelete:CASCADE"`
}

func (Eventos) TableName() string { return "eventos" }

// Tabela de ultimo evento
type UltimoEvento struct {
	UeID        uint    `gorm:"primaryKey;column:ue_id"`
	Titulo      string  `gorm:"size:75;not null"`
	FkEventosID uint    `gorm:"column:fk_eventos_id;not null"`
	Evento      Eventos `gorm:"foreignKey:FkEventosID;constraint:OnDelete:CASCADE"`
}

func (UltimoEvento) TableName() string { return "ultimo_evento" }
