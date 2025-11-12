package policy

import (
	"context"
	"errors"
	"net/http"

	roleEnum "github.com/arfanxn/welding/internal/module/role/domain/enum"
	roleRepository "github.com/arfanxn/welding/internal/module/role/domain/repository"
	"github.com/arfanxn/welding/internal/module/shared/contextkey"
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	"github.com/arfanxn/welding/internal/module/user/domain/repository"
	"github.com/arfanxn/welding/internal/module/user/usecase/dto"
	"github.com/arfanxn/welding/pkg/errorutil"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

type UserPolicy interface {
	Save(ctx context.Context, _dto *dto.SaveUser) (*entity.User, error)
	UpdatePassword(ctx context.Context, _dto *dto.UpdateUserPassword) (*entity.User, error)
	ToggleActivation(ctx context.Context, _dto *dto.ToggleActivation) (*entity.User, error)
	Destroy(ctx context.Context, _dto *dto.DestroyUser) (*entity.User, error)
}

type userPolicy struct {
	userRepository repository.UserRepository
	roleRepository roleRepository.RoleRepository
}

type NewUserPolicyParams struct {
	fx.In

	UserRepository repository.UserRepository
	RoleRepository roleRepository.RoleRepository
}

func NewUserPolicy(params NewUserPolicyParams) UserPolicy {
	return &userPolicy{
		userRepository: params.UserRepository,
		roleRepository: params.RoleRepository,
	}
}

// Save handles the policy checks and validation for saving a user.
// It ensures proper authorization and business rules are followed when creating or updating a user.
// Returns the user entity if all validations pass, or an error if any check fails.
func (p *userPolicy) Save(ctx context.Context, _dto *dto.SaveUser) (*entity.User, error) {
	authUser := ctx.Value(contextkey.UserKey).(*entity.User)

	user, err := p.getUserForSave(_dto)
	if err != nil {
		return nil, err
	}

	isSuperAdmin, err := p.checkSuperAdminStatus(user)
	if err != nil {
		return nil, err
	}

	if isSuperAdmin {
		if err := p.validateSuperAdminUpdate(authUser, user, _dto); err != nil {
			return nil, err
		}
	}

	if len(_dto.RoleIds) > 0 {
		if err := p.validateRoleAssignments(_dto.RoleIds); err != nil {
			return nil, err
		}
	}

	return user, nil
}

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

	isSuperAdmin, err := p.checkSuperAdminStatus(user)
	if err != nil {
		return nil, err
	}

	if isSuperAdmin && authUser.Id != user.Id {
		return nil, p.superAdminForbiddenError()
	}

	return user, nil
}

func (p *userPolicy) ToggleActivation(ctx context.Context, _dto *dto.ToggleActivation) (*entity.User, error) {
	user, err := p.findUser(_dto.Id)
	if err != nil {
		return nil, err
	}

	isSuperAdmin, err := p.checkSuperAdminStatus(user)
	if err != nil {
		return nil, err
	}

	if isSuperAdmin {
		return nil, p.superAdminForbiddenError()
	}

	return user, nil
}

// Destroy validates if a user can be deleted based on certain business rules
func (p *userPolicy) Destroy(_ context.Context, _dto *dto.DestroyUser) (*entity.User, error) {
	user, err := p.findUser(_dto.Id)
	if err != nil {
		return nil, err
	}

	isSuperAdmin, err := p.checkSuperAdminStatus(user)
	if err != nil {
		return nil, err
	}

	if isSuperAdmin {
		return nil, errorutil.NewHttpError(
			http.StatusForbidden,
			"User dengan role "+string(roleEnum.SuperAdmin)+" tidak dapat dihapus",
			nil,
		)
	}

	return user, nil
}

// Private helper methods

func (p *userPolicy) getUserForSave(_dto *dto.SaveUser) (*entity.User, error) {
	if _dto.Id.IsZero() {
		return entity.NewUser(), nil
	}

	user, err := p.findUser(_dto.Id.String)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (p *userPolicy) findUser(userId string) (*entity.User, error) {
	user, err := p.userRepository.Find(userId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorutil.NewHttpError(http.StatusNotFound, "User tidak ditemukan", nil)
		}
		return nil, err
	}
	return user, nil
}

func (p *userPolicy) checkSuperAdminStatus(user *entity.User) (bool, error) {
	return p.userRepository.HasRoleNames(user, []roleEnum.RoleName{roleEnum.SuperAdmin})
}

func (p *userPolicy) validateSuperAdminUpdate(authUser, targetUser *entity.User, _dto *dto.SaveUser) error {
	// Only allow self-updates for SuperAdmins
	if authUser.Id != targetUser.Id {
		return p.superAdminForbiddenError()
	}

	// Prevent role changes for SuperAdmins
	if _dto.RoleIds != nil {
		return errorutil.NewHttpError(
			http.StatusForbidden,
			"User dengan role "+string(roleEnum.SuperAdmin)+" tidak dapat diubah role",
			nil,
		)
	}

	return nil
}

func (p *userPolicy) validateRoleAssignments(roleIDs []string) error {
	roles, err := p.roleRepository.FindByIds(roleIDs)
	if err != nil {
		return err
	}

	if len(roles) != len(roleIDs) {
		return errorutil.NewHttpError(
			http.StatusBadRequest,
			"Satu atau lebih role tidak ditemukan",
			nil,
		)
	}

	for _, role := range roles {
		if role.Name == roleEnum.SuperAdmin {
			return errorutil.NewHttpError(
				http.StatusForbidden,
				"Role "+string(roleEnum.SuperAdmin)+" tidak dapat ditambahkan ke user",
				nil,
			)
		}
	}

	return nil
}

func (p *userPolicy) superAdminForbiddenError() error {
	return errorutil.NewHttpError(
		http.StatusForbidden,
		"User dengan role "+string(roleEnum.SuperAdmin)+" tidak dapat diubah",
		nil,
	)
}
