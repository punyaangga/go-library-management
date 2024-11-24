package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	IdCategory         int    `json:"IdCategory" binding:"required";gorm:"unique;not null"`
	NameProduct        string `json:"NameProduct" binding:"required";gorm:"unique;not null"`
	DesciprtionProduct string `gorm:"type:text; null"`
}

type ProductResponse struct {
	IdCategory    string `json:"StockBarcode"`
	NameProduct   string `json:"StockQty"`
	StockLocation string `json:"Description"`
}
