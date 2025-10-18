package repository

import (
	"github.com/arfanxn/welding/internal/infrastructure/database/helper"
	"github.com/arfanxn/welding/internal/module/role/domain/repository"
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	"github.com/arfanxn/welding/internal/module/shared/usecase/dto"
	"github.com/arfanxn/welding/pkg/reflectutil"
	"gorm.io/gorm"
)

var _ repository.RoleRepository = (*GormRoleRepository)(nil)

type GormRoleRepository struct {
	db *gorm.DB
}

func NewGormRoleRepository(db *gorm.DB) repository.RoleRepository {
	return &GormRoleRepository{
		db: db,
	}
}

func (r *GormRoleRepository) All() ([]*entity.Role, error) {
	var roles []*entity.Role
	if err := r.db.Find(&roles).Error; err != nil {
		return nil, err
	}
	return roles, nil
}

// query applies query filters and sorting to the database query based on the provided Query DTO.
// It supports searching by name (case-insensitive) and sorting by name in ascending or descending order.
// The modified *gorm.DB is returned with the applied scopes.
func (r *GormRoleRepository) query(db *gorm.DB, queryDto *dto.Query) *gorm.DB {
	if id, hasId := queryDto.GetFilterByCO("id", dto.QueryFilterOperatorEqual); hasId {
		db = db.Where("id = ?", id.Value)
	}

	if search, hasSearch := queryDto.GetSearch(); hasSearch {
		db = db.Where("name LIKE ?", "%"+search+"%")
	}

	if _, shouldInclude := queryDto.GetInclude("permissions"); shouldInclude {
		db = db.Preload("Permissions")
	}

	if _, shouldInclude := queryDto.GetInclude("users"); shouldInclude {
		db = db.Preload("Users")
	}

	if sort, shouldSort := queryDto.GetSortByColumn("name"); shouldSort {
		db = db.Order("name " + sort.Order)
	}

	if sort, shouldSort := queryDto.GetSortByColumn("created_at"); shouldSort {
		db = db.Order("created_at " + sort.Order)
	}

	return db
}

func (r *GormRoleRepository) Get(queryDto *dto.Query) ([]*entity.Role, error) {
	var roles []*entity.Role

	db := r.query(r.db, queryDto)

	if err := db.Find(&roles).Error; err != nil {
		return nil, err
	}

	return roles, nil
}

func (r *GormRoleRepository) Paginate(query *dto.Query) (*dto.Pagination[*entity.Role], error) {
	db := r.db.Model(&entity.Role{})

	db = r.query(db, query)

	pagination, err := helper.GormDBPaginateWithQueryDTO[*entity.Role](db, query)
	if err != nil {
		return nil, err
	}
	return pagination, nil
}

func (r *GormRoleRepository) Find(id string) (*entity.Role, error) {
	var role entity.Role
	if err := r.db.Where("id = ?", id).First(&role).Error; err != nil {
		return nil, err
	}
	return &role, nil
}

func (r *GormRoleRepository) FindByName(name string) (*entity.Role, error) {
	var role entity.Role
	if err := r.db.Where("name = ?", name).First(&role).Error; err != nil {
		return nil, err
	}
	return &role, nil
}

func (r *GormRoleRepository) Save(role *entity.Role) error {
	// Start transaction
	tx := r.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	// Save role record (without permissions to prevent M2M race conditions)
	err := tx.Omit("Permissions").Save(role).Error
	if err != nil {
		tx.Rollback()
		if helper.IsPostgresDuplicateKeyError(err) {
			return gorm.ErrDuplicatedKey
		}
		return err
	}

	// Update permissions if any provided
	if reflectutil.IsSlice(role.Permissions) {
		// Replace all role-permission associations
		// Note: role.Permissions should contain only Permission{ID: X} structs
		if err := tx.Model(role).Association("Permissions").Replace(role.Permissions); err != nil {
			tx.Rollback()
			return err
		}
	}

	// Commit transaction
	return tx.Commit().Error
}

func (r *GormRoleRepository) SaveMany(roles []*entity.Role) error {
	return r.db.CreateInBatches(roles, 100).Error
}

func (r *GormRoleRepository) Destroy(role *entity.Role) error {
	return r.db.Delete(role).Error
}
