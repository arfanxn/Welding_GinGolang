package repository

import (
	analyticEntity "github.com/arfanxn/welding/internal/module/analytic/domain/entity"
	"github.com/arfanxn/welding/pkg/query"
)

type AnalyticRepository interface {
	GetSummary(q *query.Query) (*analyticEntity.Summary, error)
	GetOrderTrends(q *query.Query) (analyticEntity.OrderTrends, error)
}
