package usecase

import (
	"context"

	"github.com/arfanxn/welding/internal/infrastructure/id"
	activityEnum "github.com/arfanxn/welding/internal/module/activity/domain/enum"
	activityDto "github.com/arfanxn/welding/internal/module/activity/usecase/dto"
	activityService "github.com/arfanxn/welding/internal/module/activity/usecase/service"
	customerRepository "github.com/arfanxn/welding/internal/module/customer/domain/repository"
	customerPolicy "github.com/arfanxn/welding/internal/module/customer/infrastructure/policy"
	"github.com/arfanxn/welding/internal/module/customer/usecase/dto"
	"github.com/arfanxn/welding/internal/module/shared/contextkey"
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	"github.com/arfanxn/welding/pkg/pagination"
	"github.com/arfanxn/welding/pkg/query"
	"github.com/arfanxn/welding/pkg/typeutil"
	"github.com/gookit/goutil"
	"go.uber.org/fx"
)

type CustomerUsecase interface {
	Paginate(ctx context.Context, q *query.Query) (*pagination.OffsetPagination[*entity.Customer], error)
	Show(ctx context.Context, q *query.Query) (*entity.Customer, error)
	Store(ctx context.Context, _dto *dto.SaveCustomer) (*entity.Customer, error)
	Update(ctx context.Context, _dto *dto.SaveCustomer) (*entity.Customer, error)
	Destroy(ctx context.Context, _dto *dto.DestroyCustomer) error
}

type customerUsecase struct {
	activityService    activityService.ActivityService
	idService          id.IdService
	customerPolicy     customerPolicy.CustomerPolicy
	customerRepository customerRepository.CustomerRepository
}

type NewCustomerUsecaseParams struct {
	fx.In

	ActivityService    activityService.ActivityService
	IdService          id.IdService
	CustomerRepository customerRepository.CustomerRepository
	CustomerPolicy     customerPolicy.CustomerPolicy
}

func NewCustomerUsecase(params NewCustomerUsecaseParams) CustomerUsecase {
	return &customerUsecase{
		activityService:    params.ActivityService,
		idService:          params.IdService,
		customerRepository: params.CustomerRepository,
		customerPolicy:     params.CustomerPolicy,
	}
}

func (u *customerUsecase) Paginate(ctx context.Context, q *query.Query) (*pagination.OffsetPagination[*entity.Customer], error) {
	op, err := u.customerRepository.Paginate(q)
	if err != nil {
		return nil, err
	}

	if _, err = u.activityService.Record(ctx, &activityDto.CreateActivity{
		Causer:      contextkey.GetUser(ctx),
		Action:      activityEnum.CustomersIndex,
		SubjectType: typeutil.Ptr(activityEnum.CustomerSubjectType),
	}); err != nil {
		return nil, err
	}

	return op, err

}

func (u *customerUsecase) Show(ctx context.Context, q *query.Query) (customer *entity.Customer, err error) {
	if customer, err = u.customerRepository.First(q); err != nil {
		return nil, err
	}

	if _, err = u.activityService.Record(ctx, &activityDto.CreateActivity{
		Causer:  contextkey.GetUser(ctx),
		Action:  activityEnum.CustomersShow,
		Subject: customer,
	}); err != nil {
		return nil, err
	}

	return customer, nil
}

func (u *customerUsecase) Store(ctx context.Context, _dto *dto.SaveCustomer) (customer *entity.Customer, err error) {
	if err = u.customerPolicy.Store(ctx, _dto); err != nil {
		return nil, err
	}

	customerId := u.idService.Generate()
	customer = &entity.Customer{}
	customer.Id = customerId
	customer.AddressId = *_dto.AddressId
	customer.Name = *_dto.Name
	customer.PhoneNumber = *_dto.PhoneNumber
	customer.Email = *_dto.Email

	if err = u.customerRepository.Save(customer); err != nil {
		return nil, err
	}

	if _, err = u.activityService.Record(ctx, &activityDto.CreateActivity{
		Causer:  contextkey.GetUser(ctx),
		Action:  activityEnum.CustomersStore,
		Subject: customer,
	}); err != nil {
		return nil, err
	}

	q := query.NewQuery().FilterById(customer.Id).Include("Address")
	if customer, err = u.customerRepository.First(q); err != nil {
		return nil, err
	}

	return
}

func (u *customerUsecase) Update(ctx context.Context, _dto *dto.SaveCustomer) (customer *entity.Customer, err error) {
	if err = u.customerPolicy.Update(ctx, _dto); err != nil {
		return nil, err
	}

	q := query.NewQuery().FilterById(*_dto.Id).Include("Address")

	if customer, err = u.customerRepository.First(q); err != nil {
		return nil, err
	}

	if !goutil.IsEmptyReal(_dto.AddressId) {
		customer.AddressId = *_dto.AddressId
	}

	if !goutil.IsEmptyReal(_dto.Name) {
		customer.Name = *_dto.Name
	}

	if !goutil.IsEmptyReal(_dto.PhoneNumber) {
		customer.PhoneNumber = *_dto.PhoneNumber
	}

	if !goutil.IsEmptyReal(_dto.Email) {
		customer.Email = *_dto.Email
	}

	if err := u.customerRepository.Save(customer); err != nil {
		return nil, err
	}

	if customer, err = u.customerRepository.First(q); err != nil {
		return nil, err
	}

	if _, err = u.activityService.Record(ctx, &activityDto.CreateActivity{
		Causer:  contextkey.GetUser(ctx),
		Action:  activityEnum.CustomersUpdate,
		Subject: customer,
	}); err != nil {
		return nil, err
	}

	if customer, err = u.customerRepository.First(q); err != nil {
		return nil, err
	}

	return
}

func (u *customerUsecase) Destroy(ctx context.Context, _dto *dto.DestroyCustomer) (err error) {
	if err = u.customerPolicy.Destroy(ctx, _dto); err != nil {
		return err
	}

	customer, err := u.customerRepository.Find(_dto.Id, nil)
	if err != nil {
		return err
	}

	err = u.customerRepository.Destroy(customer)
	if err != nil {
		return err
	}

	if _, err = u.activityService.Record(ctx, &activityDto.CreateActivity{
		Causer:  contextkey.GetUser(ctx),
		Action:  activityEnum.CustomersDestroy,
		Subject: customer,
	}); err != nil {
		return err
	}

	return nil
}
