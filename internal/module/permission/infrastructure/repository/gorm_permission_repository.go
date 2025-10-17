package repository

import (
	"github.com/arfanxn/welding/internal/infrastructure/database/helper"
	"github.com/arfanxn/welding/internal/module/permission/domain/repository"
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	"github.com/arfanxn/welding/internal/module/shared/usecase/dto"
	"gorm.io/gorm"
)

var _ repository.PermissionRepository = (*GormPermissionRepository)(nil)

type GormPermissionRepository struct {
	db *gorm.DB
}

func NewGormPermissionRepository(db *gorm.DB) repository.PermissionRepository {
	return &GormPermissionRepository{
		db: db,
	}
}

func (r *GormPermissionRepository) All() ([]*entity.Permission, error) {
	var permissions []*entity.Permission
	if err := r.db.Find(&permissions).Error; err != nil {
		return nil, err
	}
	return permissions, nil
}

// scope applies query filters and sorting to the database query based on the provided Query DTO.
// It supports searching by name (case-insensitive) and sorting by name in ascending or descending order.
// The modified *gorm.DB is returned with the applied scopes.
func (r *GormPermissionRepository) scope(db *gorm.DB, queryDto *dto.Query) *gorm.DB {
	if search, hasSearch := queryDto.GetSearch(); hasSearch {
		db = db.Where("name LIKE ?", "%"+search+"%")
	}

	if sortName, hasSortName := queryDto.GetSortByColumn("name"); hasSortName {
		db = db.Order("name " + sortName.Order)
	}

	return db
}

// Get retrieves a list of permissions based on the provided query DTO.
// It applies the scope function to the database query to filter and sort the results.
// The modified *gorm.DB is returned with the applied scopes.
func (r *GormPermissionRepository) Get(queryDto *dto.Query) ([]*entity.Permission, error) {
	var permissions []*entity.Permission

	db := r.scope(r.db, queryDto)

	if err := db.Find(&permissions).Error; err != nil {
		return nil, err
	}

	return permissions, nil
}

func (r *GormPermissionRepository) Paginate(queryDto *dto.Query) (*dto.Pagination[*entity.Permission], error) {
	db := r.db.Model(&entity.Permission{})

	db = r.scope(db, queryDto)
	pagination, err := helper.GormDBPaginateWithQueryDTO[*entity.Permission](db, queryDto)
	if err != nil {
		return nil, err
	}
	return pagination, nil
}

func (r *GormPermissionRepository) Find(id string) (*entity.Permission, error) {
	var permission entity.Permission
	if err := r.db.Where("id = ?", id).First(&permission).Error; err != nil {
		return nil, err
	}
	return &permission, nil
}

func (r *GormPermissionRepository) FindByName(name string) (*entity.Permission, error) {
	var permission entity.Permission
	if err := r.db.Where("name = ?", name).First(&permission).Error; err != nil {
		return nil, err
	}
	return &permission, nil
}

func (r *GormPermissionRepository) Save(permission *entity.Permission) error {
	return r.db.Save(permission).Error
}

func (r *GormPermissionRepository) SaveMany(permissions []*entity.Permission) error {
	return r.db.CreateInBatches(permissions, 100).Error
}
