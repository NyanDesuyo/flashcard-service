package models

import "gorm.io/gorm"

type Pack struct {
	gorm.Model
	Name   string `gorm:"not null"`
	UserID uint

	FlashCards []FlashCard `gorm:"many2many:pack_flashcards;"`
}
