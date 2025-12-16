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
	mediaEnum "github.com/arfanxn/welding/internal/module/media/domain/enum"
	mediaDto "github.com/arfanxn/welding/internal/module/media/usecase/dto"
	mediaService "github.com/arfanxn/welding/internal/module/media/usecase/service"
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

	StoreMedia(context.Context, *dto.SaveMaterialTestOrderMedia) (*entity.MaterialTestOrder, error)
	UpdateMedia(context.Context, *dto.SaveMaterialTestOrderMedia) (*entity.MaterialTestOrder, error)
	DestroyMedia(context.Context, *dto.DestroyMaterialTestOrderMedia) error
}

type materialTestOrderUsecase struct {
	activityService activityService.ActivityService
	mediaService    mediaService.MediaService
	saveMtoStep     mtoStep.SaveMaterialTestOrderStep
	mtoRepository   mtoRepository.MaterialTestOrderRepository
	mtoPolicy       mtoPolicy.MaterialTestOrderPolicy
}

type NewMaterialTestOrderUsecaseParams struct {
	fx.In

	ActivityService             activityService.ActivityService
	MediaService                mediaService.MediaService
	SaveMaterialTestOrderStep   mtoStep.SaveMaterialTestOrderStep
	MaterialTestOrderRepository mtoRepository.MaterialTestOrderRepository
	MaterialTestOrderPolicy     mtoPolicy.MaterialTestOrderPolicy
}

func NewMaterialTestOrderUsecase(params NewMaterialTestOrderUsecaseParams) MaterialTestOrderUsecase {
	return &materialTestOrderUsecase{
		activityService: params.ActivityService,
		mediaService:    params.MediaService,
		saveMtoStep:     params.SaveMaterialTestOrderStep,
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

	if mto, err = u.saveMtoStep.Handle(ctx, _dto); err != nil {
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

	if mto, err = u.saveMtoStep.Handle(ctx, _dto); err != nil {
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

func (u *materialTestOrderUsecase) StoreMedia(ctx context.Context, _dto *dto.SaveMaterialTestOrderMedia) (mto *entity.MaterialTestOrder, err error) {
	if err = u.mtoPolicy.StoreMedia(ctx, _dto); err != nil {
		return nil, err
	}

	_, err = u.mediaService.CreateFromMultipartFiles(ctx, mediaDto.CreateFromMultipartFiles{
		mediaDto.CreateFromMultipartFile{
			ModelType:      mediaEnum.ModelTypeMaterialTestOrder,
			ModelId:        *_dto.OrderId,
			CollectionName: mediaEnum.CollectionNameMaterialTestOrder,
			Name:           *_dto.Name,
			File:           _dto.File,
		},
	})
	if err != nil {
		return nil, err
	}

	mto, err = u.mtoRepository.Find(*_dto.OrderId, query.NewQuery().Include("medias"))
	if err != nil {
		return nil, err
	}

	return
}

func (u *materialTestOrderUsecase) UpdateMedia(ctx context.Context, _dto *dto.SaveMaterialTestOrderMedia) (mto *entity.MaterialTestOrder, err error) {
	/*
		TODO: update media

		if err = u.mtoPolicy.StoreMedia(ctx, _dto); err != nil {
			return nil, err
		}

		_, err = u.mediaService.CreateFromMultipartFiles(ctx, mediaDto.CreateFromMultipartFiles{
			mediaDto.CreateFromMultipartFile{
				ModelType: mediaEnum.ModelTypeMaterialTestOrder,
				ModelId:   *_dto.OrderId,
				Name:      *_dto.Name,
				File:      _dto.File,
			},
		})
		if err != nil {
			return nil, err
		}

		mto, err = u.mtoRepository.Find(*_dto.OrderId, query.NewQuery().Include("medias"))
		if err != nil {
			return nil, err
		}
	*/

	return
}

func (u *materialTestOrderUsecase) DestroyMedia(ctx context.Context, _dto *dto.DestroyMaterialTestOrderMedia) (err error) {
	/*
		if err = u.mtoPolicy.StoreMedia(ctx, _dto); err != nil {
			return nil, err
		}

		_, err = u.mediaService.CreateFromMultipartFiles(ctx, mediaDto.CreateFromMultipartFiles{
			mediaDto.CreateFromMultipartFile{
				ModelType: mediaEnum.ModelTypeMaterialTestOrder,
				ModelId:   *_dto.OrderId,
				Name:      *_dto.Name,
				File:      _dto.File,
			},
		})
		if err != nil {
			return nil, err
		}

		mto, err = u.mtoRepository.Find(*_dto.OrderId, query.NewQuery().Include("medias"))
		if err != nil {
			return nil, err
		}
	*/

	return
}
