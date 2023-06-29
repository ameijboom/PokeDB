package controllers

import (
	"fmt"
	"net/http"
	"pokedb/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// GET Series
func GetSeries(ctx *gin.Context) {
	var series []models.Series
	models.DB.Find(&series)

	ctx.JSON(http.StatusOK, gin.H{"data": series})
}

// Get Series by ID
func GetSeriesByID(ctx *gin.Context) {
	var series models.Series
	id, _ := uuid.Parse(ctx.Param("id"))

	if err := models.DB.Preload("Expansions").Where(&models.Series{ID: id}).First(&series).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
	}

	ctx.JSON(http.StatusOK, gin.H{"data": series})
}

// POST Series
func CreateSeries(ctx *gin.Context) {
	var input models.SeriesInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	count := int64(0)

	models.DB.Model(&models.Series{}).Where(models.Series{Name: input.Name}).Count(&count)

	if count != 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Entry with name \"%s\" already exists.", input.Name)})
		return
	}

	series := models.Series{
		ID:   uuid.New(),
		Name: input.Name,
	}

	models.DB.Create(&series)

	ctx.JSON(http.StatusOK, gin.H{"data": series})
}
