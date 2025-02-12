package util

import (
	"github.com/gofiber/fiber/v2"
)

// Pagination struct untuk menangani pagination
type Pagination struct {
	Page   int
	Limit  int
	Offset int
}

// NewPagination untuk membuat instance Pagination dengan default atau query params// NewPagination untuk membuat instance Pagination dari body request
func NewPagination(c *fiber.Ctx, defaultPage int, defaultLimit int) *Pagination {
	// Deklarasikan variabel untuk menampung data dari body
	var paginationRequest struct {
		Page  int `json:"page"`
		Limit int `json:"limit"`
	}

	// Ambil data pagination dari body
	if err := c.BodyParser(&paginationRequest); err != nil {
		// Jika tidak ada data di body atau error, gunakan default
		return &Pagination{
			Page:   defaultPage,
			Limit:  defaultLimit,
			Offset: (defaultPage - 1) * defaultLimit,
		}
	}

	// Jika tidak ada page atau limit di body, gunakan default
	if paginationRequest.Page < 1 {
		paginationRequest.Page = defaultPage
	}
	if paginationRequest.Limit < 1 {
		paginationRequest.Limit = defaultLimit
	}

	offset := (paginationRequest.Page - 1) * paginationRequest.Limit

	return &Pagination{
		Page:   paginationRequest.Page,
		Limit:  paginationRequest.Limit,
		Offset: offset,
	}
}

// CalculateTotalPages untuk menghitung total halaman
func CalculateTotalPages(totalRecords int64, limit int) int {
	totalPages := int(totalRecords / int64(limit))
	if totalRecords%int64(limit) > 0 {
		totalPages++
	}
	return totalPages
}
