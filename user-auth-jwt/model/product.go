package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model

	Id uint `gorm:"primaryKey"`
	Title string `gorm:"not null" json:"title"`
	Description string `gorm:"not null" json:"description"`
	Amount int `gorm:"not null" json:"amount"`
	
}