package book

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Book struct {
	ID            uuid.UUID `gorm:"primarykey"`
	Title         string
	Author        string
	PublishedDate time.Time
	ImageURL      string
	Description   string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt
}

type Books []*Book
