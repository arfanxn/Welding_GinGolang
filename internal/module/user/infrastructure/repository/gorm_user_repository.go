package repository

import (
	"strings"

	"github.com/arfanxn/welding/internal/infrastructure/database/helper"
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	"github.com/arfanxn/welding/internal/module/user/domain/repository"
	"github.com/arfanxn/welding/pkg/pagination"
	"github.com/arfanxn/welding/pkg/query"
	"github.com/arfanxn/welding/pkg/reflectutil"
	"gorm.io/gorm"
)

var _ repository.UserRepository = (*GormUserRepository)(nil)

type GormUserRepository struct {
	db *gorm.DB
}

func NewGormUserRepository(db *gorm.DB) repository.UserRepository {
	return &GormUserRepository{
		db: db,
	}
}

func (r *GormUserRepository) query(db *gorm.DB, q *query.Query) *gorm.DB {
	userTableName := entity.NewUser().TableName()
	employeeTableName := entity.NewEmployee().TableName()

	var sb strings.Builder
	sb.WriteString(userTableName + ".*")
	sb.WriteString(", ")
	sb.WriteString(employeeTableName + ".employment_identity_number")

	db = db.Select(sb.String())

	if f := q.GetFilterById(); f != nil {
		db = db.Where("id = ?", f.Value)
	}

	if search := q.GetSearch(); !search.IsZero() {
		db = db.Where("name LIKE ?", "%"+search.String+"%")
	}

	if !q.GetInclude("roles").IsZero() {
		db = db.Preload("Roles")
	}

	if !q.GetInclude("roles.permissions").IsZero() {
		db = db.Preload("Roles.Permissions")
	}

	if sort := q.GetSort("name"); sort != nil {
		db = db.Order("name " + sort.Order)
	}

	if sort := q.GetSort("created_at"); sort != nil {
		db = db.Order("created_at " + sort.Order)
	}

	db = db.Joins("LEFT JOIN " + employeeTableName + " ON " + userTableName + ".id = " + employeeTableName + ".user_id")

	return db
}

func (r *GormUserRepository) Get(q *query.Query) ([]*entity.User, error) {
	var users []*entity.User

	db := r.query(r.db, q)

	if err := db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (r *GormUserRepository) Paginate(q *query.Query) (*pagination.OffsetPagination[*entity.User], error) {
	db := r.db.Model(&entity.User{})

	db = r.query(db, q)

	paginator, err := helper.GormDBPaginateWithQuery[*entity.User](db, q)
	if err != nil {
		return nil, err
	}
	return paginator, nil
}

func (r *GormUserRepository) Find(id string) (*entity.User, error) {
	var user entity.User

	if err := r.db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *GormUserRepository) FindByEmail(email string) (*entity.User, error) {
	var user entity.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *GormUserRepository) Save(user *entity.User) error {
	// Start transaction
	tx := r.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	// Save role record (without permissions to prevent M2M race conditions)
	err := tx.Omit("Roles").Save(user).Error
	if err != nil {
		tx.Rollback()
		if helper.IsPostgresDuplicateKeyError(err) {
			return gorm.ErrDuplicatedKey
		}
		return err
	}

	// Update permissions if any provided
	if reflectutil.IsSlice(user.Roles) {
		// Replace all role-permission associations
		// Note: role.Permissions should contain only Permission{ID: X} structs
		if err := tx.Model(user).Association("Roles").Replace(user.Roles); err != nil {
			tx.Rollback()
			return err
		}
	}

	// Commit transaction
	return tx.Commit().Error
}

func (r *GormUserRepository) SaveMany(users []*entity.User) error {
	return r.db.CreateInBatches(users, 100).Error
}

func (r *GormUserRepository) Destroy(user *entity.User) error {
	return r.db.Delete(user).Error
}
