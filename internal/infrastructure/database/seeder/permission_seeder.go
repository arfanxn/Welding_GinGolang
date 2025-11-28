package seeder

import (
	"fmt"

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
	permissionDescriptions := map[enum.PermissionName]string{
		enum.UsersIndex:   "Melihat daftar user",
		enum.UsersShow:    "Melihat detail user",
		enum.UsersStore:   "Menambahkan user baru",
		enum.UsersUpdate:  "Memperbarui user",
		enum.UsersDestroy: "Menghapus user",

		enum.RolesIndex:   "Melihat daftar role",
		enum.RolesShow:    "Melihat detail role",
		enum.RolesStore:   "Menambahkan role baru",
		enum.RolesUpdate:  "Memperbarui role",
		enum.RolesDestroy: "Menghapus role",

		enum.PermissionsIndex: "Melihat daftar permission",
		enum.PermissionsShow:  "Melihat detail permission",

		enum.ActivitiesIndex: "Melihat daftar activity",
		enum.ActivitiesShow:  "Melihat detail activity",

		enum.MaterialTestMethodsIndex:   "Melihat daftar material test method",
		enum.MaterialTestMethodsShow:    "Melihat detail material test method",
		enum.MaterialTestMethodsStore:   "Menambahkan material test method baru",
		enum.MaterialTestMethodsUpdate:  "Memperbarui material test method",
		enum.MaterialTestMethodsDestroy: "Menghapus material test method",

		enum.MaterialTestMachinesIndex:   "Melihat daftar material test machine",
		enum.MaterialTestMachinesShow:    "Melihat detail material test machine",
		enum.MaterialTestMachinesStore:   "Menambahkan material test machine baru",
		enum.MaterialTestMachinesUpdate:  "Memperbarui material test machine",
		enum.MaterialTestMachinesDestroy: "Menghapus material test machine",

		enum.MaterialTestServicesIndex:   "Melihat daftar material test service",
		enum.MaterialTestServicesShow:    "Melihat detail material test service",
		enum.MaterialTestServicesStore:   "Menambahkan material test service baru",
		enum.MaterialTestServicesUpdate:  "Memperbarui material test service",
		enum.MaterialTestServicesDestroy: "Menghapus material test service",
	}

	var permissions []*entity.Permission

	for _, permissionName := range permissionNames {
		permissionDescription, ok := permissionDescriptions[permissionName]
		if !ok {
			return fmt.Errorf("permission description not configured for permission name: %s", permissionName)
		}

		permissions = append(permissions, &entity.Permission{
			Id:          s.idService.Generate(),
			Name:        permissionName,
			Description: permissionDescription,
		})
	}

	err := s.permissionRepository.SaveMany(permissions)
	if err != nil {
		return err
	}

	return nil
}
