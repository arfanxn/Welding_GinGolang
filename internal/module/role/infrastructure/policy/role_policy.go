package policy

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"regexp"

	permissionRepository "github.com/arfanxn/welding/internal/module/permission/domain/repository"
	"github.com/arfanxn/welding/internal/module/role/domain/enum"
	roleRepository "github.com/arfanxn/welding/internal/module/role/domain/repository"
	roleDto "github.com/arfanxn/welding/internal/module/role/usecase/dto"
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	"github.com/arfanxn/welding/pkg/errorutil"
	"github.com/gookit/goutil"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

type RolePolicy interface {
	Save(ctx context.Context, _dto *roleDto.SaveRole) (*entity.Role, error)
	SetDefault(ctx context.Context, _dto *roleDto.SetDefaultRole) (*entity.Role, error)
	Destroy(ctx context.Context, _dto *roleDto.DestroyRole) (*entity.Role, error)
}

type rolePolicy struct {
	roleRepository       roleRepository.RoleRepository
	permissionRepository permissionRepository.PermissionRepository
}

type NewRolePolicyParams struct {
	fx.In

	RoleRepository       roleRepository.RoleRepository
	PermissionRepository permissionRepository.PermissionRepository
}

func NewRolePolicy(params NewRolePolicyParams) RolePolicy {
	return &rolePolicy{
		roleRepository:       params.RoleRepository,
		permissionRepository: params.PermissionRepository,
	}
}

func (p *rolePolicy) Save(ctx context.Context, _dto *roleDto.SaveRole) (*entity.Role, error) {
	var (
		role *entity.Role
		err  error
	)

	if goutil.IsEmpty(_dto.Id) {
		role = entity.NewRole()
	} else {
		role, err = p.roleRepository.Find(_dto.Id)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, errorutil.NewHttpError(http.StatusNotFound, "Role tidak ditemukan", nil)
			}
			return nil, err
		}
	}

	if !role.IsSaveable() {
		return nil, errorutil.NewHttpError(
			http.StatusForbidden,
			fmt.Sprintf("Role dengan nama %s tidak dapat dibuat atau diperbarui", enum.SuperAdmin),
			nil,
		)
	}

	// Handle permissions
	if !goutil.IsEmpty(_dto.PermissionIds) {
		permissions, err := p.permissionRepository.FindByIds(_dto.PermissionIds)
		if err != nil {
			return nil, err
		}

		// Verify all requested permissions were found
		if len(permissions) != len(_dto.PermissionIds) {
			return nil, errorutil.NewHttpError(
				http.StatusBadRequest,
				"Satu atau lebih permission tidak ditemukan",
				nil,
			)
		}
	}

	return role, nil
}

func (p *rolePolicy) SetDefault(ctx context.Context, _dto *roleDto.SetDefaultRole) (*entity.Role, error) {
	role, err := p.roleRepository.Find(_dto.Id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorutil.NewHttpError(http.StatusNotFound, "Role tidak ditemukan", nil)
		}
		return nil, err
	}

	if role.IsDefault {
		return nil, errorutil.NewHttpError(http.StatusBadRequest, "Role default tidak dapat diset default", nil)
	}
	// Check if the role is unmodifiable
	if matched, _ := regexp.MatchString(enum.SuperAdmin.String(), role.Name.String()); matched {
		return nil, errorutil.NewHttpError(
			http.StatusForbidden,
			fmt.Sprintf("Role %s tidak dapat diset default", role.Name.String()),
			nil,
		)
	}

	return role, nil
}

func (p *rolePolicy) Destroy(ctx context.Context, _dto *roleDto.DestroyRole) (*entity.Role, error) {
	role, err := p.roleRepository.Find(_dto.Id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorutil.NewHttpError(http.StatusNotFound, "Role tidak ditemukan", nil)
		}
		return nil, err
	}

	if role.IsDefault {
		return nil, errorutil.NewHttpError(http.StatusForbidden, "Role default tidak dapat dihapus", nil)
	}

	if matched, _ := regexp.MatchString(string(enum.SuperAdmin), string(role.Name)); matched {
		return nil, errorutil.NewHttpError(http.StatusForbidden, fmt.Sprintf("role %s tidak dapat dihapus", role.Name.String()), nil)
	}

	return role, nil
}
