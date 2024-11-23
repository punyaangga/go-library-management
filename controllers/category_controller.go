package controllers

import (
	"libraryManagement/config"
	"libraryManagement/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddCategory(c *gin.Context) {
	var input models.Category
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create category"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Category Add Successfully"})
}
