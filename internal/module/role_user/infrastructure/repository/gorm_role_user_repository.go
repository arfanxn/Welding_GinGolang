package repository

import (
	"github.com/arfanxn/welding/internal/module/role_user/domain/repository"
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	"gorm.io/gorm"
)

var _ repository.RoleUserRepository = (*GormRoleUserRepository)(nil)

type GormRoleUserRepository struct {
	db *gorm.DB
}

func NewGormRoleUserRepository(db *gorm.DB) repository.RoleUserRepository {
	return &GormRoleUserRepository{
		db: db,
	}
}

func (r *GormRoleUserRepository) Save(roleUser *entity.RoleUser) error {
	return r.db.Save(roleUser).Error
}

func (r *GormRoleUserRepository) SaveMany(roleUsers []*entity.RoleUser) error {
	return r.db.CreateInBatches(roleUsers, 100).Error
}

func (r *GormRoleUserRepository) DestroyByUserId(userId string) error {
	return r.db.Delete(&entity.RoleUser{}, "user_id = ?", userId).Error
}

func (r *GormRoleUserRepository) Destroy(roleUser *entity.RoleUser) error {
	return r.db.Delete(roleUser).Error
}

func (r *GormRoleUserRepository) DestroyMany(roleUsers []*entity.RoleUser) error {
	return r.db.Delete(roleUsers).Error
}
