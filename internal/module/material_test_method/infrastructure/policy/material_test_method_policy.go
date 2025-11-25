// Package policy provides business rule validation and authorization logic for role management operations.
// It enforces constraints and permissions before allowing role-related actions to be executed.
package policy

import (
	"context"

	mtmRepository "github.com/arfanxn/welding/internal/module/material_test_method/domain/repository"
	"github.com/arfanxn/welding/internal/module/material_test_method/usecase/dto"
	"go.uber.org/fx"
)

type MaterialTestMethodPolicy interface {
	Store(ctx context.Context, _dto *dto.SaveMaterialTestMethod) error
	Update(ctx context.Context, _dto *dto.SaveMaterialTestMethod) error
	Destroy(ctx context.Context, _dto *dto.DestroyMaterialTestMethod) error
}

type materialTestMethodPolicy struct {
	mtmRepository mtmRepository.MaterialTestMethodRepository
}

type NewMaterialTestMethodPolicyParams struct {
	fx.In

	MtmRepository mtmRepository.MaterialTestMethodRepository
}

func NewMaterialTestMethodPolicy(params NewMaterialTestMethodPolicyParams) MaterialTestMethodPolicy {
	return &materialTestMethodPolicy{
		mtmRepository: params.MtmRepository,
	}
}

func (p *materialTestMethodPolicy) Store(ctx context.Context, _dto *dto.SaveMaterialTestMethod) error {
	return nil
}

func (p *materialTestMethodPolicy) Update(ctx context.Context, _dto *dto.SaveMaterialTestMethod) error {
	return nil
}

func (p *materialTestMethodPolicy) Destroy(ctx context.Context, _dto *dto.DestroyMaterialTestMethod) error {
	return nil
}
