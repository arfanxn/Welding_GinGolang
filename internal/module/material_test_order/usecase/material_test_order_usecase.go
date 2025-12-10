package usecase

import (
	"context"

	activityEnum "github.com/arfanxn/welding/internal/module/activity/domain/enum"
	activityDto "github.com/arfanxn/welding/internal/module/activity/usecase/dto"
	activityService "github.com/arfanxn/welding/internal/module/activity/usecase/service"
	mtoRepository "github.com/arfanxn/welding/internal/module/material_test_order/domain/repository"
	mtoPolicy "github.com/arfanxn/welding/internal/module/material_test_order/infrastructure/policy"
	"github.com/arfanxn/welding/internal/module/material_test_order/usecase/dto"
	mtoStep "github.com/arfanxn/welding/internal/module/material_test_order/usecase/step"
	"github.com/arfanxn/welding/internal/module/shared/contextkey"
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	"github.com/arfanxn/welding/pkg/pagination"
	"github.com/arfanxn/welding/pkg/query"
	"github.com/arfanxn/welding/pkg/typeutil"
	"go.uber.org/fx"
)

type MaterialTestOrderUsecase interface {
	Paginate(context.Context, *query.Query) (*pagination.OffsetPagination[*entity.MaterialTestOrder], error)
	Show(context.Context, *query.Query) (*entity.MaterialTestOrder, error)
	Store(context.Context, *dto.SaveMaterialTestOrder) (*entity.MaterialTestOrder, error)
	Update(context.Context, *dto.SaveMaterialTestOrder) (*entity.MaterialTestOrder, error)
	Destroy(context.Context, *dto.DestroyMaterialTestOrder) error
}

type materialTestOrderUsecase struct {
	activityService activityService.ActivityService
	storeMtoStep    mtoStep.StoreMaterialTestOrderStep
	updateMtoStep   mtoStep.UpdateMaterialTestOrderStep
	mtoRepository   mtoRepository.MaterialTestOrderRepository
	mtoPolicy       mtoPolicy.MaterialTestOrderPolicy
}

type NewMaterialTestOrderUsecaseParams struct {
	fx.In

	ActivityService             activityService.ActivityService
	StoreMaterialTestOrderStep  mtoStep.StoreMaterialTestOrderStep
	UpdateMaterialTestOrderStep mtoStep.UpdateMaterialTestOrderStep
	MaterialTestOrderRepository mtoRepository.MaterialTestOrderRepository
	MaterialTestOrderPolicy     mtoPolicy.MaterialTestOrderPolicy
}

func NewMaterialTestOrderUsecase(params NewMaterialTestOrderUsecaseParams) MaterialTestOrderUsecase {
	return &materialTestOrderUsecase{
		activityService: params.ActivityService,
		storeMtoStep:    params.StoreMaterialTestOrderStep,
		updateMtoStep:   params.UpdateMaterialTestOrderStep,
		mtoRepository:   params.MaterialTestOrderRepository,
		mtoPolicy:       params.MaterialTestOrderPolicy,
	}
}

func (u *materialTestOrderUsecase) Paginate(ctx context.Context, q *query.Query) (op *pagination.OffsetPagination[*entity.MaterialTestOrder], err error) {
	op, err = u.mtoRepository.Paginate(q)
	if err != nil {
		return nil, err
	}

	if _, err = u.activityService.Record(ctx, &activityDto.CreateActivity{
		Causer:      contextkey.GetUser(ctx),
		Action:      activityEnum.MaterialTestOrdersIndex,
		SubjectType: typeutil.Ptr(activityEnum.MaterialTestOrderSubjectType),
	}); err != nil {
		return nil, err
	}

	return
}

func (u *materialTestOrderUsecase) Show(ctx context.Context, q *query.Query) (mto *entity.MaterialTestOrder, err error) {
	mto, err = u.mtoRepository.First(q)
	if err != nil {
		return nil, err
	}

	if _, err = u.activityService.Record(ctx, &activityDto.CreateActivity{
		Causer:  contextkey.GetUser(ctx),
		Action:  activityEnum.MaterialTestOrdersShow,
		Subject: mto,
	}); err != nil {
		return nil, err
	}

	return
}

func (u *materialTestOrderUsecase) Store(ctx context.Context, _dto *dto.SaveMaterialTestOrder) (mto *entity.MaterialTestOrder, err error) {
	if err = u.mtoPolicy.Store(ctx, _dto); err != nil {
		return nil, err
	}

	if mto, err = u.storeMtoStep.Handle(ctx, _dto); err != nil {
		return nil, err
	}

	if _, err = u.activityService.Record(ctx, &activityDto.CreateActivity{
		Causer:  contextkey.GetUser(ctx),
		Action:  activityEnum.MaterialTestOrdersStore,
		Subject: mto,
	}); err != nil {
		return nil, err
	}

	return
}

func (u *materialTestOrderUsecase) Update(ctx context.Context, _dto *dto.SaveMaterialTestOrder) (mto *entity.MaterialTestOrder, err error) {
	if err = u.mtoPolicy.Update(ctx, _dto); err != nil {
		return nil, err
	}

	if mto, err = u.updateMtoStep.Handle(ctx, _dto); err != nil {
		return nil, err
	}

	if _, err = u.activityService.Record(ctx, &activityDto.CreateActivity{
		Causer:  contextkey.GetUser(ctx),
		Action:  activityEnum.MaterialTestOrdersUpdate,
		Subject: mto,
	}); err != nil {
		return nil, err
	}

	return
}

func (u *materialTestOrderUsecase) Destroy(ctx context.Context, _dto *dto.DestroyMaterialTestOrder) (err error) {
	if err = u.mtoPolicy.Destroy(ctx, _dto); err != nil {
		return err
	}

	mto, err := u.mtoRepository.Find(_dto.Id, nil)
	if err != nil {
		return err
	}

	err = u.mtoRepository.Destroy(mto)
	if err != nil {
		return err
	}

	if _, err = u.activityService.Record(ctx, &activityDto.CreateActivity{
		Causer:  contextkey.GetUser(ctx),
		Action:  activityEnum.MaterialTestOrdersDestroy,
		Subject: mto,
	}); err != nil {
		return err
	}
	return nil
}
