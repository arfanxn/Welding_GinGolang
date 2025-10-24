package usecase

import (
	"errors"
	"net/http"

	permissionRepository "github.com/arfanxn/welding/internal/module/permission/domain/repository"
	"github.com/arfanxn/welding/internal/module/role/domain/repository"
	roleDto "github.com/arfanxn/welding/internal/module/role/usecase/dto"
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	"github.com/arfanxn/welding/pkg/errorutil"
	"github.com/arfanxn/welding/pkg/pagination"
	"github.com/arfanxn/welding/pkg/query"
	"github.com/arfanxn/welding/pkg/reflectutil"
	"github.com/gookit/goutil"
	"github.com/oklog/ulid/v2"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

type RoleUsecase interface {
	Paginate(*query.Query) (*pagination.OffsetPagination[*entity.Role], error)
	Find(*query.Query) (*entity.Role, error)
	Save(*roleDto.SaveRole) (*entity.Role, error)
	Destroy(*roleDto.DestroyRole) error
}

var _ RoleUsecase = (*roleUsecase)(nil)

type roleUsecase struct {
	roleRepository       repository.RoleRepository
	permissionRepository permissionRepository.PermissionRepository
}

type NewRoleUsecaseParams struct {
	fx.In

	RoleRepository       repository.RoleRepository
	PermissionRepository permissionRepository.PermissionRepository
}

func NewRoleUsecase(params NewRoleUsecaseParams) RoleUsecase {
	return &roleUsecase{
		roleRepository:       params.RoleRepository,
		permissionRepository: params.PermissionRepository,
	}
}

func (u *roleUsecase) Find(q *query.Query) (*entity.Role, error) {
	roles, err := u.roleRepository.Get(q)
	if err != nil {
		return nil, err
	}
	if len(roles) == 0 {
		return nil, errorutil.NewHttpError(http.StatusNotFound, "Role not found", nil)
	}
	return roles[0], nil
}

func (u *roleUsecase) Paginate(q *query.Query) (*pagination.OffsetPagination[*entity.Role], error) {
	return u.roleRepository.Paginate(q)
}

// Save creates a new role or updates an existing one based on the provided DTO.
// It validates the role data, checks for duplicate role names, and ensures all referenced permissions exist.
//
// Parameters:
//   - _dto: The DTO containing role data to be saved
//
// Returns:
//   - *entity.Role: The saved role with updated fields
//   - error: An error if the operation fails, which could be:
//   - 404 if trying to update a non-existent role
//   - 400 if any referenced permissions don't exist
//   - 409 if a role with the same name already exists
func (u *roleUsecase) Save(_dto *roleDto.SaveRole) (*entity.Role, error) {
	var role *entity.Role
	var err error

	// Handle role creation or retrieval
	if goutil.IsEmpty(_dto.Id) {
		// Create new role
		role = entity.NewRole()
		role.Id = ulid.Make().String()
	} else {
		// Find existing role
		role, err = u.roleRepository.Find(_dto.Id)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, errorutil.NewHttpError(http.StatusNotFound, "Role not found", nil)
			}
			return nil, err
		}
	}

	role.Name = _dto.Name

	// Handle permissions
	if reflectutil.IsSlice(_dto.PermissionIds) {
		permissions, err := u.permissionRepository.FindByIds(_dto.PermissionIds)
		if err != nil {
			return nil, err
		}

		// Verify all requested permissions were found
		if len(permissions) != len(_dto.PermissionIds) {
			return nil, errorutil.NewHttpError(
				http.StatusBadRequest,
				"One or more permissions not found",
				nil,
			)
		}
		role.Permissions = permissions
	}

	// Save the role
	if err := u.roleRepository.Save(role); err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, errorutil.NewHttpError(http.StatusConflict, "Role name already exists", nil)
		}
		return nil, err
	}

	return role, nil
}

func (u *roleUsecase) Destroy(drDto *roleDto.DestroyRole) error {
	role, err := u.roleRepository.Find(drDto.Id)
	if err != nil {
		if goutil.IsNil(role) {
			return errorutil.NewHttpError(http.StatusNotFound, "Role not found", nil)
		}
		return err
	}

	return u.roleRepository.Destroy(role)
}
