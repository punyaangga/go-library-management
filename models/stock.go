package models

import "gorm.io/gorm"

type Stock struct {
	gorm.Model
	IdProduct     int    `json:"IdProduct" binding:"required";gorm:"unique;not null"`
	StockBarcode  string `json:"StockBarcode" binding:"required";gorm:"unique;not null"`
	StockQty      int    `json:"StockQty" binding:"required";gorm:"unique;not null"`
	StockLocation string `gorm:"type:text; null"`
}

type StockResponse struct {
	StockBarcode  string `json:"StockBarcode"`
	StockQty      string `json:"StockQty"`
	StockLocation string `json:"Description"`
}
