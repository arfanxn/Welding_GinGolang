package repository

import (
	"errors"

	"github.com/arfanxn/welding/internal/infrastructure/database/helper"
	"github.com/arfanxn/welding/internal/module/material_test_service/domain/repository"
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	"github.com/arfanxn/welding/internal/module/shared/domain/errorx"
	"github.com/arfanxn/welding/pkg/pagination"
	"github.com/arfanxn/welding/pkg/query"
	"gorm.io/gorm"
)

var _ repository.MaterialTestServiceRepository = (*GormMaterialTestServiceRepository)(nil)

type GormMaterialTestServiceRepository struct {
	db *gorm.DB
}

func NewGormMaterialTestServiceRepository(db *gorm.DB) repository.MaterialTestServiceRepository {
	return &GormMaterialTestServiceRepository{
		db: db,
	}
}

func (r *GormMaterialTestServiceRepository) All() (mtms []*entity.MaterialTestService, err error) {
	if err := r.db.Find(&mtms).Error; err != nil {
		return nil, err
	}
	return mtms, nil
}

func (r *GormMaterialTestServiceRepository) query(db *gorm.DB, q *query.Query) *gorm.DB {
	mtServiceTableName := entity.NewMaterialTestService().TableName()
	mtMachineTableName := entity.NewMaterialTestMachine().TableName()
	mtMethodTableName := entity.NewMaterialTestMethod().TableName()

	if q != nil {
		if id := q.GetFilterById(); id != nil {
			// Unscoped() is used to include deleted records, only unscope on id filter
			db = db.Unscoped().Where(mtServiceTableName+".id = ?", id.Value)
		}

		if q.GetInclude("machine") != nil {
			db = db.Preload("Machine", func(db *gorm.DB) *gorm.DB {
				return db.Unscoped()
			})
		}

		if q.GetInclude("method") != nil {
			db = db.Preload("Method", func(db *gorm.DB) *gorm.DB {
				return db.Unscoped()
			})
		}

		if search := q.GetSearch(); search != nil {
			s := "%" + *search + "%"

			db = db.
				Joins("LEFT JOIN " + mtMachineTableName + " ON " + mtMachineTableName + ".id = " + mtServiceTableName + ".machine_id").
				Joins("LEFT JOIN " + mtMethodTableName + " ON " + mtMethodTableName + ".id = " + mtServiceTableName + ".method_id").
				Where(
					db.Where(mtServiceTableName+".test_name ILIKE ?", s).
						Or(mtServiceTableName+".service_type ILIKE ?", s).
						Or(mtServiceTableName+".service_code ILIKE ?", s).
						Or(mtServiceTableName+".unit ILIKE ?", s).
						Or("CAST("+mtServiceTableName+".price AS TEXT) LIKE ?", s). // Cast numeric price to text for partial string matching
						Or(mtMachineTableName+".name ILIKE ?", s).
						Or(mtMethodTableName+".name ILIKE ?", s),
				)
		}

		if machineId := q.GetFilter("machine_id", query.OperatorEqual); machineId != nil {
			db = db.Where(mtServiceTableName+".machine_id = ?", machineId.Value)
		}

		if methodId := q.GetFilter("method_id", query.OperatorEqual); methodId != nil {
			db = db.Where(mtServiceTableName+".method_id = ?", methodId.Value)
		}

		if testName := q.GetFilter("test_name", query.OperatorEqual); testName != nil {
			db = db.Where(mtServiceTableName+".test_name = ?", testName.Value)
		}

		if serviceType := q.GetFilter("service_type", query.OperatorEqual); serviceType != nil {
			db = db.Where(mtServiceTableName+".service_type = ?", serviceType.Value)
		}

		if serviceCode := q.GetFilter("service_code", query.OperatorEqual); serviceCode != nil {
			db = db.Where(mtServiceTableName+".service_code = ?", serviceCode.Value)
		}

		if unit := q.GetFilter("unit", query.OperatorEqual); unit != nil {
			db = db.Where(mtServiceTableName+".unit = ?", unit.Value)
		}

		if price := q.GetFilter("price", query.OperatorLike); price != nil {
			db = db.Where("CAST("+mtServiceTableName+".price AS TEXT) LIKE ?", "%"+price.Value+"%")
		}

		if priceBetween := q.GetFilter("price", query.OperatorBetween); priceBetween != nil {
			db = db.Where(mtServiceTableName+".price BETWEEN ? AND ?", priceBetween.Values[0], priceBetween.Values[1])
		}

		if priceSort := q.GetSort("price"); priceSort != nil {
			db = db.Order(mtServiceTableName + ".price " + priceSort.Order)
		}

		if createdAt := q.GetSort("created_at"); createdAt != nil {
			db = db.Order(mtServiceTableName + ".created_at " + createdAt.Order)
		}
	}

	return db
}

func (r *GormMaterialTestServiceRepository) Get(q *query.Query) (mtms []*entity.MaterialTestService, err error) {
	db := r.query(r.db, q)

	if err = db.Find(&mtms).Error; err != nil {
		return nil, err
	}

	return mtms, nil
}

func (r *GormMaterialTestServiceRepository) Paginate(q *query.Query) (op *pagination.OffsetPagination[*entity.MaterialTestService], err error) {
	db := r.db.Model(&entity.MaterialTestService{})

	db = r.query(db, q)
	op, err = helper.GormDBPaginateWithQuery[*entity.MaterialTestService](db, q)
	if err != nil {
		return nil, err
	}
	return op, nil
}

func (r *GormMaterialTestServiceRepository) First(q *query.Query) (mtm *entity.MaterialTestService, err error) {
	db := r.query(r.db, q)

	if err := db.First(&mtm).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorx.ErrMaterialTestServiceNotFound
		}
		return nil, err
	}

	return
}

func (r *GormMaterialTestServiceRepository) Find(id string, q *query.Query) (mtm *entity.MaterialTestService, err error) {
	db := r.query(r.db, q)

	if err := db.Unscoped().Where("id = ?", id).First(&mtm).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorx.ErrMaterialTestServiceNotFound
		}
		return nil, err
	}
	return
}

func (r *GormMaterialTestServiceRepository) FindByIds(ids []string, q *query.Query) (mtms []*entity.MaterialTestService, err error) {
	db := r.query(r.db, q)

	if err := db.Where("id IN (?)", ids).Find(&mtms).Error; err != nil {
		return nil, err
	}
	if len(mtms) != len(ids) {
		return nil, errorx.ErrMaterialTestServiceNotFound
	}
	return
}

func (r *GormMaterialTestServiceRepository) CountByMachineId(machineId string) (int64, error) {
	var (
		count int64
		err   error
	)
	err = r.db.Model(&entity.MaterialTestService{}).Where("machine_id = ?", machineId).Count(&count).Error
	return count, err
}

func (r *GormMaterialTestServiceRepository) CountByMethodId(methodId string) (int64, error) {
	var (
		count int64
		err   error
	)
	err = r.db.Model(&entity.MaterialTestService{}).Where("method_id = ?", methodId).Count(&count).Error
	return count, err
}

func (r *GormMaterialTestServiceRepository) Save(mtms *entity.MaterialTestService) error {
	err := r.db.Save(mtms).Error
	if err != nil {
		if helper.IsPostgresDuplicateKeyError(err) {
			return errorx.ErrMaterialTestServiceAlreadyExists
		}
		return err
	}
	return nil
}

func (r *GormMaterialTestServiceRepository) SaveMany(mtms []*entity.MaterialTestService) error {
	return r.db.CreateInBatches(mtms, 100).Error
}

func (r *GormMaterialTestServiceRepository) Destroy(mtm *entity.MaterialTestService) error {
	return r.db.Delete(mtm).Error
}
