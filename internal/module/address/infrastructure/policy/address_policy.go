package policy

import (
	"context"

	addressRepository "github.com/arfanxn/welding/internal/module/address/domain/repository"
	"github.com/arfanxn/welding/internal/module/address/usecase/dto"
	"go.uber.org/fx"
)

type AddressPolicy interface {
	Store(ctx context.Context, _dto *dto.SaveAddress) error
	Update(ctx context.Context, _dto *dto.SaveAddress) error
	Destroy(ctx context.Context, _dto *dto.DestroyAddress) error
}

type addressPolicy struct {
	addressRepository addressRepository.AddressRepository
}

type NewAddressPolicyParams struct {
	fx.In

	AddressRepository addressRepository.AddressRepository
}

func NewAddressPolicy(params NewAddressPolicyParams) AddressPolicy {
	return &addressPolicy{
		addressRepository: params.AddressRepository,
	}
}

func (p *addressPolicy) Store(ctx context.Context, _dto *dto.SaveAddress) error {
	return nil
}

func (p *addressPolicy) Update(ctx context.Context, _dto *dto.SaveAddress) error {
	return nil
}

func (p *addressPolicy) Destroy(ctx context.Context, _dto *dto.DestroyAddress) error {
	return nil
}
