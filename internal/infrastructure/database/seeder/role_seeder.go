package seeder

import (
	"github.com/arfanxn/welding/internal/infrastructure/database/factory"
	"github.com/arfanxn/welding/internal/module/permission/domain/enum"
	permissionRepository "github.com/arfanxn/welding/internal/module/permission/domain/repository"
	permissionRoleRepository "github.com/arfanxn/welding/internal/module/permission_role/domain/repository"
	"github.com/arfanxn/welding/internal/module/role/domain/repository"
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	"github.com/samber/lo"
	"go.uber.org/fx"
)

var _ Seeder = (*RoleSeeder)(nil)

type RoleSeeder struct {
	roleRepository           repository.RoleRepository
	permissionRepository     permissionRepository.PermissionRepository
	permissionRoleRepository permissionRoleRepository.PermissionRoleRepository
}

type NewRoleSeederParams struct {
	fx.In

	RoleRepository           repository.RoleRepository
	PermissionRepository     permissionRepository.PermissionRepository
	PermissionRoleRepository permissionRoleRepository.PermissionRoleRepository
}

func NewRoleSeeder(params NewRoleSeederParams) Seeder {
	return &RoleSeeder{
		roleRepository:           params.RoleRepository,
		permissionRepository:     params.PermissionRepository,
		permissionRoleRepository: params.PermissionRoleRepository,
	}
}

func (s *RoleSeeder) Seed() error {
	superAdminRole := factory.RoleFactory.MustCreateWithOption(map[string]any{
		"Name": "super admin",
	}).(*entity.Role)
	adminRole := factory.RoleFactory.MustCreateWithOption(map[string]any{
		"Name": "admin",
	}).(*entity.Role)
	headRole := factory.RoleFactory.MustCreateWithOption(map[string]any{
		"Name": "head",
	}).(*entity.Role)
	managerRole := factory.RoleFactory.MustCreateWithOption(map[string]any{
		"Name": "manager",
	}).(*entity.Role)
	supervisorRole := factory.RoleFactory.MustCreateWithOption(map[string]any{
		"Name": "supervisor",
	}).(*entity.Role)
	engineerRole := factory.RoleFactory.MustCreateWithOption(map[string]any{
		"Name": "engineer",
	}).(*entity.Role)
	staffRole := factory.RoleFactory.MustCreateWithOption(map[string]any{
		"Name": "staff",
	}).(*entity.Role)
	operatorRole := factory.RoleFactory.MustCreateWithOption(map[string]any{
		"Name": "operator",
	}).(*entity.Role)
	customerServiceAdminRole := factory.RoleFactory.MustCreateWithOption(map[string]any{
		"Name": "customer service admin",
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
	}

	// Save all roles to the database
	err := s.roleRepository.SaveMany(roles)
	if err != nil {
		return err
	}

	// Create a map of role names to role objects for easy lookup
	roleMap := make(map[string]*entity.Role)
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
			permissionRoles = append(permissionRoles, factory.PermissionRoleFactory.MustCreateWithOption(map[string]any{
				"RoleId":       superAdminRole.Id,
				"PermissionId": permission.Id,
			}).(*entity.PermissionRole))
		}

		// Define basic permissions that should be granted to all non-super-admin roles
		// These are the minimum permissions needed for basic system navigation
		exceptSuperAdminPermissionNames := []enum.PermissionName{
			enum.UserRead,       // Allow viewing user listings
			enum.RoleRead,       // Allow viewing role definitions
			enum.PermissionRead, // Allow viewing available permissions
		}

		// Filter the complete permission list to only include the basic permissions
		// This creates a subset of permissions that will be assigned to regular roles
		exceptSuperAdminPermissions := lo.Filter(permissions, func(permission *entity.Permission, _ int) bool {
			return lo.Contains(exceptSuperAdminPermissionNames, enum.PermissionName(permission.Name))
		})

		// Get all roles except super admin to assign basic permissions
		rolesExceptSuperAdmin := lo.Filter(roles, func(role *entity.Role, _ int) bool {
			return role.Name != superAdminRole.Name
		})

		// Assign the basic permissions to each non-super-admin role
		// This ensures all roles have at least the minimum required access
		for _, role := range rolesExceptSuperAdmin {
			for _, permission := range exceptSuperAdminPermissions {
				permissionRoles = append(permissionRoles, factory.PermissionRoleFactory.MustCreateWithOption(map[string]any{
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
