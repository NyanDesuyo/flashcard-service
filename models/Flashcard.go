package models

import "gorm.io/gorm"

type Flashcard struct {
	gorm.Model
	Front  string `gorm:"not null"`
	Back   string `gorm:"not null"`
	UserId uint
}
