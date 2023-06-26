package models

import (
	"pokedb/enums"
	"time"

	"gorm.io/gorm"
)

type Card struct {
	ID               uint `gorm:"primaryKey"`
	Count            uint
	Name             string
	HP               uint
	Type             enums.PokemonType
	Stage            enums.Stage
	Variant          enums.Variant
	IsTG             bool
	TrainerGalleryID string
	Expansion        uint
	Illustrator      string
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt `gorm:"index"`
}
