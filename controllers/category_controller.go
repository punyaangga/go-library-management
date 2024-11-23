package controllers

import (
	"libraryManagement/config"
	"libraryManagement/models"
	"libraryManagement/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddCategory(c *gin.Context) {
	var input models.Category
	// Validate input
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, "Invalid input", err)
		return
	}

	// save data to database with checking data
	if err := config.DB.Create(&input).Error; err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, "Failed to create category", err)
		return
	}
	//response success
	responseData := models.Category{
		CategoryName: input.CategoryName,
		Description:  input.Description,
	}
	utils.SendSuccessResponse(c, http.StatusOK, "Category added successfully", responseData)
}
