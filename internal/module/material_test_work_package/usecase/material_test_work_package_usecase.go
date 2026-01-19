package usecase

import (
	"context"

	"github.com/arfanxn/welding/internal/infrastructure/id"
	activityEnum "github.com/arfanxn/welding/internal/module/activity/domain/enum"
	activityDto "github.com/arfanxn/welding/internal/module/activity/usecase/dto"
	activityService "github.com/arfanxn/welding/internal/module/activity/usecase/service"
	materialTestWorkPackageRepository "github.com/arfanxn/welding/internal/module/material_test_work_package/domain/repository"
	materialTestWorkPackagePolicy "github.com/arfanxn/welding/internal/module/material_test_work_package/infrastructure/policy"
	"github.com/arfanxn/welding/internal/module/material_test_work_package/usecase/dto"
	"github.com/arfanxn/welding/internal/module/shared/contextkey"
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	"github.com/arfanxn/welding/pkg/pagination"
	"github.com/arfanxn/welding/pkg/query"
	"github.com/arfanxn/welding/pkg/typeutil"
	"github.com/gookit/goutil"
	"go.uber.org/fx"
)

type MaterialTestWorkPackageUsecase interface {
	Paginate(ctx context.Context, q *query.Query) (*pagination.OffsetPagination[*entity.MaterialTestWorkPackage], error)
	Show(ctx context.Context, q *query.Query) (*entity.MaterialTestWorkPackage, error)
	Store(ctx context.Context, _dto *dto.SaveMaterialTestWorkPackage) (*entity.MaterialTestWorkPackage, error)
	Update(ctx context.Context, _dto *dto.SaveMaterialTestWorkPackage) (*entity.MaterialTestWorkPackage, error)
	Destroy(ctx context.Context, _dto *dto.DestroyMaterialTestWorkPackage) error
}

type materialTestWorkPackageUsecase struct {
	activityService                   activityService.ActivityService
	idService                         id.IdService
	materialTestWorkPackageRepository materialTestWorkPackageRepository.MaterialTestWorkPackageRepository
	materialTestWorkPackagePolicy     materialTestWorkPackagePolicy.MaterialTestWorkPackagePolicy
}

type NewMaterialTestWorkPackageUsecaseParams struct {
	fx.In

	ActivityService                   activityService.ActivityService
	IdService                         id.IdService
	MaterialTestWorkPackageRepository materialTestWorkPackageRepository.MaterialTestWorkPackageRepository
	MaterialTestWorkPackagePolicy     materialTestWorkPackagePolicy.MaterialTestWorkPackagePolicy
}

func NewMaterialTestWorkPackageUsecase(
	params NewMaterialTestWorkPackageUsecaseParams,
) MaterialTestWorkPackageUsecase {
	return &materialTestWorkPackageUsecase{
		activityService:                   params.ActivityService,
		idService:                         params.IdService,
		materialTestWorkPackageRepository: params.MaterialTestWorkPackageRepository,
		materialTestWorkPackagePolicy:     params.MaterialTestWorkPackagePolicy,
	}
}

func (u *materialTestWorkPackageUsecase) Paginate(ctx context.Context, q *query.Query) (op *pagination.OffsetPagination[*entity.MaterialTestWorkPackage], err error) {
	if op, err = u.materialTestWorkPackageRepository.Paginate(q); err != nil {
		return nil, err
	}

	if _, err = u.activityService.Record(ctx, &activityDto.CreateActivity{
		Causer:      contextkey.GetUser(ctx),
		Action:      activityEnum.MaterialTestWorkPackagesIndex,
		SubjectType: typeutil.Ptr(activityEnum.MaterialTestWorkPackageSubjectType),
	}); err != nil {
		return nil, err
	}

	return op, nil
}

func (u *materialTestWorkPackageUsecase) Show(ctx context.Context, q *query.Query) (mtm *entity.MaterialTestWorkPackage, err error) {
	if mtm, err = u.materialTestWorkPackageRepository.First(q); err != nil {
		return nil, err
	}

	if _, err = u.activityService.Record(ctx, &activityDto.CreateActivity{
		Causer:  contextkey.GetUser(ctx),
		Action:  activityEnum.MaterialTestWorkPackagesShow,
		Subject: mtm,
	}); err != nil {
		return nil, err
	}

	return mtm, nil
}

func (u *materialTestWorkPackageUsecase) Store(ctx context.Context, _dto *dto.SaveMaterialTestWorkPackage) (mtm *entity.MaterialTestWorkPackage, err error) {
	if err = u.materialTestWorkPackagePolicy.Store(ctx, _dto); err != nil {
		return nil, err
	}

	mtm = &entity.MaterialTestWorkPackage{}
	mtm.Id = u.idService.Generate()
	mtm.Name = *_dto.Name
	if !goutil.IsEmptyReal(_dto.Description) {
		mtm.Description = _dto.Description
	}

	if err = u.materialTestWorkPackageRepository.Save(mtm); err != nil {
		return nil, err
	}

	if _, err = u.activityService.Record(ctx, &activityDto.CreateActivity{
		Causer:  contextkey.GetUser(ctx),
		Action:  activityEnum.MaterialTestWorkPackagesStore,
		Subject: mtm,
	}); err != nil {
		return nil, err
	}

	return
}

func (u *materialTestWorkPackageUsecase) Update(ctx context.Context, _dto *dto.SaveMaterialTestWorkPackage) (mtm *entity.MaterialTestWorkPackage, err error) {
	if err = u.materialTestWorkPackagePolicy.Update(ctx, _dto); err != nil {
		return nil, err
	}

	if mtm, err = u.materialTestWorkPackageRepository.Find(*_dto.Id, nil); err != nil {
		return nil, err
	}

	if !goutil.IsEmptyReal(_dto.Name) {
		mtm.Name = *_dto.Name
	}

	if !goutil.IsEmptyReal(_dto.Description) {
		mtm.Description = _dto.Description
	}

	if err := u.materialTestWorkPackageRepository.Save(mtm); err != nil {
		return nil, err
	}

	if _, err = u.activityService.Record(ctx, &activityDto.CreateActivity{
		Causer:  contextkey.GetUser(ctx),
		Action:  activityEnum.MaterialTestWorkPackagesUpdate,
		Subject: mtm,
	}); err != nil {
		return nil, err
	}

	return
}

func (u *materialTestWorkPackageUsecase) Destroy(ctx context.Context, _dto *dto.DestroyMaterialTestWorkPackage) (err error) {
	if err = u.materialTestWorkPackagePolicy.Destroy(ctx, _dto); err != nil {
		return err
	}

	mtm, err := u.materialTestWorkPackageRepository.Find(_dto.Id, nil)
	if err != nil {
		return err
	}

	err = u.materialTestWorkPackageRepository.Destroy(mtm)
	if err != nil {
		return err
	}

	if _, err = u.activityService.Record(ctx, &activityDto.CreateActivity{
		Causer:  contextkey.GetUser(ctx),
		Action:  activityEnum.MaterialTestWorkPackagesDestroy,
		Subject: mtm,
	}); err != nil {
		return err
	}

	return nil
}
