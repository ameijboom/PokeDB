package main

import (
	"pokedb/models"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load("../.env")
	models.ConnectDatabase()
}
