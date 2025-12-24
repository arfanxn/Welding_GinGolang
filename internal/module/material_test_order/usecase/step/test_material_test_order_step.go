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

type TestMaterialTestOrderStep interface {
	Handle(ctx context.Context, dto *dto.TestMaterialTestOrder) (*entity.MaterialTestOrder, error)
}

type testMaterialTestOrderStep struct {
	mtoRepository mtoRepository.MaterialTestOrderRepository
}

type NewTestMaterialTestOrderStepParams struct {
	fx.In

	MaterialTestOrderRepository mtoRepository.MaterialTestOrderRepository
}

func NewTestMaterialTestOrderStep(params NewTestMaterialTestOrderStepParams) TestMaterialTestOrderStep {
	return &testMaterialTestOrderStep{
		mtoRepository: params.MaterialTestOrderRepository,
	}
}

func (s *testMaterialTestOrderStep) Handle(ctx context.Context, _dto *dto.TestMaterialTestOrder) (
	mto *entity.MaterialTestOrder, err error,
) {
	mto, err = s.mtoRepository.Find(_dto.Id, nil)
	if err != nil {
		return
	}

	mto.Status = mtoEnum.MaterialTestOrderStatusTesting
	mto.TestingAt = typeutil.Ptr(time.Now())

	if err = s.mtoRepository.Save(mto); err != nil {
		return
	}

	return
}
