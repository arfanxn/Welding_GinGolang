package step

import (
	"context"

	"github.com/arfanxn/welding/internal/infrastructure/id"
	mtoRepository "github.com/arfanxn/welding/internal/module/material_test_order/domain/repository"
	"github.com/arfanxn/welding/internal/module/material_test_order/usecase/dto"
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
)

type UpdateMaterialTestOrderStep interface {
	Handle(ctx context.Context, dto *dto.SaveMaterialTestOrder) (*entity.MaterialTestOrder, error)
}

type updateMaterialTestOrderStep struct {
	idService     id.IdService
	mtoRepository mtoRepository.MaterialTestOrderRepository
}

func NewUpdateMaterialTestOrderStep(
	idService id.IdService,
	mtoRepository mtoRepository.MaterialTestOrderRepository,
) UpdateMaterialTestOrderStep {
	return &updateMaterialTestOrderStep{
		idService:     idService,
		mtoRepository: mtoRepository,
	}
}

func (s *updateMaterialTestOrderStep) Handle(ctx context.Context, _dto *dto.SaveMaterialTestOrder) (mto *entity.MaterialTestOrder, err error) {
	// TODO: update this
	// q := query.NewQuery().FilterById(*_dto.Id).Include("Machine").Include("Method")
	// if mts, err = s.mtoRepository.First(q); err != nil {
	// 	return nil, err
	// }

	// if !goutil.IsEmptyReal(_dto.MachineId) {
	// 	mts.MachineId = *_dto.MachineId
	// }

	// if !goutil.IsEmptyReal(_dto.MethodId) {
	// 	mts.MethodId = *_dto.MethodId
	// }

	// if !goutil.IsEmptyReal(_dto.ServiceType) {
	// 	mts.ServiceType = *_dto.ServiceType
	// }

	// if !goutil.IsEmptyReal(_dto.ServiceCode) {
	// 	mts.ServiceCode = *_dto.ServiceCode
	// }

	// if !goutil.IsEmptyReal(_dto.Unit) {
	// 	mts.Unit = *_dto.Unit
	// }

	// if !goutil.IsEmptyReal(_dto.Price) {
	// 	mts.Price = *_dto.Price
	// }

	// if err := s.mtoRepository.Save(mts); err != nil {
	// 	return nil, err
	// }

	// // refetch
	// q = query.NewQuery().FilterById(*_dto.Id).Include("Machine").Include("Method")
	// if mts, err = s.mtoRepository.First(q); err != nil {
	// 	return nil, err
	// }

	return
}
