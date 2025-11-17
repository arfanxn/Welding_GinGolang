package action

import (
	"context"
	"errors"
	"time"

	"github.com/arfanxn/welding/internal/module/code/domain/enum"
	codeRepository "github.com/arfanxn/welding/internal/module/code/domain/repository"
	roleRepository "github.com/arfanxn/welding/internal/module/role/domain/repository"
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	"github.com/arfanxn/welding/internal/module/shared/domain/errorx"
	userRepository "github.com/arfanxn/welding/internal/module/user/domain/repository"
	"github.com/arfanxn/welding/internal/module/user/usecase/dto"
	"gorm.io/gorm"
)

type RegisterUserAction interface {
	Handle(ctx context.Context, _dto *dto.Register) (*entity.User, error)
}

type registerUserAction struct {
	saveUserAction SaveUserAction

	userRepository userRepository.UserRepository
	codeRepository codeRepository.CodeRepository
	roleRepository roleRepository.RoleRepository
}

func NewRegisterUserAction(
	saveUserAction SaveUserAction,

	userRepository userRepository.UserRepository,
	codeRepository codeRepository.CodeRepository,
	roleRepository roleRepository.RoleRepository,
) RegisterUserAction {
	return &registerUserAction{
		saveUserAction: saveUserAction,

		userRepository: userRepository,
		codeRepository: codeRepository,
		roleRepository: roleRepository,
	}
}

// Handle registers a new user based on the provided registration DTO.
// It handles two registration scenarios:
// - Invitation-based registration: Uses a valid invitation code to assign specific roles
// - Default registration: Assigns the default role to new users
//
// The function validates invitation codes (if provided), determines appropriate roles,
// creates the user account, and marks invitation codes as used.
//
// Parameters:
//   - ctx: Context for the operation
//   - _dto: Register DTO containing user registration data
//
// Returns:
//   - *entity.User: The newly registered user with all associations
//   - error: Any error encountered during the registration process
func (a *registerUserAction) Handle(
	ctx context.Context,
	_dto *dto.Register,
) (*entity.User, error) {
	// Initialize variables for role determination and invitation handling
	var (
		err                  error
		roleIds              []string
		isWithInvitationCode bool = _dto.InvitationCode != nil
		code                 *entity.Code
	)

	// Handle invitation-based registration
	if isWithInvitationCode {
		// Find invitation code by type and value
		code, err = a.codeRepository.FindByTypeAndValue(enum.UserRegisterInvitation, *_dto.InvitationCode)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, errorx.ErrCodeNotFound
			}
			return nil, err
		}

		// Validate invitation code hasn't been used
		if code.IsUsed() {
			return nil, errorx.ErrCodeAlreadyUsed
		}

		// Validate invitation code hasn't expired
		if code.IsExpired() {
			return nil, errorx.ErrCodeExpired
		}

		// Extract role ID from invitation code metadata
		codeMeta, err := code.GetMeta()
		if err != nil {
			return nil, err
		}

		roleId := codeMeta["role_id"].(string)
		roleIds = []string{roleId}
	} else {
		// Handle default registration without invitation code
		defaultRole, err := a.roleRepository.FindDefault()
		if err != nil {
			return nil, err
		}
		roleIds = []string{defaultRole.Id}
	}

	// Set activation time to immediately activate the new user
	activatedAt := time.Now()

	// Create the user account using the save user action
	user, err := a.saveUserAction.Handle(ctx, &dto.SaveUser{
		Name:                     &_dto.Name,
		PhoneNumber:              &_dto.PhoneNumber,
		Email:                    &_dto.Email,
		Password:                 &_dto.Password,
		ActivatedAt:              &activatedAt,
		RoleIds:                  roleIds,
		EmploymentIdentityNumber: _dto.EmploymentIdentityNumber,
	})

	// Mark invitation code as used if one was provided
	if isWithInvitationCode {
		code.MarkUsed()
		if err := a.codeRepository.Save(code); err != nil {
			return nil, err
		}
	}

	return user, err
}
