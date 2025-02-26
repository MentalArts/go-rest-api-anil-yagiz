package models

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title          string   `json:"title" binding:"required"`
	ISBN           string   `json:"isbn" binding:"required" gorm:"unique"`
	PublicationYear int      `json:"publication_year"`
	Description    string   `json:"description"`
	AuthorID       uint     `json:"author_id" binding:"required"`
	Author         Author   `json:"author,omitempty" gorm:"foreignKey:AuthorID"`
	Reviews        []Review `json:"reviews,omitempty" gorm:"foreignKey:BookID"`
} 