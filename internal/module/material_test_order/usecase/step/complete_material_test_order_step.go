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

type CompleteMaterialTestOrderStep interface {
	Handle(ctx context.Context, dto *dto.CompleteMaterialTestOrder) (*entity.MaterialTestOrder, error)
}

type completeMaterialTestOrderStep struct {
	mtoRepository mtoRepository.MaterialTestOrderRepository
}

type NewCompleteMaterialTestOrderStepParams struct {
	fx.In

	MaterialTestOrderRepository mtoRepository.MaterialTestOrderRepository
}

func NewCompleteMaterialTestOrderStep(params NewCompleteMaterialTestOrderStepParams) CompleteMaterialTestOrderStep {
	return &completeMaterialTestOrderStep{
		mtoRepository: params.MaterialTestOrderRepository,
	}
}

func (s *completeMaterialTestOrderStep) Handle(ctx context.Context, _dto *dto.CompleteMaterialTestOrder) (
	mto *entity.MaterialTestOrder, err error,
) {
	mto, err = s.mtoRepository.Find(_dto.Id, nil)
	if err != nil {
		return
	}

	mto.Status = mtoEnum.MaterialTestOrderStatusCompleted
	mto.CompletedAt = typeutil.Ptr(time.Now())

	if err = s.mtoRepository.Save(mto); err != nil {
		return
	}

	return
}
