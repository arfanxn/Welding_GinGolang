package repository

import (
	"errors"

	"github.com/arfanxn/welding/internal/infrastructure/database/helper"
	"github.com/arfanxn/welding/internal/module/address/domain/repository"
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	"github.com/arfanxn/welding/internal/module/shared/domain/errorx"
	"github.com/arfanxn/welding/pkg/pagination"
	"github.com/arfanxn/welding/pkg/query"
	"gorm.io/gorm"
)

type GormAddressRepository struct {
	db *gorm.DB
}

func NewGormAddressRepository(db *gorm.DB) repository.AddressRepository {
	return &GormAddressRepository{db: db}
}

func (r *GormAddressRepository) query(db *gorm.DB, q *query.Query) *gorm.DB {
	addressTableName := entity.NewAddress().TableName()

	if q != nil {
		if id := q.GetFilterById(); id != nil {
			db = db.Where(addressTableName+".id = ?", id.Value)
		}

		if search := q.GetSearch(); search != nil {
			db = db.Where(addressTableName+".full_address ILIKE ?", "%"+*search+"%")
		}

		if sort := q.GetSort("full_address"); sort != nil {
			db = db.Order(addressTableName + ".full_address " + sort.Order)
		}

		if sort := q.GetSort("created_at"); sort != nil {
			db = db.Order(addressTableName + ".created_at " + sort.Order)
		}
	}

	return db
}

func (r *GormAddressRepository) Get(q *query.Query) (addresses []*entity.Address, err error) {
	db := r.query(r.db, q)

	if err = db.Find(&addresses).Error; err != nil {
		return nil, err
	}

	return addresses, nil
}

func (r *GormAddressRepository) Paginate(q *query.Query) (op *pagination.OffsetPagination[*entity.Address], err error) {
	db := r.db.Model(&entity.Address{})

	db = r.query(db, q)
	op, err = helper.GormDBPaginateWithQuery[*entity.Address](db, q)
	if err != nil {
		return nil, err
	}
	return op, nil
}

func (r *GormAddressRepository) First(q *query.Query) (address *entity.Address, err error) {
	db := r.query(r.db, q)

	if err := db.First(&address).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorx.ErrAddressNotFound
		}
		return nil, err
	}

	return
}

func (r *GormAddressRepository) Find(id string, q *query.Query) (address *entity.Address, err error) {
	db := r.query(r.db, q)

	if err := db.Where("id = ?", id).First(&address).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorx.ErrAddressNotFound
		}
		return nil, err
	}
	return
}

func (r *GormAddressRepository) Save(addresses *entity.Address) error {
	err := r.db.Save(addresses).Error
	if err != nil {
		if helper.IsPostgresDuplicateKeyError(err) {
			return errorx.ErrAddressAlreadyExists
		}
		return err
	}
	return nil
}

func (r *GormAddressRepository) SaveMany(addresses []*entity.Address) error {
	return r.db.CreateInBatches(addresses, 100).Error
}

func (r *GormAddressRepository) Destroy(address *entity.Address) error {
	return r.db.Delete(address).Error
}
