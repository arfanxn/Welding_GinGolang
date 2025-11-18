package repository

import (
	"github.com/arfanxn/welding/internal/infrastructure/database/helper"
	"github.com/arfanxn/welding/internal/module/permission_role/domain/repository"
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	"github.com/arfanxn/welding/internal/module/shared/domain/errorx"
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
	err := r.db.Save(permissionRole).Error
	if err != nil {
		if helper.IsPostgresDuplicateKeyError(err) {
			return errorx.ErrPermissionRoleAlreadyExists
		}
		return err
	}

	return nil
}

func (r *GormPermissionRoleRepository) SaveMany(permissionRoles []*entity.PermissionRole) error {
	return r.db.CreateInBatches(permissionRoles, 100).Error
}

func (r *GormPermissionRoleRepository) DestroyByRoleId(roleId string) error {
	return r.db.Where("role_id = ?", roleId).Delete(&entity.PermissionRole{}).Error
}

func (r *GormPermissionRoleRepository) Destroy(permissionRole *entity.PermissionRole) error {
	return r.db.Delete(permissionRole).Error
}
