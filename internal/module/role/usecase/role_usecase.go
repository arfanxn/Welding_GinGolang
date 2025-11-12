package usecase

import (
	"context"
	"errors"
	"net/http"

	permissionRepository "github.com/arfanxn/welding/internal/module/permission/domain/repository"
	"github.com/arfanxn/welding/internal/module/role/domain/repository"
	"github.com/arfanxn/welding/internal/module/role/infrastructure/policy"
	roleDto "github.com/arfanxn/welding/internal/module/role/usecase/dto"
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	"github.com/arfanxn/welding/pkg/errorutil"
	"github.com/arfanxn/welding/pkg/pagination"
	"github.com/arfanxn/welding/pkg/query"
	"github.com/gookit/goutil"
	"github.com/samber/lo"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

type RoleUsecase interface {
	Paginate(context.Context, *query.Query) (*pagination.OffsetPagination[*entity.Role], error)
	Find(context.Context, *query.Query) (*entity.Role, error)
	Save(context.Context, *roleDto.SaveRole) (*entity.Role, error)
	SetDefault(context.Context, *roleDto.SetDefaultRole) (*entity.Role, error)
	Destroy(context.Context, *roleDto.DestroyRole) error
}

var _ RoleUsecase = (*roleUsecase)(nil)

type roleUsecase struct {
	rolePolicy           policy.RolePolicy
	roleRepository       repository.RoleRepository
	permissionRepository permissionRepository.PermissionRepository
}

type NewRoleUsecaseParams struct {
	fx.In

	RolePolicy           policy.RolePolicy
	RoleRepository       repository.RoleRepository
	PermissionRepository permissionRepository.PermissionRepository
}

func NewRoleUsecase(params NewRoleUsecaseParams) RoleUsecase {
	return &roleUsecase{
		rolePolicy:           params.RolePolicy,
		roleRepository:       params.RoleRepository,
		permissionRepository: params.PermissionRepository,
	}
}

func (u *roleUsecase) Find(ctx context.Context, q *query.Query) (*entity.Role, error) {
	roles, err := u.roleRepository.Get(q)
	if err != nil {
		return nil, err
	}
	if len(roles) == 0 {
		return nil, errorutil.NewHttpError(http.StatusNotFound, "Role tidak ditemukan", nil)
	}
	return roles[0], nil
}

func (u *roleUsecase) Paginate(ctx context.Context, q *query.Query) (*pagination.OffsetPagination[*entity.Role], error) {
	return u.roleRepository.Paginate(q)
}

func (u *roleUsecase) Save(ctx context.Context, _dto *roleDto.SaveRole) (*entity.Role, error) {
	role, err := u.rolePolicy.Save(ctx, _dto)
	if err != nil {
		return nil, err
	}

	if !goutil.IsEmpty(_dto.Name.String()) {
		role.Name = _dto.Name
	}

	if !goutil.IsEmpty(_dto.PermissionIds) {
		role.Permissions = lo.Map(_dto.PermissionIds, func(permId string, _ int) *entity.Permission {
			return &entity.Permission{Id: permId}
		})
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

func (u *roleUsecase) SetDefault(ctx context.Context, _dto *roleDto.SetDefaultRole) (*entity.Role, error) {
	role, err := u.rolePolicy.SetDefault(ctx, _dto)
	if err != nil {
		return nil, err
	}

	// Save the role
	if err := u.roleRepository.SetDefault(role); err != nil {
		return nil, err
	}

	return role, nil
}

func (u *roleUsecase) Destroy(ctx context.Context, _dto *roleDto.DestroyRole) error {
	role, err := u.rolePolicy.Destroy(ctx, _dto)
	if err != nil {
		return err
	}

	return u.roleRepository.Destroy(role)
}
