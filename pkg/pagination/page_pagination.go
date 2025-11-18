package pagination

import (
	"net/url"
	"strconv"

	"github.com/arfanxn/welding/pkg/boolutil"
	"github.com/guregu/null/v6"
)

type PagePagination[T any] struct {
	CurrentPage  int         `json:"current_page"`
	PerPage      int         `json:"per_page"`
	HasPrevPage  bool        `json:"has_prev_page"`
	HasNextPage  bool        `json:"has_next_page"`
	PrevPage     null.Int    `json:"prev_page"`
	NextPage     null.Int    `json:"next_page"`
	LastPage     int         `json:"last_page"`
	FirstPageUrl string      `json:"first_page_url"`
	PrevPageUrl  null.String `json:"prev_page_url"`
	NextPageUrl  null.String `json:"next_page_url"`
	LastPageUrl  string      `json:"last_page_url"`
	Items        []T         `json:"items"`
	TotalItems   int         `json:"total_items"`
}

func buildPagePaginationUrl(u url.URL, page int) *url.URL {
	q := u.Query()
	q.Set("page", strconv.Itoa(page))
	u.RawQuery = q.Encode()
	return &u
}

func NewPagePagination[T any](
	totalItems int,
	items []T,
	currentPage int,
	perPage int,
	u url.URL,
) *PagePagination[T] {
	lastPage := max((int(totalItems)+perPage-1)/perPage, 1)

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

	firstPageUrl := buildPagePaginationUrl(u, 1).String()
	prevPageUrl := boolutil.Ternary(hasPrevPage, null.StringFrom(buildPagePaginationUrl(u, currentPage-1).String()), null.String{})
	nextPageUrl := boolutil.Ternary(hasNextPage, null.StringFrom(buildPagePaginationUrl(u, currentPage+1).String()), null.String{})
	lastPageUrl := buildPagePaginationUrl(u, lastPage).String()

	return &PagePagination[T]{
		CurrentPage:  currentPage,
		PerPage:      perPage,
		HasPrevPage:  hasPrevPage,
		HasNextPage:  hasNextPage,
		PrevPage:     prevPage,
		NextPage:     nextPage,
		LastPage:     lastPage,
		FirstPageUrl: firstPageUrl,
		PrevPageUrl:  prevPageUrl,
		NextPageUrl:  nextPageUrl,
		LastPageUrl:  lastPageUrl,
		Items:        items,
		TotalItems:   totalItems,
	}
}

// Convert OffsetPagination to PagePagination
func PPFromOP[T any](op *OffsetPagination[T], u url.URL) *PagePagination[T] {
	offset := op.Offset
	hasOffset := offset >= 0
	perPage := op.Limit
	hasPerPage := perPage >= 0
	currentPage := 1
	if hasOffset && hasPerPage && perPage > 0 {
		currentPage = (offset / perPage) + 1
	}

	return NewPagePagination(op.TotalItems, op.Items, currentPage, perPage, u)
}
