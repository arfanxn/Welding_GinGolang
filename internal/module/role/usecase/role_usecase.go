package usecase

import (
	"context"

	activityEnum "github.com/arfanxn/welding/internal/module/activity/domain/enum"
	activityDto "github.com/arfanxn/welding/internal/module/activity/usecase/dto"
	activityService "github.com/arfanxn/welding/internal/module/activity/usecase/service"
	permissionRepository "github.com/arfanxn/welding/internal/module/permission/domain/repository"
	"github.com/arfanxn/welding/internal/module/role/domain/repository"
	"github.com/arfanxn/welding/internal/module/role/infrastructure/policy"
	"github.com/arfanxn/welding/internal/module/role/usecase/dto"

	"github.com/arfanxn/welding/internal/module/role/usecase/step"
	"github.com/arfanxn/welding/internal/module/shared/contextkey"
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	"github.com/arfanxn/welding/pkg/pagination"
	"github.com/arfanxn/welding/pkg/query"
	"github.com/arfanxn/welding/pkg/typeutil"
	"go.uber.org/fx"
)

type RoleUsecase interface {
	Paginate(context.Context, *query.Query) (*pagination.OffsetPagination[*entity.Role], error)
	Show(context.Context, *query.Query) (*entity.Role, error)
	Store(context.Context, *dto.SaveRole) (*entity.Role, error)
	Update(context.Context, *dto.SaveRole) (*entity.Role, error)
	SetDefault(context.Context, *dto.SetDefaultRole) (*entity.Role, error)
	Destroy(context.Context, *dto.DestroyRole) error
}

var _ RoleUsecase = (*roleUsecase)(nil)

type roleUsecase struct {
	activityService      activityService.ActivityService
	storeRoleStep        step.StoreRoleStep
	updateRoleStep       step.UpdateRoleStep
	rolePolicy           policy.RolePolicy
	roleRepository       repository.RoleRepository
	permissionRepository permissionRepository.PermissionRepository
}

type NewRoleUsecaseParams struct {
	fx.In

	ActivityService      activityService.ActivityService
	StoreRoleStep        step.StoreRoleStep
	UpdateRoleStep       step.UpdateRoleStep
	RolePolicy           policy.RolePolicy
	RoleRepository       repository.RoleRepository
	PermissionRepository permissionRepository.PermissionRepository
}

func NewRoleUsecase(params NewRoleUsecaseParams) RoleUsecase {
	return &roleUsecase{
		activityService:      params.ActivityService,
		storeRoleStep:        params.StoreRoleStep,
		updateRoleStep:       params.UpdateRoleStep,
		rolePolicy:           params.RolePolicy,
		roleRepository:       params.RoleRepository,
		permissionRepository: params.PermissionRepository,
	}
}

func (u *roleUsecase) Show(ctx context.Context, q *query.Query) (role *entity.Role, err error) {
	role, err = u.roleRepository.First(q)
	if err != nil {
		return nil, err
	}

	if _, err = u.activityService.Record(ctx, &activityDto.CreateActivity{
		Causer:  contextkey.GetUser(ctx),
		Action:  activityEnum.RolesShow,
		Subject: role,
	}); err != nil {
		return nil, err
	}

	return
}

func (u *roleUsecase) Paginate(ctx context.Context, q *query.Query) (op *pagination.OffsetPagination[*entity.Role], err error) {
	op, err = u.roleRepository.Paginate(q)
	if err != nil {
		return nil, err
	}

	if _, err = u.activityService.Record(ctx, &activityDto.CreateActivity{
		Causer:      contextkey.GetUser(ctx),
		Action:      activityEnum.RolesIndex,
		SubjectType: typeutil.Ptr(activityEnum.RoleSubjectType),
	}); err != nil {
		return nil, err
	}

	return
}

func (u *roleUsecase) Store(ctx context.Context, _dto *dto.SaveRole) (role *entity.Role, err error) {
	if err = u.rolePolicy.Store(ctx, _dto); err != nil {
		return nil, err
	}

	role, err = u.storeRoleStep.Handle(ctx, _dto)
	if err != nil {
		return nil, err
	}

	if _, err = u.activityService.Record(ctx, &activityDto.CreateActivity{
		Causer:  contextkey.GetUser(ctx),
		Action:  activityEnum.RolesStore,
		Subject: role,
	}); err != nil {
		return nil, err
	}

	return
}

func (u *roleUsecase) Update(ctx context.Context, _dto *dto.SaveRole) (role *entity.Role, err error) {
	if err = u.rolePolicy.Update(ctx, _dto); err != nil {
		return nil, err
	}

	role, err = u.updateRoleStep.Handle(ctx, _dto)
	if err != nil {
		return nil, err
	}

	if _, err = u.activityService.Record(ctx, &activityDto.CreateActivity{
		Causer:  contextkey.GetUser(ctx),
		Action:  activityEnum.RolesUpdate,
		Subject: role,
	}); err != nil {
		return nil, err
	}

	return
}

func (u *roleUsecase) SetDefault(ctx context.Context, _dto *dto.SetDefaultRole) (*entity.Role, error) {
	if err := u.rolePolicy.SetDefault(ctx, _dto); err != nil {
		return nil, err
	}

	role, err := u.roleRepository.Find(_dto.Id, nil)
	if err != nil {
		return nil, err
	}

	if err := u.roleRepository.SetDefault(role); err != nil {
		return nil, err
	}

	if _, err = u.activityService.Record(ctx, &activityDto.CreateActivity{
		Causer:  contextkey.GetUser(ctx),
		Action:  activityEnum.RolesSetDefault,
		Subject: role,
	}); err != nil {
		return nil, err
	}

	return role, nil
}

func (u *roleUsecase) Destroy(ctx context.Context, _dto *dto.DestroyRole) error {
	if err := u.rolePolicy.Destroy(ctx, _dto); err != nil {
		return err
	}

	role, err := u.roleRepository.Find(_dto.Id, nil)
	if err != nil {
		return err
	}

	err = u.roleRepository.Destroy(role)
	if err != nil {
		return err
	}

	if _, err = u.activityService.Record(ctx, &activityDto.CreateActivity{
		Causer:  contextkey.GetUser(ctx),
		Action:  activityEnum.RolesDestroy,
		Subject: role,
	}); err != nil {
		return err
	}

	return nil
}
