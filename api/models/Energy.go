package models

import (
	"pokedb/enums"

	"gorm.io/gorm"
)

type Energy struct {
	gorm.Model
	Type   enums.PokemonType
	Amount uint
}
