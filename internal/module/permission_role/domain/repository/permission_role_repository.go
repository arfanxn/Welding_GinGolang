package repository

import "github.com/arfanxn/welding/internal/module/shared/domain/entity"

type PermissionRoleRepository interface {
	Save(permissionRole *entity.PermissionRole) error
	SaveMany(permissionRoles []*entity.PermissionRole) error
}
