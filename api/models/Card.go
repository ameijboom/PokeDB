package models

import (
	"pokedb/enums"
	"time"

	"gorm.io/gorm"
)

type Card struct {
	ID               uint              `gorm:"primaryKey" json:"id"`
	Count            uint              `json:"count"`
	Name             string            `json:"name"`
	HP               uint              `json:"hp"`
	Type             enums.PokemonType `json:"type"`
	Stage            enums.Stage       `json:"stage"`
	Variant          enums.Variant     `json:"variant"`
	IsTG             bool              `json:"is_tg"`
	TrainerGalleryID string            `json:"trainer_gallery_id"`
	Expansion        uint              `json:"-"`
	Illustrator      string            `json:"illustrator"`
	CreatedAt        time.Time         `json:"created_at"`
	UpdatedAt        time.Time         `json:"updated_at"`
	DeletedAt        gorm.DeletedAt    `gorm:"index" json:"deleted_at"`
}
