package policy

import (
	"context"

	materialTestWorkPackageRepository "github.com/arfanxn/welding/internal/module/material_test_work_package/domain/repository"
	"github.com/arfanxn/welding/internal/module/material_test_work_package/usecase/dto"
	"go.uber.org/fx"
)

type MaterialTestWorkPackagePolicy interface {
	Store(ctx context.Context, _dto *dto.SaveMaterialTestWorkPackage) error
	Update(ctx context.Context, _dto *dto.SaveMaterialTestWorkPackage) error
	Destroy(ctx context.Context, _dto *dto.DestroyMaterialTestWorkPackage) error
}

type materialTestWorkPackagePolicy struct {
	materialTestWorkPackageRepository materialTestWorkPackageRepository.MaterialTestWorkPackageRepository
}

type NewMaterialTestWorkPackagePolicyParams struct {
	fx.In

	MaterialTestWorkPackageRepository materialTestWorkPackageRepository.MaterialTestWorkPackageRepository
}

func NewMaterialTestWorkPackagePolicy(params NewMaterialTestWorkPackagePolicyParams) MaterialTestWorkPackagePolicy {
	return &materialTestWorkPackagePolicy{
		materialTestWorkPackageRepository: params.MaterialTestWorkPackageRepository,
	}
}

func (p *materialTestWorkPackagePolicy) Store(ctx context.Context, _dto *dto.SaveMaterialTestWorkPackage) error {
	return nil
}

func (p *materialTestWorkPackagePolicy) Update(ctx context.Context, _dto *dto.SaveMaterialTestWorkPackage) error {
	return nil
}

func (p *materialTestWorkPackagePolicy) Destroy(ctx context.Context, _dto *dto.DestroyMaterialTestWorkPackage) error {
	return nil
}
