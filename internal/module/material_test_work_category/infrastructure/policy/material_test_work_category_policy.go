package policy

import (
	"context"

	materialTestWorkCategoryRepository "github.com/arfanxn/welding/internal/module/material_test_work_category/domain/repository"
	"github.com/arfanxn/welding/internal/module/material_test_work_category/usecase/dto"
	"go.uber.org/fx"
)

type MaterialTestWorkCategoryPolicy interface {
	Store(ctx context.Context, _dto *dto.SaveMaterialTestWorkCategory) error
	Update(ctx context.Context, _dto *dto.SaveMaterialTestWorkCategory) error
	Destroy(ctx context.Context, _dto *dto.DestroyMaterialTestWorkCategory) error
}

type materialTestWorkCategoryPolicy struct {
	materialTestWorkCategoryRepository materialTestWorkCategoryRepository.MaterialTestWorkCategoryRepository
}

type NewMaterialTestWorkCategoryPolicyParams struct {
	fx.In

	MaterialTestWorkCategoryRepository materialTestWorkCategoryRepository.MaterialTestWorkCategoryRepository
}

func NewMaterialTestWorkCategoryPolicy(params NewMaterialTestWorkCategoryPolicyParams) MaterialTestWorkCategoryPolicy {
	return &materialTestWorkCategoryPolicy{
		materialTestWorkCategoryRepository: params.MaterialTestWorkCategoryRepository,
	}
}

func (p *materialTestWorkCategoryPolicy) Store(ctx context.Context, _dto *dto.SaveMaterialTestWorkCategory) error {
	return nil
}

func (p *materialTestWorkCategoryPolicy) Update(ctx context.Context, _dto *dto.SaveMaterialTestWorkCategory) error {
	return nil
}

func (p *materialTestWorkCategoryPolicy) Destroy(ctx context.Context, _dto *dto.DestroyMaterialTestWorkCategory) error {
	return nil
}
