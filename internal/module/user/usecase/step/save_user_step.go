package step

import (
	"context"

	"github.com/arfanxn/welding/internal/infrastructure/id"
	"github.com/arfanxn/welding/internal/infrastructure/security"
	employeeRepository "github.com/arfanxn/welding/internal/module/employee/domain/repository"
	roleRepository "github.com/arfanxn/welding/internal/module/role/domain/repository"
	roleUserRepository "github.com/arfanxn/welding/internal/module/role_user/domain/repository"
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	userRepository "github.com/arfanxn/welding/internal/module/user/domain/repository"
	"github.com/arfanxn/welding/internal/module/user/usecase/dto"
	"github.com/arfanxn/welding/pkg/query"
	"github.com/gookit/goutil"
	"github.com/guregu/null/v6"
	"github.com/samber/lo"
	"go.uber.org/fx"
)

type SaveUserStep interface {
	Handle(ctx context.Context, _dto *dto.SaveUser) (*entity.User, error)
}

type saveUserStep struct {
	passwordService security.PasswordService
	idService       id.IdService

	userRepository     userRepository.UserRepository
	employeeRepository employeeRepository.EmployeeRepository
	roleRepository     roleRepository.RoleRepository
	roleUserRepository roleUserRepository.RoleUserRepository
}

type NewSaveUserStepParams struct {
	fx.In

	IdService       id.IdService
	PasswordService security.PasswordService

	UserRepository     userRepository.UserRepository
	EmployeeRepository employeeRepository.EmployeeRepository
	RoleRepository     roleRepository.RoleRepository
	RoleUserRepository roleUserRepository.RoleUserRepository
}

func NewSaveUserStep(params NewSaveUserStepParams) SaveUserStep {
	return &saveUserStep{
		passwordService: params.PasswordService,
		idService:       params.IdService,

		userRepository:     params.UserRepository,
		employeeRepository: params.EmployeeRepository,
		roleRepository:     params.RoleRepository,
		roleUserRepository: params.RoleUserRepository,
	}
}

// Handle saves or updates a user based on the provided DTO.
// It handles both creating new users and updating existing users, including:
// - Basic user information (name, phone, email, password)
// - Account activation/deactivation status
// - User role assignments
// - Employee association with employment identity number
//
// Parameters:
//   - ctx: Context for the operation
//   - _dto: SaveUser DTO containing user data to save
//
// Returns:
//   - *entity.User: The saved/updated user with all associations
//   - error: Any error encountered during the operation
func (s *saveUserStep) Handle(ctx context.Context, _dto *dto.SaveUser) (*entity.User, error) {
	// Initialize query and include relationships
	var (
		q      = query.NewQuery()
		user   *entity.User
		userId string
		err    error
	)

	// Conditionally include Employee relationship only when employee data is provided
	// This optimizes database queries by avoiding unnecessary JOIN operations
	if _dto.EmploymentIdentityNumber != nil {
		q.Include("Employee")
	}

	// Conditionally include Roles relationship only when role assignments are provided
	// This optimizes database queries by avoiding unnecessary JOIN operations
	if _dto.RoleIds != nil {
		q.Include("Roles")
	}

	// Retrieve existing user or create new one
	if !goutil.IsEmptyReal(_dto.Id) {
		// Update scenario: fetch existing user
		userId = *_dto.Id
		q = q.FilterById(userId)
		user, err = s.userRepository.First(q)
		if err != nil {
			return nil, err
		}
	} else {
		// Create scenario: initialize new user with generated ID
		userId = s.idService.Generate()
		q = q.FilterById(userId)
		user = &entity.User{Id: userId}
	}

	// Update basic user information if provided
	if !goutil.IsEmptyReal(_dto.Name) {
		user.Name = *_dto.Name
	}
	if !goutil.IsEmptyReal(_dto.PhoneNumber) {
		user.PhoneNumber = *_dto.PhoneNumber
	}

	// Handle email update - reset email verification if email changed
	if !goutil.IsEmptyReal(_dto.Email) {
		if user.Email != *_dto.Email {
			user.EmailVerifiedAt = null.TimeFromPtr(nil)
		}
		user.Email = *_dto.Email
	}

	// Handle password update with hashing
	if !goutil.IsEmptyReal(_dto.Password) {
		user.Password, err = s.passwordService.Hash(*_dto.Password)
		if err != nil {
			return nil, err
		}
	}

	// Handle account activation/deactivation
	if !goutil.IsEmptyReal(_dto.ActivatedAt) {
		user.ActivatedAt = null.TimeFromPtr(_dto.ActivatedAt)
		user.DeactivatedAt = null.TimeFromPtr(nil)
	}
	if !goutil.IsEmptyReal(_dto.DeactivatedAt) {
		user.ActivatedAt = null.TimeFromPtr(nil)
		user.DeactivatedAt = null.TimeFromPtr(_dto.DeactivatedAt)
	}

	// Save user basic information
	if err := s.userRepository.Save(user); err != nil {
		return nil, err
	}

	// Handle role assignments - replace all existing roles with new ones
	if _dto.RoleIds != nil {
		// Remove all existing role associations for this user
		if err := s.roleUserRepository.DestroyByUserId(user.Id); err != nil {
			return nil, err
		}

		if !goutil.IsEmptyReal(_dto.RoleIds[0]) {
			// Create new role associations
			rus := lo.Map(_dto.RoleIds, func(roleId string, _ int) *entity.RoleUser {
				return &entity.RoleUser{RoleId: roleId, UserId: user.Id}
			})
			if err := s.roleUserRepository.SaveMany(rus); err != nil {
				return nil, err
			}
		}
	}

	// Handle employee association based on employment identity number
	if _dto.EmploymentIdentityNumber != nil {
		if !(goutil.IsEmptyReal(_dto.EmploymentIdentityNumber)) {
			// Create or update employee record
			if user.Employee == nil {
				user.Employee = &entity.Employee{UserId: user.Id}
			}
			user.Employee.EmploymentIdentityNumber = *_dto.EmploymentIdentityNumber
			if err := s.employeeRepository.Save(user.Employee); err != nil {
				return nil, err
			}
		} else {
			// Remove employee record if employment identity number is empty
			if err := s.employeeRepository.DestroyByUserId(user.Id); err != nil {
				return nil, err
			}
			user.Employee = nil
		}
	}

	// Fetch complete user with all associations to return
	user, err = s.userRepository.First(q)
	if err != nil {
		return nil, err
	}

	return user, nil
}
