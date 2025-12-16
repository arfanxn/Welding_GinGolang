package policy

import (
	"context"

	mtoRepository "github.com/arfanxn/welding/internal/module/material_test_order/domain/repository"
	"github.com/arfanxn/welding/internal/module/material_test_order/usecase/dto"
	mtsRepository "github.com/arfanxn/welding/internal/module/material_test_service/domain/repository"
	mtwcRepository "github.com/arfanxn/welding/internal/module/material_test_work_category/domain/repository"
	userRepository "github.com/arfanxn/welding/internal/module/user/domain/repository"
	"github.com/gookit/goutil"
	"go.uber.org/fx"
)

type MaterialTestOrderPolicy interface {
	Store(ctx context.Context, _dto *dto.SaveMaterialTestOrder) error
	Update(ctx context.Context, _dto *dto.SaveMaterialTestOrder) error
	Destroy(ctx context.Context, _dto *dto.DestroyMaterialTestOrder) error

	StoreMedia(ctx context.Context, _dto *dto.SaveMaterialTestOrderMedia) error
	UpdateMedia(ctx context.Context, _dto *dto.SaveMaterialTestOrderMedia) error
	DestroyMedia(ctx context.Context, _dto *dto.SaveMaterialTestOrderMedia) error
}

type materialTestOrderPolicy struct {
	mtoRepository  mtoRepository.MaterialTestOrderRepository
	mtsRepository  mtsRepository.MaterialTestServiceRepository
	mtwcRepository mtwcRepository.MaterialTestWorkCategoryRepository
	userRepository userRepository.UserRepository
}

type NewMaterialTestOrderPolicyParams struct {
	fx.In

	MaterialTestOrderRepository        mtoRepository.MaterialTestOrderRepository
	MaterialTestServiceRepository      mtsRepository.MaterialTestServiceRepository
	MaterialTestWorkCategoryRepository mtwcRepository.MaterialTestWorkCategoryRepository
	UserRepository                     userRepository.UserRepository
}

func NewMaterialTestOrderPolicy(params NewMaterialTestOrderPolicyParams) MaterialTestOrderPolicy {
	return &materialTestOrderPolicy{
		mtoRepository:  params.MaterialTestOrderRepository,
		mtsRepository:  params.MaterialTestServiceRepository,
		mtwcRepository: params.MaterialTestWorkCategoryRepository,
		userRepository: params.UserRepository,
	}
}

func (p *materialTestOrderPolicy) Store(ctx context.Context, _dto *dto.SaveMaterialTestOrder) error {
	if err := p.validateWorkCategory(*_dto.WorkCategoryId); err != nil {
		return err
	}

	if _dto.OwnerUserIds != nil {
		if !goutil.IsEmptyReal(_dto.OwnerUserIds[0]) {
			if err := p.validateUsers(_dto.OwnerUserIds); err != nil {
				return err
			}
		}
	}

	if _dto.OrderedServices != nil {
		if !goutil.IsEmptyReal(_dto.OrderedServices[0]) {
			serviceIds := []string{}
			for _, orderedServiceDto := range _dto.OrderedServices {
				serviceIds = append(serviceIds, orderedServiceDto.ServiceId)
			}

			if err := p.validateServices(serviceIds); err != nil {
				return err
			}

		}
	}

	return nil
}

func (p *materialTestOrderPolicy) Update(ctx context.Context, _dto *dto.SaveMaterialTestOrder) error {
	if !goutil.IsEmptyReal(_dto.WorkCategoryId) {
		if err := p.validateWorkCategory(*_dto.WorkCategoryId); err != nil {
			return err
		}
	}

	if _dto.OwnerUserIds != nil {
		if !goutil.IsEmptyReal(_dto.OwnerUserIds[0]) {
			if err := p.validateUsers(_dto.OwnerUserIds); err != nil {
				return err
			}
		}
	}

	if _dto.OrderedServices != nil {
		if !goutil.IsEmptyReal(_dto.OrderedServices[0]) {
			serviceIds := []string{}
			for _, orderedServiceDto := range _dto.OrderedServices {
				serviceIds = append(serviceIds, orderedServiceDto.ServiceId)
			}

			if err := p.validateServices(serviceIds); err != nil {
				return err
			}

		}
	}

	return nil
}

func (p *materialTestOrderPolicy) Destroy(ctx context.Context, _dto *dto.DestroyMaterialTestOrder) error {
	return nil
}

func (p *materialTestOrderPolicy) StoreMedia(ctx context.Context, _dto *dto.SaveMaterialTestOrderMedia) error {
	return nil
}

func (p *materialTestOrderPolicy) UpdateMedia(ctx context.Context, _dto *dto.SaveMaterialTestOrderMedia) error {
	return nil
}

func (p *materialTestOrderPolicy) DestroyMedia(ctx context.Context, _dto *dto.SaveMaterialTestOrderMedia) error {
	return nil
}

func (p *materialTestOrderPolicy) validateServices(serviceIds []string) error {
	_, err := p.mtsRepository.FindByIds(serviceIds, nil)
	if err != nil {
		return err
	}
	return nil
}

func (p *materialTestOrderPolicy) validateWorkCategory(WorkCategoryId string) error {
	_, err := p.mtwcRepository.Find(WorkCategoryId, nil)
	if err != nil {
		return err
	}
	return nil
}

func (p *materialTestOrderPolicy) validateUsers(userIds []string) error {
	_, err := p.userRepository.FindByIds(userIds, nil)
	if err != nil {
		return err
	}
	return nil
}
