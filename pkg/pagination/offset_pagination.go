package pagination

type OffsetPagination[T any] struct {
	Offset     int `json:"offset"`
	Limit      int `json:"limit"`
	Items      []T `json:"items"`
	TotalItems int `json:"total_items"`
}

func NewOffsetPagination[T any](
	offset int,
	limit int,
	totalItems int,
	items []T,
) *OffsetPagination[T] {
	return &OffsetPagination[T]{
		Offset:     offset,
		Limit:      limit,
		Items:      items,
		TotalItems: totalItems,
	}
}
