package helper

import (
	"errors"

	"github.com/arfanxn/welding/pkg/pagination"
	"github.com/arfanxn/welding/pkg/query"
	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
)

// IsPostgresDuplicateKeyError checks if the error is a PostgreSQL duplicate key error
func IsPostgresDuplicateKeyError(err error) bool {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) && pgErr.Code == "23505" {
		return true
	}
	return false
}

func GormDBPaginateWithQuery[T any](db *gorm.DB, q *query.Query) (*pagination.OffsetPagination[T], error) {
	// First, get the total count of matching records (before applying pagination)
	var totalItems int64
	if err := db.Count(&totalItems).Error; err != nil {
		return nil, err
	}

	// Get pagination parameters from the query DTO
	offset := q.GetOffset()
	hasOffset := offset >= 0
	limit := q.GetLimit()
	hasLimit := limit >= 0

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

	return pagination.NewOffsetPagination(offset, limit, int(totalItems), items), nil
}
