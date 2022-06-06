package models

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title string `json:"title" binding:"required" gorm:"size:255"`
	Price int    `json:"price" binding:"required,number"`
}
