package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Series struct {
	ID         uuid.UUID      `gorm:"primaryKey" json:"id"`
	Name       string         `json:"name"`
	Expansions []Expansion    `json:"expansions"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
