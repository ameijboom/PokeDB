package models

import "gorm.io/gorm"

type Ability struct {
	gorm.Model
	Name        string
	Description string
}
