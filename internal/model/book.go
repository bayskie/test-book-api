package model

import (
	"database/sql"
	"time"
)

type Book struct {
	ID        uint         `gorm:"primarykey" json:"id"`
	Title     string       `json:"title" validate:"required,min=2"`
	Author    string       `json:"author" validate:"required"`
	Year      int          `json:"year" validate:"required,gte=0,lte=3000"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	DeletedAt sql.NullTime `gorm:"index" json:"-"`
}
