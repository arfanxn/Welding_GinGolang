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

type ApprovePaymentMaterialTestOrderStep interface {
	Handle(ctx context.Context, dto *dto.ApprovePaymentMaterialTestOrder) (*entity.MaterialTestOrder, error)
}

type approvePaymentMaterialTestOrderStep struct {
	mtoRepository mtoRepository.MaterialTestOrderRepository
}

type NewApprovePaymentMaterialTestOrderStepParams struct {
	fx.In

	MaterialTestOrderRepository mtoRepository.MaterialTestOrderRepository
}

func NewApprovePaymentMaterialTestOrderStep(params NewApprovePaymentMaterialTestOrderStepParams) ApprovePaymentMaterialTestOrderStep {
	return &approvePaymentMaterialTestOrderStep{
		mtoRepository: params.MaterialTestOrderRepository,
	}
}

func (s *approvePaymentMaterialTestOrderStep) Handle(ctx context.Context, _dto *dto.ApprovePaymentMaterialTestOrder) (
	mto *entity.MaterialTestOrder, err error,
) {
	mto, err = s.mtoRepository.Find(_dto.Id, nil)
	if err != nil {
		return
	}

	mto.Status = mtoEnum.MaterialTestOrderStatusPaymentApproved
	mto.PaymentApprovedAt = typeutil.Ptr(time.Now())
	mto.PaymentRejectedAt = nil

	if err = s.mtoRepository.Save(mto); err != nil {
		return
	}

	return
}
