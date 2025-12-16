package seeder

import (
	codeEnum "github.com/arfanxn/welding/internal/module/code/domain/enum"
	codeRepository "github.com/arfanxn/welding/internal/module/code/domain/repository"
	roleEnum "github.com/arfanxn/welding/internal/module/role/domain/enum"
	roleRepository "github.com/arfanxn/welding/internal/module/role/domain/repository"
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	userRepository "github.com/arfanxn/welding/internal/module/user/domain/repository"
	"github.com/arfanxn/welding/pkg/query"
	"github.com/arfanxn/welding/pkg/types"
	"github.com/bluele/factory-go/factory"
	"github.com/guregu/null/v6"
	"go.uber.org/fx"
)

type CodeSeeder struct {
	codeFactory    *factory.Factory
	codeRepository codeRepository.CodeRepository
	userRepository userRepository.UserRepository
	roleRepository roleRepository.RoleRepository
}

type NewCodeSeederParams struct {
	fx.In

	CodeFactory    *factory.Factory `name:"code_factory"`
	UserRepository userRepository.UserRepository
	CodeRepository codeRepository.CodeRepository
	RoleRepository roleRepository.RoleRepository
}

func NewCodeSeeder(
	params NewCodeSeederParams,
) Seeder {
	return &CodeSeeder{
		codeFactory:    params.CodeFactory,
		codeRepository: params.CodeRepository,
		userRepository: params.UserRepository,
		roleRepository: params.RoleRepository,
	}
}

func (s *CodeSeeder) Seed() error {
	user, err := s.userRepository.First(nil)
	if err != nil {
		return err
	}

	roleEngineer, err := s.roleRepository.First(query.NewQuery().Filter("name", query.OperatorEqual, roleEnum.Engineer))
	if err != nil {
		return err
	}

	codeFactory := s.codeFactory

	codes := []*entity.Code{
		codeFactory.MustCreateWithOption(map[string]any{
			"CodeableId":   null.StringFromPtr(nil),
			"CodeableType": null.StringFromPtr(nil),
			"Type":         codeEnum.UserRegisterInvitation,
			"Meta": types.JSONMap{
				"role_id": roleEngineer.Id,
			},
		}).(*entity.Code),
		codeFactory.MustCreateWithOption(map[string]any{
			"CodeableId":   null.StringFrom(user.Email),
			"CodeableType": null.StringFrom("email"),
			"Type":         codeEnum.UserEmailVerification,
		}).(*entity.Code),
		codeFactory.MustCreateWithOption(map[string]any{
			"CodeableId":   null.StringFrom(user.Email),
			"CodeableType": null.StringFrom("email"),
			"Type":         codeEnum.UserResetPassword,
		}).(*entity.Code),
	}

	if err := s.codeRepository.SaveMany(codes); err != nil {
		return err
	}

	return nil
}
