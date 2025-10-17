package seeder

import (
	"time"

	"github.com/arfanxn/welding/internal/module/permission/domain/enum"
	"github.com/arfanxn/welding/internal/module/permission/domain/repository"
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	"github.com/guregu/null/v6"
	"github.com/oklog/ulid/v2"
)

var _ Seeder = (*PermissionSeeder)(nil)

type PermissionSeeder struct {
	permissionRepository repository.PermissionRepository
}

func NewPermissionSeeder(permissionRepository repository.PermissionRepository) Seeder {
	return &PermissionSeeder{permissionRepository: permissionRepository}
}

func (s *PermissionSeeder) Seed() error {
	permissionNames := enum.PermissionNames
	var permissions []*entity.Permission

	for _, permissionName := range permissionNames {
		permissions = append(permissions, &entity.Permission{
			Id:        ulid.Make().String(),
			Name:      permissionName.String(),
			CreatedAt: time.Now(),
			UpdatedAt: null.Time{},
			DeletedAt: null.Time{},
		})
	}

	err := s.permissionRepository.SaveMany(permissions)
	if err != nil {
		return err
	}

	return nil
}
