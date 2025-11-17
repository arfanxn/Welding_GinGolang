package repository

import (
	"errors"

	"github.com/arfanxn/welding/internal/infrastructure/database/helper"
	"github.com/arfanxn/welding/internal/module/role/domain/repository"
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	"github.com/arfanxn/welding/internal/module/shared/domain/errorx"
	"github.com/arfanxn/welding/pkg/pagination"
	"github.com/arfanxn/welding/pkg/query"
	"github.com/gookit/goutil"
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
func (r *GormRoleRepository) query(db *gorm.DB, q *query.Query) *gorm.DB {
	if id := q.GetFilterById(); id != nil {
		db = db.Where("id = ?", id.Value)
	}

	if search := q.GetSearch(); !search.IsZero() {
		db = db.Where("name LIKE ?", "%"+search.String+"%")
	}

	if !q.GetInclude("permissions").IsZero() {
		db = db.Preload("Permissions")
	}

	if !q.GetInclude("users").IsZero() {
		db = db.Preload("Users")
	}

	if sort := q.GetSort("name"); sort != nil {
		db = db.Order("name " + sort.Order)
	}

	if sort := q.GetSort("created_at"); sort != nil {
		db = db.Order("created_at " + sort.Order)
	}

	return db
}

func (r *GormRoleRepository) Get(q *query.Query) ([]*entity.Role, error) {
	var roles []*entity.Role

	db := r.query(r.db, q)

	if err := db.Find(&roles).Error; err != nil {
		return nil, err
	}

	return roles, nil
}

func (r *GormRoleRepository) Paginate(q *query.Query) (*pagination.OffsetPagination[*entity.Role], error) {
	db := r.db.Model(&entity.Role{})

	db = r.query(db, q)

	pagination, err := helper.GormDBPaginateWithQuery[*entity.Role](db, q)
	if err != nil {
		return nil, err
	}
	return pagination, nil
}

func (r *GormRoleRepository) Find(id string) (*entity.Role, error) {
	var role entity.Role
	if err := r.db.Where("id = ?", id).First(&role).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorx.ErrRoleNotFound
		}
		return nil, err
	}
	return &role, nil
}

func (r *GormRoleRepository) FindDefault() (*entity.Role, error) {
	var role entity.Role
	if err := r.db.Where("is_default = ?", true).First(&role).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorx.ErrRoleDefaultNotConfigured
		}
		return nil, err
	}
	return &role, nil
}

func (r *GormRoleRepository) FindByIds(ids []string) ([]*entity.Role, error) {
	var roles []*entity.Role
	if err := r.db.Where("id IN (?)", ids).Find(&roles).Error; err != nil {
		if len(roles) != len(ids) {
			return nil, errorx.ErrRolesNotFound
		}
		return nil, err
	}
	return roles, nil
}

func (r *GormRoleRepository) FindByName(name string) (*entity.Role, error) {
	var role entity.Role
	if err := r.db.Where("name = ?", name).First(&role).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorx.ErrRoleNotFound
		}
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
			return errorx.ErrRoleAlreadyExists
		}
		return err
	}

	// Update permissions if any provided
	if !goutil.IsEmpty(role.Permissions) {
		// Replace all role-permission associations
		// Note: role.Permissions should contain only Permission{ID: X} structs
		if err := tx.Model(role).Association("Permissions").Replace(role.Permissions); err != nil {
			tx.Rollback()
			return err
		}
	}

	// Commit transaction
	if err := tx.Commit().Error; err != nil {
		return err
	}

	db := r.db.Model(&entity.Role{})
	if !goutil.IsEmpty(role.Permissions) {
		db = db.Preload("Permissions")
	}

	if err := db.First(role, "id = ?", role.Id).Error; err != nil {
		return err
	}

	return nil
}

func (r *GormRoleRepository) SetDefault(role *entity.Role) error {
	tx := r.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	// Set all roles to not default
	if err := tx.Model(&entity.Role{}).Where("id != ?", role.Id).Update("is_default", false).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Set the specified role as default
	role.IsDefault = true
	if err := tx.Save(role).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (r *GormRoleRepository) SaveMany(roles []*entity.Role) error {
	return r.db.CreateInBatches(roles, 100).Error
}

func (r *GormRoleRepository) Destroy(role *entity.Role) error {
	return r.db.Delete(role).Error
}
