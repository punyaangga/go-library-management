package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	CategoryName string `gorm:"unique;not null"`
	Description  string `gorm:"type:text; null"`
}
