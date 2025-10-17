package request

import (
	"net/http"
	"regexp"
	"strings"

	"github.com/arfanxn/welding/internal/module/shared/usecase/dto"
	"github.com/arfanxn/welding/pkg/errors"
	"github.com/creasty/defaults"
)

var _ Request = (*Query)(nil)

type Query struct {
	Page    int `form:"page" json:"page" default:"1"`
	PerPage int `form:"per_page" json:"per_page" default:"10"`

	Search string `form:"search" json:"search"`

	Joins   []string `form:"join" json:"join" default:"[]"`
	Filters []string `form:"filter" json:"filter" default:"[]"`
	Sorts   []string `form:"sort" json:"sort" default:"[]"`
}

// NewQuery creates and initializes a new Query instance with default values.
// The defaults are set using struct tags and the function will panic if
// setting defaults fails. This is typically used during request handling
// where default values for pagination and other query parameters are required.
func NewQuery() *Query {
	q := &Query{}
	defaults.MustSet(q)
	return q
}

// TODO: Implement request.Query validation
func (r *Query) Validate() error {
	return nil
}

func (q *Query) ToQueryDTO() (*dto.Query, error) {
	queryDto := &dto.Query{}

	queryDto.Offset = (q.Page - 1) * q.PerPage
	queryDto.Limit = q.PerPage
	queryDto.Search = q.Search
	queryDto.Joins = q.Joins

	for _, f := range q.Filters {
		column, operator, value, err := splitQueryFilter(f)
		if err != nil {
			return nil, err
		}
		queryDto.Filters = append(queryDto.Filters, dto.QueryFilter{
			Column:   column,
			Operator: operator,
			Value:    value,
		})
	}

	for _, s := range q.Sorts {
		column, order, err := splitQuerySort(s)
		if err != nil {
			return nil, err
		}
		queryDto.Sorts = append(queryDto.Sorts, dto.QuerySort{
			Column: column,
			Order:  order,
		})
	}

	return queryDto, nil
}

func (q *Query) MustToQueryDTO() *dto.Query {
	queryDto, err := q.ToQueryDTO()
	if err != nil {
		panic(err)
	}
	return queryDto
}

var filterRegex = func() *regexp.Regexp {
	queryFilterOperators := strings.Join(dto.QueryFilterOperators, "|")

	return regexp.MustCompile(
		// Capture group 1: The column name (any sequence of non-whitespace characters not containing the operator)
		`^(\S+?)` +
			// Capture group 2: The operator
			`\s*(` + queryFilterOperators + `)\s*` +
			// Capture group 3: The value (everything else, non-greedy)
			`(.+)$`,
	)
}()

func splitQueryFilter(filter string) (column, operator, value string, err error) {
	// Find all submatches
	matches := filterRegex.FindStringSubmatch(filter)

	// Check if a match was found and has the expected number of submatches (4: full match + 3 groups)
	if len(matches) < 4 {
		return "", "", "",
			errors.NewHttpError(http.StatusBadRequest, "invalid filter format", nil)
	}

	// Assign the captured groups
	column = strings.TrimSpace(matches[1])
	operator = strings.TrimSpace(matches[2])
	value = strings.TrimSpace(matches[3])

	return column, operator, value, nil
}

var sortRegex = regexp.MustCompile(`^\s*(-?)([^\s]+)\s*$`)

func splitQuerySort(sort string) (column, order string, err error) {
	matches := sortRegex.FindStringSubmatch(sort)
	if len(matches) != 3 {
		return "", "", errors.NewHttpError(http.StatusBadRequest, "invalid sort format", nil)
	}

	column = strings.TrimSpace(matches[2])
	if column == "" {
		return "", "", errors.NewHttpError(http.StatusBadRequest, "sort column cannot be empty", nil)
	}

	order = "ASC"
	if matches[1] == "-" {
		order = "DESC"
	}

	return column, order, nil
}
