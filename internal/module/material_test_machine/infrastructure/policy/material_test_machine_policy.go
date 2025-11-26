// Package policy provides business rule validation and authorization logic for role management operations.
// It enforces constraints and permissions before allowing role-related actions to be executed.
package policy

import (
	"context"

	mtmRepository "github.com/arfanxn/welding/internal/module/material_test_machine/domain/repository"
	"github.com/arfanxn/welding/internal/module/material_test_machine/usecase/dto"
	"go.uber.org/fx"
)

type MaterialTestMachinePolicy interface {
	Store(ctx context.Context, _dto *dto.SaveMaterialTestMachine) error
	Update(ctx context.Context, _dto *dto.SaveMaterialTestMachine) error
	Destroy(ctx context.Context, _dto *dto.DestroyMaterialTestMachine) error
}

type materialTestMachinePolicy struct {
	mtmRepository mtmRepository.MaterialTestMachineRepository
}

type NewMaterialTestMachinePolicyParams struct {
	fx.In

	MtmRepository mtmRepository.MaterialTestMachineRepository
}

func NewMaterialTestMachinePolicy(params NewMaterialTestMachinePolicyParams) MaterialTestMachinePolicy {
	return &materialTestMachinePolicy{
		mtmRepository: params.MtmRepository,
	}
}

func (p *materialTestMachinePolicy) Store(ctx context.Context, _dto *dto.SaveMaterialTestMachine) error {
	return nil
}

func (p *materialTestMachinePolicy) Update(ctx context.Context, _dto *dto.SaveMaterialTestMachine) error {
	return nil
}

func (p *materialTestMachinePolicy) Destroy(ctx context.Context, _dto *dto.DestroyMaterialTestMachine) error {
	return nil
}
