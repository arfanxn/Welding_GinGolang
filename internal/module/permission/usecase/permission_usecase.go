package usecase

import (
	"context"

	activityEnum "github.com/arfanxn/welding/internal/module/activity/domain/enum"
	activityDto "github.com/arfanxn/welding/internal/module/activity/usecase/dto"
	activityService "github.com/arfanxn/welding/internal/module/activity/usecase/service"
	"github.com/arfanxn/welding/internal/module/permission/domain/repository"
	"github.com/arfanxn/welding/internal/module/shared/contextkey"
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	"github.com/arfanxn/welding/pkg/pagination"
	"github.com/arfanxn/welding/pkg/query"
	"github.com/arfanxn/welding/pkg/typeutil"
)

type PermissionUsecase interface {
	Paginate(ctx context.Context, q *query.Query) (*pagination.OffsetPagination[*entity.Permission], error)
}

type permissionUsecase struct {
	activityService      activityService.ActivityService
	permissionRepository repository.PermissionRepository
}

func NewPermissionUsecase(
	activityService activityService.ActivityService,
	permissionRepository repository.PermissionRepository) PermissionUsecase {
	return &permissionUsecase{
		activityService:      activityService,
		permissionRepository: permissionRepository,
	}
}

func (u *permissionUsecase) Paginate(ctx context.Context, q *query.Query) (*pagination.OffsetPagination[*entity.Permission], error) {
	op, err := u.permissionRepository.Paginate(q)
	if err != nil {
		return nil, err
	}

	if _, err = u.activityService.Record(ctx, &activityDto.CreateActivity{
		Causer:      contextkey.GetUser(ctx),
		Action:      activityEnum.CodesCreateUserRegisterInvitation,
		SubjectType: typeutil.Ptr(activityEnum.PermissionSubjectType),
	}); err != nil {
		return nil, err
	}

	return op, nil
}
