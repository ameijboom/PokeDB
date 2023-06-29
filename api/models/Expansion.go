package models

import (
	"pokedb/enums"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Expansion struct {
	ID                  uint                 `gorm:"primaryKey" json:"id"`
	SeriesId            uuid.UUID            `json:"series_id"`
	TrainerGalleryID    string               `json:"trainer_gallery_id"`
	SetNumber           uint                 `json:"set_number"`
	Name                string               `json:"name"`
	Abbreviation        string               `json:"abbreviation"`
	ExpansionType       enums.Expansion_Type `json:"expansion_type"`
	CardCount           uint                 `json:"card_count"`
	TrainerGalleryCount uint                 `json:"trainer_gallery_count"`
	SecretCardCount     uint                 `json:"secret_card_count"`
	Cards               []Card               `gorm:"foreignKey:Expansion" json:"cards"`
	ReleaseDate         time.Time            `json:"release_date"`
	CreatedAt           time.Time            `json:"created_at"`
	UpdatedAt           time.Time            `json:"updated_at"`
	DeletedAt           gorm.DeletedAt       `gorm:"index" json:"deleted_at"`
}
