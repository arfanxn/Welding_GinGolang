package usecase

import (
	"context"
	"time"

	"github.com/arfanxn/welding/internal/infrastructure/http/jwt"
	"github.com/arfanxn/welding/internal/infrastructure/logger"
	"github.com/arfanxn/welding/internal/infrastructure/security"
	"github.com/arfanxn/welding/internal/module/code/domain/enum"
	codeRepository "github.com/arfanxn/welding/internal/module/code/domain/repository"
	roleRepository "github.com/arfanxn/welding/internal/module/role/domain/repository"
	"github.com/arfanxn/welding/internal/module/shared/contextkey"
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	"github.com/arfanxn/welding/internal/module/shared/domain/errorx"
	"github.com/arfanxn/welding/internal/module/user/domain/repository"
	"github.com/arfanxn/welding/internal/module/user/infrastructure/policy"
	"github.com/arfanxn/welding/internal/module/user/usecase/dto"
	"github.com/arfanxn/welding/internal/module/user/usecase/step"
	"github.com/arfanxn/welding/pkg/pagination"
	"github.com/arfanxn/welding/pkg/query"
	"github.com/guregu/null/v6"
	"go.uber.org/fx"
)

var _ UserUsecase = (*userUsecase)(nil)

type UserUsecase interface {
	Register(ctx context.Context, _dto *dto.Register) (*entity.User, error)
	VerifyEmail(ctx context.Context, verifyDto *dto.VerifyEmail) (*entity.User, error)
	ResetPassword(ctx context.Context, _dto *dto.ResetPassword) (*entity.User, error)
	Login(ctx context.Context, loginDto *dto.Login) (*dto.LoginResult, error)
	Show(ctx context.Context, q *query.Query) (*entity.User, error)
	Paginate(ctx context.Context, q *query.Query) (*pagination.OffsetPagination[*entity.User], error)
	Store(ctx context.Context, _dto *dto.SaveUser) (*entity.User, error)
	Update(ctx context.Context, _dto *dto.SaveUser) (*entity.User, error)
	UpdateMePassword(ctx context.Context, _dto *dto.UpdateUserMePassword) (*entity.User, error)
	// ! Deprecated
	// UpdatePassword(ctx context.Context, _dto *dto.UpdateUserPassword) (*entity.User, error)
	ToggleActivation(ctx context.Context, _dto *dto.ToggleActivation) (*entity.User, error)
	Destroy(ctx context.Context, _dto *dto.DestroyUser) error
}

type userUsecase struct {
	registerUserStep step.RegisterUserStep
	saveUserStep     step.SaveUserStep

	userPolicy     policy.UserPolicy
	userRepository repository.UserRepository
	roleRepository roleRepository.RoleRepository
	codeRepository codeRepository.CodeRepository

	jwtService      jwt.JWTService
	passwordService security.PasswordService
	logger          *logger.Logger
}

type NewUserUsecaseParams struct {
	fx.In

	RegisterUserStep step.RegisterUserStep
	SaveUserStep     step.SaveUserStep

	UserPolicy     policy.UserPolicy
	UserRepository repository.UserRepository
	RoleRepository roleRepository.RoleRepository
	CodeRepository codeRepository.CodeRepository

	JWTService      jwt.JWTService
	PasswordService security.PasswordService
	Logger          *logger.Logger
}

func NewUserUsecase(params NewUserUsecaseParams) UserUsecase {
	return &userUsecase{
		registerUserStep: params.RegisterUserStep,
		saveUserStep:     params.SaveUserStep,

		userPolicy:     params.UserPolicy,
		userRepository: params.UserRepository,
		roleRepository: params.RoleRepository,
		codeRepository: params.CodeRepository,

		jwtService:      params.JWTService,
		passwordService: params.PasswordService,
		logger:          params.Logger,
	}
}

// Register handles user registration with optional invitation code validation
// 1. If an invitation code is provided, it validates the code and extracts role information
// 2. Creates a new user with the provided details
// 3. If an invitation code was used, marks it as used
func (u *userUsecase) Register(ctx context.Context, _dto *dto.Register) (*entity.User, error) {
	user, err := u.registerUserStep.Handle(ctx, _dto)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userUsecase) VerifyEmail(ctx context.Context, _dto *dto.VerifyEmail) (*entity.User, error) {
	code, err := u.codeRepository.FindByCodeableAndTypeAndValue(
		_dto.Email,
		"email",
		enum.UserEmailVerification,
		_dto.Code,
	)
	if err != nil {
		return nil, err
	}

	if code.IsUsed() {
		return nil, errorx.ErrCodeAlreadyUsed
	}

	if code.IsExpired() {
		return nil, errorx.ErrCodeExpired
	}

	user, err := u.userRepository.FindByEmail(_dto.Email)
	if err != nil {
		return nil, err
	}

	if user.IsEmailVerified() {
		return nil, errorx.ErrUserEmailAlreadyVerified
	}

	user.EmailVerifiedAt = null.TimeFrom(time.Now())
	if err := u.userRepository.Save(user); err != nil {
		return nil, err
	}

	code.UsedAt = null.TimeFrom(time.Now())
	if err := u.codeRepository.Save(code); err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userUsecase) ResetPassword(ctx context.Context, _dto *dto.ResetPassword) (*entity.User, error) {
	code, err := u.codeRepository.FindByCodeableAndTypeAndValue(
		_dto.Email,
		"email",
		enum.UserResetPassword,
		_dto.Code,
	)
	if err != nil {
		return nil, err
	}

	if code.IsUsed() {
		return nil, errorx.ErrCodeAlreadyUsed
	}

	if code.IsExpired() {
		return nil, errorx.ErrCodeExpired
	}

	user, err := u.userRepository.FindByEmail(_dto.Email)
	if err != nil {
		return nil, err
	}

	user.Password, err = u.passwordService.Hash(_dto.Password)
	if err != nil {
		return nil, err
	}

	if err := u.userRepository.Save(user); err != nil {
		return nil, err
	}

	code.UsedAt = null.TimeFrom(time.Now())
	if err := u.codeRepository.Save(code); err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userUsecase) Login(ctx context.Context, loginDto *dto.Login) (*dto.LoginResult, error) {
	user, err := u.userRepository.FindByEmail(loginDto.Email)
	if err != nil {
		return nil, errorx.ErrUserPasswordIncorrect
	}

	if err = u.passwordService.Check(user.Password, loginDto.Password); err != nil {
		return nil, errorx.ErrUserPasswordIncorrect
	}

	token, err := u.jwtService.CreateToken(user.Id)
	if err != nil {
		return nil, errorx.ErrUserPasswordIncorrect
	}

	return &dto.LoginResult{
		User:  user,
		Token: token,
	}, nil
}

func (u *userUsecase) Show(ctx context.Context, q *query.Query) (*entity.User, error) {
	user, err := u.userRepository.First(q)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userUsecase) Paginate(ctx context.Context, q *query.Query) (*pagination.OffsetPagination[*entity.User], error) {
	return u.userRepository.Paginate(q)
}

func (u *userUsecase) Store(ctx context.Context, _dto *dto.SaveUser) (*entity.User, error) {
	if err := u.userPolicy.Store(ctx, _dto); err != nil {
		return nil, err
	}

	return u.saveUserStep.Handle(ctx, _dto)
}

func (u *userUsecase) Update(ctx context.Context, _dto *dto.SaveUser) (*entity.User, error) {
	if err := u.userPolicy.Update(ctx, _dto); err != nil {
		return nil, err
	}

	return u.saveUserStep.Handle(ctx, _dto)
}

func (u *userUsecase) UpdateMePassword(ctx context.Context, _dto *dto.UpdateUserMePassword) (*entity.User, error) {
	userId := ctx.Value(contextkey.UserIdKey).(string)

	err := u.userPolicy.UpdateMePassword(ctx, _dto)
	if err != nil {
		return nil, err
	}

	return u.saveUserStep.Handle(ctx, &dto.SaveUser{
		Id:       &userId,
		Password: &_dto.Password,
	})
}

/*
! Deprecated
func (u *userUsecase) UpdatePassword(ctx context.Context, _dto *dto.UpdateUserPassword) (*entity.User, error) {
	user, err := u.userPolicy.UpdatePassword(ctx, _dto)
	if err != nil {
		return nil, err
	}

	user.Password, err = u.passwordService.Hash(_dto.Password)
	if err != nil {
		return nil, err
	}

	err = u.userRepository.Save(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
*/

func (u *userUsecase) ToggleActivation(ctx context.Context, _dto *dto.ToggleActivation) (*entity.User, error) {
	if err := u.userPolicy.ToggleActivation(ctx, _dto); err != nil {
		return nil, err
	}

	user, err := u.userRepository.Find(_dto.Id)
	if err != nil {
		return nil, err
	}

	return u.userRepository.ToggleActivation(user)
}

func (u *userUsecase) Destroy(ctx context.Context, _dto *dto.DestroyUser) error {
	if err := u.userPolicy.Destroy(ctx, _dto); err != nil {
		return err
	}

	user, err := u.userRepository.Find(_dto.Id)
	if err != nil {
		return err
	}

	return u.userRepository.Destroy(user)
}
