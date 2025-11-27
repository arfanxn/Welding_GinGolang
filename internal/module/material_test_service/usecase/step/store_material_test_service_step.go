package step

import (
	"context"

	"github.com/arfanxn/welding/internal/infrastructure/id"
	mtsRepository "github.com/arfanxn/welding/internal/module/material_test_service/domain/repository"
	"github.com/arfanxn/welding/internal/module/material_test_service/usecase/dto"
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	"github.com/arfanxn/welding/pkg/query"
)

type StoreMaterialTestServiceStep interface {
	Handle(ctx context.Context, dto *dto.SaveMaterialTestService) (*entity.MaterialTestService, error)
}

type storeMaterialTestServiceStep struct {
	idService     id.IdService
	mtsRepository mtsRepository.MaterialTestServiceRepository
}

func NewStoreMaterialTestServiceStep(
	idService id.IdService,
	mtsRepository mtsRepository.MaterialTestServiceRepository,
) StoreMaterialTestServiceStep {
	return &storeMaterialTestServiceStep{
		idService:     idService,
		mtsRepository: mtsRepository,
	}
}

func (s *storeMaterialTestServiceStep) Handle(ctx context.Context, _dto *dto.SaveMaterialTestService) (mts *entity.MaterialTestService, err error) {
	mts = &entity.MaterialTestService{}
	mts.Id = s.idService.Generate()
	mts.MachineId = *_dto.MachineId
	mts.MethodId = *_dto.MethodId
	mts.ServiceType = *_dto.ServiceType
	mts.ServiceCode = *_dto.ServiceCode
	mts.Unit = *_dto.Unit
	mts.Price = *_dto.Price

	if err := s.mtsRepository.Save(mts); err != nil {
		return nil, err
	}

	q := query.NewQuery().FilterById(mts.Id).Include("Machine").Include("Method")
	if mts, err = s.mtsRepository.First(q); err != nil {
		return nil, err
	}

	return
}
