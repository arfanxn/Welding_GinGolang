// Package policy provides business rule validation and authorization logic for role management operations.
// It enforces constraints and permissions before allowing role-related actions to be executed.
package policy

import (
	"context"

	permissionRepository "github.com/arfanxn/welding/internal/module/permission/domain/repository"
	"github.com/arfanxn/welding/internal/module/role/domain/enum"
	roleRepository "github.com/arfanxn/welding/internal/module/role/domain/repository"
	roleDto "github.com/arfanxn/welding/internal/module/role/usecase/dto"
	"github.com/arfanxn/welding/internal/module/shared/domain/errorx"
	"github.com/gookit/goutil"
	"go.uber.org/fx"
)

// RolePolicy defines the contract for role-related business rule validation.
// It provides methods to validate role operations before they are executed by the use case layer.
type RolePolicy interface {
	Store(ctx context.Context, _dto *roleDto.SaveRole) error
	Update(ctx context.Context, _dto *roleDto.SaveRole) error
	SetDefault(ctx context.Context, _dto *roleDto.SetDefaultRole) error
	Destroy(ctx context.Context, _dto *roleDto.DestroyRole) error
}

// rolePolicy implements the RolePolicy interface with concrete business rule validation logic.
// It uses repositories to access role and permission data for validation purposes.
type rolePolicy struct {
	// roleRepository provides access to role data for validation operations
	roleRepository roleRepository.RoleRepository
	// permissionRepository provides access to permission data for validation operations
	permissionRepository permissionRepository.PermissionRepository
}

// NewRolePolicyParams defines the dependency injection parameters for creating a new RolePolicy instance.
// It uses fx.In for automatic dependency injection with the Uber FX framework.
type NewRolePolicyParams struct {
	fx.In

	// RoleRepository dependency for role data access
	RoleRepository roleRepository.RoleRepository
	// PermissionRepository dependency for permission data access
	PermissionRepository permissionRepository.PermissionRepository
}

// NewRolePolicy creates a new instance of rolePolicy with the provided dependencies.
// It implements the RolePolicy interface and is typically used with dependency injection.
func NewRolePolicy(params NewRolePolicyParams) RolePolicy {
	return &rolePolicy{
		roleRepository:       params.RoleRepository,
		permissionRepository: params.PermissionRepository,
	}
}

// Store validates the business rules for creating a new role.
// It prevents creation of restricted roles and validates permission assignments.
// Returns an error if any validation rule is violated.
func (p *rolePolicy) Store(ctx context.Context, _dto *roleDto.SaveRole) error {
	// Prevent creation of SuperAdmin role as it's a system-managed role
	if *_dto.Name == enum.SuperAdmin {
		return errorx.ErrRoleSuperAdminStoreForbidden
	}

	// Validate that all specified permission IDs exist and are assignable
	if err := p.validatePermissionAssignments(_dto.PermissionIds); err != nil {
		return err
	}

	// All validations passed, role can be created
	return nil
}

// Update validates the business rules for updating an existing role.
// It verifies the role exists, prevents modification of restricted roles,
// and validates permission assignments.
// Returns an error if any validation rule is violated.
func (p *rolePolicy) Update(ctx context.Context, _dto *roleDto.SaveRole) error {
	// Verify that the role exists before attempting to update it
	role, err := p.roleRepository.Find(*_dto.Id)
	if err != nil {
		return err
	}

	// Prevent modification of SuperAdmin role as it's a system-managed role
	if role.IsSuperAdmin() {
		return errorx.ErrRoleSuperAdminUpdateForbidden
	}

	// Validate that all specified permission IDs exist and are assignable
	if err := p.validatePermissionAssignments(_dto.PermissionIds); err != nil {
		return err
	}

	// All validations passed, role can be updated
	return nil
}

// SetDefault validates the business rules for setting a role as the default role.
// It verifies the role exists, prevents setting already default roles,
// and prevents setting restricted system roles as default.
// Returns an error if any validation rule is violated.
func (p *rolePolicy) SetDefault(ctx context.Context, _dto *roleDto.SetDefaultRole) error {
	// Verify that the role exists before attempting to set it as default
	role, err := p.roleRepository.Find(_dto.Id)
	if err != nil {
		return err
	}

	// Prevent setting a role that's already the default (idempotent operation check)
	if role.IsDefault {
		return errorx.ErrRoleAlreadyDefault
	}
	// Prevent setting SuperAdmin role as default
	if role.IsSuperAdmin() {
		return errorx.ErrRoleSuperAdminSetDefaultForbidden
	}

	// All validations passed, role can be set as default
	return nil
}

// Destroy validates the business rules for deleting a role.
// It verifies the role exists, prevents deletion of default roles,
// and prevents deletion of restricted system roles.
// Returns an error if any validation rule is violated.
func (p *rolePolicy) Destroy(ctx context.Context, _dto *roleDto.DestroyRole) error {
	// Verify that the role exists before attempting to delete it
	role, err := p.roleRepository.Find(_dto.Id)
	if err != nil {
		return err
	}

	// Prevent deletion of default roles to maintain system integrity
	if role.IsDefault {
		return errorx.ErrRoleDefaultDestroyForbidden
	}

	// Prevent deletion of system-managed role (like SuperAdmin)
	if role.IsSuperAdmin() {
		return errorx.ErrRoleSuperAdminDestroyForbidden
	}

	// All validations passed, role can be deleted
	return nil
}

// ==================================================
// Private helper methods
// ==================================================

// validatePermissionAssignments validates that all specified permission IDs exist in the system.
// This ensures that permission assignments are valid and prevents orphaned references.
// Returns an error if any permission ID is not found, or nil if all are valid.
func (p *rolePolicy) validatePermissionAssignments(permissionIds []string) error {
	// Skip validation if no permissions are specified (optional assignment)
	if goutil.IsEmpty(permissionIds) {
		return nil
	}

	if goutil.IsEmpty(permissionIds[0]) {
		return nil
	}

	// Attempt to fetch all specified permissions to verify they exist
	_, err := p.permissionRepository.FindByIds(permissionIds)
	if err != nil {
		return err
	}

	// All permission IDs are valid
	return nil
}
