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

type RejectPaymentMaterialTestOrderStep interface {
	Handle(ctx context.Context, dto *dto.RejectPaymentMaterialTestOrder) (*entity.MaterialTestOrder, error)
}

type rejectPaymentMaterialTestOrderStep struct {
	mtoRepository mtoRepository.MaterialTestOrderRepository
}

type NewRejectPaymentMaterialTestOrderStepParams struct {
	fx.In

	MaterialTestOrderRepository mtoRepository.MaterialTestOrderRepository
}

func NewRejectPaymentMaterialTestOrderStep(params NewRejectPaymentMaterialTestOrderStepParams) RejectPaymentMaterialTestOrderStep {
	return &rejectPaymentMaterialTestOrderStep{
		mtoRepository: params.MaterialTestOrderRepository,
	}
}

func (s *rejectPaymentMaterialTestOrderStep) Handle(ctx context.Context, _dto *dto.RejectPaymentMaterialTestOrder) (
	mto *entity.MaterialTestOrder, err error,
) {
	mto, err = s.mtoRepository.Find(_dto.Id, nil)
	if err != nil {
		return
	}

	mto.Status = mtoEnum.MaterialTestOrderStatusPaymentRejected
	mto.PaymentRejectedAt = typeutil.Ptr(time.Now())

	if err = s.mtoRepository.Save(mto); err != nil {
		return
	}

	return
}
