package service

import (
	permissionEnum "github.com/arfanxn/welding/internal/module/permission/domain/enum"
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	"github.com/samber/lo"
)

type PermissionService interface {
	CheckByNames(permissions any, requiredPermNames ...permissionEnum.PermissionName) (bool, error)
}

type permissionService struct {
}

func NewPermissionService() PermissionService {
	return &permissionService{}
}

// CheckByNames checks if all required permission names exist in the provided permissions.
// It accepts permissions in different formats: []*entity.Permission, []entity.Permission, or []permissionEnum.PermissionName.
// Returns (true, nil) if all required permissions are found, (false, nil) if any permission is missing,
// and (false, error) if there's an error during the check.
func (s *permissionService) CheckByNames(permissionsAny any, requiredPermNames ...permissionEnum.PermissionName) (bool, error) {
	permissions := []*entity.Permission{}

	if ps, ok := permissionsAny.([]*entity.Permission); ok {
		permissions = ps
	} else if ps, ok := permissionsAny.([]entity.Permission); ok {
		for _, p := range ps {
			permissions = append(permissions, &p)
		}
	} else if pNames, ok := permissionsAny.([]permissionEnum.PermissionName); ok {
		for _, pName := range pNames {
			permissions = append(permissions, &entity.Permission{
				Name: pName,
			})
		}
	}

	for _, requiredPermName := range requiredPermNames {
		found := lo.ContainsBy(permissions, func(p *entity.Permission) bool {
			return p.Name == requiredPermName
		})
		if !found {
			return false, nil
		}
	}

	return true, nil
}
