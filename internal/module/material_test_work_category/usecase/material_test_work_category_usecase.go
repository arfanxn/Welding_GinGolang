package usecase

import (
	"context"

	"github.com/arfanxn/welding/internal/infrastructure/id"
	activityEnum "github.com/arfanxn/welding/internal/module/activity/domain/enum"
	activityDto "github.com/arfanxn/welding/internal/module/activity/usecase/dto"
	activityService "github.com/arfanxn/welding/internal/module/activity/usecase/service"
	materialTestWorkCategoryRepository "github.com/arfanxn/welding/internal/module/material_test_work_category/domain/repository"
	materialTestWorkCategoryPolicy "github.com/arfanxn/welding/internal/module/material_test_work_category/infrastructure/policy"
	"github.com/arfanxn/welding/internal/module/material_test_work_category/usecase/dto"
	"github.com/arfanxn/welding/internal/module/shared/contextkey"
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	"github.com/arfanxn/welding/pkg/pagination"
	"github.com/arfanxn/welding/pkg/query"
	"github.com/arfanxn/welding/pkg/typeutil"
	"github.com/gookit/goutil"
	"go.uber.org/fx"
)

type MaterialTestWorkCategoryUsecase interface {
	Paginate(ctx context.Context, q *query.Query) (*pagination.OffsetPagination[*entity.MaterialTestWorkCategory], error)
	Show(ctx context.Context, q *query.Query) (*entity.MaterialTestWorkCategory, error)
	Store(ctx context.Context, _dto *dto.SaveMaterialTestWorkCategory) (*entity.MaterialTestWorkCategory, error)
	Update(ctx context.Context, _dto *dto.SaveMaterialTestWorkCategory) (*entity.MaterialTestWorkCategory, error)
	Destroy(ctx context.Context, _dto *dto.DestroyMaterialTestWorkCategory) error
}

type materialTestWorkCategoryUsecase struct {
	activityService                    activityService.ActivityService
	idService                          id.IdService
	materialTestWorkCategoryRepository materialTestWorkCategoryRepository.MaterialTestWorkCategoryRepository
	materialTestWorkCategoryPolicy     materialTestWorkCategoryPolicy.MaterialTestWorkCategoryPolicy
}

type NewMaterialTestWorkCategoryUsecaseParams struct {
	fx.In

	ActivityService                    activityService.ActivityService
	IdService                          id.IdService
	MaterialTestWorkCategoryRepository materialTestWorkCategoryRepository.MaterialTestWorkCategoryRepository
	MaterialTestWorkCategoryPolicy     materialTestWorkCategoryPolicy.MaterialTestWorkCategoryPolicy
}

func NewMaterialTestWorkCategoryUsecase(
	params NewMaterialTestWorkCategoryUsecaseParams,
) MaterialTestWorkCategoryUsecase {
	return &materialTestWorkCategoryUsecase{
		activityService:                    params.ActivityService,
		idService:                          params.IdService,
		materialTestWorkCategoryRepository: params.MaterialTestWorkCategoryRepository,
		materialTestWorkCategoryPolicy:     params.MaterialTestWorkCategoryPolicy,
	}
}

func (u *materialTestWorkCategoryUsecase) Paginate(ctx context.Context, q *query.Query) (op *pagination.OffsetPagination[*entity.MaterialTestWorkCategory], err error) {
	if op, err = u.materialTestWorkCategoryRepository.Paginate(q); err != nil {
		return nil, err
	}

	if _, err = u.activityService.Record(ctx, &activityDto.CreateActivity{
		Causer:      contextkey.GetUser(ctx),
		Action:      activityEnum.MaterialTestWorkCategoriesIndex,
		SubjectType: typeutil.Ptr(activityEnum.MaterialTestWorkCategorySubjectType),
	}); err != nil {
		return nil, err
	}

	return op, nil
}

func (u *materialTestWorkCategoryUsecase) Show(ctx context.Context, q *query.Query) (mtm *entity.MaterialTestWorkCategory, err error) {
	if mtm, err = u.materialTestWorkCategoryRepository.First(q); err != nil {
		return nil, err
	}

	if _, err = u.activityService.Record(ctx, &activityDto.CreateActivity{
		Causer:  contextkey.GetUser(ctx),
		Action:  activityEnum.MaterialTestWorkCategoriesShow,
		Subject: mtm,
	}); err != nil {
		return nil, err
	}

	return mtm, nil
}

func (u *materialTestWorkCategoryUsecase) Store(ctx context.Context, _dto *dto.SaveMaterialTestWorkCategory) (mtm *entity.MaterialTestWorkCategory, err error) {
	if err = u.materialTestWorkCategoryPolicy.Store(ctx, _dto); err != nil {
		return nil, err
	}

	mtm = &entity.MaterialTestWorkCategory{}
	mtm.Id = u.idService.Generate()
	mtm.Name = *_dto.Name
	if !goutil.IsEmptyReal(_dto.Description) {
		mtm.Description = _dto.Description
	}

	if err = u.materialTestWorkCategoryRepository.Save(mtm); err != nil {
		return nil, err
	}

	if _, err = u.activityService.Record(ctx, &activityDto.CreateActivity{
		Causer:  contextkey.GetUser(ctx),
		Action:  activityEnum.MaterialTestWorkCategoriesStore,
		Subject: mtm,
	}); err != nil {
		return nil, err
	}

	return
}

func (u *materialTestWorkCategoryUsecase) Update(ctx context.Context, _dto *dto.SaveMaterialTestWorkCategory) (mtm *entity.MaterialTestWorkCategory, err error) {
	if err = u.materialTestWorkCategoryPolicy.Update(ctx, _dto); err != nil {
		return nil, err
	}

	if mtm, err = u.materialTestWorkCategoryRepository.Find(*_dto.Id, nil); err != nil {
		return nil, err
	}

	if !goutil.IsEmptyReal(_dto.Name) {
		mtm.Name = *_dto.Name
	}

	if !goutil.IsEmptyReal(_dto.Description) {
		mtm.Description = _dto.Description
	}

	if err := u.materialTestWorkCategoryRepository.Save(mtm); err != nil {
		return nil, err
	}

	if _, err = u.activityService.Record(ctx, &activityDto.CreateActivity{
		Causer:  contextkey.GetUser(ctx),
		Action:  activityEnum.MaterialTestWorkCategoriesUpdate,
		Subject: mtm,
	}); err != nil {
		return nil, err
	}

	return
}

func (u *materialTestWorkCategoryUsecase) Destroy(ctx context.Context, _dto *dto.DestroyMaterialTestWorkCategory) (err error) {
	if err = u.materialTestWorkCategoryPolicy.Destroy(ctx, _dto); err != nil {
		return err
	}

	mtm, err := u.materialTestWorkCategoryRepository.Find(_dto.Id, nil)
	if err != nil {
		return err
	}

	err = u.materialTestWorkCategoryRepository.Destroy(mtm)
	if err != nil {
		return err
	}

	if _, err = u.activityService.Record(ctx, &activityDto.CreateActivity{
		Causer:  contextkey.GetUser(ctx),
		Action:  activityEnum.MaterialTestWorkCategoriesDestroy,
		Subject: mtm,
	}); err != nil {
		return err
	}

	return nil
}
