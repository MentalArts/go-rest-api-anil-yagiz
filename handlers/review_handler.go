package handlers

import (
	"mentalartsapi/dto"
	"mentalartsapi/models"
	"mentalartsapi/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// CreateReview godoc
// @Summary Create a new review for a book
// @Description Create a new review for a book with the input payload
// @Tags reviews
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Param review body dto.ReviewRequest true "Review data"
// @Success 201 {object} models.Review
// @Failure 400 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /api/v1/books/{id}/reviews [post]
func CreateReview(c *gin.Context) {
	bookID := c.Param("id")
	var reviewRequest dto.ReviewRequest
	var review models.Review

	// Check if book exists
	var book models.Book
	if err := db.First(&book, bookID).Error; err != nil {
		c.JSON(http.StatusNotFound, dto.ErrorResponse{Error: "book not found"})
		return
	}

	if err := c.ShouldBindJSON(&reviewRequest); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
		return
	}

	review.Rating = reviewRequest.Rating
	review.Comment = reviewRequest.Comment
	review.DatePosted = time.Now()
	review.BookID = book.ID

	result := db.Create(&review)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: result.Error.Error()})
		return
	}

	// Load relations for response
	db.Preload("Book").First(&review, review.ID)

	c.JSON(http.StatusCreated, review)
}

// GetBookReviews godoc
// @Summary Get all reviews for a book
// @Description Get all reviews for a book with pagination
// @Tags reviews
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Param page query int false "Page number"
// @Param page_size query int false "Page size"
// @Success 200 {object} map[string]interface{}
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /api/v1/books/{id}/reviews [get]
func GetBookReviews(c *gin.Context) {
	bookID := c.Param("id")
	var reviews []models.Review
	var totalCount int64

	// Check if book exists
	var book models.Book
	if err := db.First(&book, bookID).Error; err != nil {
		c.JSON(http.StatusNotFound, dto.ErrorResponse{Error: "book not found"})
		return
	}

	pagination := utils.ParsePaginationQuery(c)
	
	// Count total records
	if err := db.Model(&models.Review{}).Where("book_id = ?", bookID).Count(&totalCount).Error; err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	// Get paginated reviews
	if err := utils.Paginate(db, &pagination).
		Where("book_id = ?", bookID).
		Find(&reviews).Error; err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	// Create pagination response
	paginationResponse := utils.CreatePaginationResponse(totalCount, pagination)

	// Create final response
	response := gin.H{
		"data":       reviews,
		"pagination": paginationResponse,
	}

	c.JSON(http.StatusOK, response)
}

// UpdateReview godoc
// @Summary Update a review
// @Description Update a review with the input payload
// @Tags reviews
// @Accept json
// @Produce json
// @Param id path int true "Review ID"
// @Param review body dto.ReviewRequest true "Review data"
// @Success 200 {object} models.Review
// @Failure 404 {object} dto.ErrorResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /api/v1/reviews/{id} [put]
func UpdateReview(c *gin.Context) {
	id := c.Param("id")
	var reviewRequest dto.ReviewRequest
	var review models.Review

	if err := db.First(&review, id).Error; err != nil {
		c.JSON(http.StatusNotFound, dto.ErrorResponse{Error: "review not found"})
		return
	}

	if err := c.ShouldBindJSON(&reviewRequest); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
		return
	}

	review.Rating = reviewRequest.Rating
	review.Comment = reviewRequest.Comment

	if err := db.Save(&review).Error; err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "internal server error"})
		return
	}

	c.JSON(http.StatusOK, review)
}

// DeleteReview godoc
// @Summary Delete a review
// @Description Delete a review by ID
// @Tags reviews
// @Accept json
// @Produce json
// @Param id path int true "Review ID"
// @Success 200 {object} dto.Response
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /api/v1/reviews/{id} [delete]
func DeleteReview(c *gin.Context) {
	id := c.Param("id")
	var review models.Review

	if err := db.First(&review, id).Error; err != nil {
		c.JSON(http.StatusNotFound, dto.ErrorResponse{Error: "review not found"})
		return
	}

	if err := db.Delete(&review).Error; err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.Response{Msg: "review deleted successfully"})
} 