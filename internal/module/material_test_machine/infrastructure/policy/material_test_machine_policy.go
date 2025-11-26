// Package policy provides business rule validation and authorization logic for role management operations.
// It enforces constraints and permissions before allowing role-related actions to be executed.
package policy

import (
	"context"

	mtmRepository "github.com/arfanxn/welding/internal/module/material_test_machine/domain/repository"
	"github.com/arfanxn/welding/internal/module/material_test_machine/usecase/dto"
	mtsRepository "github.com/arfanxn/welding/internal/module/material_test_service/domain/repository"
	"github.com/arfanxn/welding/internal/module/shared/domain/errorx"
	"go.uber.org/fx"
)

type MaterialTestMachinePolicy interface {
	Store(ctx context.Context, _dto *dto.SaveMaterialTestMachine) error
	Update(ctx context.Context, _dto *dto.SaveMaterialTestMachine) error
	Destroy(ctx context.Context, _dto *dto.DestroyMaterialTestMachine) error
}

type materialTestMachinePolicy struct {
	mtmRepository mtmRepository.MaterialTestMachineRepository
	mtsRepository mtsRepository.MaterialTestServiceRepository
}

type NewMaterialTestMachinePolicyParams struct {
	fx.In

	MaterialTestMachineRepository mtmRepository.MaterialTestMachineRepository
	MaterialTestServiceRepository mtsRepository.MaterialTestServiceRepository
}

func NewMaterialTestMachinePolicy(params NewMaterialTestMachinePolicyParams) MaterialTestMachinePolicy {
	return &materialTestMachinePolicy{
		mtmRepository: params.MaterialTestMachineRepository,
		mtsRepository: params.MaterialTestServiceRepository,
	}
}

func (p *materialTestMachinePolicy) Store(ctx context.Context, _dto *dto.SaveMaterialTestMachine) error {
	return nil
}

func (p *materialTestMachinePolicy) Update(ctx context.Context, _dto *dto.SaveMaterialTestMachine) error {
	return nil
}

func (p *materialTestMachinePolicy) Destroy(ctx context.Context, _dto *dto.DestroyMaterialTestMachine) error {
	count, err := p.mtsRepository.CountByMachineId(_dto.Id)
	if err != nil {
		return err
	}

	if count > 0 {
		return errorx.ErrMaterialTestMachineInUseDestroyForbidden
	}

	return nil
}
