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
