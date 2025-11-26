package usecase

import (
	"context"

	activityEnum "github.com/arfanxn/welding/internal/module/activity/domain/enum"
	activityDto "github.com/arfanxn/welding/internal/module/activity/usecase/dto"
	activityService "github.com/arfanxn/welding/internal/module/activity/usecase/service"
	mtsRepository "github.com/arfanxn/welding/internal/module/material_test_service/domain/repository"
	mtsPolicy "github.com/arfanxn/welding/internal/module/material_test_service/infrastructure/policy"
	"github.com/arfanxn/welding/internal/module/material_test_service/usecase/dto"
	"github.com/arfanxn/welding/internal/module/material_test_service/usecase/step"
	"github.com/arfanxn/welding/internal/module/shared/contextkey"
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	"github.com/arfanxn/welding/pkg/pagination"
	"github.com/arfanxn/welding/pkg/query"
	"github.com/arfanxn/welding/pkg/typeutil"
	"go.uber.org/fx"
)

type MaterialTestServiceUsecase interface {
	Paginate(ctx context.Context, q *query.Query) (*pagination.OffsetPagination[*entity.MaterialTestService], error)
	Show(ctx context.Context, q *query.Query) (*entity.MaterialTestService, error)
	Store(ctx context.Context, _dto *dto.SaveMaterialTestService) (*entity.MaterialTestService, error)
	Update(ctx context.Context, _dto *dto.SaveMaterialTestService) (*entity.MaterialTestService, error)
	Destroy(ctx context.Context, _dto *dto.DestroyMaterialTestService) error
}

type mtsUsecase struct {
	activityService activityService.ActivityService
	storeMtsStep    step.StoreMaterialTestServiceStep
	updateMtsStep   step.UpdateMaterialTestServiceStep
	mtsRepository   mtsRepository.MaterialTestServiceRepository
	mtsPolicy       mtsPolicy.MaterialTestServicePolicy
}

type NewMaterialTestServiceUsecaseParams struct {
	fx.In

	ActivityService               activityService.ActivityService
	StoreMaterialTestServiceStep  step.StoreMaterialTestServiceStep
	UpdateMaterialTestServiceStep step.UpdateMaterialTestServiceStep
	MaterialTestServiceRepository mtsRepository.MaterialTestServiceRepository
	MaterialTestServicePolicy     mtsPolicy.MaterialTestServicePolicy
}

func NewMaterialTestServiceUsecase(
	params NewMaterialTestServiceUsecaseParams,
) MaterialTestServiceUsecase {
	return &mtsUsecase{
		activityService: params.ActivityService,
		storeMtsStep:    params.StoreMaterialTestServiceStep,
		updateMtsStep:   params.UpdateMaterialTestServiceStep,
		mtsRepository:   params.MaterialTestServiceRepository,
		mtsPolicy:       params.MaterialTestServicePolicy,
	}
}

func (u *mtsUsecase) Paginate(ctx context.Context, q *query.Query) (op *pagination.OffsetPagination[*entity.MaterialTestService], err error) {
	if op, err = u.mtsRepository.Paginate(q); err != nil {
		return nil, err
	}

	if _, err = u.activityService.Record(ctx, &activityDto.CreateActivity{
		Causer:      contextkey.GetUser(ctx),
		Action:      activityEnum.MaterialTestServicesIndex,
		SubjectType: typeutil.Ptr(activityEnum.MaterialTestServiceSubjectType),
	}); err != nil {
		return nil, err
	}

	return op, nil
}

func (u *mtsUsecase) Show(ctx context.Context, q *query.Query) (mts *entity.MaterialTestService, err error) {
	if mts, err = u.mtsRepository.First(q); err != nil {
		return nil, err
	}

	if _, err = u.activityService.Record(ctx, &activityDto.CreateActivity{
		Causer:  contextkey.GetUser(ctx),
		Action:  activityEnum.MaterialTestServicesShow,
		Subject: mts,
	}); err != nil {
		return nil, err
	}

	return mts, nil
}

func (u *mtsUsecase) Store(ctx context.Context, _dto *dto.SaveMaterialTestService) (mts *entity.MaterialTestService, err error) {
	if err = u.mtsPolicy.Store(ctx, _dto); err != nil {
		return nil, err
	}

	if mts, err = u.storeMtsStep.Handle(ctx, _dto); err != nil {
		return nil, err
	}

	if _, err = u.activityService.Record(ctx, &activityDto.CreateActivity{
		Causer:  contextkey.GetUser(ctx),
		Action:  activityEnum.MaterialTestServicesStore,
		Subject: mts,
	}); err != nil {
		return nil, err
	}

	return
}

func (u *mtsUsecase) Update(ctx context.Context, _dto *dto.SaveMaterialTestService) (mts *entity.MaterialTestService, err error) {
	if err = u.mtsPolicy.Update(ctx, _dto); err != nil {
		return nil, err
	}

	if mts, err = u.updateMtsStep.Handle(ctx, _dto); err != nil {
		return nil, err
	}

	if _, err = u.activityService.Record(ctx, &activityDto.CreateActivity{
		Causer:  contextkey.GetUser(ctx),
		Action:  activityEnum.MaterialTestServicesUpdate,
		Subject: mts,
	}); err != nil {
		return nil, err
	}

	return
}

func (u *mtsUsecase) Destroy(ctx context.Context, _dto *dto.DestroyMaterialTestService) (err error) {
	if err = u.mtsPolicy.Destroy(ctx, _dto); err != nil {
		return err
	}

	mts, err := u.mtsRepository.Find(_dto.Id)
	if err != nil {
		return err
	}

	err = u.mtsRepository.Destroy(mts)
	if err != nil {
		return err
	}

	if _, err = u.activityService.Record(ctx, &activityDto.CreateActivity{
		Causer:  contextkey.GetUser(ctx),
		Action:  activityEnum.MaterialTestServicesDestroy,
		Subject: mts,
	}); err != nil {
		return err
	}

	return nil
}
