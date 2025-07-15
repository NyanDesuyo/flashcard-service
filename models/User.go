package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UUID     uuid.UUID `gorm:"type:uuid;unique;not null"`
	Username string    `gorm:"unique;not null"`
	Password string    `gorm:"gorm:not null"`

	Flashcards []FlashCard
	Packs      []Pack
}
