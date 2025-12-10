package repository

import (
	"errors"

	"github.com/arfanxn/welding/internal/infrastructure/database/helper"
	"github.com/arfanxn/welding/internal/module/material_test_order/domain/repository"
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	"github.com/arfanxn/welding/internal/module/shared/domain/errorx"
	"github.com/arfanxn/welding/pkg/pagination"
	"github.com/arfanxn/welding/pkg/query"
	"gorm.io/gorm"
)

type GormMaterialTestOrderRepository struct {
	db *gorm.DB
}

func NewGormMaterialTestOrderRepository(db *gorm.DB) repository.MaterialTestOrderRepository {
	return &GormMaterialTestOrderRepository{
		db: db,
	}
}

func (r *GormMaterialTestOrderRepository) query(db *gorm.DB, q *query.Query) (*gorm.DB, error) {
	mtoTableName := entity.NewMaterialTestOrder().TableName()

	if q != nil {
		if f := q.GetFilterById(); f != nil {
			db = db.Where(mtoTableName+".id = ?", f.Value)
		}

		if search := q.GetSearch(); search != nil {
			// TODO: add more search fields
			db = db.Where("CAST("+mtoTableName+".number AS TEXT) ILIKE ?", "%"+*search+"%")
		}

		if number := q.GetFilter("number", query.OperatorEqual); number != nil {
			db = db.Where(mtoTableName+".number = ?", number.Value)
		}

		if createdAtBetween := q.GetFilter("created_at", query.OperatorBetween); createdAtBetween != nil {
			db = db.Where(mtoTableName+".created_at BETWEEN ? AND ?", createdAtBetween.Values[0], createdAtBetween.Values[1])
		}

		if sort := q.GetSort("created_at"); sort != nil {
			db = db.Order(mtoTableName + ".created_at " + sort.Order)
		}
	}

	return db, nil
}

func (r *GormMaterialTestOrderRepository) Get(q *query.Query) ([]*entity.MaterialTestOrder, error) {
	var mtos []*entity.MaterialTestOrder

	db, err := r.query(r.db, q)
	if err != nil {
		return nil, err
	}

	if err := db.Find(&mtos).Error; err != nil {
		return nil, err
	}

	return mtos, nil
}

func (r *GormMaterialTestOrderRepository) Paginate(q *query.Query) (*pagination.OffsetPagination[*entity.MaterialTestOrder], error) {
	db := r.db.Model(&entity.MaterialTestOrder{})

	db, err := r.query(db, q)
	if err != nil {
		return nil, err
	}

	op, err := helper.GormDBPaginateWithQuery[*entity.MaterialTestOrder](db, q)
	if err != nil {
		return nil, err
	}
	return op, nil
}

func (r *GormMaterialTestOrderRepository) First(q *query.Query) (*entity.MaterialTestOrder, error) {
	var mto *entity.MaterialTestOrder

	db, err := r.query(r.db, q)
	if err != nil {
		return nil, err
	}

	if err := db.First(&mto).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorx.ErrMaterialTestOrderNotFound
		}
		return nil, err
	}

	return mto, nil
}

func (r *GormMaterialTestOrderRepository) Find(id string, q *query.Query) (*entity.MaterialTestOrder, error) {
	var mto entity.MaterialTestOrder

	db, err := r.query(r.db, q)
	if err != nil {
		return nil, err
	}

	if err := db.Where("id = ?", id).First(&mto).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorx.ErrMaterialTestOrderNotFound
		}
		return nil, err
	}
	return &mto, nil
}

func (r *GormMaterialTestOrderRepository) CountByCustomerId(customerId string) (count int64, err error) {
	db := r.db.Model(&entity.MaterialTestOrder{})

	if err = db.Where("customer_id = ?", customerId).Count(&count).Error; err != nil {
		return 0, err
	}
	return
}

func (r *GormMaterialTestOrderRepository) Save(mto *entity.MaterialTestOrder) error {
	return r.db.Save(mto).Error
}

func (r *GormMaterialTestOrderRepository) SaveMany(mtos []*entity.MaterialTestOrder) error {
	return r.db.CreateInBatches(mtos, 100).Error
}

func (r *GormMaterialTestOrderRepository) Destroy(mto *entity.MaterialTestOrder) error {
	return r.db.Delete(mto).Error
}
