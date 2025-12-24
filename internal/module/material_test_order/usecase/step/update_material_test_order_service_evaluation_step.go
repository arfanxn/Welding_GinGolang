package step

import (
	"context"

	mtoRepository "github.com/arfanxn/welding/internal/module/material_test_order/domain/repository"
	"github.com/arfanxn/welding/internal/module/material_test_order/usecase/dto"
	mtoseRepository "github.com/arfanxn/welding/internal/module/material_test_order_service_evaluation/domain/repository"
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	"github.com/arfanxn/welding/pkg/query"
	"go.uber.org/fx"
)

type UpdateMaterialTestOrderServiceEvaluationStep interface {
	Handle(ctx context.Context, dto *dto.UpdateMaterialTestOrderServiceEvaluation) (*entity.MaterialTestOrder, error)
}

type updateMaterialTestOrderServiceEvaluationStep struct {
	mtoRepository   mtoRepository.MaterialTestOrderRepository
	mtoseRepository mtoseRepository.MaterialTestOrderServiceEvaluationRepository
}

type NewUpdateMaterialTestOrderServiceEvaluationStepParams struct {
	fx.In

	MaterialTestOrderRepository                  mtoRepository.MaterialTestOrderRepository
	MaterialTestOrderServiceEvaluationRepository mtoseRepository.MaterialTestOrderServiceEvaluationRepository
}

func NewUpdateMaterialTestOrderServiceEvaluationStep(params NewUpdateMaterialTestOrderServiceEvaluationStepParams) UpdateMaterialTestOrderServiceEvaluationStep {
	return &updateMaterialTestOrderServiceEvaluationStep{
		mtoRepository:   params.MaterialTestOrderRepository,
		mtoseRepository: params.MaterialTestOrderServiceEvaluationRepository,
	}
}

func (s *updateMaterialTestOrderServiceEvaluationStep) Handle(ctx context.Context, _dto *dto.UpdateMaterialTestOrderServiceEvaluation) (
	mto *entity.MaterialTestOrder, err error,
) {
	mtose, err := s.mtoseRepository.Find(_dto.OrderServiceEvaluationId, nil)
	if err != nil {
		return nil, err
	}

	mtose.IsEquipmentAvailable = _dto.IsEquipmentAvailable
	mtose.IsPersonnelAvailable = _dto.IsPersonnelAvailable
	mtose.IsTimeAvailable = _dto.IsTimeAvailable
	mtose.IsTestReady = _dto.IsTestReady
	mtose.IsSubcontractLabAvailable = _dto.IsSubcontractLabAvailable

	err = s.mtoseRepository.Save(mtose)
	if err != nil {
		return nil, err
	}

	mto, err = s.mtoRepository.Find(_dto.Id, query.NewQuery().Include("ordered_services.evaluation"))
	if err != nil {
		return nil, err
	}

	return
}
