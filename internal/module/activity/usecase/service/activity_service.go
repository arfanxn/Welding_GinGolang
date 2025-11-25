package service

import (
	"context"
	"fmt"

	"github.com/arfanxn/welding/internal/infrastructure/id"
	activityEnum "github.com/arfanxn/welding/internal/module/activity/domain/enum"
	activityRepository "github.com/arfanxn/welding/internal/module/activity/domain/repository"
	"github.com/arfanxn/welding/internal/module/activity/usecase/dto"
	"github.com/arfanxn/welding/internal/module/shared/contextkey"
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	"github.com/arfanxn/welding/internal/module/shared/domain/errorx"
	"github.com/arfanxn/welding/pkg/reflectutil"
	"github.com/arfanxn/welding/pkg/types"
	"github.com/arfanxn/welding/pkg/typeutil"
	"github.com/iancoleman/strcase"
	"github.com/samber/lo"
	"go.uber.org/fx"
)

type ActivityService interface {
	Create(ctx context.Context, _dto *dto.CreateActivity) (*entity.Activity, error)
	Record(ctx context.Context, _dto *dto.CreateActivity) (*entity.Activity, error)
}

type activityService struct {
	idService          id.IdService
	activityRepository activityRepository.ActivityRepository
}

type NewActivityServiceParams struct {
	fx.In

	IdService          id.IdService
	ActivityRepository activityRepository.ActivityRepository
}

func NewActivityService(params NewActivityServiceParams) ActivityService {
	return &activityService{
		idService:          params.IdService,
		activityRepository: params.ActivityRepository,
	}
}

func (s *activityService) Create(
	ctx context.Context, _dto *dto.CreateActivity,
) (*entity.Activity, error) {
	var (
		id              string                      = s.idService.Generate()
		action          activityEnum.ActivityAction = _dto.Action
		causerIdPtr     *string
		causerTypePtr   *activityEnum.CauserType
		causerIpAddrPtr *string
		subjectIdPtr    *string
		subjectTypePtr  *activityEnum.SubjectType
		properties      types.JSONMap = _dto.Properties
	)

	if _dto.CauserId != nil || _dto.CauserType != nil {
		causerIdPtr = _dto.CauserId
		causerTypePtr = _dto.CauserType
	} else if _dto.Causer != nil {
		fmt.Println("it went here 1")
		if cId, ok := reflectutil.GetStructField(_dto.Causer, "Id"); ok {
			causerId := cId.(string)
			fmt.Println("it went here 2")
			causerIdPtr = &causerId
			fmt.Println("it went here 3")
			cType := reflectutil.GetStructName(_dto.Causer)
			cType = strcase.ToSnake(cType)
			acType := activityEnum.CauserType(cType)
			if !lo.Contains(activityEnum.CauserTypes, activityEnum.CauserType(acType)) {
				return nil, errorx.ErrActivityInvalidCauserType
			}
			causerTypePtr = &acType
		}
	} else if userId, ok := ctx.Value(contextkey.UserIdKey).(string); ok {
		causerIdPtr = &userId
		causerTypePtr = typeutil.Ptr(activityEnum.UserCauserType)
	} else {
		causerIdPtr = nil
		causerTypePtr = nil
	}

	if _dto.CauserIpAddress != nil {
		causerIpAddrPtr = _dto.CauserIpAddress
	} else if userIpAddr, ok := ctx.Value(contextkey.ClientIpKey).(string); ok {
		causerIpAddrPtr = typeutil.Ptr(userIpAddr)
	} else {
		causerIpAddrPtr = nil
	}

	if _dto.SubjectId != nil || _dto.SubjectType != nil {
		subjectIdPtr = _dto.SubjectId
		subjectTypePtr = _dto.SubjectType
	} else if _dto.Subject != nil {
		if sId, ok := reflectutil.GetStructField(_dto.Subject, "Id"); ok {
			subjectId := sId.(string)
			subjectIdPtr = &subjectId
			sType := reflectutil.GetStructName(_dto.Subject)
			sType = strcase.ToSnake(sType)
			asType := activityEnum.SubjectType(sType)
			if !lo.Contains(activityEnum.SubjectTypes, activityEnum.SubjectType(sType)) {
				return nil, errorx.ErrActivityInvalidSubjectType
			}
			subjectTypePtr = &asType
		}
	} else {
		subjectIdPtr = nil
		subjectTypePtr = nil
	}

	activity := entity.NewActivity()
	activity.Id = id
	activity.CauserId = causerIdPtr
	activity.CauserType = causerTypePtr
	activity.CauserIpAddress = causerIpAddrPtr
	activity.Action = action
	activity.SubjectId = subjectIdPtr
	activity.SubjectType = subjectTypePtr
	activity.Properties = properties
	return activity, nil
}

func (s *activityService) Record(ctx context.Context, _dto *dto.CreateActivity) (*entity.Activity, error) {
	activity, err := s.Create(ctx, _dto)
	if err != nil {
		return nil, err
	}

	if err := s.activityRepository.Save(activity); err != nil {
		return nil, err
	}

	return activity, nil
}
