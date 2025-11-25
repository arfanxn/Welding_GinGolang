package repository

import (
	"errors"

	"github.com/arfanxn/welding/internal/infrastructure/database/helper"
	"github.com/arfanxn/welding/internal/module/material_test_method/domain/repository"
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	"github.com/arfanxn/welding/internal/module/shared/domain/errorx"
	"github.com/arfanxn/welding/pkg/pagination"
	"github.com/arfanxn/welding/pkg/query"
	"gorm.io/gorm"
)

var _ repository.MaterialTestMethodRepository = (*GormMaterialTestMethodRepository)(nil)

type GormMaterialTestMethodRepository struct {
	db *gorm.DB
}

func NewGormMaterialTestMethodRepository(db *gorm.DB) repository.MaterialTestMethodRepository {
	return &GormMaterialTestMethodRepository{
		db: db,
	}
}

func (r *GormMaterialTestMethodRepository) All() (mtms []*entity.MaterialTestMethod, err error) {
	if err := r.db.Find(&mtms).Error; err != nil {
		return nil, err
	}
	return mtms, nil
}

func (r *GormMaterialTestMethodRepository) query(db *gorm.DB, q *query.Query) *gorm.DB {
	if q != nil {
		if id := q.GetFilterById(); id != nil {
			db = db.Where("id = ?", id.Value)
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

func (r *GormMaterialTestMethodRepository) Get(q *query.Query) (mtms []*entity.MaterialTestMethod, err error) {
	db := r.query(r.db, q)

	if err = db.Find(&mtms).Error; err != nil {
		return nil, err
	}

	return mtms, nil
}

func (r *GormMaterialTestMethodRepository) Paginate(q *query.Query) (op *pagination.OffsetPagination[*entity.MaterialTestMethod], err error) {
	db := r.db.Model(&entity.MaterialTestMethod{})

	db = r.query(db, q)
	op, err = helper.GormDBPaginateWithQuery[*entity.MaterialTestMethod](db, q)
	if err != nil {
		return nil, err
	}
	return op, nil
}

func (r *GormMaterialTestMethodRepository) First(q *query.Query) (mtm *entity.MaterialTestMethod, err error) {
	db := r.query(r.db, q)

	if err := db.First(&mtm).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorx.ErrMaterialTestMethodNotFound
		}
		return nil, err
	}

	return
}

func (r *GormMaterialTestMethodRepository) Find(id string) (mtm *entity.MaterialTestMethod, err error) {
	if err := r.db.Where("id = ?", id).First(&mtm).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorx.ErrMaterialTestMethodNotFound
		}
		return nil, err
	}
	return
}

func (r *GormMaterialTestMethodRepository) FindByName(name string) (mtm *entity.MaterialTestMethod, err error) {
	if err := r.db.Where("name = ?", name).First(&mtm).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorx.ErrMaterialTestMethodNotFound
		}
		return nil, err
	}
	return
}

func (r *GormMaterialTestMethodRepository) FindByIds(ids []string) (mtms []*entity.MaterialTestMethod, err error) {
	if err := r.db.Where("id IN (?)", ids).Find(&mtms).Error; err != nil {
		return nil, err
	}
	if len(mtms) != len(ids) {
		return nil, errorx.ErrMaterialTestMethodNotFound
	}
	return
}

func (r *GormMaterialTestMethodRepository) Save(mtms *entity.MaterialTestMethod) error {
	err := r.db.Save(mtms).Error
	if err != nil {
		if helper.IsPostgresDuplicateKeyError(err) {
			return errorx.ErrMaterialTestMethodAlreadyExists
		}
		return err
	}
	return nil
}

func (r *GormMaterialTestMethodRepository) SaveMany(mtms []*entity.MaterialTestMethod) error {
	return r.db.CreateInBatches(mtms, 100).Error
}

func (r *GormMaterialTestMethodRepository) Destroy(mtm *entity.MaterialTestMethod) error {
	return r.db.Delete(mtm).Error
}
