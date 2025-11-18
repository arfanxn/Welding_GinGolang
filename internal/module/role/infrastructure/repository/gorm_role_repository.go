package repository

import (
	"errors"

	"github.com/arfanxn/welding/internal/infrastructure/database/helper"
	"github.com/arfanxn/welding/internal/module/role/domain/repository"
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	"github.com/arfanxn/welding/internal/module/shared/domain/errorx"
	"github.com/arfanxn/welding/pkg/pagination"
	"github.com/arfanxn/welding/pkg/query"
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
	roleTableName := entity.NewRole().TableName()

	if id := q.GetFilterById(); id != nil {
		db = db.Where(roleTableName+".id = ?", id.Value)
	}

	if search := q.GetSearch(); search != nil {
		db = db.Where(roleTableName+".name ILIKE ?", "%"+*search+"%")
	}

	if q.GetInclude("permissions") != nil {
		db = db.Preload("Permissions")
	}

	if q.GetInclude("users") != nil {
		db = db.Preload("Users")
	}

	if sort := q.GetSort("name"); sort != nil {
		db = db.Order(roleTableName + ".name " + sort.Order)
	}

	if sort := q.GetSort("created_at"); sort != nil {
		db = db.Order(roleTableName + ".created_at " + sort.Order)
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

func (r *GormRoleRepository) First(q *query.Query) (*entity.Role, error) {
	roles, err := r.Get(q)
	if err != nil {
		return nil, err
	}
	if len(roles) == 0 {
		return nil, errorx.ErrRoleNotFound
	}
	role := roles[0]
	return role, nil
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
		return nil, err
	}
	if len(roles) != len(ids) {
		return nil, errorx.ErrRolesNotFound
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
	err := r.db.Omit("Permissions").Save(role).Error
	if err != nil {
		if helper.IsPostgresDuplicateKeyError(err) {
			return errorx.ErrRoleAlreadyExists
		}
		return err
	}
	return nil
}

func (r *GormRoleRepository) SetDefault(role *entity.Role) error {
	// Set all roles to not default
	if err := r.db.Model(&entity.Role{}).Where("id != ?", role.Id).Update("is_default", false).Error; err != nil {
		return err
	}

	// Set the specified role as default
	role.IsDefault = true
	if err := r.db.Save(role).Error; err != nil {
		return err
	}

	return nil
}

func (r *GormRoleRepository) SaveMany(roles []*entity.Role) error {
	return r.db.CreateInBatches(roles, 100).Error
}

func (r *GormRoleRepository) Destroy(role *entity.Role) error {
	return r.db.Delete(role).Error
}
