package dto

import "strings"

const (
	QueryFilterOperatorEqual              = "=="
	QueryFilterOperatorNotEqual           = "!="
	QueryFilterOperatorGreaterThan        = ">"
	QueryFilterOperatorLessThan           = "<"
	QueryFilterOperatorGreaterThanOrEqual = ">="
	QueryFilterOperatorLessThanOrEqual    = "<="
	QueryFilterOperatorLike               = "LIKE"
	QueryFilterOperatorILike              = "ILIKE"
	QueryFilterOperatorIn                 = "IN"
	QueryFilterOperatorNotIn              = "NOT IN"
	QueryFilterOperatorBetween            = "BETWEEN"
	QueryFilterOperatorNotBetween         = "NOT BETWEEN"
)

var QueryFilterOperators = []string{
	QueryFilterOperatorEqual,
	QueryFilterOperatorNotEqual,
	QueryFilterOperatorGreaterThan,
	QueryFilterOperatorLessThan,
	QueryFilterOperatorGreaterThanOrEqual,
	QueryFilterOperatorLessThanOrEqual,
	QueryFilterOperatorLike,
	QueryFilterOperatorILike,
	QueryFilterOperatorIn,
	QueryFilterOperatorNotIn,
	QueryFilterOperatorBetween,
	QueryFilterOperatorNotBetween,
}

type QueryFilter struct {
	Column   string `json:"column"`
	Operator string `json:"operator"`
	Value    string `json:"value"`
}

const (
	SortOrderAsc  string = "ASC"
	SortOrderDesc string = "DESC"
)

type QuerySort struct {
	Column string `json:"column"`
	Order  string `json:"order"`
}

type Query struct {
	Offset  int           `json:"offset"`
	Limit   int           `json:"limit"`
	Search  string        `json:"search"`
	Filters []QueryFilter `json:"filters"`
	Sorts   []QuerySort   `json:"sorts"`
	Joins   []string      `json:"joins"`
}

// GetOffset returns the offset value and a boolean indicating if it's set (greater than 0)
func (q *Query) GetOffset() (int, bool) {
	has := q.Offset > 0
	return q.Offset, has
}

// GetLimit returns the limit value and a boolean indicating if it's set (greater than 0)
func (q *Query) GetLimit() (int, bool) {
	has := q.Limit > 0
	return q.Limit, has
}

// GetSearch returns the search term and a boolean indicating if it's not empty
func (q *Query) GetSearch() (string, bool) {
	has := q.Search != ""
	return q.Search, has
}

// GetFilterByCO searches for a filter by column name and operator (case-insensitive)
// and returns the matching filter along with a boolean indicating if found.
// The search is performed using case-insensitive comparison for both column name and operator.
func (q *Query) GetFilterByCO(column string, operator string) (*QueryFilter, bool) {
	for _, filter := range q.Filters {
		if strings.EqualFold(filter.Column, column) && strings.EqualFold(filter.Operator, operator) {
			return &filter, true
		}
	}
	return nil, false
}

// GetSortByColumn searches for a sort by column name (case-insensitive) and returns
// the sort and a boolean indicating if found
func (q *Query) GetSortByColumn(column string) (*QuerySort, bool) {
	for _, sort := range q.Sorts {
		if strings.EqualFold(sort.Column, column) {
			return &sort, true
		}
	}
	return nil, false
}

// GetJoin searches for a join by name (case-insensitive) and returns
// the join string and a boolean indicating if found
func (q *Query) GetJoin(join string) (string, bool) {
	for _, j := range q.Joins {
		if strings.EqualFold(j, join) {
			return j, true
		}
	}
	return "", false
}
