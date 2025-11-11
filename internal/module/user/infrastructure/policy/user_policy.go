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
	"github.com/gookit/goutil"
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

func NewUserPolicy(
	params NewUserPolicyParams,
) UserPolicy {
	return &userPolicy{
		userRepository: params.UserRepository,
		roleRepository: params.RoleRepository,
	}
}

func (p *userPolicy) Save(ctx context.Context, _dto *dto.SaveUser) (*entity.User, error) {
	// Initialize variables to hold user data and error
	var (
		user *entity.User
		err  error
	)

	// 1. Handle user creation or retrieval
	if _dto.Id.IsZero() {
		// Create new user if ID is zero (new user)
		user = entity.NewUser()
	} else {
		// Find existing user by ID
		user, err = p.userRepository.Find(_dto.Id.String)
		if err != nil {
			// Return 404 if user not found
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, errorutil.NewHttpError(http.StatusNotFound, "User tidak ditemukan", nil)
			}
			// Return other errors as-is
			return nil, err
		}

		// 2. Check if user has SuperAdmin role
		isSuperAdmin, err := p.userRepository.HasRoleNames(user, []roleEnum.RoleName{roleEnum.SuperAdmin})
		if err != nil {
			return nil, err
		}
		if isSuperAdmin {
			return nil, errorutil.NewHttpError(
				http.StatusForbidden,
				"User dengan role "+string(roleEnum.SuperAdmin)+" tidak dapat diubah",
				nil,
			)
		}
	}

	// 3. Validate new roles being assigned
	if !goutil.IsEmpty(_dto.RoleIds) {
		roles, err := p.roleRepository.FindByIds(_dto.RoleIds)
		if err != nil {
			return nil, err
		}

		if len(roles) != len(_dto.RoleIds) {
			return nil, errorutil.NewHttpError(
				http.StatusBadRequest,
				"One or more roles not found",
				nil,
			)
		}

		// Check if any of the new roles is SuperAdmin
		for _, role := range roles {
			if role.Name == roleEnum.SuperAdmin {
				// Prevent assignment of SuperAdmin role
				return nil, errorutil.NewHttpError(
					http.StatusForbidden,
					"Role "+string(roleEnum.SuperAdmin)+" tidak dapat ditambahkan ke user",
					nil,
				)
			}
		}

		user.Roles = roles
	} else {
		defaultRole, err := p.roleRepository.FindDefault()
		if err != nil {
			// TODO: return custom error on repository instead of gorm's error
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, errorutil.NewHttpError(http.StatusBadRequest, "Default role is not configured", nil)
			}
			return nil, err
		}
		user.Roles = []*entity.Role{defaultRole}
	}

	// Return the user if all validations pass
	return user, nil
}

// UpdatePassword checks if the authenticated user is authorized to update another user's password.
func (p *userPolicy) UpdatePassword(
	ctx context.Context,
	_dto *dto.UpdateUserPassword,
) (*entity.User, error) {
	// Get the currently authenticated user from context
	authUser := ctx.Value(contextkey.UserKey).(*entity.User)

	// Find the target user whose password is being updated
	user, err := p.userRepository.Find(_dto.Id)
	if err != nil {
		// Return 404 if user is not found
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorutil.NewHttpError(http.StatusNotFound, "User tidak ditemukan", nil)
		}
		return nil, err
	}

	// Check if the target user is a SuperAdmin
	isSuperAdmin, err := p.userRepository.HasRoleNames(user, []roleEnum.RoleName{roleEnum.SuperAdmin})
	if err != nil {
		return nil, err
	}

	// If the target user is a SuperAdmin, only allow them to update their own password
	// Prevent other users (even other SuperAdmins) from updating a SuperAdmin's password
	if isSuperAdmin && authUser.Id != user.Id {
		return nil, errorutil.NewHttpError(
			http.StatusForbidden,
			"User dengan role "+string(roleEnum.SuperAdmin)+" tidak dapat diubah",
			nil,
		)
	}

	// If all checks pass, return the user to proceed with password update
	return user, nil
}

func (p *userPolicy) ToggleActivation(ctx context.Context, _dto *dto.ToggleActivation) (*entity.User, error) {
	user, err := p.userRepository.Find(_dto.Id)
	if err != nil {
		// TODO: return custom error on repository instead of gorm's error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorutil.NewHttpError(http.StatusNotFound, "User tidak ditemukan", nil)
		}
		return nil, err
	}

	isSuperAdmin, err := p.userRepository.HasRoleNames(user, []roleEnum.RoleName{roleEnum.SuperAdmin})
	if err != nil {
		return nil, err
	}
	if isSuperAdmin {
		return nil, errorutil.NewHttpError(
			http.StatusForbidden,
			"User dengan role "+string(roleEnum.SuperAdmin)+" tidak dapat diubah",
			nil,
		)
	}

	return user, nil
}

// Destroy validates if a user can be deleted based on certain business rules
func (p *userPolicy) Destroy(ctx context.Context, _dto *dto.DestroyUser) (*entity.User, error) {
	// 1. Attempt to find the user by ID
	user, err := p.userRepository.Find(_dto.Id)

	if err != nil {
		// 2. If user not found, return 404 error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorutil.NewHttpError(http.StatusNotFound, "User tidak ditemukan", nil)
		}
		// 3. For other errors, return them as-is
		return nil, err
	}

	// 4. Check if user has any protected roles
	isSuperAdmin, err := p.userRepository.HasRoleNames(user, []roleEnum.RoleName{roleEnum.SuperAdmin})
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

	// 6. If all validations pass, return the user (allowing the delete operation to proceed)
	return user, nil
}
