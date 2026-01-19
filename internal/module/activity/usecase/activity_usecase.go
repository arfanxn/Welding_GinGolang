package usecase

import (
	"context"

	activityRepository "github.com/arfanxn/welding/internal/module/activity/domain/repository"
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	"github.com/arfanxn/welding/pkg/pagination"
	"github.com/arfanxn/welding/pkg/query"
	"go.uber.org/fx"
)

type ActivityUsecase interface {
	Paginate(ctx context.Context, q *query.Query) (*pagination.OffsetPagination[*entity.Activity], error)
	Show(ctx context.Context, q *query.Query) (*entity.Activity, error)
}

type activityUsecase struct {
	activityRepository activityRepository.ActivityRepository
}

type NewActivityUsecaseParams struct {
	fx.In
	ActivityRepository activityRepository.ActivityRepository
}

func NewActivityUsecase(params NewActivityUsecaseParams) ActivityUsecase {
	return &activityUsecase{
		activityRepository: params.ActivityRepository,
	}
}

func (u *activityUsecase) Paginate(ctx context.Context, q *query.Query) (op *pagination.OffsetPagination[*entity.Activity], err error) {
	op, err = u.activityRepository.Paginate(q)
	if err != nil {
		return nil, err
	}

	return
}

func (u *activityUsecase) Show(ctx context.Context, q *query.Query) (activity *entity.Activity, err error) {
	activity, err = u.activityRepository.First(q)
	if err != nil {
		return nil, err
	}
	return
}
