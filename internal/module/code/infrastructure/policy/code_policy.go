package policy

import (
	"context"

	codeRepository "github.com/arfanxn/welding/internal/module/code/domain/repository"
	"github.com/arfanxn/welding/internal/module/code/usecase/dto"
	roleRepository "github.com/arfanxn/welding/internal/module/role/domain/repository"
	"github.com/arfanxn/welding/internal/module/shared/domain/errorx"
	"go.uber.org/fx"
)

type CodePolicy interface {
	CreateUserRegisterInvitation(ctx context.Context, _dto *dto.CreateUserRegisterInvitation) error
}

type codePolicy struct {
	codeRepository codeRepository.CodeRepository
	roleRepository roleRepository.RoleRepository
}

type NewCodePolicyParams struct {
	fx.In

	CodeRepository codeRepository.CodeRepository
	RoleRepository roleRepository.RoleRepository
}

func NewCodePolicy(params NewCodePolicyParams) CodePolicy {
	return &codePolicy{
		codeRepository: params.CodeRepository,
		roleRepository: params.RoleRepository,
	}
}

// CreateUserRegisterInvitation validates the user registration invitation request.
// It performs the following validations:
// 1. Checks if the specified role exists in the system
// 2. Ensures the role is not a super admin role (super admin invitations are not allowed)
//
// Parameters:
//   - ctx: Context for request-scoped values, cancellation signals, and deadlines
//   - _dto: Data transfer object containing the role ID for the invitation
//
// Returns:
//   - error: Returns nil if validation passes, otherwise returns an appropriate HTTP error
func (p *codePolicy) CreateUserRegisterInvitation(ctx context.Context, _dto *dto.CreateUserRegisterInvitation) error {
	// Retrieve the role to validate its existence and type
	role, err := p.roleRepository.Find(_dto.RoleId)
	if err != nil {
		return err
	}

	// Prevent creating invitations for super admin roles
	if role.IsSuperAdmin() {
		return errorx.ErrUserSuperAdminAssignmentForbidden
	}

	// Return nil if all validations pass
	return nil
}
