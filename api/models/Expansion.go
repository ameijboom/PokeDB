package models

import (
	. "pokedb/enums"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Expansion struct {
	ID                  uint `gorm:"primaryKey"`
	SeriesId            uuid.UUID
	TrainerGalleryID    string
	SetNumber           uint
	Name                string
	Abbreviation        string
	ExpansionType       Expansion_Type
	CardCount           uint
	TrainerGalleryCount uint
	SecretCardCount     uint
	Cards               []Card `gorm:"foreignKey:Expansion"`
	ReleaseDate         time.Time
	CreatedAt           time.Time
	UpdatedAt           time.Time
	DeletedAt           gorm.DeletedAt `gorm:"index"`
}
