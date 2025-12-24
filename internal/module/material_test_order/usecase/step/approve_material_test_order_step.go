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

type ApproveMaterialTestOrderStep interface {
	Handle(ctx context.Context, dto *dto.ApproveMaterialTestOrder) (*entity.MaterialTestOrder, error)
}

type approveMaterialTestOrderStep struct {
	mtoRepository mtoRepository.MaterialTestOrderRepository
}

type NewApproveMaterialTestOrderStepParams struct {
	fx.In

	MaterialTestOrderRepository mtoRepository.MaterialTestOrderRepository
}

func NewApproveMaterialTestOrderStep(params NewApproveMaterialTestOrderStepParams) ApproveMaterialTestOrderStep {
	return &approveMaterialTestOrderStep{
		mtoRepository: params.MaterialTestOrderRepository,
	}
}

func (s *approveMaterialTestOrderStep) Handle(ctx context.Context, _dto *dto.ApproveMaterialTestOrder) (
	mto *entity.MaterialTestOrder, err error,
) {
	mto, err = s.mtoRepository.Find(_dto.Id, nil)
	if err != nil {
		return
	}

	mto.Status = mtoEnum.MaterialTestOrderStatusAwaitingPayment
	mto.ApprovedAt = typeutil.Ptr(time.Now())
	mto.RejectedAt = nil

	if err = s.mtoRepository.Save(mto); err != nil {
		return
	}

	return
}
