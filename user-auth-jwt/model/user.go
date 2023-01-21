package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Id uint `gorm:"primaryKey"`
	Username string `gorm:"uniqueIndex;not null" json:"username"`
	Email string `gorm:"uniqueIndex;not null" json:"email"`
	Password string `gorm:"not null" json:"password"`
	
}