package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	CategoryName string `json:categoryName; gorm:"unique;not null"`
	Description  string `json:description;gorm:"type:text; null"`
}
