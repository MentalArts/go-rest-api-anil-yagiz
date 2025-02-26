package models

import (
	"time"

	"gorm.io/gorm"
)

type Author struct {
	gorm.Model
	Name      string    `json:"name" binding:"required"`
	Biography string    `json:"biography"`
	BirthDate time.Time `json:"birth_date"`
	Books     []Book    `json:"books,omitempty" gorm:"foreignKey:AuthorID"`
}
