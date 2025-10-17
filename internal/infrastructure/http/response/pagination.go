package response

import (
	"net/url"
	"strconv"

	"github.com/arfanxn/welding/internal/module/shared/usecase/dto"
	"github.com/gin-gonic/gin"
	"github.com/guregu/null/v6"
)

type Pagination[T any] struct {
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

func NewPaginationFromContextPaginationDTO[T any](
	c *gin.Context, paginationDto *dto.Pagination[T],
) *Pagination[T] {
	var (
		url          *url.URL = c.Request.URL
		firstPageUrl string
		prevPageUrl  null.String
		nextPageUrl  null.String
		lastPageUrl  string
	)

	q := url.Query()
	q.Set("page", "1")
	url.RawQuery = q.Encode()
	firstPageUrl = url.String()

	if paginationDto.PrevPage.Valid {
		q := url.Query()
		q.Set("page", strconv.FormatInt(paginationDto.PrevPage.Int64, 10))
		url.RawQuery = q.Encode()
		prevPageUrl = null.StringFrom(url.String())
	}

	if paginationDto.NextPage.Valid {
		q := url.Query()
		q.Set("page", strconv.FormatInt(paginationDto.NextPage.Int64, 10))
		url.RawQuery = q.Encode()
		nextPageUrl = null.StringFrom(url.String())
	}

	if paginationDto.LastPage > 0 {
		q := url.Query()
		q.Set("page", strconv.Itoa(paginationDto.LastPage))
		url.RawQuery = q.Encode()
		lastPageUrl = url.String()
	}

	return &Pagination[T]{
		CurrentPage:  paginationDto.CurrentPage,
		PerPage:      paginationDto.PerPage,
		HasPrevPage:  paginationDto.HasPrevPage,
		HasNextPage:  paginationDto.HasNextPage,
		PrevPage:     paginationDto.PrevPage,
		NextPage:     paginationDto.NextPage,
		LastPage:     paginationDto.LastPage,
		FirstPageUrl: firstPageUrl,
		PrevPageUrl:  prevPageUrl,
		NextPageUrl:  nextPageUrl,
		LastPageUrl:  lastPageUrl,
		Items:        paginationDto.Items,
		TotalItems:   paginationDto.TotalItems,
	}
}
