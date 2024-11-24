package models

import "time"

type ExpiredToken struct {
	Id        uint      `gorm:"primaryKey" json:"id"`
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"created_at"`
}
