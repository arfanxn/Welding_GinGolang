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

	MaterialTestMethodRepository mtmRepository.MaterialTestMethodRepository
}

func NewMaterialTestMethodPolicy(params NewMaterialTestMethodPolicyParams) MaterialTestMethodPolicy {
	return &materialTestMethodPolicy{
		mtmRepository: params.MaterialTestMethodRepository,
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
