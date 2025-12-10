package usecase

import (
	"context"

	"github.com/arfanxn/welding/internal/infrastructure/id"
	activityEnum "github.com/arfanxn/welding/internal/module/activity/domain/enum"
	activityDto "github.com/arfanxn/welding/internal/module/activity/usecase/dto"
	activityService "github.com/arfanxn/welding/internal/module/activity/usecase/service"
	addressRepository "github.com/arfanxn/welding/internal/module/address/domain/repository"
	addressPolicy "github.com/arfanxn/welding/internal/module/address/infrastructure/policy"
	"github.com/arfanxn/welding/internal/module/address/usecase/dto"
	"github.com/arfanxn/welding/internal/module/shared/contextkey"
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	"github.com/arfanxn/welding/pkg/pagination"
	"github.com/arfanxn/welding/pkg/query"
	"github.com/arfanxn/welding/pkg/typeutil"
	"github.com/gookit/goutil"
	"go.uber.org/fx"
)

type AddressUsecase interface {
	Paginate(ctx context.Context, q *query.Query) (*pagination.OffsetPagination[*entity.Address], error)
	Show(ctx context.Context, q *query.Query) (*entity.Address, error)
	Store(ctx context.Context, _dto *dto.SaveAddress) (*entity.Address, error)
	Update(ctx context.Context, _dto *dto.SaveAddress) (*entity.Address, error)
	Destroy(ctx context.Context, _dto *dto.DestroyAddress) error
}

type addressUsecase struct {
	activityService   activityService.ActivityService
	idService         id.IdService
	addressRepository addressRepository.AddressRepository
	addressPolicy     addressPolicy.AddressPolicy
}

type NewAddressUsecaseParams struct {
	fx.In

	ActivityService   activityService.ActivityService
	IdService         id.IdService
	AddressRepository addressRepository.AddressRepository
	AddressPolicy     addressPolicy.AddressPolicy
}

func NewAddressUsecase(params NewAddressUsecaseParams) AddressUsecase {
	return &addressUsecase{
		activityService:   params.ActivityService,
		idService:         params.IdService,
		addressRepository: params.AddressRepository,
		addressPolicy:     params.AddressPolicy,
	}
}

func (u *addressUsecase) Paginate(ctx context.Context, q *query.Query) (*pagination.OffsetPagination[*entity.Address], error) {
	op, err := u.addressRepository.Paginate(q)
	if err != nil {
		return nil, err
	}

	if _, err = u.activityService.Record(ctx, &activityDto.CreateActivity{
		Causer:      contextkey.GetUser(ctx),
		Action:      activityEnum.AddressesIndex,
		SubjectType: typeutil.Ptr(activityEnum.AddressSubjectType),
	}); err != nil {
		return nil, err
	}

	return op, err

}

func (u *addressUsecase) Show(ctx context.Context, q *query.Query) (address *entity.Address, err error) {
	if address, err = u.addressRepository.First(q); err != nil {
		return nil, err
	}

	if _, err = u.activityService.Record(ctx, &activityDto.CreateActivity{
		Causer:  contextkey.GetUser(ctx),
		Action:  activityEnum.AddressesShow,
		Subject: address,
	}); err != nil {
		return nil, err
	}

	return address, nil
}

func (u *addressUsecase) Store(ctx context.Context, _dto *dto.SaveAddress) (address *entity.Address, err error) {
	if err = u.addressPolicy.Store(ctx, _dto); err != nil {
		return nil, err
	}

	address = &entity.Address{}
	address.Id = u.idService.Generate()
	address.FullAddress = *_dto.FullAddress

	if err = u.addressRepository.Save(address); err != nil {
		return nil, err
	}

	if _, err = u.activityService.Record(ctx, &activityDto.CreateActivity{
		Causer:  contextkey.GetUser(ctx),
		Action:  activityEnum.AddressesStore,
		Subject: address,
	}); err != nil {
		return nil, err
	}

	return
}

func (u *addressUsecase) Update(ctx context.Context, _dto *dto.SaveAddress) (address *entity.Address, err error) {
	if err = u.addressPolicy.Update(ctx, _dto); err != nil {
		return nil, err
	}

	if address, err = u.addressRepository.Find(*_dto.Id, nil); err != nil {
		return nil, err
	}

	if !goutil.IsEmptyReal(_dto.FullAddress) {
		address.FullAddress = *_dto.FullAddress
	}

	if err := u.addressRepository.Save(address); err != nil {
		return nil, err
	}

	if _, err = u.activityService.Record(ctx, &activityDto.CreateActivity{
		Causer:  contextkey.GetUser(ctx),
		Action:  activityEnum.AddressesUpdate,
		Subject: address,
	}); err != nil {
		return nil, err
	}

	return
}

func (u *addressUsecase) Destroy(ctx context.Context, _dto *dto.DestroyAddress) (err error) {
	if err = u.addressPolicy.Destroy(ctx, _dto); err != nil {
		return err
	}

	address, err := u.addressRepository.Find(_dto.Id, nil)
	if err != nil {
		return err
	}

	err = u.addressRepository.Destroy(address)
	if err != nil {
		return err
	}

	if _, err = u.activityService.Record(ctx, &activityDto.CreateActivity{
		Causer:  contextkey.GetUser(ctx),
		Action:  activityEnum.AddressesDestroy,
		Subject: address,
	}); err != nil {
		return err
	}

	return nil
}
