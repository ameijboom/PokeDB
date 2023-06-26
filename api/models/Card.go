package models

import (
	"pokedb/enums"
	"time"

	"gorm.io/gorm"
)

type Card struct {
	ID               uint `gorm:"primaryKey"`
	Name             string
	HP               uint
	Type             enums.PokemonType
	Stage            enums.Stage
	Ability          Ability
	Attacks          []Attack
	Weakness         Energy
	Resistance       Energy
	Retreat          Energy
	Variant          enums.Variant
	IsTG             bool
	TrainerGalleryID string
	Expansion        Expansion
	Illustrator      string
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt `gorm:"index"`
}
