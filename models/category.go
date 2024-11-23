package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	CategoryName string `json:"CategoryName" binding:"required";gorm:"unique;not null"`
	Description  string `gorm:"type:text; null"`
}

type CategoryResponse struct {
	CategoryName string `json:"CategoryName"`
	Description  string `json:"Description"`
}
