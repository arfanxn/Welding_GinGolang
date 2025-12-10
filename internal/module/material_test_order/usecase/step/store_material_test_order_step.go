package step

import (
	"context"

	"github.com/arfanxn/welding/internal/infrastructure/id"
	addressRepository "github.com/arfanxn/welding/internal/module/address/domain/repository"
	customerRepository "github.com/arfanxn/welding/internal/module/customer/domain/repository"
	mtoRepository "github.com/arfanxn/welding/internal/module/material_test_order/domain/repository"
	"github.com/arfanxn/welding/internal/module/material_test_order/usecase/dto"
	mtwpRepository "github.com/arfanxn/welding/internal/module/material_test_work_package/domain/repository"
	mediaRepository "github.com/arfanxn/welding/internal/module/media/domain/repository"
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	"github.com/arfanxn/welding/pkg/query"
	"go.uber.org/fx"
)

type StoreMaterialTestOrderStep interface {
	Handle(ctx context.Context, dto *dto.SaveMaterialTestOrder) (*entity.MaterialTestOrder, error)
}

type storeMaterialTestServiceStep struct {
	idService          id.IdService
	mtoRepository      mtoRepository.MaterialTestOrderRepository
	mtwpRepository     mtwpRepository.MaterialTestWorkPackageRepository
	mediaRepository    mediaRepository.MediaRepository
	addressRepository  addressRepository.AddressRepository
	customerRepository customerRepository.CustomerRepository
}

type NewStoreMaterialTestOrderStepParams struct {
	fx.In

	IdService                         id.IdService
	MaterialTestOrderRepository       mtoRepository.MaterialTestOrderRepository
	MaterialTestWorkPackageRepository mtwpRepository.MaterialTestWorkPackageRepository
	MediaRepository                   mediaRepository.MediaRepository
	AddressRepository                 addressRepository.AddressRepository
	CustomerRepository                customerRepository.CustomerRepository
}

func NewStoreMaterialTestOrderStep(params NewStoreMaterialTestOrderStepParams) StoreMaterialTestOrderStep {
	return &storeMaterialTestServiceStep{
		idService:          params.IdService,
		mtoRepository:      params.MaterialTestOrderRepository,
		mtwpRepository:     params.MaterialTestWorkPackageRepository,
		mediaRepository:    params.MediaRepository,
		addressRepository:  params.AddressRepository,
		customerRepository: params.CustomerRepository,
	}
}

func (s *storeMaterialTestServiceStep) Handle(ctx context.Context, _dto *dto.SaveMaterialTestOrder) (
	mto *entity.MaterialTestOrder, err error,
) {
	mto = &entity.MaterialTestOrder{}
	mto.Id = s.idService.Generate()

	if err := s.mtoRepository.Save(mto); err != nil {
		return nil, err
	}

	q := query.NewQuery().FilterById(mto.Id)
	// TODO: add required relations
	if mto, err = s.mtoRepository.First(q); err != nil {
		return nil, err
	}

	return
}
