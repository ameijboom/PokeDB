package models

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("POKEDB_POSTGRES_HOST"),
		os.Getenv("POKEDB_POSTGRES_USER"),
		os.Getenv("POKEDB_POSTGRES_PASS"),
		os.Getenv("POKEDB_POSTGRES_DB"),
		os.Getenv("POKEDB_POSTGRES_PORT"))

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect to database")
	}

	db.AutoMigrate(&Series{})
	db.AutoMigrate(&Expansion{})
	db.AutoMigrate(&Card{})

	if err != nil {
		panic("failed to migrate models")
	}

	DB = db
}
