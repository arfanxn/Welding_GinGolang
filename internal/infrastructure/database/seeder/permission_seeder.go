package seeder

import (
	"github.com/arfanxn/welding/internal/infrastructure/id"
	"github.com/arfanxn/welding/internal/module/permission/domain/enum"
	"github.com/arfanxn/welding/internal/module/permission/domain/repository"
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	"go.uber.org/fx"
)

var _ Seeder = (*PermissionSeeder)(nil)

type PermissionSeeder struct {
	idService            id.IdService
	permissionRepository repository.PermissionRepository
}

type NewPermissionSeederParams struct {
	fx.In

	IdService            id.IdService
	PermissionRepository repository.PermissionRepository
}

func NewPermissionSeeder(
	params NewPermissionSeederParams,
) Seeder {
	return &PermissionSeeder{
		idService:            params.IdService,
		permissionRepository: params.PermissionRepository,
	}
}

func (s *PermissionSeeder) Seed() error {
	permissionNames := enum.PermissionNames
	var permissions []*entity.Permission

	for _, permissionName := range permissionNames {
		permissions = append(permissions, &entity.Permission{
			Id:   s.idService.Generate(),
			Name: permissionName,
		})
	}

	err := s.permissionRepository.SaveMany(permissions)
	if err != nil {
		return err
	}

	return nil
}
