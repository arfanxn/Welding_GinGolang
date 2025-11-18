package policy

import (
	"context"

	"github.com/arfanxn/welding/internal/infrastructure/security"
	roleEnum "github.com/arfanxn/welding/internal/module/role/domain/enum"
	roleRepository "github.com/arfanxn/welding/internal/module/role/domain/repository"
	"github.com/arfanxn/welding/internal/module/shared/contextkey"
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	"github.com/arfanxn/welding/internal/module/shared/domain/errorx"
	"github.com/arfanxn/welding/internal/module/user/domain/repository"
	"github.com/arfanxn/welding/internal/module/user/usecase/dto"
	"github.com/gookit/goutil"
	"go.uber.org/fx"
)

type UserPolicy interface {
	Store(ctx context.Context, _dto *dto.SaveUser) error
	Update(ctx context.Context, _dto *dto.SaveUser) error
	UpdateMePassword(ctx context.Context, _dto *dto.UpdateUserMePassword) error
	// ! Deprecated
	// UpdatePassword(ctx context.Context, _dto *dto.UpdateUserPassword) (*entity.User, error)
	ToggleActivation(ctx context.Context, _dto *dto.ToggleActivation) error
	Destroy(ctx context.Context, _dto *dto.DestroyUser) error
}

type userPolicy struct {
	passwordService security.PasswordService

	userRepository repository.UserRepository
	roleRepository roleRepository.RoleRepository
}

type NewUserPolicyParams struct {
	fx.In

	PasswordService security.PasswordService

	UserRepository repository.UserRepository
	RoleRepository roleRepository.RoleRepository
}

func NewUserPolicy(params NewUserPolicyParams) UserPolicy {
	return &userPolicy{
		passwordService: params.PasswordService,

		userRepository: params.UserRepository,
		roleRepository: params.RoleRepository,
	}
}

func (p *userPolicy) Store(ctx context.Context, _dto *dto.SaveUser) error {
	if err := p.validateRoleAssignments(_dto.RoleIds); err != nil {
		return err
	}
	return nil
}

func (p *userPolicy) Update(ctx context.Context, _dto *dto.SaveUser) error {
	authUser := ctx.Value(contextkey.UserKey).(*entity.User)

	targetUser, err := p.findUser(*_dto.Id)
	if err != nil {
		return err
	}

	isTargetUserSuperAdmin, err := p.isSuperAdmin(targetUser)
	if err != nil {
		return err
	}

	if isTargetUserSuperAdmin {
		// Only allow self-updates for SuperAdmins
		if authUser.Id != targetUser.Id {
			return errorx.ErrUserSuperAdminUpdateForbidden
		}

		// Prevent role changes for SuperAdmins
		if _dto.RoleIds != nil {
			return errorx.ErrUserSuperAdminRoleChangeForbidden
		}
	}

	if err := p.validateRoleAssignments(_dto.RoleIds); err != nil {
		return err
	}
	return nil
}

func (p *userPolicy) UpdateMePassword(ctx context.Context, _dto *dto.UpdateUserMePassword) error {
	authUser := ctx.Value(contextkey.UserKey).(*entity.User)

	if err := p.passwordService.Check(authUser.Password, _dto.CurrentPassword); err != nil {
		return errorx.ErrUserPasswordIncorrect
	}

	return nil
}

/*
! Deprecated
// UpdatePassword checks if the authenticated user is authorized to update another user's password.
func (p *userPolicy) UpdatePassword(
	ctx context.Context,
	_dto *dto.UpdateUserPassword,
) (*entity.User, error) {
	authUser := ctx.Value(contextkey.UserKey).(*entity.User)

	user, err := p.findUser(_dto.Id)
	if err != nil {
		return nil, err
	}

	isSuperAdmin, err := p.isSuperAdmin(user)
	if err != nil {
		return nil, err
	}

	if isSuperAdmin && authUser.Id != user.Id {
		return nil, errorx.ErrUserSuperAdminUpdateForbidden
	}

	return user, nil
}
*/

func (p *userPolicy) ToggleActivation(ctx context.Context, _dto *dto.ToggleActivation) error {
	user, err := p.findUser(_dto.Id)
	if err != nil {
		return err
	}

	isSuperAdmin, err := p.isSuperAdmin(user)
	if err != nil {
		return err
	}

	if isSuperAdmin {
		return errorx.ErrUserSuperAdminUpdateForbidden
	}

	return nil
}

// Destroy validates if a user can be deleted based on certain business rules
func (p *userPolicy) Destroy(_ context.Context, _dto *dto.DestroyUser) error {
	user, err := p.findUser(_dto.Id)
	if err != nil {
		return err
	}

	isSuperAdmin, err := p.isSuperAdmin(user)
	if err != nil {
		return err
	}

	if isSuperAdmin {
		return errorx.ErrUserSuperAdminUpdateForbidden
	}

	return nil
}

// ==================================================
// Private helper methods
// ==================================================

func (p *userPolicy) findUser(userId string) (*entity.User, error) {
	return p.userRepository.Find(userId)
}

func (p *userPolicy) isSuperAdmin(user *entity.User) (bool, error) {
	return p.userRepository.HasRoleNames(user, []roleEnum.RoleName{roleEnum.SuperAdmin})
}

func (p *userPolicy) validateRoleAssignments(roleIDs []string) error {
	if goutil.IsEmpty(roleIDs) {
		return nil
	}

	if goutil.IsEmpty(roleIDs[0]) {
		return nil
	}

	roles, err := p.roleRepository.FindByIds(roleIDs)
	if err != nil {
		return err
	}

	for _, role := range roles {
		if role.Name == roleEnum.SuperAdmin {
			return errorx.ErrUserSuperAdminAssignmentForbidden
		}
	}

	return nil
}
