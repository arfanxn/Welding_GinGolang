package dto

import "github.com/guregu/null/v6"

type Pagination[T any] struct {
	CurrentPage int      `json:"current_page"`
	PerPage     int      `json:"per_page"`
	HasPrevPage bool     `json:"has_prev_page"`
	HasNextPage bool     `json:"has_next_page"`
	PrevPage    null.Int `json:"prev_page"`
	NextPage    null.Int `json:"next_page"`
	LastPage    int      `json:"last_page"`
	Items       []T      `json:"items"`
	TotalItems  int      `json:"total_items"`
}
