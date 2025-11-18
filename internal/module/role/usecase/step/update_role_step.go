package step

import (
	"context"

	"github.com/arfanxn/welding/internal/infrastructure/id"
	permissionRoleRepository "github.com/arfanxn/welding/internal/module/permission_role/domain/repository"
	roleRepository "github.com/arfanxn/welding/internal/module/role/domain/repository"
	"github.com/arfanxn/welding/internal/module/role/usecase/dto"
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	"github.com/arfanxn/welding/pkg/query"
	"github.com/gookit/goutil"
	"github.com/samber/lo"
)

type UpdateRoleStep interface {
	Handle(ctx context.Context, dto *dto.SaveRole) (*entity.Role, error)
}

type updateRoleStep struct {
	idService                id.IdService
	roleRepository           roleRepository.RoleRepository
	permissionRoleRepository permissionRoleRepository.PermissionRoleRepository
}

func NewUpdateRoleStep(
	idService id.IdService,
	roleRepository roleRepository.RoleRepository,
	permissionRoleRepository permissionRoleRepository.PermissionRoleRepository,
) UpdateRoleStep {
	return &updateRoleStep{
		idService:                idService,
		roleRepository:           roleRepository,
		permissionRoleRepository: permissionRoleRepository,
	}
}

func (s *updateRoleStep) Handle(ctx context.Context, _dto *dto.SaveRole) (*entity.Role, error) {
	q := query.NewQuery().FilterById(*_dto.Id)

	if _dto.PermissionIds != nil {
		q.Include("Permissions")
	}

	role, err := s.roleRepository.First(q)
	if err != nil {
		return nil, err
	}

	if !goutil.IsEmptyReal(_dto.Name) {
		role.Name = *_dto.Name
	}

	if err := s.roleRepository.Save(role); err != nil {
		return nil, err
	}

	if _dto.PermissionIds != nil {
		// Remove all existing role associations for this user
		if err := s.permissionRoleRepository.DestroyByRoleId(role.Id); err != nil {
			return nil, err
		}

		if !goutil.IsEmptyReal(_dto.PermissionIds[0]) {
			prs := lo.Map(_dto.PermissionIds, func(permId string, _ int) *entity.PermissionRole {
				return &entity.PermissionRole{RoleId: role.Id, PermissionId: permId}
			})

			if err := s.permissionRoleRepository.SaveMany(prs); err != nil {
				return nil, err
			}
		}
	}

	role, err = s.roleRepository.First(q)
	if err != nil {
		return nil, err
	}

	return role, nil
}
