package dto

import "time"

type Response struct {
	Msg string `json:"message"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type User struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

// Pagination request
type PaginationQuery struct {
	Page     int `form:"page" json:"page" binding:"min=1"`
	PageSize int `form:"page_size" json:"page_size" binding:"min=1,max=100"`
}

// Pagination response
type Pagination struct {
	TotalRecords int64 `json:"total_records"`
	TotalPages   int   `json:"total_pages"`
	Page         int   `json:"page"`
	PageSize     int   `json:"page_size"`
	HasMore      bool  `json:"has_more"`
}

// Author DTO
type AuthorRequest struct {
	Name      string    `json:"name" binding:"required"`
	Biography string    `json:"biography"`
	BirthDate time.Time `json:"birth_date"`
}

// Book DTO
type BookRequest struct {
	Title          string `json:"title" binding:"required"`
	ISBN           string `json:"isbn" binding:"required"`
	PublicationYear int    `json:"publication_year"`
	Description    string `json:"description"`
	AuthorID       uint   `json:"author_id" binding:"required"`
}

// Review DTO
type ReviewRequest struct {
	Rating  int    `json:"rating" binding:"required,min=1,max=5"`
	Comment string `json:"comment"`
}
