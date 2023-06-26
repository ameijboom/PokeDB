package models

import "gorm.io/gorm"

type Attack struct {
	gorm.Model
	Name        string
	Description string
	Damage      uint
	Cost        Energy
}
