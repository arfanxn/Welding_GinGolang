package policy

import (
	"context"

	mtmRepository "github.com/arfanxn/welding/internal/module/material_test_method/domain/repository"
	"github.com/arfanxn/welding/internal/module/material_test_method/usecase/dto"
	mtsRepository "github.com/arfanxn/welding/internal/module/material_test_service/domain/repository"
	"github.com/arfanxn/welding/internal/module/shared/domain/errorx"
	"go.uber.org/fx"
)

type MaterialTestMethodPolicy interface {
	Store(ctx context.Context, _dto *dto.SaveMaterialTestMethod) error
	Update(ctx context.Context, _dto *dto.SaveMaterialTestMethod) error
	Destroy(ctx context.Context, _dto *dto.DestroyMaterialTestMethod) error
}

type materialTestMethodPolicy struct {
	mtmRepository mtmRepository.MaterialTestMethodRepository
	mtsRepository mtsRepository.MaterialTestServiceRepository
}

type NewMaterialTestMethodPolicyParams struct {
	fx.In

	MaterialTestMethodRepository  mtmRepository.MaterialTestMethodRepository
	MaterialTestServiceRepository mtsRepository.MaterialTestServiceRepository
}

func NewMaterialTestMethodPolicy(params NewMaterialTestMethodPolicyParams) MaterialTestMethodPolicy {
	return &materialTestMethodPolicy{
		mtmRepository: params.MaterialTestMethodRepository,
		mtsRepository: params.MaterialTestServiceRepository,
	}
}

func (p *materialTestMethodPolicy) Store(ctx context.Context, _dto *dto.SaveMaterialTestMethod) error {
	return nil
}

func (p *materialTestMethodPolicy) Update(ctx context.Context, _dto *dto.SaveMaterialTestMethod) error {
	return nil
}

func (p *materialTestMethodPolicy) Destroy(ctx context.Context, _dto *dto.DestroyMaterialTestMethod) error {
	count, err := p.mtsRepository.CountByMethodId(_dto.Id)
	if err != nil {
		return err
	}

	if count > 0 {
		return errorx.ErrMaterialTestMethodInUseDestroyForbidden
	}

	return nil
}
