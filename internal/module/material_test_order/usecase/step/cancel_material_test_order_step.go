package step

import (
	"context"
	"time"

	mtoEnum "github.com/arfanxn/welding/internal/module/material_test_order/domain/enum"
	mtoRepository "github.com/arfanxn/welding/internal/module/material_test_order/domain/repository"
	"github.com/arfanxn/welding/internal/module/material_test_order/usecase/dto"
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	"github.com/arfanxn/welding/pkg/typeutil"
	"go.uber.org/fx"
)

type CancelMaterialTestOrderStep interface {
	Handle(ctx context.Context, dto *dto.CancelMaterialTestOrder) (*entity.MaterialTestOrder, error)
}

type cancelMaterialTestOrderStep struct {
	mtoRepository mtoRepository.MaterialTestOrderRepository
}

type NewCancelMaterialTestOrderStepParams struct {
	fx.In

	MaterialTestOrderRepository mtoRepository.MaterialTestOrderRepository
}

func NewCancelMaterialTestOrderStep(params NewCancelMaterialTestOrderStepParams) CancelMaterialTestOrderStep {
	return &cancelMaterialTestOrderStep{
		mtoRepository: params.MaterialTestOrderRepository,
	}
}

func (s *cancelMaterialTestOrderStep) Handle(ctx context.Context, _dto *dto.CancelMaterialTestOrder) (
	mto *entity.MaterialTestOrder, err error,
) {
	mto, err = s.mtoRepository.Find(_dto.Id, nil)
	if err != nil {
		return
	}

	mto.Status = mtoEnum.MaterialTestOrderStatusCancelled
	mto.CancelledAt = typeutil.Ptr(time.Now())

	if err = s.mtoRepository.Save(mto); err != nil {
		return
	}

	return
}
