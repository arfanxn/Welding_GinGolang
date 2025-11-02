package usecase

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/arfanxn/welding/internal/infrastructure/http/jwt"
	"github.com/arfanxn/welding/internal/infrastructure/logger"
	"github.com/arfanxn/welding/internal/module/code/domain/enum"
	codeRepository "github.com/arfanxn/welding/internal/module/code/domain/repository"
	roleRepository "github.com/arfanxn/welding/internal/module/role/domain/repository"
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	"github.com/arfanxn/welding/internal/module/user/domain/repository"
	"github.com/arfanxn/welding/internal/module/user/infrastructure/policy"
	"github.com/arfanxn/welding/internal/module/user/usecase/dto"
	"github.com/arfanxn/welding/pkg/errorutil"
	"github.com/arfanxn/welding/pkg/pagination"
	"github.com/arfanxn/welding/pkg/query"
	"github.com/guregu/null/v6"
	"go.uber.org/fx"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var _ UserUsecase = (*userUsecase)(nil)

type UserUsecase interface {
	Register(ctx context.Context, _dto *dto.Register) (*entity.User, error)
	VerifyEmail(ctx context.Context, verifyDto *dto.VerifyEmail) (*entity.User, error)
	ResetPassword(ctx context.Context, _dto *dto.ResetPassword) (*entity.User, error)
	Login(ctx context.Context, loginDto *dto.Login) (*dto.LoginResult, error)
	Find(ctx context.Context, q *query.Query) (*entity.User, error)
	Paginate(ctx context.Context, q *query.Query) (*pagination.OffsetPagination[*entity.User], error)
	Save(ctx context.Context, _dto *dto.SaveUser) (*entity.User, error)
	ToggleActivation(ctx context.Context, _dto *dto.ToggleActivation) (*entity.User, error)
	Destroy(ctx context.Context, _dto *dto.DestroyUser) error
}

type userUsecase struct {
	userPolicy     policy.UserPolicy
	userRepository repository.UserRepository
	roleRepository roleRepository.RoleRepository
	codeRepository codeRepository.CodeRepository
	jwtService     jwt.JWTService
	logger         *logger.Logger
}

type NewUserUsecaseParams struct {
	fx.In

	UserPolicy     policy.UserPolicy
	UserRepository repository.UserRepository
	RoleRepository roleRepository.RoleRepository
	CodeRepository codeRepository.CodeRepository
	JWTService     jwt.JWTService
	Logger         *logger.Logger
}

func NewUserUsecase(params NewUserUsecaseParams) UserUsecase {
	return &userUsecase{
		userPolicy:     params.UserPolicy,
		userRepository: params.UserRepository,
		roleRepository: params.RoleRepository,
		codeRepository: params.CodeRepository,
		jwtService:     params.JWTService,
		logger:         params.Logger,
	}
}

// Register handles user registration with optional invitation code validation
// 1. If an invitation code is provided, it validates the code and extracts role information
// 2. Creates a new user with the provided details
// 3. If an invitation code was used, marks it as used
func (u *userUsecase) Register(ctx context.Context, _dto *dto.Register) (*entity.User, error) {
	var (
		code    *entity.Code
		err     error
		roleIds = []string{} // Initialize empty role IDs slice
	)

	// 1. Handle invitation code if provided
	if !_dto.InvitationCode.IsZero() {
		// 1.1 Find the code in the repository
		code, err = u.codeRepository.FindByTypeAndValue(enum.UserRegisterInvitation, _dto.InvitationCode.String)
		if err != nil {
			// 1.2 Handle code not found error
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, errorutil.NewHttpError(http.StatusNotFound, "Kode undangan tidak ditemukan", nil)
			}
			return nil, err
		}

		// 1.3 Validate code status
		if code.IsUsed() {
			return nil, errorutil.NewHttpError(http.StatusBadRequest, "Kode undangan sudah digunakan", nil)
		}

		// 1.4 Check code expiration
		if code.IsExpired() {
			return nil, errorutil.NewHttpError(http.StatusBadRequest, "Kode undangan sudah kadaluarsa", nil)
		}

		// 1.5 Extract role information from code metadata
		codeMeta, err := code.GetMeta()
		if err != nil {
			return nil, err
		}

		// 1.6 Get role ID from code metadata and assign to user
		roleId := codeMeta["role_id"].(string)
		roleIds = []string{roleId}
	}

	// 2. Create user with provided details
	saveUserDto := &dto.SaveUser{
		Name:                     _dto.Name,
		PhoneNumber:              _dto.PhoneNumber,
		Email:                    _dto.Email,
		Password:                 _dto.Password,
		RoleIds:                  roleIds, // Includes role from invitation if provided
		EmploymentIdentityNumber: _dto.EmploymentIdentityNumber,
	}

	// 2.1 Save the new user
	user, err := u.Save(ctx, saveUserDto)
	if err != nil {
		return nil, err
	}

	// 3. If invitation code was used, mark it as used
	if !_dto.InvitationCode.IsZero() {
		code.UsedAt = null.TimeFrom(time.Now())
		if err := u.codeRepository.Save(code); err != nil {
			return nil, err
		}
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
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorutil.NewHttpError(http.StatusBadRequest, "Kode verifikasi salah", nil)
		}
		return nil, err
	}

	if code.IsUsed() {
		return nil, errorutil.NewHttpError(http.StatusBadRequest, "Kode verifikasi sudah digunakan", nil)
	}

	if code.IsExpired() {
		return nil, errorutil.NewHttpError(http.StatusBadRequest, "Kode verifikasi sudah kadaluarsa", nil)
	}

	user, err := u.userRepository.FindByEmail(_dto.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorutil.NewHttpError(http.StatusNotFound, "User tidak ditemukan", nil)
		}
		return nil, err
	}

	if user.IsEmailVerified() {
		return nil, errorutil.NewHttpError(http.StatusBadRequest, "Email sudah diverifikasi", nil)
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
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorutil.NewHttpError(http.StatusBadRequest, "Kode reset password salah", nil)
		}
		return nil, err
	}

	if code.IsUsed() {
		return nil, errorutil.NewHttpError(http.StatusBadRequest, "Kode reset password sudah digunakan", nil)
	}

	if code.IsExpired() {
		return nil, errorutil.NewHttpError(http.StatusBadRequest, "Kode reset password sudah kadaluarsa", nil)
	}

	user, err := u.userRepository.FindByEmail(_dto.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorutil.NewHttpError(http.StatusNotFound, "User tidak ditemukan", nil)
		}
		return nil, err
	}

	err = user.SetPassword(_dto.Password)
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
		return nil, unauthorizedError
	}

	return &dto.LoginResult{
		User:  user,
		Token: token,
	}, nil
}

func (u *userUsecase) Find(ctx context.Context, q *query.Query) (*entity.User, error) {
	users, err := u.userRepository.Get(q)
	if err != nil {
		return nil, err
	}
	if len(users) == 0 {
		return nil, errorutil.NewHttpError(http.StatusNotFound, "User tidak ditemukan", nil)
	}
	return users[0], nil
}

func (u *userUsecase) Paginate(ctx context.Context, q *query.Query) (*pagination.OffsetPagination[*entity.User], error) {
	return u.userRepository.Paginate(q)
}

func (u *userUsecase) Save(ctx context.Context, _dto *dto.SaveUser) (*entity.User, error) {
	user, err := u.userPolicy.Save(ctx, _dto)
	if err != nil {
		return nil, err
	}

	user.Name = _dto.Name
	user.PhoneNumber = _dto.PhoneNumber
	if user.Email != _dto.Email {
		user.EmailVerifiedAt = null.TimeFromPtr(nil)
	}
	user.Email = _dto.Email
	if err := user.SetPassword(_dto.Password); err != nil {
		return nil, err
	}
	user.ActivatedAt = null.TimeFrom(time.Now())
	user.DeactivatedAt = null.TimeFromPtr(nil) // Set to null

	if _dto.EmploymentIdentityNumber.Valid {
		user.Employee = &entity.Employee{
			UserId:                   user.Id,
			EmploymentIdentityNumber: _dto.EmploymentIdentityNumber.String,
		}
	} else {
		user.Employee = nil
	}

	if err := u.userRepository.Save(user); err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, errorutil.NewHttpError(http.StatusConflict, "User already exists", nil)
		}
		return nil, err
	}

	return user, nil
}

func (u *userUsecase) ToggleActivation(ctx context.Context, _dto *dto.ToggleActivation) (*entity.User, error) {
	user, err := u.userPolicy.ToggleActivation(ctx, _dto)
	if err != nil {
		return nil, err
	}

	return u.userRepository.ToggleActivation(user)
}

func (u *userUsecase) Destroy(ctx context.Context, _dto *dto.DestroyUser) error {
	user, err := u.userPolicy.Destroy(ctx, _dto)
	if err != nil {
		return err
	}

	return u.userRepository.Destroy(user)
}
