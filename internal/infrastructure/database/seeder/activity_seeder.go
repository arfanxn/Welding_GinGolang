package seeder

import (
	"github.com/arfanxn/welding/internal/infrastructure/id"
	activityEnum "github.com/arfanxn/welding/internal/module/activity/domain/enum"
	activityRepository "github.com/arfanxn/welding/internal/module/activity/domain/repository"
	codeEnum "github.com/arfanxn/welding/internal/module/code/domain/enum"
	codeRepository "github.com/arfanxn/welding/internal/module/code/domain/repository"
	permissionRepository "github.com/arfanxn/welding/internal/module/permission/domain/repository"
	roleEnum "github.com/arfanxn/welding/internal/module/role/domain/enum"
	roleRepository "github.com/arfanxn/welding/internal/module/role/domain/repository"
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	userRepository "github.com/arfanxn/welding/internal/module/user/domain/repository"
	"github.com/arfanxn/welding/pkg/query"
	"github.com/arfanxn/welding/pkg/typeutil"
	"github.com/bluele/factory-go/factory"
	"github.com/brianvoe/gofakeit/v7"
	"go.uber.org/fx"
)

var _ Seeder = (*ActivitySeeder)(nil)

type ActivitySeeder struct {
	idService            id.IdService
	activityFactory      *factory.Factory
	activityRepository   activityRepository.ActivityRepository
	codeRepository       codeRepository.CodeRepository
	userRepository       userRepository.UserRepository
	roleRepository       roleRepository.RoleRepository
	permissionRepository permissionRepository.PermissionRepository
}

type NewActivitySeederParams struct {
	fx.In

	IdService            id.IdService
	ActivityFactory      *factory.Factory `name:"activity_factory"`
	ActivityRepository   activityRepository.ActivityRepository
	CodeRepository       codeRepository.CodeRepository
	UserRepository       userRepository.UserRepository
	RoleRepository       roleRepository.RoleRepository
	PermissionRepository permissionRepository.PermissionRepository
}

func NewActivitySeeder(
	params NewActivitySeederParams,
) Seeder {
	return &ActivitySeeder{
		idService:            params.IdService,
		activityFactory:      params.ActivityFactory,
		activityRepository:   params.ActivityRepository,
		codeRepository:       params.CodeRepository,
		userRepository:       params.UserRepository,
		roleRepository:       params.RoleRepository,
		permissionRepository: params.PermissionRepository,
	}
}

func (s *ActivitySeeder) Seed() error {
	user, err := s.userRepository.First(nil)
	if err != nil {
		return err
	}

	roleEngineer, err := s.roleRepository.First(query.NewQuery().Filter("name", query.OperatorEqual, roleEnum.Engineer.String()))
	if err != nil {
		return err
	}

	activityFactory := s.activityFactory

	activities := []*entity.Activity{
		// users.register
		activityFactory.MustCreateWithOption(map[string]any{
			"Id":          "01KAJM7WQ4PAYGVCQRRGZ4MRFE",
			"CauserId":    &user.Id,
			"CauserType":  typeutil.Ptr(activityEnum.UserCauserType),
			"Action":      activityEnum.UsersRegister,
			"SubjectId":   &user.Id,
			"SubjectType": typeutil.Ptr(activityEnum.UserSubjectType),
		}).(*entity.Activity),

		// users.reset_password
		activityFactory.MustCreateWithOption(map[string]any{
			"Id":          "01KAJM7WQ5RZ6ZNQGKMND9VQK1",
			"CauserId":    &user.Id,
			"CauserType":  typeutil.Ptr(activityEnum.UserCauserType),
			"Action":      activityEnum.UsersResetPassword,
			"SubjectId":   &user.Id,
			"SubjectType": typeutil.Ptr(activityEnum.UserSubjectType),
		}).(*entity.Activity),

		// users.verify_email
		activityFactory.MustCreateWithOption(map[string]any{
			"Id":          "01KAJM7WQ56RVWNV80P74DH7RV",
			"CauserId":    &user.Id,
			"CauserType":  typeutil.Ptr(activityEnum.UserCauserType),
			"Action":      activityEnum.UsersVerifyEmail,
			"SubjectId":   &user.Id,
			"SubjectType": typeutil.Ptr(activityEnum.UserSubjectType),
		}).(*entity.Activity),

		// users.login
		activityFactory.MustCreateWithOption(map[string]any{
			"Id":          "01KAJM7WQ56WCBAR6BX8Y9GT20",
			"CauserId":    &user.Id,
			"CauserType":  typeutil.Ptr(activityEnum.UserCauserType),
			"Action":      activityEnum.UsersLogin,
			"SubjectId":   &user.Id,
			"SubjectType": typeutil.Ptr(activityEnum.UserSubjectType),
		}).(*entity.Activity),

		// users.logout
		activityFactory.MustCreateWithOption(map[string]any{
			"Id":          "01KAJM7WQ5HPN8HKEPTE5N1W8F",
			"CauserId":    &user.Id,
			"CauserType":  typeutil.Ptr(activityEnum.UserCauserType),
			"Action":      activityEnum.UsersLogout,
			"SubjectId":   &user.Id,
			"SubjectType": typeutil.Ptr(activityEnum.UserSubjectType),
		}).(*entity.Activity),

		// users.me
		activityFactory.MustCreateWithOption(map[string]any{
			"Id":          "01KAJM7WQ5JMGKT0D04C0ZZ44Q",
			"CauserId":    &user.Id,
			"CauserType":  typeutil.Ptr(activityEnum.UserCauserType),
			"Action":      activityEnum.UsersMe,
			"SubjectId":   &user.Id,
			"SubjectType": typeutil.Ptr(activityEnum.UserSubjectType),
		}).(*entity.Activity),

		// users.update_me_profile
		activityFactory.MustCreateWithOption(map[string]any{
			"Id":          "01KAJM7WQ5Q5DXWTR6MGD6XY3W",
			"CauserId":    &user.Id,
			"CauserType":  typeutil.Ptr(activityEnum.UserCauserType),
			"Action":      activityEnum.UsersUpdateMeProfile,
			"SubjectId":   &user.Id,
			"SubjectType": typeutil.Ptr(activityEnum.UserSubjectType),
		}).(*entity.Activity),

		// users.update_me_password
		activityFactory.MustCreateWithOption(map[string]any{
			"Id":          "01KAJM7WQ5ZQ4MDC9F5FT0NPGH",
			"CauserId":    &user.Id,
			"CauserType":  typeutil.Ptr(activityEnum.UserCauserType),
			"Action":      activityEnum.UsersUpdateMePassword,
			"SubjectId":   &user.Id,
			"SubjectType": typeutil.Ptr(activityEnum.UserSubjectType),
		}).(*entity.Activity),

		// users.index
		activityFactory.MustCreateWithOption(map[string]any{
			"Id":          "01KAJM7WQ5CTBDX17AP705DYM4",
			"CauserId":    &user.Id,
			"CauserType":  typeutil.Ptr(activityEnum.UserCauserType),
			"Action":      activityEnum.UsersIndex,
			"SubjectType": typeutil.Ptr(activityEnum.UserSubjectType),
		}).(*entity.Activity),

		// users.show
		activityFactory.MustCreateWithOption(map[string]any{
			"Id":          "01KAJM7WQ5X07KEJTA9MR165CY",
			"CauserId":    &user.Id,
			"CauserType":  typeutil.Ptr(activityEnum.UserCauserType),
			"Action":      activityEnum.UsersShow,
			"SubjectId":   &user.Id,
			"SubjectType": typeutil.Ptr(activityEnum.UserSubjectType),
		}).(*entity.Activity),

		// users.store
		activityFactory.MustCreateWithOption(map[string]any{
			"Id":          "01KAJM7WQ5VDWW98RSE01DWHCX",
			"CauserId":    &user.Id,
			"CauserType":  typeutil.Ptr(activityEnum.UserCauserType),
			"Action":      activityEnum.UsersStore,
			"SubjectId":   &user.Id,
			"SubjectType": typeutil.Ptr(activityEnum.UserSubjectType),
		}).(*entity.Activity),

		// users.update
		activityFactory.MustCreateWithOption(map[string]any{
			"Id":          "01KAJM7WQ5KP9GPR49WC0HTYN0",
			"CauserId":    &user.Id,
			"CauserType":  typeutil.Ptr(activityEnum.UserCauserType),
			"Action":      activityEnum.UsersUpdate,
			"SubjectId":   &user.Id,
			"SubjectType": typeutil.Ptr(activityEnum.UserSubjectType),
		}).(*entity.Activity),

		// users.toggle_activation
		activityFactory.MustCreateWithOption(map[string]any{
			"Id":          "01KAJM7WQ58T7XVXZWQ8NEFACX",
			"CauserId":    &user.Id,
			"CauserType":  typeutil.Ptr(activityEnum.UserCauserType),
			"Action":      activityEnum.UsersToggleActivation,
			"SubjectId":   &user.Id,
			"SubjectType": typeutil.Ptr(activityEnum.UserSubjectType),
		}).(*entity.Activity),

		// users.destroy
		activityFactory.MustCreateWithOption(map[string]any{
			"Id":          "01KAJM7WQ5FBS6PR89M1M4EQJ9",
			"CauserId":    &user.Id,
			"CauserType":  typeutil.Ptr(activityEnum.UserCauserType),
			"Action":      activityEnum.UsersDestroy,
			"SubjectId":   &user.Id,
			"SubjectType": typeutil.Ptr(activityEnum.UserSubjectType),
		}).(*entity.Activity),

		// permissions.index
		activityFactory.MustCreateWithOption(map[string]any{
			"Id":          "01KAJM7WQ59N6SN9KEJR0GCE5D",
			"CauserId":    &user.Id,
			"CauserType":  typeutil.Ptr(activityEnum.UserCauserType),
			"Action":      activityEnum.PermissionsIndex,
			"SubjectType": typeutil.Ptr(activityEnum.PermissionSubjectType),
		}).(*entity.Activity),

		// roles.index
		activityFactory.MustCreateWithOption(map[string]any{
			"Id":          "01KAJM7WQ5MV0HRQFMGJHJ6C1Z",
			"CauserId":    &user.Id,
			"CauserType":  typeutil.Ptr(activityEnum.UserCauserType),
			"Action":      activityEnum.RolesIndex,
			"SubjectType": typeutil.Ptr(activityEnum.RoleSubjectType),
		}).(*entity.Activity),

		// roles.show
		activityFactory.MustCreateWithOption(map[string]any{
			"Id":          "01KAJM7WQ5400N5XH805H28CE1",
			"CauserId":    &user.Id,
			"CauserType":  typeutil.Ptr(activityEnum.UserCauserType),
			"Action":      activityEnum.RolesShow,
			"SubjectId":   &roleEngineer.Id,
			"SubjectType": typeutil.Ptr(activityEnum.RoleSubjectType),
		}).(*entity.Activity),

		// roles.store
		activityFactory.MustCreateWithOption(map[string]any{
			"Id":          "01KAJM7WQ5K4X7Y63E1NEZQNVE",
			"CauserId":    &user.Id,
			"CauserType":  typeutil.Ptr(activityEnum.UserCauserType),
			"Action":      activityEnum.RolesStore,
			"SubjectId":   &roleEngineer.Id,
			"SubjectType": typeutil.Ptr(activityEnum.RoleSubjectType),
		}).(*entity.Activity),

		// roles.update
		activityFactory.MustCreateWithOption(map[string]any{
			"Id":          "01KAJM7WQ550NK895D8P4YAMC2",
			"CauserId":    &user.Id,
			"CauserType":  typeutil.Ptr(activityEnum.UserCauserType),
			"Action":      activityEnum.RolesUpdate,
			"SubjectId":   &roleEngineer.Id,
			"SubjectType": typeutil.Ptr(activityEnum.RoleSubjectType),
		}).(*entity.Activity),

		// roles.set_default
		activityFactory.MustCreateWithOption(map[string]any{
			"Id":          "01KAJM7WQ5ETSK8FFEKQ6KJENF",
			"CauserId":    &user.Id,
			"CauserType":  typeutil.Ptr(activityEnum.UserCauserType),
			"Action":      activityEnum.RolesSetDefault,
			"SubjectId":   &roleEngineer.Id,
			"SubjectType": typeutil.Ptr(activityEnum.RoleSubjectType),
		}).(*entity.Activity),

		// roles.destroy
		activityFactory.MustCreateWithOption(map[string]any{
			"Id":          "01KAJM7WQ5R98D72WR5H1KPKAR",
			"CauserId":    &user.Id,
			"CauserType":  typeutil.Ptr(activityEnum.UserCauserType),
			"Action":      activityEnum.RolesDestroy,
			"SubjectId":   &roleEngineer.Id,
			"SubjectType": typeutil.Ptr(activityEnum.RoleSubjectType),
		}).(*entity.Activity),
	}

	{
		code, err := s.codeRepository.FindByType(codeEnum.UserRegisterInvitation)
		if err != nil {
			return err
		}
		activities = append(activities,
			// codes.create_user_register_invitation
			activityFactory.MustCreateWithOption(map[string]any{
				"Id":          "01KAJM7WQ5F8T0N5FFFRHM0JP2",
				"CauserId":    &user.Id,
				"CauserType":  typeutil.Ptr(activityEnum.UserCauserType),
				"Action":      activityEnum.CodesCreateUserRegisterInvitation,
				"SubjectId":   &code.Id,
				"SubjectType": typeutil.Ptr(activityEnum.CodeSubjectType),
			}).(*entity.Activity),
		)
	}

	{
		code, err := s.codeRepository.FindByType(codeEnum.UserEmailVerification)
		if err != nil {
			return err
		}
		activities = append(activities,
			// codes.create_user_email_verification
			activityFactory.MustCreateWithOption(map[string]any{
				"Id":          "01KAJM7WQ5NCVXT09B6MBH0FH8",
				"CauserId":    typeutil.Ptr(gofakeit.Email()),
				"CauserType":  typeutil.Ptr(activityEnum.EmailCauserType),
				"Action":      activityEnum.CodesCreateUserEmailVerification,
				"SubjectId":   &code.Id,
				"SubjectType": typeutil.Ptr(activityEnum.CodeSubjectType),
			}).(*entity.Activity),
		)
	}

	{
		code, err := s.codeRepository.FindByType(codeEnum.UserResetPassword)
		if err != nil {
			return err
		}
		activities = append(activities,
			// codes.create_user_reset_password
			activityFactory.MustCreateWithOption(map[string]any{
				"Id":          "01KAJM7WQ5WB1RA7DK1FAVCH5J",
				"CauserId":    typeutil.Ptr(gofakeit.Email()),
				"CauserType":  typeutil.Ptr(activityEnum.EmailCauserType),
				"Action":      activityEnum.CodesCreateUserResetPassword,
				"SubjectId":   &code.Id,
				"SubjectType": typeutil.Ptr(activityEnum.CodeSubjectType),
			}).(*entity.Activity),
		)
	}

	if err := s.activityRepository.SaveMany(activities); err != nil {
		return err
	}
	return nil
}
