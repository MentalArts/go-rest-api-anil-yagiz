package utils

import (
	"mentalartsapi/dto"
	"math"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

const (
	DefaultPage     = 1
	DefaultPageSize = 10
)

// ParsePaginationQuery extracts pagination parameters from the request
func ParsePaginationQuery(c *gin.Context) dto.PaginationQuery {
	pagination := dto.PaginationQuery{
		Page:     DefaultPage,
		PageSize: DefaultPageSize,
	}

	// Try to bind query parameters
	if page := c.Query("page"); page != "" {
		if pageInt, err := c.GetQuery("page"); err {
			if pageVal, err := c.GetInt("page"); err == nil && pageVal > 0 {
				pagination.Page = pageVal
			}
		}
	}

	if pageSize := c.Query("page_size"); pageSize != "" {
		if pageSizeInt, err := c.GetQuery("page_size"); err {
			if pageSizeVal, err := c.GetInt("page_size"); err == nil && pageSizeVal > 0 {
				if pageSizeVal > 100 {
					pagination.PageSize = 100 // Maximum page size
				} else {
					pagination.PageSize = pageSizeVal
				}
			}
		}
	}

	return pagination
}

// Paginate applies pagination to a GORM query
func Paginate(query *gorm.DB, pagination *dto.PaginationQuery) *gorm.DB {
	offset := (pagination.Page - 1) * pagination.PageSize
	return query.Offset(offset).Limit(pagination.PageSize)
}

// CreatePaginationResponse creates a pagination response
func CreatePaginationResponse(totalRecords int64, pagination dto.PaginationQuery) dto.Pagination {
	totalPages := int(math.Ceil(float64(totalRecords) / float64(pagination.PageSize)))
	
	return dto.Pagination{
		TotalRecords: totalRecords,
		TotalPages:   totalPages,
		Page:         pagination.Page,
		PageSize:     pagination.PageSize,
		HasMore:      pagination.Page < totalPages,
	}
} 