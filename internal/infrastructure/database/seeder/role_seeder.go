package seeder

import (
	"github.com/arfanxn/welding/internal/infrastructure/id"
	permissionEnum "github.com/arfanxn/welding/internal/module/permission/domain/enum"
	permissionRepository "github.com/arfanxn/welding/internal/module/permission/domain/repository"
	permissionRoleRepository "github.com/arfanxn/welding/internal/module/permission_role/domain/repository"
	"github.com/arfanxn/welding/internal/module/role/domain/enum"
	"github.com/arfanxn/welding/internal/module/role/domain/repository"
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	"github.com/bluele/factory-go/factory"
	"github.com/samber/lo"
	"go.uber.org/fx"
)

var _ Seeder = (*RoleSeeder)(nil)

type RoleSeeder struct {
	idService                id.IdService
	roleFactory              *factory.Factory
	permissionRoleFactory    *factory.Factory
	roleRepository           repository.RoleRepository
	permissionRepository     permissionRepository.PermissionRepository
	permissionRoleRepository permissionRoleRepository.PermissionRoleRepository
}

type NewRoleSeederParams struct {
	fx.In

	IdService                id.IdService
	RoleFactory              *factory.Factory `name:"role_factory"`
	PermissionRoleFactory    *factory.Factory `name:"permission_role_factory"`
	RoleRepository           repository.RoleRepository
	PermissionRepository     permissionRepository.PermissionRepository
	PermissionRoleRepository permissionRoleRepository.PermissionRoleRepository
}

func NewRoleSeeder(params NewRoleSeederParams) Seeder {
	return &RoleSeeder{
		idService:                params.IdService,
		roleFactory:              params.RoleFactory,
		permissionRoleFactory:    params.PermissionRoleFactory,
		roleRepository:           params.RoleRepository,
		permissionRepository:     params.PermissionRepository,
		permissionRoleRepository: params.PermissionRoleRepository,
	}
}

func (s *RoleSeeder) Seed() error {
	roleFactory := s.roleFactory
	permissionRoleFactory := s.permissionRoleFactory

	superAdminRole := roleFactory.MustCreateWithOption(map[string]any{
		"Id":   "01K7KJ8RYMGS1FWJRZD000TCW0",
		"Name": enum.SuperAdmin,
	}).(*entity.Role)
	adminRole := roleFactory.MustCreateWithOption(map[string]any{
		"Id":   "01K7KJ8RYN8QJMYKA1WNBHQ22K",
		"Name": enum.Admin,
	}).(*entity.Role)
	headRole := roleFactory.MustCreateWithOption(map[string]any{
		"Id":   "01K7KJ8RYN8QJMYKA1WQQQYZA6",
		"Name": enum.Head,
	}).(*entity.Role)
	managerRole := roleFactory.MustCreateWithOption(map[string]any{
		"Id":   "01K7KJ8RYN8QJMYKA1WTZNZVYT",
		"Name": enum.Manager,
	}).(*entity.Role)
	supervisorRole := roleFactory.MustCreateWithOption(map[string]any{
		"Id":   "01K7KJ8RYN8QJMYKA1WYFRKSYN",
		"Name": enum.Supervisor,
	}).(*entity.Role)
	engineerRole := roleFactory.MustCreateWithOption(map[string]any{
		"Id":   "01K7KJ8RYN8QJMYKA1X0VNR31G",
		"Name": enum.Engineer,
	}).(*entity.Role)
	staffRole := roleFactory.MustCreateWithOption(map[string]any{
		"Id":   "01K7KJ8RYN8QJMYKA1X405T2FF",
		"Name": enum.Staff,
	}).(*entity.Role)
	operatorRole := roleFactory.MustCreateWithOption(map[string]any{
		"Id":   "01K7KJ8RYN8QJMYKA1X78V9C0T",
		"Name": enum.Operator,
	}).(*entity.Role)
	customerServiceAdminRole := roleFactory.MustCreateWithOption(map[string]any{
		"Id":   "01K7KJ8RYN8QJMYKA1X7GKVWKP",
		"Name": enum.CustomerServiceAdmin,
	}).(*entity.Role)
	customerRole := roleFactory.MustCreateWithOption(map[string]any{
		"Id":        "01K8SW84Z6T1FM90KES0G0BMM1",
		"IsDefault": true, // Set as default role
		"Name":      enum.DefaultRoleName,
	}).(*entity.Role)

	roles := []*entity.Role{
		superAdminRole,
		adminRole,
		headRole,
		managerRole,
		supervisorRole,
		engineerRole,
		staffRole,
		operatorRole,
		customerServiceAdminRole,
		customerRole,
	}

	// Save all roles to the database
	err := s.roleRepository.SaveMany(roles)
	if err != nil {
		return err
	}

	// Create a map of role names to role objects for easy lookup
	roleMap := make(map[enum.RoleName]*entity.Role)
	for _, role := range roles {
		roleMap[role.Name] = role
	}

	{
		// ========== PermissionRole ==========
		// This section handles the assignment of permissions to roles.
		// Each role is granted specific permissions based on their access level.

		// Fetch all available permissions from the database
		permissions, err := s.permissionRepository.All()
		if err != nil {
			return err
		}

		// Initialize slice to store all permission-role relationships that will be created
		var permissionRoles []*entity.PermissionRole

		// Super admin gets all permissions without any restrictions
		superAdminPermissions := permissions

		// Assign all available permissions to super admin role
		// This ensures super admins have full access to all features
		for _, permission := range superAdminPermissions {
			permissionRoles = append(permissionRoles, permissionRoleFactory.MustCreateWithOption(map[string]any{
				"RoleId":       superAdminRole.Id,
				"PermissionId": permission.Id,
			}).(*entity.PermissionRole))
		}

		// Define basic permissions that should be granted to all non-super-admin roles
		// These are the minimum permissions needed for basic system navigation
		exceptSuperAdminPermissionNames := []permissionEnum.PermissionName{
			permissionEnum.UsersIndex,
			permissionEnum.RolesIndex,
			permissionEnum.PermissionsIndex,
		}

		// Filter the complete permission list to only include the basic permissions
		// This creates a subset of permissions that will be assigned to regular roles
		exceptSuperAdminPermissions := lo.Filter(permissions, func(permission *entity.Permission, _ int) bool {
			return lo.Contains(exceptSuperAdminPermissionNames, permissionEnum.PermissionName(permission.Name))
		})

		// Get all roles except super admin to assign basic permissions
		rolesExceptSuperAdmin := lo.Filter(roles, func(role *entity.Role, _ int) bool {
			return role.Name != superAdminRole.Name
		})

		// Assign the basic permissions to each non-super-admin role
		// This ensures all roles have at least the minimum required access
		for _, role := range rolesExceptSuperAdmin {
			for _, permission := range exceptSuperAdminPermissions {
				permissionRoles = append(permissionRoles, permissionRoleFactory.MustCreateWithOption(map[string]any{
					"RoleId":       role.Id,
					"PermissionId": permission.Id,
				}).(*entity.PermissionRole))
			}
		}

		// Save all permission-role relationships to the database in a single batch
		// This is more efficient than saving each relationship individually
		err = s.permissionRoleRepository.SaveMany(permissionRoles)
		if err != nil {
			return err
		}
	}

	return nil
}
