package usecase

import (
	"context"

	permissionRepository "github.com/arfanxn/welding/internal/module/permission/domain/repository"
	"github.com/arfanxn/welding/internal/module/role/domain/repository"
	"github.com/arfanxn/welding/internal/module/role/infrastructure/policy"
	"github.com/arfanxn/welding/internal/module/role/usecase/dto"
	roleDto "github.com/arfanxn/welding/internal/module/role/usecase/dto"
	"github.com/arfanxn/welding/internal/module/role/usecase/step"
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	"github.com/arfanxn/welding/pkg/pagination"
	"github.com/arfanxn/welding/pkg/query"
	"go.uber.org/fx"
)

type RoleUsecase interface {
	Paginate(context.Context, *query.Query) (*pagination.OffsetPagination[*entity.Role], error)
	Show(context.Context, *query.Query) (*entity.Role, error)
	Store(context.Context, *roleDto.SaveRole) (*entity.Role, error)
	Update(context.Context, *roleDto.SaveRole) (*entity.Role, error)
	SetDefault(context.Context, *roleDto.SetDefaultRole) (*entity.Role, error)
	Destroy(context.Context, *roleDto.DestroyRole) error
}

var _ RoleUsecase = (*roleUsecase)(nil)

type roleUsecase struct {
	storeRoleStep        step.StoreRoleStep
	updateRoleStep       step.UpdateRoleStep
	rolePolicy           policy.RolePolicy
	roleRepository       repository.RoleRepository
	permissionRepository permissionRepository.PermissionRepository
}

type NewRoleUsecaseParams struct {
	fx.In

	StoreRoleStep        step.StoreRoleStep
	UpdateRoleStep       step.UpdateRoleStep
	RolePolicy           policy.RolePolicy
	RoleRepository       repository.RoleRepository
	PermissionRepository permissionRepository.PermissionRepository
}

func NewRoleUsecase(params NewRoleUsecaseParams) RoleUsecase {
	return &roleUsecase{
		storeRoleStep:        params.StoreRoleStep,
		updateRoleStep:       params.UpdateRoleStep,
		rolePolicy:           params.RolePolicy,
		roleRepository:       params.RoleRepository,
		permissionRepository: params.PermissionRepository,
	}
}

func (u *roleUsecase) Show(ctx context.Context, q *query.Query) (*entity.Role, error) {
	return u.roleRepository.First(q)
}

func (u *roleUsecase) Paginate(ctx context.Context, q *query.Query) (*pagination.OffsetPagination[*entity.Role], error) {
	return u.roleRepository.Paginate(q)
}

func (u *roleUsecase) Store(ctx context.Context, _dto *dto.SaveRole) (*entity.Role, error) {
	if err := u.rolePolicy.Store(ctx, _dto); err != nil {
		return nil, err
	}

	return u.storeRoleStep.Handle(ctx, _dto)
}

func (u *roleUsecase) Update(ctx context.Context, _dto *roleDto.SaveRole) (*entity.Role, error) {
	if err := u.rolePolicy.Update(ctx, _dto); err != nil {
		return nil, err
	}

	return u.updateRoleStep.Handle(ctx, _dto)
}

func (u *roleUsecase) SetDefault(ctx context.Context, _dto *roleDto.SetDefaultRole) (*entity.Role, error) {
	if err := u.rolePolicy.SetDefault(ctx, _dto); err != nil {
		return nil, err
	}

	role, err := u.roleRepository.Find(_dto.Id)
	if err != nil {
		return nil, err
	}

	if err := u.roleRepository.SetDefault(role); err != nil {
		return nil, err
	}

	return role, nil
}

func (u *roleUsecase) Destroy(ctx context.Context, _dto *roleDto.DestroyRole) error {
	if err := u.rolePolicy.Destroy(ctx, _dto); err != nil {
		return err
	}

	role, err := u.roleRepository.Find(_dto.Id)
	if err != nil {
		return err
	}

	return u.roleRepository.Destroy(role)
}
