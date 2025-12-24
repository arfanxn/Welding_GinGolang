package step

import (
	"context"

	mtoRepository "github.com/arfanxn/welding/internal/module/material_test_order/domain/repository"
	"github.com/arfanxn/welding/internal/module/material_test_order/usecase/dto"
	mediaRepository "github.com/arfanxn/welding/internal/module/media/domain/repository"
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	"github.com/arfanxn/welding/pkg/query"
	"github.com/gookit/goutil"
	"go.uber.org/fx"
)

type UpdateMaterialTestOrderMediaStep interface {
	Handle(ctx context.Context, dto *dto.SaveMaterialTestOrderMedia) (*entity.MaterialTestOrder, error)
}

type updateMaterialTestOrderMediaStep struct {
	mtoRepository   mtoRepository.MaterialTestOrderRepository
	mediaRepository mediaRepository.MediaRepository
}

type NewUpdateMaterialTestOrderMediaStepParams struct {
	fx.In

	MaterialTestOrderRepository mtoRepository.MaterialTestOrderRepository
	MediaRepository             mediaRepository.MediaRepository
}

func NewUpdateMaterialTestOrderMediaStep(params NewUpdateMaterialTestOrderMediaStepParams) UpdateMaterialTestOrderMediaStep {
	return &updateMaterialTestOrderMediaStep{
		mtoRepository:   params.MaterialTestOrderRepository,
		mediaRepository: params.MediaRepository,
	}
}

func (s *updateMaterialTestOrderMediaStep) Handle(ctx context.Context, _dto *dto.SaveMaterialTestOrderMedia) (
	mto *entity.MaterialTestOrder, err error,
) {
	media, err := s.mediaRepository.Find(*_dto.MediaId, nil)
	if err != nil {
		return nil, err
	}

	if !goutil.IsEmptyReal(_dto.Name) {
		media.Name = *_dto.Name
	}

	err = s.mediaRepository.Save(media)
	if err != nil {
		return nil, err
	}

	mto, err = s.mtoRepository.Find(*_dto.Id, query.NewQuery().Include("medias"))
	if err != nil {
		return nil, err
	}

	return
}
