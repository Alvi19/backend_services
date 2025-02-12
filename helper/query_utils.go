package helper

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type PaginatedResponse struct {
	Total    int         `json:"total"`
	Page     int         `json:"page"`
	PageSize int         `json:"page_size"`
	Data     interface{} `json:"data"`
}

type FiltersResponse struct {
	Total    int
	Page     int
	PageSize int
	Query    *gorm.DB
}

func GetFilteredDataWithPagination(
	recorddata interface{},
	db *gorm.DB,
	c *fiber.Ctx,
	filters []string,
) (*FiltersResponse, error) {
	var total int64

	// Ambil nilai dari query string
	page := c.QueryInt("page", 0)             // Default 0 jika tidak ada
	pageSize := c.QueryInt("page_size", 0)    // Default 0 jika tidak ada
	orderBy := c.Query("order_by", "id DESC") // Default sorting

	// Query awal
	query := db.Model(recorddata)

	for _, field := range filters {
		value := c.Query(field)
		if value != "" {
			query = query.Where(field+" = ?", value)
		}
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, err
	}

	if page > 0 && pageSize > 0 {
		query = query.Offset((page - 1) * pageSize).Limit(pageSize)
	}
	if orderBy != "" {
		query = query.Order(orderBy)
	}

	response := FiltersResponse{
		Total:    int(total),
		Page:     page,
		PageSize: pageSize,
		Query:    query,
	}

	return &response, nil
}
