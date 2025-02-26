package handlers

import (
	"mentalartsapi/dto"
	"mentalartsapi/models"
	"mentalartsapi/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateBook godoc
// @Summary Create a new book
// @Description Create a new book with the input payload
// @Tags books
// @Accept json
// @Produce json
// @Param book body dto.BookRequest true "Book data"
// @Success 201 {object} models.Book
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /api/v1/books [post]
func CreateBook(c *gin.Context) {
	var bookRequest dto.BookRequest
	var book models.Book

	if err := c.ShouldBindJSON(&bookRequest); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
		return
	}

	// Check if author exists
	var author models.Author
	if err := db.First(&author, bookRequest.AuthorID).Error; err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "author not found"})
		return
	}

	book.Title = bookRequest.Title
	book.ISBN = bookRequest.ISBN
	book.PublicationYear = bookRequest.PublicationYear
	book.Description = bookRequest.Description
	book.AuthorID = bookRequest.AuthorID

	result := db.Create(&book)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: result.Error.Error()})
		return
	}

	// Load relations for response
	db.Preload("Author").First(&book, book.ID)

	c.JSON(http.StatusCreated, book)
}

// GetAllBooks godoc
// @Summary Get all books
// @Description Get all books with pagination
// @Tags books
// @Accept json
// @Produce json
// @Param page query int false "Page number"
// @Param page_size query int false "Page size"
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} dto.ErrorResponse
// @Router /api/v1/books [get]
func GetAllBooks(c *gin.Context) {
	var books []models.Book
	var totalCount int64

	pagination := utils.ParsePaginationQuery(c)
	
	// Count total records
	if err := db.Model(&models.Book{}).Count(&totalCount).Error; err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	// Get paginated books with their author
	if err := utils.Paginate(db, &pagination).
		Preload("Author").
		Find(&books).Error; err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	// Create pagination response
	paginationResponse := utils.CreatePaginationResponse(totalCount, pagination)

	// Create final response
	response := gin.H{
		"data":       books,
		"pagination": paginationResponse,
	}

	c.JSON(http.StatusOK, response)
}

// GetBook godoc
// @Summary Get a book
// @Description Get a book by ID with author and reviews
// @Tags books
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Success 200 {object} models.Book
// @Failure 404 {object} dto.ErrorResponse
// @Router /api/v1/books/{id} [get]
func GetBook(c *gin.Context) {
	id := c.Param("id")
	var book models.Book

	if err := db.Preload("Author").Preload("Reviews").First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, dto.ErrorResponse{Error: "book not found"})
		return
	}

	c.JSON(http.StatusOK, book)
}

// UpdateBook godoc
// @Summary Update a book
// @Description Update a book with the input payload
// @Tags books
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Param book body dto.BookRequest true "Book data"
// @Success 200 {object} models.Book
// @Failure 404 {object} dto.ErrorResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /api/v1/books/{id} [put]
func UpdateBook(c *gin.Context) {
	id := c.Param("id")
	var bookRequest dto.BookRequest
	var book models.Book

	if err := db.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, dto.ErrorResponse{Error: "book not found"})
		return
	}

	if err := c.ShouldBindJSON(&bookRequest); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
		return
	}

	// Check if author exists
	if err := db.First(&models.Author{}, bookRequest.AuthorID).Error; err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "author not found"})
		return
	}

	book.Title = bookRequest.Title
	book.ISBN = bookRequest.ISBN
	book.PublicationYear = bookRequest.PublicationYear
	book.Description = bookRequest.Description
	book.AuthorID = bookRequest.AuthorID

	if err := db.Save(&book).Error; err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "internal server error"})
		return
	}

	// Load relations for response
	db.Preload("Author").First(&book, book.ID)

	c.JSON(http.StatusOK, book)
}

// DeleteBook godoc
// @Summary Delete a book
// @Description Delete a book by ID
// @Tags books
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Success 200 {object} dto.Response
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /api/v1/books/{id} [delete]
func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	var book models.Book

	if err := db.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, dto.ErrorResponse{Error: "book not found"})
		return
	}

	if err := db.Delete(&book).Error; err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.Response{Msg: "book deleted successfully"})
} 