package repository

import (
	"errors"

	"github.com/arfanxn/welding/internal/infrastructure/database/helper"
	"github.com/arfanxn/welding/internal/module/material_test_work_category/domain/repository"
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	"github.com/arfanxn/welding/internal/module/shared/domain/errorx"
	"github.com/arfanxn/welding/pkg/pagination"
	"github.com/arfanxn/welding/pkg/query"
	"gorm.io/gorm"
)

var _ repository.MaterialTestWorkCategoryRepository = (*GormMaterialTestWorkCategoryRepository)(nil)

type GormMaterialTestWorkCategoryRepository struct {
	db *gorm.DB
}

func NewGormMaterialTestWorkCategoryRepository(db *gorm.DB) repository.MaterialTestWorkCategoryRepository {
	return &GormMaterialTestWorkCategoryRepository{
		db: db,
	}
}

// query applies query filters and sorting to the database query based on the provided Query DTO.
// It supports searching by name (case-insensitive) and sorting by name in ascending or descending order.
// The modified *gorm.DB is returned with the applied scopes.
func (r *GormMaterialTestWorkCategoryRepository) query(db *gorm.DB, q *query.Query) *gorm.DB {
	mtWorkPackageTableName := entity.NewMaterialTestWorkCategory().TableName()

	if q != nil {
		if id := q.GetFilterById(); id != nil {
			db = db.Where(mtWorkPackageTableName+".id = ?", id.Value)
		}

		if search := q.GetSearch(); search != nil {
			s := "%" + *search + "%"
			db = db.Where(mtWorkPackageTableName+".name ILIKE ?", s).Or(mtWorkPackageTableName+".description ILIKE ?", s)
		}

		if sort := q.GetSort("name"); sort != nil {
			db = db.Order(mtWorkPackageTableName + ".name " + sort.Order)
		}

		if sort := q.GetSort("created_at"); sort != nil {
			db = db.Order(mtWorkPackageTableName + ".created_at " + sort.Order)
		}
	}

	return db
}

func (r *GormMaterialTestWorkCategoryRepository) Get(q *query.Query) ([]*entity.MaterialTestWorkCategory, error) {
	var materialTestWorkCategories []*entity.MaterialTestWorkCategory

	db := r.query(r.db, q)

	if err := db.Find(&materialTestWorkCategories).Error; err != nil {
		return nil, err
	}

	return materialTestWorkCategories, nil
}

func (r *GormMaterialTestWorkCategoryRepository) Paginate(q *query.Query) (*pagination.OffsetPagination[*entity.MaterialTestWorkCategory], error) {
	db := r.db.Model(&entity.MaterialTestWorkCategory{})

	db = r.query(db, q)

	pagination, err := helper.GormDBPaginateWithQuery[*entity.MaterialTestWorkCategory](db, q)
	if err != nil {
		return nil, err
	}
	return pagination, nil
}

func (r *GormMaterialTestWorkCategoryRepository) First(q *query.Query) (*entity.MaterialTestWorkCategory, error) {
	var materialTestWorkCategory *entity.MaterialTestWorkCategory

	db := r.query(r.db, q)

	if err := db.First(&materialTestWorkCategory).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorx.ErrMaterialTestWorkCategoryNotFound
		}
		return nil, err
	}

	return materialTestWorkCategory, nil
}

func (r *GormMaterialTestWorkCategoryRepository) Find(id string, q *query.Query) (*entity.MaterialTestWorkCategory, error) {
	var materialTestWorkCategory entity.MaterialTestWorkCategory

	db := r.query(r.db, q)

	if err := db.Where("id = ?", id).First(&materialTestWorkCategory).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorx.ErrMaterialTestWorkCategoryNotFound
		}
		return nil, err
	}
	return &materialTestWorkCategory, nil
}

func (r *GormMaterialTestWorkCategoryRepository) Save(materialTestWorkCategory *entity.MaterialTestWorkCategory) error {
	err := r.db.Save(materialTestWorkCategory).Error
	if err != nil {
		if helper.IsPostgresDuplicateKeyError(err) {
			return errorx.ErrMaterialTestWorkCategoryAlreadyExists
		}
		return err
	}
	return nil
}

func (r *GormMaterialTestWorkCategoryRepository) SaveMany(materialTestWorkCategories []*entity.MaterialTestWorkCategory) error {
	return r.db.CreateInBatches(materialTestWorkCategories, 100).Error
}

func (r *GormMaterialTestWorkCategoryRepository) Destroy(materialTestWorkCategory *entity.MaterialTestWorkCategory) error {
	return r.db.Delete(materialTestWorkCategory).Error
}
