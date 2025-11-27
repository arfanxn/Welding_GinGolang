package step

import (
	"context"

	"github.com/arfanxn/welding/internal/infrastructure/id"
	mtsRepository "github.com/arfanxn/welding/internal/module/material_test_service/domain/repository"
	"github.com/arfanxn/welding/internal/module/material_test_service/usecase/dto"
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	"github.com/arfanxn/welding/pkg/query"
	"github.com/gookit/goutil"
)

type UpdateMaterialTestServiceStep interface {
	Handle(ctx context.Context, dto *dto.SaveMaterialTestService) (*entity.MaterialTestService, error)
}

type updateMaterialTestServiceStep struct {
	idService     id.IdService
	mtsRepository mtsRepository.MaterialTestServiceRepository
}

func NewUpdateMaterialTestServiceStep(
	idService id.IdService,
	mtsRepository mtsRepository.MaterialTestServiceRepository,
) UpdateMaterialTestServiceStep {
	return &updateMaterialTestServiceStep{
		idService:     idService,
		mtsRepository: mtsRepository,
	}
}

func (s *updateMaterialTestServiceStep) Handle(ctx context.Context, _dto *dto.SaveMaterialTestService) (mts *entity.MaterialTestService, err error) {
	q := query.NewQuery().FilterById(*_dto.Id).Include("Machine").Include("Method")
	if mts, err = s.mtsRepository.First(q); err != nil {
		return nil, err
	}

	if !goutil.IsEmptyReal(_dto.MachineId) {
		mts.MachineId = *_dto.MachineId
	}

	if !goutil.IsEmptyReal(_dto.MethodId) {
		mts.MethodId = *_dto.MethodId
	}

	if !goutil.IsEmptyReal(_dto.ServiceType) {
		mts.ServiceType = *_dto.ServiceType
	}

	if !goutil.IsEmptyReal(_dto.ServiceCode) {
		mts.ServiceCode = *_dto.ServiceCode
	}

	if !goutil.IsEmptyReal(_dto.Unit) {
		mts.Unit = *_dto.Unit
	}

	if !goutil.IsEmptyReal(_dto.Price) {
		mts.Price = *_dto.Price
	}

	if err := s.mtsRepository.Save(mts); err != nil {
		return nil, err
	}

	// refetch
	q = query.NewQuery().FilterById(*_dto.Id).Include("Machine").Include("Method")
	if mts, err = s.mtsRepository.First(q); err != nil {
		return nil, err
	}

	return
}
