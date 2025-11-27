package policy

import (
	"context"

	mtMachineRepository "github.com/arfanxn/welding/internal/module/material_test_machine/domain/repository"
	mtMethodRepository "github.com/arfanxn/welding/internal/module/material_test_method/domain/repository"
	mtServiceRepository "github.com/arfanxn/welding/internal/module/material_test_service/domain/repository"
	"github.com/arfanxn/welding/internal/module/material_test_service/usecase/dto"
	"github.com/gookit/goutil"
	"go.uber.org/fx"
)

type MaterialTestServicePolicy interface {
	Store(ctx context.Context, _dto *dto.SaveMaterialTestService) error
	Update(ctx context.Context, _dto *dto.SaveMaterialTestService) error
	Destroy(ctx context.Context, _dto *dto.DestroyMaterialTestService) error
}

type materialTestServicePolicy struct {
	mtMethodRepository  mtMethodRepository.MaterialTestMethodRepository
	mtMachineRepository mtMachineRepository.MaterialTestMachineRepository
	mtServiceRepository mtServiceRepository.MaterialTestServiceRepository
}

type NewMaterialTestServicePolicyParams struct {
	fx.In

	MaterialTestMethodRepository  mtMethodRepository.MaterialTestMethodRepository
	MaterialTestMachineRepository mtMachineRepository.MaterialTestMachineRepository
	MaterialTestServiceRepository mtServiceRepository.MaterialTestServiceRepository
}

func NewMaterialTestServicePolicy(params NewMaterialTestServicePolicyParams) MaterialTestServicePolicy {
	return &materialTestServicePolicy{
		mtMethodRepository:  params.MaterialTestMethodRepository,
		mtMachineRepository: params.MaterialTestMachineRepository,
		mtServiceRepository: params.MaterialTestServiceRepository,
	}
}

func (p *materialTestServicePolicy) Store(ctx context.Context, _dto *dto.SaveMaterialTestService) error {
	if err := p.validateMachine(*_dto.MachineId); err != nil {
		return err
	}
	if err := p.validateMethod(*_dto.MethodId); err != nil {
		return err
	}
	return nil
}

func (p *materialTestServicePolicy) Update(ctx context.Context, _dto *dto.SaveMaterialTestService) error {
	if !goutil.IsEmptyReal(_dto.MachineId) {
		if err := p.validateMachine(*_dto.MachineId); err != nil {
			return err
		}
	}

	if !goutil.IsEmptyReal(_dto.MethodId) {
		if err := p.validateMethod(*_dto.MethodId); err != nil {
			return err
		}
	}

	return nil
}

func (p *materialTestServicePolicy) Destroy(ctx context.Context, _dto *dto.DestroyMaterialTestService) error {
	return nil
}

// ==================================================
// Private helper methods
// ==================================================

func (p *materialTestServicePolicy) validateMethod(methodId string) error {
	_, err := p.mtMethodRepository.Find(methodId)
	if err != nil {
		return err
	}
	return nil
}

func (p *materialTestServicePolicy) validateMachine(machineId string) error {
	_, err := p.mtMachineRepository.Find(machineId)
	if err != nil {
		return err
	}
	return nil
}
