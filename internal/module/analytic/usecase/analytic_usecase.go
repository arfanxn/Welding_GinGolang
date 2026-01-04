package usecase

import (
	analyticEntity "github.com/arfanxn/welding/internal/module/analytic/domain/entity"
	analyticRepository "github.com/arfanxn/welding/internal/module/analytic/domain/repository"
	"github.com/arfanxn/welding/pkg/query"
	"go.uber.org/fx"
)

type AnalyticUsecase interface {
	GetSummary(q *query.Query) (summary *analyticEntity.Summary, err error)
	GetOrderTrends(q *query.Query) (orderTrends analyticEntity.OrderTrends, err error)
}

type analyticUsecase struct {
	analyticRepository analyticRepository.AnalyticRepository
}

type NewAnalyticUsecaseParams struct {
	fx.In

	AnalyticRepository analyticRepository.AnalyticRepository
}

func NewAnalyticUsecase(params NewAnalyticUsecaseParams) AnalyticUsecase {
	return &analyticUsecase{
		analyticRepository: params.AnalyticRepository,
	}
}

func (u *analyticUsecase) GetSummary(q *query.Query) (summary *analyticEntity.Summary, err error) {
	summary, err = u.analyticRepository.GetSummary(q)
	return
}

func (u *analyticUsecase) GetOrderTrends(q *query.Query) (orderTrends analyticEntity.OrderTrends, err error) {
	orderTrends, err = u.analyticRepository.GetOrderTrends(q)
	return
}
