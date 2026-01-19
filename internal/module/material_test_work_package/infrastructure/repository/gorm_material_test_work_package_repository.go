package repository

import (
	"errors"

	"github.com/arfanxn/welding/internal/infrastructure/database/helper"
	"github.com/arfanxn/welding/internal/module/material_test_work_package/domain/repository"
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	"github.com/arfanxn/welding/internal/module/shared/domain/errorx"
	"github.com/arfanxn/welding/pkg/pagination"
	"github.com/arfanxn/welding/pkg/query"
	"gorm.io/gorm"
)

var _ repository.MaterialTestWorkPackageRepository = (*GormMaterialTestWorkPackageRepository)(nil)

type GormMaterialTestWorkPackageRepository struct {
	db *gorm.DB
}

func NewGormMaterialTestWorkPackageRepository(db *gorm.DB) repository.MaterialTestWorkPackageRepository {
	return &GormMaterialTestWorkPackageRepository{
		db: db,
	}
}

// query applies query filters and sorting to the database query based on the provided Query DTO.
// It supports searching by name (case-insensitive) and sorting by name in ascending or descending order.
// The modified *gorm.DB is returned with the applied scopes.
func (r *GormMaterialTestWorkPackageRepository) query(db *gorm.DB, q *query.Query) *gorm.DB {
	mtWorkPackageTableName := entity.NewMaterialTestWorkPackage().TableName()

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

func (r *GormMaterialTestWorkPackageRepository) Get(q *query.Query) ([]*entity.MaterialTestWorkPackage, error) {
	var materialTestWorkPackages []*entity.MaterialTestWorkPackage

	db := r.query(r.db, q)

	if err := db.Find(&materialTestWorkPackages).Error; err != nil {
		return nil, err
	}

	return materialTestWorkPackages, nil
}

func (r *GormMaterialTestWorkPackageRepository) Paginate(q *query.Query) (*pagination.OffsetPagination[*entity.MaterialTestWorkPackage], error) {
	db := r.db.Model(&entity.MaterialTestWorkPackage{})

	db = r.query(db, q)

	pagination, err := helper.GormDBPaginateWithQuery[*entity.MaterialTestWorkPackage](db, q)
	if err != nil {
		return nil, err
	}
	return pagination, nil
}

func (r *GormMaterialTestWorkPackageRepository) First(q *query.Query) (*entity.MaterialTestWorkPackage, error) {
	var materialTestWorkPackage *entity.MaterialTestWorkPackage

	db := r.query(r.db, q)

	if err := db.First(&materialTestWorkPackage).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorx.ErrMaterialTestWorkPackageNotFound
		}
		return nil, err
	}

	return materialTestWorkPackage, nil
}

func (r *GormMaterialTestWorkPackageRepository) Find(id string, q *query.Query) (*entity.MaterialTestWorkPackage, error) {
	var materialTestWorkPackage entity.MaterialTestWorkPackage

	db := r.query(r.db, q)

	if err := db.Where("id = ?", id).First(&materialTestWorkPackage).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorx.ErrMaterialTestWorkPackageNotFound
		}
		return nil, err
	}
	return &materialTestWorkPackage, nil
}

func (r *GormMaterialTestWorkPackageRepository) Save(materialTestWorkPackage *entity.MaterialTestWorkPackage) error {
	err := r.db.Save(materialTestWorkPackage).Error
	if err != nil {
		if helper.IsPostgresDuplicateKeyError(err) {
			return errorx.ErrMaterialTestWorkPackageAlreadyExists
		}
		return err
	}
	return nil
}

func (r *GormMaterialTestWorkPackageRepository) SaveMany(materialTestWorkPackages []*entity.MaterialTestWorkPackage) error {
	return r.db.CreateInBatches(materialTestWorkPackages, 100).Error
}

func (r *GormMaterialTestWorkPackageRepository) Destroy(materialTestWorkPackage *entity.MaterialTestWorkPackage) error {
	return r.db.Delete(materialTestWorkPackage).Error
}
