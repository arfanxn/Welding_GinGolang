package step

import (
	"context"

	mtoRepository "github.com/arfanxn/welding/internal/module/material_test_order/domain/repository"
	"github.com/arfanxn/welding/internal/module/material_test_order/usecase/dto"
	mediaEnum "github.com/arfanxn/welding/internal/module/media/domain/enum"
	mediaDto "github.com/arfanxn/welding/internal/module/media/usecase/dto"
	mediaService "github.com/arfanxn/welding/internal/module/media/usecase/service"
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	"github.com/arfanxn/welding/pkg/query"
	"go.uber.org/fx"
)

type StoreMaterialTestOrderMediaStep interface {
	Handle(ctx context.Context, dto *dto.SaveMaterialTestOrderMedia) (*entity.MaterialTestOrder, error)
}

type storeMaterialTestOrderMediaStep struct {
	mtoRepository mtoRepository.MaterialTestOrderRepository
	mediaService  mediaService.MediaService
}

type NewStoreMaterialTestOrderMediaStepParams struct {
	fx.In

	MaterialTestOrderRepository mtoRepository.MaterialTestOrderRepository
	MediaService                mediaService.MediaService
}

func NewStoreMaterialTestOrderMediaStep(params NewStoreMaterialTestOrderMediaStepParams) StoreMaterialTestOrderMediaStep {
	return &storeMaterialTestOrderMediaStep{
		mtoRepository: params.MaterialTestOrderRepository,
		mediaService:  params.MediaService,
	}
}

func (s *storeMaterialTestOrderMediaStep) Handle(ctx context.Context, _dto *dto.SaveMaterialTestOrderMedia) (
	mto *entity.MaterialTestOrder, err error,
) {
	_, err = s.mediaService.CreateFromMultipartFiles(ctx, mediaDto.CreateFromMultipartFiles{
		mediaDto.CreateFromMultipartFile{
			ModelType:      mediaEnum.ModelTypeMaterialTestOrder,
			ModelId:        *_dto.Id,
			CollectionName: mediaEnum.CollectionNameMaterialTestOrder,
			Name:           *_dto.Name,
			File:           _dto.File,
		},
	})
	if err != nil {
		return nil, err
	}

	mto, err = s.mtoRepository.Find(*_dto.Id, query.NewQuery().Include("medias"))
	if err != nil {
		return nil, err
	}

	return
}
