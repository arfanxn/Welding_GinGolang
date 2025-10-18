package response

import (
	"net/url"
	"strconv"

	"github.com/arfanxn/welding/internal/module/shared/usecase/dto"
	"github.com/arfanxn/welding/pkg/boolutil"
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
	// Get the current request URL
	scheme := boolutil.Ternary(
		c.Request.TLS != nil || c.GetHeader("X-Forwarded-Proto") == "https",
		"https",
		"http")
	host := c.Request.Host
	path := c.Request.URL.Path

	// Helper function to build full URL with pagination
	buildPageUrl := func(page int) string {
		u := url.URL{
			Scheme:   scheme,
			Host:     host,
			Path:     path,
			RawQuery: "",
		}
		q := u.Query()
		q.Set("page", strconv.Itoa(page))
		u.RawQuery = q.Encode()
		return u.String()
	}

	// Build pagination URLs
	firstPageUrl := buildPageUrl(1)
	var prevPageUrl null.String
	var nextPageUrl null.String
	var lastPageUrl string

	if paginationDto.PrevPage.Valid {
		prevPageUrl = null.StringFrom(buildPageUrl(int(paginationDto.PrevPage.Int64)))
	}

	if paginationDto.NextPage.Valid {
		nextPageUrl = null.StringFrom(buildPageUrl(int(paginationDto.NextPage.Int64)))
	}

	if paginationDto.LastPage > 0 {
		lastPageUrl = buildPageUrl(paginationDto.LastPage)
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
