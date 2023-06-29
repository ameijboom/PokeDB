package controllers

import (
	"net/http"
	"pokedb/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// GET Series
func GetSeries(ctx *gin.Context) {
	var series models.Series
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
