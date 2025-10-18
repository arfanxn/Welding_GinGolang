package helper

import (
	"errors"

	"github.com/arfanxn/welding/internal/module/shared/usecase/dto"
	"github.com/guregu/null/v6"
	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
)

// GormDBPaginateWithQueryDTO handles pagination for GORM queries using a Query DTO
// It returns a Pagination struct containing the paginated results and metadata
func GormDBPaginateWithQueryDTO[T any](db *gorm.DB, queryDto *dto.Query) (*dto.Pagination[T], error) {
	// First, get the total count of matching records (before applying pagination)
	var totalItems int64
	if err := db.Count(&totalItems).Error; err != nil {
		return nil, err
	}

	// Get pagination parameters from the query DTO
	offset, hasOffset := queryDto.GetOffset()
	limit, hasLimit := queryDto.GetLimit()

	// Apply offset if specified
	if hasOffset {
		db = db.Offset(offset)
	}

	// Apply limit if specified
	if hasLimit {
		db = db.Limit(limit)
	}

	// Execute the query to get paginated results
	var items []T
	if err := db.Find(&items).Error; err != nil {
		return nil, err
	}

	// Calculate current page number (1-based)
	currentPage := 1
	if hasOffset && hasLimit && limit > 0 {
		currentPage = (offset / limit) + 1
	}

	// Calculate total number of pages (ceiling of totalItems/limit)
	lastPage := max((int(totalItems)+limit-1)/limit, 1)

	// Calculate previous and next page numbers
	var prevPage, nextPage null.Int

	// Only set previous page if not on first page
	if currentPage > 1 {
		prevPage = null.IntFrom(int64(currentPage - 1))
	}

	// Only set next page if not on last page
	if currentPage < lastPage {
		nextPage = null.IntFrom(int64(currentPage + 1))
	}

	// Determine if previous/next pages exist
	hasPrevPage := currentPage > 1
	hasNextPage := currentPage < lastPage

	// Return pagination result with metadata and items
	return &dto.Pagination[T]{
		CurrentPage: currentPage,     // Current page number (1-based)
		PerPage:     limit,           // Number of items per page
		HasPrevPage: hasPrevPage,     // Whether there's a previous page
		HasNextPage: hasNextPage,     // Whether there's a next page
		PrevPage:    prevPage,        // Previous page number (nullable)
		NextPage:    nextPage,        // Next page number (nullable)
		LastPage:    lastPage,        // Last page number
		Items:       items,           // Paginated items
		TotalItems:  int(totalItems), // Total number of items across all pages
	}, nil
}

// IsPostgresDuplicateKeyError checks if the error is a PostgreSQL duplicate key error
func IsPostgresDuplicateKeyError(err error) bool {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) && pgErr.Code == "23505" {
		return true
	}
	return false
}
