package query

import (
	"regexp"
	"strings"

	"github.com/creasty/defaults"
	"github.com/guregu/null/v6"
)

const (
	OperatorEqual              = "=="
	OperatorNotEqual           = "!="
	OperatorGreaterThan        = ">"
	OperatorLessThan           = "<"
	OperatorGreaterThanOrEqual = ">="
	OperatorLessThanOrEqual    = "<="
	OperatorLike               = "LIKE"
	OperatorILike              = "ILIKE"
	OperatorIn                 = "IN"
	OperatorNotIn              = "NOT IN"
	OperatorBetween            = "BETWEEN"
	OperatorNotBetween         = "NOT BETWEEN"
)

var Operators = []string{
	OperatorEqual,
	OperatorNotEqual,
	OperatorGreaterThan,
	OperatorLessThan,
	OperatorGreaterThanOrEqual,
	OperatorLessThanOrEqual,
	OperatorLike,
	OperatorILike,
	OperatorIn,
	OperatorNotIn,
	OperatorBetween,
	OperatorNotBetween,
}

type Filter struct {
	Column   string `json:"column"`
	Operator string `json:"operator"`
	Value    string `json:"value"`
}

const (
	OrderAsc  string = "ASC"
	OrderDesc string = "DESC"
)

type Sort struct {
	Column string `json:"column"`
	Order  string `json:"order"`
}

// Query represents the request parameters for querying resources with pagination,
// filtering, sorting, and related data loading capabilities.
type Query struct {
	// Page specifies the page number for pagination (1-based index).
	// Example: ?page=2
	Page int `form:"page" json:"page" default:"1"`

	// PerPage specifies the number of items per page.
	// Example: ?per_page=20
	PerPage int `form:"per_page" json:"per_page" default:"10"`

	// Search is a free-text search query.
	// Implementation depends on the specific endpoint.
	// Example: ?search=admin
	Search null.String `form:"search" json:"search"`

	// Includes specifies related models to be loaded (eager loading).
	// Format: "relation" or "relation1,relation2"
	// Example: ?include=permissions&include=users
	Includes []string `form:"include" json:"include" default:"[]"`

	// Filters specifies conditions to filter the results.
	// Format: "field operator value"
	// Supported operators: ==, !=, >, >=, <, <=, IN, NOT IN, etc.
	// Examples:
	// - ?filter=name==admin
	// - ?filter=created_at>2023-01-01
	// - ?filter=statusINactive,pending
	Filters []string `form:"filter" json:"filter" default:"[]"`

	// Sorts specifies the order of results.
	// Prefix field with - for descending order.
	// Examples:
	// - ?sort=name (ascending)
	// - ?sort=-created_at (descending)
	// - ?sort=name&sort=-created_at (multiple sorts)
	Sorts []string `form:"sort" json:"sort" default:"[]"`
}

func NewQuery() *Query {
	q := &Query{}
	defaults.MustSet(q)
	return q
}

func (q *Query) FilterById(id string) *Query {
	q.Filters = append(q.Filters, "id=="+id)
	return q
}

func (q *Query) GetFilterById() *Filter {
	return q.GetFilter("id", OperatorEqual)
}

func (q *Query) GetPage() int {
	return q.Page
}

func (q *Query) GetPerPage() int {
	return q.PerPage
}

func (q *Query) GetLimit() int {
	return q.PerPage
}

func (q *Query) GetOffset() int {
	return (q.Page - 1) * q.PerPage
}

func (q *Query) GetSearch() null.String {
	return q.Search
}

func (q *Query) GetInclude(include string) null.String {
	// Create a case-insensitive regex pattern that matches the exact include string
	// with optional surrounding whitespace
	pattern := `(?i)^\s*` + regexp.QuoteMeta(include) + `\s*$`
	re := regexp.MustCompile(pattern)

	for _, i := range q.Includes {
		if re.MatchString(i) {
			return null.StringFrom(i)
		}
	}
	return null.String{}
}

func (q *Query) GetFilter(column string, operator string) *Filter {
	// Escape special regex characters in the operator
	escapedOp := regexp.QuoteMeta(operator)

	// Create a regex pattern that matches:
	// ^\s*<column>\s*<escapedOp>\s*(.*)$
	// Where \s* matches any whitespace
	pattern := `^\s*` + regexp.QuoteMeta(column) + `\s*` + escapedOp + `\s*(.*)$`
	re := regexp.MustCompile(pattern)

	for _, filter := range q.Filters {
		matches := re.FindStringSubmatch(filter)
		if len(matches) > 1 {
			return &Filter{
				Column:   column,
				Operator: operator,
				Value:    strings.TrimSpace(matches[1]),
			}
		}
	}
	return nil
}

func (q *Query) GetSort(column string) *Sort {
	// Create a regex pattern that matches:
	// ^-?<column>$
	// Where -? means optional - for descending sort
	pattern := `^\s*(-?)` + regexp.QuoteMeta(column) + `\s*$`
	re := regexp.MustCompile(pattern)

	for _, sort := range q.Sorts {
		matches := re.FindStringSubmatch(sort)
		if len(matches) > 0 {
			order := OrderAsc
			if matches[1] == "-" {
				order = OrderDesc
			}
			return &Sort{
				Column: column,
				Order:  order,
			}
		}
	}
	return nil
}
