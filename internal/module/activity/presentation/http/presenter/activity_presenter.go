package presenter

import (
	"context"
	"net/url"

	activityEnum "github.com/arfanxn/welding/internal/module/activity/domain/enum"
	activity_viewmodel "github.com/arfanxn/welding/internal/module/activity/presentation/http/viewmodel"
	activityService "github.com/arfanxn/welding/internal/module/activity/usecase/service"
	"github.com/arfanxn/welding/internal/module/shared/contextkey"
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	"github.com/arfanxn/welding/internal/module/shared/domain/errorx"
	"github.com/arfanxn/welding/pkg/pagination"
	"github.com/gookit/goutil"
	"go.uber.org/fx"
)

// ActivityPresenter defines the interface for converting activity domain entities to view models.
// It handles the transformation of activity data between the domain layer and the presentation layer.
type ActivityPresenter interface {
	// FromEntityToViewModel converts a single activity entity to its corresponding view model.
	// It takes a context and an activity entity, and returns the view model representation
	// along with any error that occurred during conversion.
	FromEntityToViewModel(ctx context.Context, activity *entity.Activity) (*activity_viewmodel.ActivityViewModel, error)

	// FromEntitiesToViewModels converts a slice of activity entities to their corresponding view models.
	// It processes multiple activities in bulk and returns a slice of view models.
	// If any error occurs during conversion, it returns the error immediately.
	FromEntitiesToViewModels(ctx context.Context, activities []*entity.Activity) (
		[]*activity_viewmodel.ActivityViewModel,
		error,
	)

	// FromEntityOffsetPaginationToViewModelOffsetPagination converts a paginated result of activity entities
	// to a paginated result of view models using offset-based pagination.
	// This is useful for APIs that need to return paginated lists of activities.
	FromEntityOffsetPaginationToViewModelOffsetPagination(
		ctx context.Context,
		op *pagination.OffsetPagination[*entity.Activity],
	) (*pagination.OffsetPagination[*activity_viewmodel.ActivityViewModel], error)

	// FromEntityOffsetPaginationToViewModelPagePagination converts a paginated result of activity entities
	// to a page-based pagination result of view models.
	// This is useful for web interfaces that display activities with page numbers.
	FromEntityOffsetPaginationToViewModelPagePagination(
		ctx context.Context,
		op *pagination.OffsetPagination[*entity.Activity],
	) (*pagination.PagePagination[*activity_viewmodel.ActivityViewModel], error)
}

type activityPresenter struct {
	activityService activityService.ActivityService
}

type NewActivityPresenterParams struct {
	fx.In

	ActivityService activityService.ActivityService
}

func NewActivityPresenter(params NewActivityPresenterParams) ActivityPresenter {
	return &activityPresenter{
		activityService: params.ActivityService,
	}
}

func (s *activityPresenter) resolveActionDescription(action activityEnum.ActivityAction) (string, error) {
	var description string
	switch action {
	case activityEnum.UsersRegister:
		description = "User mendaftar"
	case activityEnum.UsersResetPassword:
		description = "User mengatur ulang kata sandi"
	case activityEnum.UsersVerifyEmail:
		description = "User memverifikasi email"
	case activityEnum.UsersLogin:
		description = "User masuk"
	case activityEnum.UsersLogout:
		description = "User keluar"
	case activityEnum.UsersUpdateMeProfile:
		description = "User memperbarui profil mereka"
	case activityEnum.UsersUpdateMePassword:
		description = "User memperbarui password mereka"
	case activityEnum.UsersIndex:
		description = "User melihat index users"
	case activityEnum.UsersShow:
		description = "User melihat detail user"
	case activityEnum.UsersStore:
		description = "User menambahkan user"
	case activityEnum.UsersUpdate:
		description = "User memperbarui user"
	case activityEnum.UsersToggleActivation:
		description = "User mengaktifkan/menonaktifkan user"
	case activityEnum.UsersDestroy:
		description = "User menghapus user"
	case activityEnum.PermissionsIndex:
		description = "User melihat index permissions"
	case activityEnum.RolesIndex:
		description = "User melihat index roles"
	case activityEnum.RolesShow:
		description = "User melihat detail role"
	case activityEnum.RolesStore:
		description = "User menambahkan role"
	case activityEnum.RolesUpdate:
		description = "User memperbarui role"
	case activityEnum.RolesSetDefault:
		description = "User mengatur default role"
	case activityEnum.RolesDestroy:
		description = "User menghapus role"
	case activityEnum.CodesCreateUserRegisterInvitation:
		description = "User membuat kode undangan untuk pendaftaran user"
	case activityEnum.CodesCreateUserEmailVerification:
		description = "User membuat kode verifikasi email"
	case activityEnum.CodesCreateUserResetPassword:
		description = "User membuat kode reset password"
	default:
	}

	if goutil.IsEmpty(description) {
		return "", errorx.ErrActivityInvalidAction
	}

	return description, nil
}

func (p *activityPresenter) FromEntityToViewModel(ctx context.Context, activity *entity.Activity) (*activity_viewmodel.ActivityViewModel, error) {
	activity_vm := &activity_viewmodel.ActivityViewModel{}

	activity_vm.Id = activity.Id
	activity_vm.CauserId = activity.CauserId
	activity_vm.CauserType = activity.CauserType
	activity_vm.CauserIpAddress = activity.CauserIpAddress
	activity_vm.Action = activity.Action

	description, err := p.resolveActionDescription(activity.Action)
	if err != nil {
		return nil, err
	}
	activity_vm.Description = description

	activity_vm.SubjectId = activity.SubjectId
	activity_vm.SubjectType = activity.SubjectType
	activity_vm.Properties = activity.Properties
	activity_vm.CreatedAt = activity.CreatedAt
	activity_vm.UpdatedAt = activity.UpdatedAt

	// causer
	switch {
	case activity.CauserUser != nil:
		activity_vm.Causer = activity.CauserUser
	}

	// subject
	switch {
	case activity.SubjectUser != nil:
		activity_vm.Subject = activity.SubjectUser
	case activity.SubjectRole != nil:
		activity_vm.Subject = activity.SubjectRole
	case activity.SubjectPermission != nil:
		activity_vm.Subject = activity.SubjectPermission
	case activity.SubjectCode != nil:
		activity_vm.Subject = activity.SubjectCode
	}

	return activity_vm, nil
}

func (p *activityPresenter) FromEntitiesToViewModels(ctx context.Context, activities []*entity.Activity) ([]*activity_viewmodel.ActivityViewModel, error) {
	var viewModels []*activity_viewmodel.ActivityViewModel
	for _, activity := range activities {
		viewModel, err := p.FromEntityToViewModel(ctx, activity)
		if err != nil {
			return nil, err
		}
		viewModels = append(viewModels, viewModel)
	}

	return viewModels, nil
}

func (p *activityPresenter) FromEntityOffsetPaginationToViewModelOffsetPagination(ctx context.Context, op *pagination.OffsetPagination[*entity.Activity]) (*pagination.OffsetPagination[*activity_viewmodel.ActivityViewModel], error) {
	viewModels, err := p.FromEntitiesToViewModels(ctx, op.Items)
	if err != nil {
		return nil, err
	}

	return pagination.NewOffsetPagination(op.Offset, op.Limit, op.TotalItems, viewModels), nil
}

func (p *activityPresenter) FromEntityOffsetPaginationToViewModelPagePagination(ctx context.Context, op *pagination.OffsetPagination[*entity.Activity]) (*pagination.PagePagination[*activity_viewmodel.ActivityViewModel], error) {
	convertedOP, err := p.FromEntityOffsetPaginationToViewModelOffsetPagination(ctx, op)
	if err != nil {
		return nil, err
	}

	url, _ := ctx.Value(contextkey.RequestURLKey).(url.URL)

	return pagination.FromOPToPP(convertedOP, url), nil
}
