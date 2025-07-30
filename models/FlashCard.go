package models

import "gorm.io/gorm"

type FlashCard struct {
	gorm.Model
	Front string `gorm:"not null"`
	Back  string `gorm:"not null"`

	UserId uint
}
