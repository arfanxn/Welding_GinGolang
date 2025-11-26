package usecase

import (
	"context"

	"github.com/arfanxn/welding/internal/infrastructure/id"
	activityEnum "github.com/arfanxn/welding/internal/module/activity/domain/enum"
	activityDto "github.com/arfanxn/welding/internal/module/activity/usecase/dto"
	activityService "github.com/arfanxn/welding/internal/module/activity/usecase/service"
	mtmRepository "github.com/arfanxn/welding/internal/module/material_test_machine/domain/repository"
	materialTestMachinePolicy "github.com/arfanxn/welding/internal/module/material_test_machine/infrastructure/policy"
	"github.com/arfanxn/welding/internal/module/material_test_machine/usecase/dto"
	"github.com/arfanxn/welding/internal/module/shared/contextkey"
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	"github.com/arfanxn/welding/pkg/pagination"
	"github.com/arfanxn/welding/pkg/query"
	"github.com/arfanxn/welding/pkg/typeutil"
	"github.com/gookit/goutil"
	"go.uber.org/fx"
)

type MaterialTestMachineUsecase interface {
	Paginate(ctx context.Context, q *query.Query) (*pagination.OffsetPagination[*entity.MaterialTestMachine], error)
	Show(ctx context.Context, q *query.Query) (*entity.MaterialTestMachine, error)
	Store(ctx context.Context, _dto *dto.SaveMaterialTestMachine) (*entity.MaterialTestMachine, error)
	Update(ctx context.Context, _dto *dto.SaveMaterialTestMachine) (*entity.MaterialTestMachine, error)
	Destroy(ctx context.Context, _dto *dto.DestroyMaterialTestMachine) error
}

type mtmUsecase struct {
	activityService activityService.ActivityService
	idService       id.IdService
	mtmRepository   mtmRepository.MaterialTestMachineRepository
	mtmPolicy       materialTestMachinePolicy.MaterialTestMachinePolicy
}

type NewMaterialTestMachineUsecaseParams struct {
	fx.In

	ActivityService activityService.ActivityService
	IdService       id.IdService
	MtmRepository   mtmRepository.MaterialTestMachineRepository
	MtmPolicy       materialTestMachinePolicy.MaterialTestMachinePolicy
}

func NewMaterialTestMachineUsecase(
	params NewMaterialTestMachineUsecaseParams,
) MaterialTestMachineUsecase {
	return &mtmUsecase{
		activityService: params.ActivityService,
		idService:       params.IdService,
		mtmRepository:   params.MtmRepository,
		mtmPolicy:       params.MtmPolicy,
	}
}

func (u *mtmUsecase) Paginate(ctx context.Context, q *query.Query) (op *pagination.OffsetPagination[*entity.MaterialTestMachine], err error) {
	if op, err = u.mtmRepository.Paginate(q); err != nil {
		return nil, err
	}

	if _, err = u.activityService.Record(ctx, &activityDto.CreateActivity{
		Causer:      contextkey.GetUser(ctx),
		Action:      activityEnum.MaterialTestMachinesIndex,
		SubjectType: typeutil.Ptr(activityEnum.MaterialTestMachineSubjectType),
	}); err != nil {
		return nil, err
	}

	return op, nil
}

func (u *mtmUsecase) Show(ctx context.Context, q *query.Query) (mtm *entity.MaterialTestMachine, err error) {
	if mtm, err = u.mtmRepository.First(q); err != nil {
		return nil, err
	}

	if _, err = u.activityService.Record(ctx, &activityDto.CreateActivity{
		Causer:  contextkey.GetUser(ctx),
		Action:  activityEnum.MaterialTestMachinesShow,
		Subject: mtm,
	}); err != nil {
		return nil, err
	}

	return mtm, nil
}

func (u *mtmUsecase) Store(ctx context.Context, _dto *dto.SaveMaterialTestMachine) (mtm *entity.MaterialTestMachine, err error) {
	if err = u.mtmPolicy.Store(ctx, _dto); err != nil {
		return nil, err
	}

	mtm = &entity.MaterialTestMachine{}
	mtm.Id = u.idService.Generate()
	mtm.Name = *_dto.Name
	if !goutil.IsEmptyReal(_dto.Description) {
		mtm.Description = _dto.Description
	}

	if err = u.mtmRepository.Save(mtm); err != nil {
		return nil, err
	}

	if _, err = u.activityService.Record(ctx, &activityDto.CreateActivity{
		Causer:  contextkey.GetUser(ctx),
		Action:  activityEnum.MaterialTestMachinesStore,
		Subject: mtm,
	}); err != nil {
		return nil, err
	}

	return
}

func (u *mtmUsecase) Update(ctx context.Context, _dto *dto.SaveMaterialTestMachine) (mtm *entity.MaterialTestMachine, err error) {
	if err = u.mtmPolicy.Update(ctx, _dto); err != nil {
		return nil, err
	}

	if mtm, err = u.mtmRepository.Find(*_dto.Id); err != nil {
		return nil, err
	}

	if !goutil.IsEmptyReal(_dto.Name) {
		mtm.Name = *_dto.Name
	}

	if !goutil.IsEmptyReal(_dto.Description) {
		mtm.Description = _dto.Description
	}

	if err := u.mtmRepository.Save(mtm); err != nil {
		return nil, err
	}

	if _, err = u.activityService.Record(ctx, &activityDto.CreateActivity{
		Causer:  contextkey.GetUser(ctx),
		Action:  activityEnum.MaterialTestMachinesUpdate,
		Subject: mtm,
	}); err != nil {
		return nil, err
	}

	return
}

func (u *mtmUsecase) Destroy(ctx context.Context, _dto *dto.DestroyMaterialTestMachine) (err error) {
	if err = u.mtmPolicy.Destroy(ctx, _dto); err != nil {
		return err
	}

	mtm, err := u.mtmRepository.Find(_dto.Id)
	if err != nil {
		return err
	}

	err = u.mtmRepository.Destroy(mtm)
	if err != nil {
		return err
	}

	if _, err = u.activityService.Record(ctx, &activityDto.CreateActivity{
		Causer:  contextkey.GetUser(ctx),
		Action:  activityEnum.MaterialTestMachinesDestroy,
		Subject: mtm,
	}); err != nil {
		return err
	}

	return nil
}
