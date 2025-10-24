package usecase

import (
	"errors"
	"net/http"

	"github.com/arfanxn/welding/internal/infrastructure/http/jwt"
	"github.com/arfanxn/welding/internal/infrastructure/logger"
	roleRepository "github.com/arfanxn/welding/internal/module/role/domain/repository"
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	"github.com/arfanxn/welding/internal/module/user/domain/repository"
	"github.com/arfanxn/welding/internal/module/user/usecase/dto"
	"github.com/arfanxn/welding/pkg/errorutil"
	"github.com/arfanxn/welding/pkg/pagination"
	"github.com/arfanxn/welding/pkg/query"
	"github.com/arfanxn/welding/pkg/reflectutil"
	"github.com/gookit/goutil"
	"github.com/oklog/ulid/v2"
	"go.uber.org/fx"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var _ UserUsecase = (*userUsecase)(nil)

type UserUsecase interface {
	Register(user *entity.User) error
	Login(loginDto *dto.Login) (*dto.LoginResult, error)
	Find(q *query.Query) (*entity.User, error)
	Paginate(q *query.Query) (*pagination.OffsetPagination[*entity.User], error)
	Save(_dto *dto.SaveUser) (*entity.User, error)
	Destroy(_dto *dto.DestroyUser) error
}

type userUsecase struct {
	userRepository repository.UserRepository
	roleRepository roleRepository.RoleRepository
	jwtService     jwt.JWTService
	logger         *logger.Logger
}

type NewUserUsecaseParams struct {
	fx.In

	UserRepository repository.UserRepository
	RoleRepository roleRepository.RoleRepository
	JWTService     jwt.JWTService
	Logger         *logger.Logger
}

func NewUserUsecase(params NewUserUsecaseParams) UserUsecase {
	return &userUsecase{
		userRepository: params.UserRepository,
		roleRepository: params.RoleRepository,
		jwtService:     params.JWTService,
		logger:         params.Logger,
	}
}

func (u *userUsecase) Register(user *entity.User) error {
	err := u.userRepository.Save(user)
	return err
}

func (u *userUsecase) Login(loginDto *dto.Login) (*dto.LoginResult, error) {
	unauthorizedError := errorutil.NewHttpError(http.StatusUnauthorized, "Email atau password salah", nil)
	user, err := u.userRepository.FindByEmail(loginDto.Email)
	if err != nil {
		return nil, unauthorizedError
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginDto.Password)); err != nil {
		return nil, unauthorizedError
	}

	token, err := u.jwtService.CreateToken(user.Id)
	if err != nil {
		u.logger.Error("failed to create JWT token")
		return nil, unauthorizedError
	}

	return &dto.LoginResult{
		User:  user,
		Token: token,
	}, nil
}

func (u *userUsecase) Find(q *query.Query) (*entity.User, error) {
	users, err := u.userRepository.Get(q)
	if err != nil {
		return nil, err
	}
	if len(users) == 0 {
		return nil, errorutil.NewHttpError(http.StatusNotFound, "User not found", nil)
	}
	return users[0], nil
}

func (u *userUsecase) Paginate(q *query.Query) (*pagination.OffsetPagination[*entity.User], error) {
	return u.userRepository.Paginate(q)
}

func (u *userUsecase) Save(_dto *dto.SaveUser) (*entity.User, error) {
	var user *entity.User
	var err error

	if _dto.Id.IsZero() {
		user = entity.NewUser()
		user.Id = ulid.Make().String()
	} else {
		user, err = u.userRepository.Find(_dto.Id.String)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, errorutil.NewHttpError(http.StatusNotFound, "User not found", nil)
			}
			return nil, err
		}
	}

	user.Name = _dto.Name
	user.PhoneNumber = _dto.PhoneNumber
	user.Email = _dto.Email
	// Hash the user password
	hashedPasswordBytes, err := bcrypt.GenerateFromPassword([]byte(_dto.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user.Password = string(hashedPasswordBytes)

	if reflectutil.IsSlice(_dto.RoleIds) {
		roles, err := u.roleRepository.FindByIds(_dto.RoleIds)
		if err != nil {
			return nil, err
		}
		// Verify all requested roles were found
		if len(roles) != len(_dto.RoleIds) {
			return nil, errorutil.NewHttpError(
				http.StatusBadRequest,
				"One or more roles not found",
				nil,
			)
		}
		user.Roles = roles
	}

	if err := u.userRepository.Save(user); err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, errorutil.NewHttpError(http.StatusConflict, "User already exists", nil)
		}
		return nil, err
	}

	return user, nil
}

func (u *userUsecase) Destroy(_dto *dto.DestroyUser) error {
	user, err := u.userRepository.Find(_dto.Id)
	if err != nil {
		if goutil.IsNil(user) {
			return errorutil.NewHttpError(http.StatusNotFound, "User not found", nil)
		}
		return err
	}

	return u.userRepository.Destroy(user)
}
