package action

import (
	"context"
	"errors"
	"net/http"

	roleRepository "github.com/arfanxn/welding/internal/module/role/domain/repository"
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	userRepository "github.com/arfanxn/welding/internal/module/user/domain/repository"
	"github.com/arfanxn/welding/internal/module/user/usecase/dto"
	"github.com/arfanxn/welding/pkg/errorutil"
	"github.com/gookit/goutil"
	"github.com/guregu/null/v6"
	"github.com/samber/lo"
	"gorm.io/gorm"
)

type SaveUserAction interface {
	Handle(ctx context.Context, user *entity.User, _dto *dto.SaveUser) (*entity.User, error)
}

type saveUserAction struct {
	userRepository userRepository.UserRepository
	roleRepository roleRepository.RoleRepository
}

func NewSaveUserAction(
	userRepository userRepository.UserRepository,
	roleRepository roleRepository.RoleRepository,
) SaveUserAction {
	return &saveUserAction{
		userRepository: userRepository,
		roleRepository: roleRepository,
	}
}

func (a *saveUserAction) Handle(ctx context.Context, user *entity.User, _dto *dto.SaveUser) (*entity.User, error) {
	var err error

	if user == nil {
		if _dto.Id.Valid {
			user, err = a.userRepository.Find(_dto.Id.String)
			if err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					return nil, errorutil.NewHttpError(http.StatusNotFound, "User not found", nil)
				}
				return nil, err
			}
		} else {
			user = entity.NewUser()
		}
	}

	if !goutil.IsEmpty(_dto.Name) {
		user.Name = _dto.Name
	}
	if !goutil.IsEmpty(_dto.PhoneNumber) {
		user.PhoneNumber = _dto.PhoneNumber
	}
	if !goutil.IsEmpty(_dto.Email) {
		if user.Email != _dto.Email {
			user.EmailVerifiedAt = null.TimeFromPtr(nil)
		}
		user.Email = _dto.Email
	}
	if !goutil.IsEmpty(_dto.Password) {
		if err := user.SetPassword(_dto.Password); err != nil {
			return nil, err
		}
	}
	if _dto.ActivatedAt.Valid {
		user.ActivatedAt = _dto.ActivatedAt
		user.DeactivatedAt = null.TimeFromPtr(nil)
	}
	if _dto.DeactivatedAt.Valid {
		user.ActivatedAt = null.TimeFromPtr(nil)
		user.DeactivatedAt = _dto.DeactivatedAt
	}

	user.SetEmploymentIdentityNumber(_dto.EmploymentIdentityNumber)

	if len(_dto.RoleIds) > 0 {
		user.Roles = lo.Map(_dto.RoleIds, func(roleId string, _ int) *entity.Role {
			return &entity.Role{Id: roleId}

		})
	}

	if err := a.userRepository.Save(user); err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, errorutil.NewHttpError(http.StatusConflict, "User already exists", nil)
		}
		return nil, err
	}

	return user, nil
}
