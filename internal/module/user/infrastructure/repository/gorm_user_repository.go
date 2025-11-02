package repository

import (
	"strings"
	"time"

	"github.com/arfanxn/welding/internal/infrastructure/database/helper"
	permissionEnum "github.com/arfanxn/welding/internal/module/permission/domain/enum"
	roleEnum "github.com/arfanxn/welding/internal/module/role/domain/enum"
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	"github.com/arfanxn/welding/internal/module/user/domain/repository"
	"github.com/arfanxn/welding/pkg/pagination"
	"github.com/arfanxn/welding/pkg/query"
	"github.com/arfanxn/welding/pkg/reflectutil"
	"github.com/gookit/goutil"
	"github.com/guregu/null/v6"
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

// HasPermissionNames checks if a user has all the specified permissions.
// The permissions parameter is an array of permission names.
func (r *GormUserRepository) HasPermissionNames(user *entity.User, permissionNames []permissionEnum.PermissionName) (bool, error) {
	if len(permissionNames) == 0 {
		return true, nil
	}

	db := r.db.Model(&entity.User{}).
		Joins("JOIN role_user ON role_user.user_id = users.id").
		Joins("JOIN roles ON roles.id = role_user.role_id").
		Joins("JOIN permission_role ON permission_role.role_id = roles.id").
		Joins("JOIN permissions ON permissions.id = permission_role.permission_id").
		Where("users.id = ?", user.Id).
		Where("permissions.name IN (?)", permissionNames)

	var count int64
	err := db.Distinct("permission_role.permission_id").
		Count(&count).Error

	if err != nil {
		return false, err
	}

	// Only return true if the user has all the specified permissions
	return count == int64(len(permissionNames)), nil
}

// HasRoleNames checks if a user has all the specified roles.
// The roleNames parameter is an array of role names.
func (r *GormUserRepository) HasRoleNames(user *entity.User, roleNames []roleEnum.RoleName) (bool, error) {
	if len(roleNames) == 0 {
		return true, nil
	}

	db := r.db.Model(&entity.User{}).
		Joins("JOIN role_user ON role_user.user_id = users.id").
		Joins("JOIN roles ON roles.id = role_user.role_id").
		Where("users.id = ?", user.Id).
		Where("roles.name IN (?)", roleNames)

	var count int64
	err := db.Distinct("role_user.role_id").
		Count(&count).Error

	if err != nil {
		return false, err
	}

	// Only return true if the user has all the specified roles
	return count == int64(len(roleNames)), nil
}

func (r *GormUserRepository) ToggleActivation(user *entity.User) (*entity.User, error) {
	if user.ActivatedAt.Valid {
		user.ActivatedAt = null.TimeFromPtr(nil)
		user.DeactivatedAt = null.TimeFrom(time.Now())
	} else {
		user.ActivatedAt = null.TimeFrom(time.Now())
		user.DeactivatedAt = null.TimeFromPtr(nil)
	}

	if err := r.db.Save(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (r *GormUserRepository) Save(user *entity.User) error {
	// 1. Start database transaction
	tx := r.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	// 2. Defer a rollback in case anything fails
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		}
	}()

	// 3. Save user record (without roles and employee to prevent M2M race conditions)
	err := tx.Omit("Roles", "Employee").Save(user).Error
	if err != nil {
		tx.Rollback()
		if helper.IsPostgresDuplicateKeyError(err) {
			return gorm.ErrDuplicatedKey
		}
		return err
	}

	// 4. Handle employee association if it exists
	if !goutil.IsNil(user.Employee) {
		// 4.1 Set the UserID if not set
		if goutil.IsEmpty(user.Employee.UserId) {
			user.Employee.UserId = user.Id
		}

		// 4.2 Save the employee record
		if err := tx.Save(user.Employee).Error; err != nil {
			tx.Rollback()
			return err
		}

		// 4.3 Refresh the employee association
		if err := tx.Model(user).Association("Employee").Replace(user.Employee); err != nil {
			tx.Rollback()
			return err
		}
	} else {
		// 4.4 Delete the employee record if it doesn't exist
		tx.Model(&user).Association("Employee").Delete(&user.Employee)

	}

	// 5. Update roles if any provided
	if reflectutil.IsSlice(user.Roles) && len(user.Roles) > 0 {
		if err := tx.Model(user).Association("Roles").Replace(user.Roles); err != nil {
			tx.Rollback()
			return err
		}
	}

	// 6. Commit the transaction if everything is successful
	// Note: No need to refresh user as GORM's Save and Replace methods
	// already update the passed struct with the latest database values
	return tx.Commit().Error
}

func (r *GormUserRepository) SaveMany(users []*entity.User) error {
	return r.db.CreateInBatches(users, 100).Error
}

func (r *GormUserRepository) Destroy(user *entity.User) error {
	return r.db.Delete(user).Error
}
