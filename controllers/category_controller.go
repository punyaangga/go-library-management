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
	responseData := models.CategoryResponse{
		CategoryName: input.CategoryName,
		Description:  input.Description,
	}
	utils.SendSuccessResponse(c, http.StatusOK, "Category add successfully", responseData)
}

func UpdateCategory(c *gin.Context) {
	var category models.Category
	var input models.Category
	categoryID := c.Param("id")
	// Get data from database
	if err := config.DB.First(&category, categoryID).Error; err != nil {
		utils.SendErrorResponse(c, http.StatusNotFound, "Category not found", err)
		return
	}

	// Validate input
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, "Invalid input", err)
		return
	}

	// Update data in database
	if err := config.DB.Model(&category).Updates(input).Error; err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, "Failed to update category", err)
		return
	}

	//response success
	responseData := models.CategoryResponse{
		CategoryName: input.CategoryName,
		Description:  input.Description,
	}
	utils.SendSuccessResponse(c, http.StatusOK, "Category updated successfully", responseData)
}
