package models

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title       string
	Publisher   string
	Author      string
	PublishYear string
	UserID      uint
}
