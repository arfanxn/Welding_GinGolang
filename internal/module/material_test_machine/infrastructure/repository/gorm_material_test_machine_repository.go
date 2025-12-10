package repository

import (
	"errors"

	"github.com/arfanxn/welding/internal/infrastructure/database/helper"
	"github.com/arfanxn/welding/internal/module/material_test_machine/domain/repository"
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	"github.com/arfanxn/welding/internal/module/shared/domain/errorx"
	"github.com/arfanxn/welding/pkg/pagination"
	"github.com/arfanxn/welding/pkg/query"
	"gorm.io/gorm"
)

var _ repository.MaterialTestMachineRepository = (*GormMaterialTestMachineRepository)(nil)

type GormMaterialTestMachineRepository struct {
	db *gorm.DB
}

func NewGormMaterialTestMachineRepository(db *gorm.DB) repository.MaterialTestMachineRepository {
	return &GormMaterialTestMachineRepository{
		db: db,
	}
}

func (r *GormMaterialTestMachineRepository) All() (mtms []*entity.MaterialTestMachine, err error) {
	if err := r.db.Find(&mtms).Error; err != nil {
		return nil, err
	}
	return mtms, nil
}

func (r *GormMaterialTestMachineRepository) query(db *gorm.DB, q *query.Query) *gorm.DB {
	if q != nil {
		if id := q.GetFilterById(); id != nil {
			// Unscoped() is used to include deleted records, only unscope on id filter
			db = db.Unscoped().Where("id = ?", id.Value)
		}

		if search := q.GetSearch(); search != nil {
			db = db.Where("name ILIKE ?", "%"+*search+"%").Or("description ILIKE ?", "%"+*search+"%")
		}

		if name := q.GetFilter("name", query.OperatorEqual); name != nil {
			db = db.Where("name = ?", name.Value)
		}

		if sort := q.GetSort("name"); sort != nil {
			db = db.Order("name " + sort.Order)
		}

		if createdAt := q.GetSort("created_at"); createdAt != nil {
			db = db.Order("created_at " + createdAt.Order)
		}
	}

	return db
}

func (r *GormMaterialTestMachineRepository) Get(q *query.Query) (mtms []*entity.MaterialTestMachine, err error) {
	db := r.query(r.db, q)

	if err = db.Find(&mtms).Error; err != nil {
		return nil, err
	}

	return mtms, nil
}

func (r *GormMaterialTestMachineRepository) Paginate(q *query.Query) (op *pagination.OffsetPagination[*entity.MaterialTestMachine], err error) {
	db := r.db.Model(&entity.MaterialTestMachine{})

	db = r.query(db, q)
	op, err = helper.GormDBPaginateWithQuery[*entity.MaterialTestMachine](db, q)
	if err != nil {
		return nil, err
	}
	return op, nil
}

func (r *GormMaterialTestMachineRepository) First(q *query.Query) (mtm *entity.MaterialTestMachine, err error) {
	db := r.query(r.db, q)

	if err := db.First(&mtm).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorx.ErrMaterialTestMachineNotFound
		}
		return nil, err
	}

	return
}

func (r *GormMaterialTestMachineRepository) Find(id string, q *query.Query) (mtm *entity.MaterialTestMachine, err error) {
	db := r.query(r.db, q)

	if err := db.Unscoped().Where("id = ?", id).First(&mtm).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorx.ErrMaterialTestMachineNotFound
		}
		return nil, err
	}
	return
}

func (r *GormMaterialTestMachineRepository) FindByName(name string, q *query.Query) (mtm *entity.MaterialTestMachine, err error) {
	db := r.query(r.db, q)

	if err := db.Where("name = ?", name).First(&mtm).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorx.ErrMaterialTestMachineNotFound
		}
		return nil, err
	}
	return
}

func (r *GormMaterialTestMachineRepository) FindByIds(ids []string, q *query.Query) (mtms []*entity.MaterialTestMachine, err error) {
	db := r.query(r.db, q)

	if err := db.Where("id IN (?)", ids).Find(&mtms).Error; err != nil {
		return nil, err
	}
	if len(mtms) != len(ids) {
		return nil, errorx.ErrMaterialTestMachineNotFound
	}
	return
}

func (r *GormMaterialTestMachineRepository) Save(mtms *entity.MaterialTestMachine) error {
	err := r.db.Save(mtms).Error
	if err != nil {
		if helper.IsPostgresDuplicateKeyError(err) {
			return errorx.ErrMaterialTestMachineAlreadyExists
		}
		return err
	}
	return nil
}

func (r *GormMaterialTestMachineRepository) SaveMany(mtms []*entity.MaterialTestMachine) error {
	return r.db.CreateInBatches(mtms, 100).Error
}

func (r *GormMaterialTestMachineRepository) Destroy(mtm *entity.MaterialTestMachine) error {
	return r.db.Delete(mtm).Error
}
