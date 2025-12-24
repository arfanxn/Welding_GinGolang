package step

import (
	"context"

	mtoRepository "github.com/arfanxn/welding/internal/module/material_test_order/domain/repository"
	"github.com/arfanxn/welding/internal/module/material_test_order/usecase/dto"
	mediaService "github.com/arfanxn/welding/internal/module/media/usecase/service"
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	"github.com/arfanxn/welding/pkg/query"
	"go.uber.org/fx"
)

type DestroyMaterialTestOrderMediaStep interface {
	Handle(ctx context.Context, dto *dto.DestroyMaterialTestOrderMedia) (*entity.MaterialTestOrder, error)
}

type destroyMaterialTestOrderMediaStep struct {
	mtoRepository mtoRepository.MaterialTestOrderRepository
	mediaService  mediaService.MediaService
}

type NewDestroyMaterialTestOrderMediaStepParams struct {
	fx.In

	MaterialTestOrderRepository mtoRepository.MaterialTestOrderRepository
	MediaService                mediaService.MediaService
}

func NewDestroyMaterialTestOrderMediaStep(params NewDestroyMaterialTestOrderMediaStepParams) DestroyMaterialTestOrderMediaStep {
	return &destroyMaterialTestOrderMediaStep{
		mtoRepository: params.MaterialTestOrderRepository,
		mediaService:  params.MediaService,
	}
}

func (s *destroyMaterialTestOrderMediaStep) Handle(ctx context.Context, _dto *dto.DestroyMaterialTestOrderMedia) (
	mto *entity.MaterialTestOrder, err error,
) {
	err = s.mediaService.DestroyByQuery(ctx, query.NewQuery().FilterById(_dto.MediaId))
	if err != nil {
		return nil, err
	}

	q := query.NewQuery().Include("medias")
	mto, err = s.mtoRepository.Find(_dto.Id, q)
	if err != nil {
		return
	}

	return
}
