package policy

import (
	"context"

	material_test_order_repository "github.com/arfanxn/welding/internal/module/material_test_order/domain/repository"
	"github.com/arfanxn/welding/internal/module/material_test_order/usecase/dto"
	"go.uber.org/fx"
)

type MaterialTestOrderPolicy interface {
	Store(ctx context.Context, _dto *dto.SaveMaterialTestOrder) error
	Update(ctx context.Context, _dto *dto.SaveMaterialTestOrder) error
	Destroy(ctx context.Context, _dto *dto.DestroyMaterialTestOrder) error
}

type materialTestOrderPolicy struct {
	materialTestOrderRepository material_test_order_repository.MaterialTestOrderRepository
}

type NewMaterialTestOrderPolicyParams struct {
	fx.In

	MaterialTestOrderRepository material_test_order_repository.MaterialTestOrderRepository
}

func NewMaterialTestOrderPolicy(params NewMaterialTestOrderPolicyParams) MaterialTestOrderPolicy {
	return &materialTestOrderPolicy{
		materialTestOrderRepository: params.MaterialTestOrderRepository,
	}
}

func (p *materialTestOrderPolicy) Store(ctx context.Context, _dto *dto.SaveMaterialTestOrder) error {
	return nil
}

func (p *materialTestOrderPolicy) Update(ctx context.Context, _dto *dto.SaveMaterialTestOrder) error {
	return nil
}

func (p *materialTestOrderPolicy) Destroy(ctx context.Context, _dto *dto.DestroyMaterialTestOrder) error {
	return nil
}
