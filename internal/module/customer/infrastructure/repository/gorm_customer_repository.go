package repository

import (
	"errors"

	"github.com/arfanxn/welding/internal/infrastructure/database/helper"
	"github.com/arfanxn/welding/internal/module/customer/domain/repository"
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	"github.com/arfanxn/welding/internal/module/shared/domain/errorx"
	"github.com/arfanxn/welding/pkg/pagination"
	"github.com/arfanxn/welding/pkg/query"
	"gorm.io/gorm"
)

type GormCustomerRepository struct {
	db *gorm.DB
}

func NewGormCustomerRepository(db *gorm.DB) repository.CustomerRepository {
	return &GormCustomerRepository{
		db: db,
	}
}

func (r *GormCustomerRepository) query(db *gorm.DB, q *query.Query) *gorm.DB {
	customerTableName := entity.NewCustomer().TableName()
	addressTableName := entity.NewAddress().TableName()

	if q != nil {
		if id := q.GetFilterById(); id != nil {
			db = db.Where(customerTableName+".id = ?", id.Value)
		}

		if q.GetInclude("address") != nil {
			db = db.Preload("Address")
		}

		db = db.Joins("INNER JOIN " + addressTableName + " ON " +
			customerTableName + ".address_id = " + addressTableName + ".id")

		if search := q.GetSearch(); search != nil {
			s := "%" + *search + "%"

			db = db.
				Where(customerTableName+".name ILIKE ?", s).
				Or(customerTableName+".phone_number ILIKE ?", s).
				Or(customerTableName+".email ILIKE ?", s).
				Or(addressTableName+".full_address ILIKE ?", s)
		}

		if sort := q.GetSort("name"); sort != nil {
			db = db.Order(customerTableName + ".name " + sort.Order)
		}

		if sort := q.GetSort("phone_number"); sort != nil {
			db = db.Order(customerTableName + ".phone_number " + sort.Order)
		}

		if sort := q.GetSort("email"); sort != nil {
			db = db.Order(customerTableName + ".email " + sort.Order)
		}

		if sort := q.GetSort("full_address"); sort != nil {
			db = db.Order(addressTableName + ".full_address " + sort.Order)
		}

		if sort := q.GetSort("created_at"); sort != nil {
			db = db.Order(customerTableName + ".created_at " + sort.Order)
		}
	}

	return db
}

func (r *GormCustomerRepository) Get(q *query.Query) (customers []*entity.Customer, err error) {
	db := r.query(r.db, q)

	if err = db.Find(&customers).Error; err != nil {
		return nil, err
	}

	return customers, nil
}

func (r *GormCustomerRepository) Paginate(q *query.Query) (op *pagination.OffsetPagination[*entity.Customer], err error) {
	db := r.db.Model(&entity.Customer{})

	db = r.query(db, q)
	op, err = helper.GormDBPaginateWithQuery[*entity.Customer](db, q)
	if err != nil {
		return nil, err
	}
	return op, nil
}

func (r *GormCustomerRepository) First(q *query.Query) (customer *entity.Customer, err error) {
	db := r.query(r.db, q)

	if err := db.First(&customer).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorx.ErrCustomerNotFound
		}
		return nil, err
	}

	return
}

func (r *GormCustomerRepository) Find(id string, q *query.Query) (customer *entity.Customer, err error) {
	db := r.query(r.db, q)

	if err := db.Where("id = ?", id).First(&customer).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorx.ErrCustomerNotFound
		}
		return nil, err
	}
	return
}

func (r *GormCustomerRepository) Save(customers *entity.Customer) error {
	err := r.db.Omit("Address").Save(customers).Error
	if err != nil {
		if helper.IsPostgresDuplicateKeyError(err) {
			return errorx.ErrCustomerAlreadyExists
		}
		return err
	}

	return nil
}

func (r *GormCustomerRepository) SaveMany(customers []*entity.Customer) error {
	return r.db.CreateInBatches(customers, 100).Error
}

func (r *GormCustomerRepository) Destroy(customer *entity.Customer) error {
	return r.db.Delete(customer).Error
}
