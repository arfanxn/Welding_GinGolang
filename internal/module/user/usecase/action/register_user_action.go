package action

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/arfanxn/welding/internal/module/code/domain/enum"
	codeRepository "github.com/arfanxn/welding/internal/module/code/domain/repository"
	roleRepository "github.com/arfanxn/welding/internal/module/role/domain/repository"
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	userRepository "github.com/arfanxn/welding/internal/module/user/domain/repository"
	"github.com/arfanxn/welding/internal/module/user/usecase/dto"
	"github.com/arfanxn/welding/pkg/errorutil"
	"github.com/guregu/null/v6"
	"gorm.io/gorm"
)

type RegisterUserAction interface {
	Handle(ctx context.Context, user *entity.User, _dto *dto.Register) (*entity.User, error)
}

type registerUserAction struct {
	saveUserAction SaveUserAction

	userRepository userRepository.UserRepository
	codeRepository codeRepository.CodeRepository
	roleRepository roleRepository.RoleRepository
}

func NewRegisterUserAction(
	saveUserAction SaveUserAction,

	userRepository userRepository.UserRepository,
	codeRepository codeRepository.CodeRepository,
	roleRepository roleRepository.RoleRepository,
) RegisterUserAction {
	return &registerUserAction{
		saveUserAction: saveUserAction,

		userRepository: userRepository,
		codeRepository: codeRepository,
		roleRepository: roleRepository,
	}
}

func (a *registerUserAction) Handle(
	ctx context.Context,
	user *entity.User,
	_dto *dto.Register,
) (*entity.User, error) {
	var (
		err                  error
		roleIds              []string
		isWithInvitationCode bool = _dto.InvitationCode.Valid
		code                 *entity.Code
	)

	if isWithInvitationCode {
		code, err = a.codeRepository.FindByTypeAndValue(enum.UserRegisterInvitation, _dto.InvitationCode.String)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, errorutil.NewHttpError(http.StatusNotFound, "Kode undangan tidak ditemukan", nil)
			}
			return nil, err
		}

		if code.IsUsed() {
			return nil, errorutil.NewHttpError(http.StatusBadRequest, "Kode undangan sudah digunakan", nil)
		}

		if code.IsExpired() {
			return nil, errorutil.NewHttpError(http.StatusBadRequest, "Kode undangan sudah kadaluarsa", nil)
		}

		codeMeta, err := code.GetMeta()
		if err != nil {
			return nil, err
		}

		roleId := codeMeta["role_id"].(string)
		roleIds = []string{roleId}
	} else {
		defaultRole, err := a.roleRepository.FindDefault()
		if err != nil {
			// TODO: return custom error on repository instead of gorm's error
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, errorutil.NewHttpError(http.StatusBadRequest, "Default role is not configured", nil)
			}
			return nil, err
		}
		roleIds = []string{defaultRole.Id}
	}

	if user == nil {
		user = entity.NewUser()
	}

	user, err = a.saveUserAction.Handle(ctx, user, &dto.SaveUser{
		Name:                     _dto.Name,
		PhoneNumber:              _dto.PhoneNumber,
		Email:                    _dto.Email,
		Password:                 _dto.Password,
		ActivatedAt:              null.TimeFrom(time.Now()),
		RoleIds:                  roleIds,
		EmploymentIdentityNumber: _dto.EmploymentIdentityNumber,
	})

	if isWithInvitationCode {
		code.MarkUsed()
		if err := a.codeRepository.Save(code); err != nil {
			return nil, err
		}
	}

	return user, err
}
