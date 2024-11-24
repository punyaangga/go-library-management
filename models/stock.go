package models

import (
	"libraryManagement/config"

	"gorm.io/gorm"
)

type Stock struct {
	gorm.Model
	IdProduct     int    `json:"IdProduct" binding:"required";gorm:"unique;not null"`
	StockBarcode  string `json:"StockBarcode" binding:"required";gorm:"unique;not null"`
	StockQty      int    `json:"StockQty" binding:"required";gorm:"unique;not null"`
	StockLocation string `gorm:"type:text; null"`
}

type StockResponse struct {
	StockBarcode  string `json:"StockBarcode"`
	StockQty      int    `json:"StockQty"`
	StockLocation string `json:"Description"`
}

// Method to check if product exists
func (p *Product) ProductExists(idProduct int) bool {

	var product Product
	// Try to find the product by id
	if err := config.DB.First(&product, idProduct).Error; err != nil {
		return false // Product not found
	}
	return true // Product exists
}
