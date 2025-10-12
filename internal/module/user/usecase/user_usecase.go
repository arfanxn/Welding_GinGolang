package usecase

import (
	"net/http"

	"github.com/arfanxn/welding/internal/infrastructure/http/jwt"
	"github.com/arfanxn/welding/internal/infrastructure/logger"
	"github.com/arfanxn/welding/internal/module/user/domain/entity"
	"github.com/arfanxn/welding/internal/module/user/domain/repository"
	"github.com/arfanxn/welding/internal/module/user/usecase/dto"
	"github.com/arfanxn/welding/pkg/errors"
	"go.uber.org/fx"
	"golang.org/x/crypto/bcrypt"
)

var _ UserUsecase = (*userUsecase)(nil)

type UserUsecase interface {
	Login(*dto.Login) (*dto.LoginResult, error)
	Register(user *entity.User) error
}

type userUsecase struct {
	userRepository repository.UserRepository
	jwtService     jwt.JWTService
	logger         *logger.Logger
}

type NewUserUsecaseParams struct {
	fx.In

	UserRepository repository.UserRepository
	JWTService     jwt.JWTService
	Logger         *logger.Logger
}

func NewUserUsecase(params NewUserUsecaseParams) UserUsecase {
	return &userUsecase{
		userRepository: params.UserRepository,
		jwtService:     params.JWTService,
		logger:         params.Logger,
	}
}

func (u *userUsecase) Register(user *entity.User) error {
	err := u.userRepository.Store(user)
	return err
}

func (u *userUsecase) Login(loginDto *dto.Login) (*dto.LoginResult, error) {
	unauthorizedError := errors.NewHttpError(http.StatusUnauthorized, "Email atau password salah", nil)
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
