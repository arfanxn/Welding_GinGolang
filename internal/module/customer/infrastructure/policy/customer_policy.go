package policy

import (
	"context"

	addressRepository "github.com/arfanxn/welding/internal/module/address/domain/repository"
	customerRepository "github.com/arfanxn/welding/internal/module/customer/domain/repository"
	"github.com/arfanxn/welding/internal/module/customer/usecase/dto"
	"github.com/gookit/goutil"
	"go.uber.org/fx"
)

type CustomerPolicy interface {
	Store(ctx context.Context, _dto *dto.SaveCustomer) error
	Update(ctx context.Context, _dto *dto.SaveCustomer) error
	Destroy(ctx context.Context, _dto *dto.DestroyCustomer) error
}

type customerPolicy struct {
	customerRepository customerRepository.CustomerRepository
	addressRepository  addressRepository.AddressRepository
}

type NewCustomerPolicyParams struct {
	fx.In

	CustomerRepository customerRepository.CustomerRepository
	AddressRepository  addressRepository.AddressRepository
}

func NewCustomerPolicy(params NewCustomerPolicyParams) CustomerPolicy {
	return &customerPolicy{
		customerRepository: params.CustomerRepository,
		addressRepository:  params.AddressRepository,
	}
}

func (p *customerPolicy) Store(ctx context.Context, _dto *dto.SaveCustomer) error {
	if err := p.validateAddressAssignments(_dto.AddressId); err != nil {
		return err
	}

	return nil
}

func (p *customerPolicy) Update(ctx context.Context, _dto *dto.SaveCustomer) error {
	if err := p.validateAddressAssignments(_dto.AddressId); err != nil {
		return err
	}

	return nil
}

func (p *customerPolicy) Destroy(ctx context.Context, _dto *dto.DestroyCustomer) error {
	return nil
}

func (p *customerPolicy) validateAddressAssignments(addressId *string) error {
	if goutil.IsEmptyReal(addressId) {
		return nil
	}

	_, err := p.addressRepository.Find(*addressId, nil)
	if err != nil {
		return err
	}

	return nil
}
