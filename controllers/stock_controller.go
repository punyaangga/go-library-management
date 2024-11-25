package controllers

import (
	"libraryManagement/config"
	"libraryManagement/models"
	"libraryManagement/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddStock(c *gin.Context) {
	var input models.Stock
	//  Validate input
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, "Invalid input", err)
		return
	}

	// Check if the product exists using the ProductExists method
	product := models.Product{}
	if !product.ProductExists(input.IdProduct) {
		utils.SendErrorResponse(c, http.StatusBadRequest, "Product not available", nil)
		return
	}

	// save data to database with checking data
	if err := config.DB.Create(&input).Error; err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, "Failed to create Stock", err)
		return
	}
	//response success
	responseData := models.StockResponse{
		StockBarcode:  input.StockBarcode,
		StockQty:      input.StockQty,
		StockLocation: input.StockLocation,
	}
	utils.SendSuccessResponse(c, http.StatusOK, "Stock add successfully", responseData)
}

func UpdateStock(c *gin.Context) {
	var input models.Stock
	StockID := c.Param("id")

	// Get data from database
	if err := config.DB.First(&input, StockID).Error; err != nil {
		utils.SendErrorResponse(c, http.StatusNotFound, "Stock not found", err)
		return
	}

	// Validate input
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, "Invalid input", err)
		return
	}

	// Update data in database
	if err := config.DB.Model(&input).Updates(input).Error; err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, "Failed to update stock", err)
		return
	}

	//response success
	responseData := models.StockResponse{
		StockBarcode:  input.StockBarcode,
		StockQty:      input.StockQty,
		StockLocation: input.StockLocation,
	}
	utils.SendSuccessResponse(c, http.StatusOK, "Stock updated successfully", responseData)
}

func GetStocks(c *gin.Context) {
	var id *string
	// get ID form URL
	idParam := c.Param("id")
	if idParam != "" {
		id = &idParam // Save ID to pointer
	}

	stocks, err := models.GetStocksWithProducts(config.DB, id)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusNotFound, err.Error(), nil)
		return
	}

	utils.SendSuccessResponse(c, http.StatusOK, "Stocks retrieved successfully", stocks)
}
