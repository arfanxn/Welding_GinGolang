package repository

import (
	"github.com/arfanxn/welding/internal/module/permission_role/domain/repository"
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	"gorm.io/gorm"
)

var _ repository.PermissionRoleRepository = (*GormPermissionRoleRepository)(nil)

type GormPermissionRoleRepository struct {
	db *gorm.DB
}

func NewGormPermissionRoleRepository(db *gorm.DB) repository.PermissionRoleRepository {
	return &GormPermissionRoleRepository{db: db}
}

func (r *GormPermissionRoleRepository) Save(permissionRole *entity.PermissionRole) error {
	return r.db.Save(permissionRole).Error
}

func (r *GormPermissionRoleRepository) SaveMany(permissionRoles []*entity.PermissionRole) error {
	return r.db.CreateInBatches(permissionRoles, 100).Error
}
