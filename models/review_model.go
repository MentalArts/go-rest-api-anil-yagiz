package models

import (
	"time"

	"gorm.io/gorm"
)

type Review struct {
	gorm.Model
	Rating     int       `json:"rating" binding:"required,min=1,max=5"`
	Comment    string    `json:"comment"`
	DatePosted time.Time `json:"date_posted" gorm:"default:CURRENT_TIMESTAMP"`
	BookID     uint      `json:"book_id" binding:"required"`
	Book       Book      `json:"book,omitempty" gorm:"foreignKey:BookID"`
} 