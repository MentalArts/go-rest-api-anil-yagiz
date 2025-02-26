package handlers

import (
	"mentalartsapi/dto"
	"mentalartsapi/models"
	"mentalartsapi/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB(database *gorm.DB) {
	db = database
}

// CreateAuthor godoc
// @Summary Create a new author
// @Description Create a new author with the input payload
// @Tags authors
// @Accept json
// @Produce json
// @Param author body dto.AuthorRequest true "Author data"
// @Success 201 {object} models.Author
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /api/v1/authors [post]
func CreateAuthor(c *gin.Context) {
	var authorRequest dto.AuthorRequest
	var author models.Author

	if err := c.ShouldBindJSON(&authorRequest); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
		return
	}

	author.Name = authorRequest.Name
	author.Biography = authorRequest.Biography
	author.BirthDate = authorRequest.BirthDate

	result := db.Create(&author)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, author)
}

// GetAllAuthors godoc
// @Summary Get all authors
// @Description Get all authors with pagination
// @Tags authors
// @Accept json
// @Produce json
// @Param page query int false "Page number"
// @Param page_size query int false "Page size"
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} dto.ErrorResponse
// @Router /api/v1/authors [get]
func GetAllAuthors(c *gin.Context) {
	var authors []models.Author
	var totalCount int64

	pagination := utils.ParsePaginationQuery(c)
	
	// Count total records
	if err := db.Model(&models.Author{}).Count(&totalCount).Error; err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	// Get paginated authors with their books
	if err := utils.Paginate(db, &pagination).
		Preload("Books").
		Find(&authors).Error; err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	// Create pagination response
	paginationResponse := utils.CreatePaginationResponse(totalCount, pagination)

	// Create final response
	response := gin.H{
		"data":       authors,
		"pagination": paginationResponse,
	}

	c.JSON(http.StatusOK, response)
}

// GetAuthor godoc
// @Summary Get a single author
// @Description Get a single author by ID with their books
// @Tags authors
// @Accept json
// @Produce json
// @Param id path int true "Author ID"
// @Success 200 {object} models.Author
// @Failure 404 {object} dto.ErrorResponse
// @Router /api/v1/authors/{id} [get]
func GetAuthor(c *gin.Context) {
	id := c.Param("id")
	var author models.Author

	if err := db.Preload("Books").First(&author, id).Error; err != nil {
		c.JSON(http.StatusNotFound, dto.ErrorResponse{Error: "author not found"})
		return
	}

	c.JSON(http.StatusOK, author)
}

// UpdateAuthor godoc
// @Summary Update an author
// @Description Update an author with the input payload
// @Tags authors
// @Accept json
// @Produce json
// @Param id path int true "Author ID"
// @Param author body dto.AuthorRequest true "Author data"
// @Success 200 {object} models.Author
// @Failure 404 {object} dto.ErrorResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /api/v1/authors/{id} [put]
func UpdateAuthor(c *gin.Context) {
	id := c.Param("id")
	var authorRequest dto.AuthorRequest
	var author models.Author

	if err := db.First(&author, id).Error; err != nil {
		c.JSON(http.StatusNotFound, dto.ErrorResponse{Error: "author not found"})
		return
	}

	if err := c.ShouldBindJSON(&authorRequest); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
		return
	}

	author.Name = authorRequest.Name
	author.Biography = authorRequest.Biography
	author.BirthDate = authorRequest.BirthDate

	if err := db.Save(&author).Error; err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "internal server error"})
		return
	}

	c.JSON(http.StatusOK, author)
}

// DeleteAuthor godoc
// @Summary Delete an author
// @Description Delete an author by ID
// @Tags authors
// @Accept json
// @Produce json
// @Param id path int true "Author ID"
// @Success 200 {object} dto.Response
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /api/v1/authors/{id} [delete]
func DeleteAuthor(c *gin.Context) {
	id := c.Param("id")
	var author models.Author

	if err := db.First(&author, id).Error; err != nil {
		c.JSON(http.StatusNotFound, dto.ErrorResponse{Error: err.Error()})
		return
	}

	if err := db.Delete(&author).Error; err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.Response{Msg: "author deleted successfully"})
}
