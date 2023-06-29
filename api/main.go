package main

import (
	"pokedb/controllers"
	"pokedb/models"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load("../.env")
	router := gin.Default()
	models.ConnectDatabase()

	router.GET("/api/series", controllers.GetSeries)
	router.GET("/api/series/:id", controllers.GetSeriesByID)
	router.POST("/api/series", controllers.CreateSeries)

	router.Run()
}
