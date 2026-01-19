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

type RejectMaterialTestOrderStep interface {
	Handle(ctx context.Context, dto *dto.RejectMaterialTestOrder) (*entity.MaterialTestOrder, error)
}

type rejectMaterialTestOrderStep struct {
	mtoRepository mtoRepository.MaterialTestOrderRepository
}

type NewRejectMaterialTestOrderStepParams struct {
	fx.In

	MaterialTestOrderRepository mtoRepository.MaterialTestOrderRepository
}

func NewRejectMaterialTestOrderStep(params NewRejectMaterialTestOrderStepParams) RejectMaterialTestOrderStep {
	return &rejectMaterialTestOrderStep{
		mtoRepository: params.MaterialTestOrderRepository,
	}
}

func (s *rejectMaterialTestOrderStep) Handle(ctx context.Context, _dto *dto.RejectMaterialTestOrder) (
	mto *entity.MaterialTestOrder, err error,
) {
	mto, err = s.mtoRepository.Find(_dto.Id, nil)
	if err != nil {
		return
	}

	mto.Status = mtoEnum.MaterialTestOrderStatusRejected
	mto.RejectedAt = typeutil.Ptr(time.Now())

	if err = s.mtoRepository.Save(mto); err != nil {
		return
	}

	return
}
