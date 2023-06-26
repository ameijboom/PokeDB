package main

import (
	"fmt"
	"os"
	"pokedb/enums"
	"pokedb/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	godotenv.Load("../.env")
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

	db.AutoMigrate(&models.Series{})
	db.AutoMigrate(&models.Expansion{})
	db.AutoMigrate(&models.Card{})

	result := db.Create(&models.Card{
		ID:          138,
		Count:       1,
		Name:        "Lugia V",
		HP:          220,
		Type:        enums.Colourless,
		Stage:       enums.Basic,
		Variant:     enums.V,
		IsTG:        false,
		Illustrator: "Mitsuhiro Arita",
		Expansion:   195,
	})

	fmt.Println(result.Error)
	fmt.Println(result.RowsAffected)
}
