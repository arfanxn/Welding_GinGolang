package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/arfanxn/welding/internal/infrastructure/id"
	"github.com/arfanxn/welding/internal/infrastructure/logger"
	"github.com/arfanxn/welding/internal/infrastructure/mail"
	activityEnum "github.com/arfanxn/welding/internal/module/activity/domain/enum"
	activityDto "github.com/arfanxn/welding/internal/module/activity/usecase/dto"
	activityService "github.com/arfanxn/welding/internal/module/activity/usecase/service"
	"github.com/arfanxn/welding/internal/module/code/domain/enum"
	"github.com/arfanxn/welding/internal/module/code/domain/repository"
	"github.com/arfanxn/welding/internal/module/code/infrastructure/policy"
	"github.com/arfanxn/welding/internal/module/code/usecase/dto"
	codeService "github.com/arfanxn/welding/internal/module/code/usecase/service"
	roleRepository "github.com/arfanxn/welding/internal/module/role/domain/repository"
	"github.com/arfanxn/welding/internal/module/shared/contextkey"
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	"github.com/arfanxn/welding/pkg/types"
	"github.com/arfanxn/welding/pkg/typeutil"
	"github.com/guregu/null/v6"
	"go.uber.org/zap"
)

type CodeUsecase interface {
	CreateUserRegisterInvitation(ctx context.Context, _dto *dto.CreateUserRegisterInvitation) (*entity.Code, error)
	CreateUserEmailVerification(ctx context.Context, _dto *dto.CreateUserEmailVerification) (*entity.Code, error)
	CreateUserResetPassword(ctx context.Context, _dto *dto.CreateUserResetPassword) (*entity.Code, error)
}

type codeUsecase struct {
	idService       id.IdService
	codeService     codeService.CodeService
	activityService activityService.ActivityService
	logger          *logger.Logger
	codePolicy      policy.CodePolicy
	codeRepository  repository.CodeRepository
	roleRepository  roleRepository.RoleRepository
	mailService     mail.MailService
}

func NewCodeUsecase(
	idService id.IdService,
	codeService codeService.CodeService,
	activityService activityService.ActivityService,
	logger *logger.Logger,
	codePolicy policy.CodePolicy,
	codeRepository repository.CodeRepository,
	roleRepository roleRepository.RoleRepository,
	mailService mail.MailService,
) CodeUsecase {
	return &codeUsecase{
		idService:       idService,
		codeService:     codeService,
		activityService: activityService,
		logger:          logger,
		codePolicy:      codePolicy,
		codeRepository:  codeRepository,
		roleRepository:  roleRepository,
		mailService:     mailService,
	}
}

// CreateUserRegisterInvitation generates a new user registration invitation code with the specified role and expiration.
// It performs the following steps:
// 1. Validates the invitation request using the code policy
// 2. Verifies the existence of the specified role
// 3. Creates a new invitation code with type UserRegisterInvitation
// 4. Associates the role ID with the code using metadata
// 5. Sets the expiration time for the invitation
// 6. Persists the code to the repository
// 7. Records the activity log for the invitation creation
//
// Parameters:
//   - ctx: Context for request-scoped values, cancellation signals, and deadlines
//   - _dto: Data transfer object containing invitation details (role ID and expiration time)
//
// Returns:
//   - *entity.Code: The created invitation code with associated metadata
//   - error: Returns an error if:
//   - The role is not found
//   - The invitation code generation fails
//   - There's an error persisting the code
//   - The activity log cannot be recorded
func (s *codeUsecase) CreateUserRegisterInvitation(ctx context.Context, _dto *dto.CreateUserRegisterInvitation) (*entity.Code, error) {
	var err error

	// Validate the invitation request
	err = s.codePolicy.CreateUserRegisterInvitation(ctx, _dto)
	if err != nil {
		return nil, err
	}

	// Create new invitation code
	code := &entity.Code{}
	code.Id = s.idService.Generate()
	code.Value = s.codeService.Generate()
	code.Type = enum.UserRegisterInvitation
	// Store role ID in metadata for later reference
	code.Meta = types.JSONMap{
		"role_id": _dto.RoleId,
	}
	// Set when the invitation will expire
	code.ExpiredAt = _dto.ExpiredAt

	// Save the invitation code to the repository
	err = s.codeRepository.Save(code)
	if err != nil {
		return nil, err
	}

	if _, err = s.activityService.Record(ctx, &activityDto.CreateActivity{
		Causer:  contextkey.GetUser(ctx),
		Action:  activityEnum.CodesCreateUserRegisterInvitation,
		Subject: code,
	}); err != nil {
		return nil, err
	}

	return code, nil
}

func (s *codeUsecase) CreateUserEmailVerification(ctx context.Context, _dto *dto.CreateUserEmailVerification) (*entity.Code, error) {
	var err error

	code := &entity.Code{}
	code.Id = s.idService.Generate()
	code.Value = s.codeService.Generate()
	code.Type = enum.UserEmailVerification
	code.CodeableId = null.StringFrom(_dto.Email)
	code.CodeableType = null.StringFrom("email")
	code.ExpiredAt = time.Now().Add(time.Minute * 30)

	err = s.codeRepository.Save(code)
	if err != nil {
		return nil, err
	}

	// TODO: move this to a job queue, and monitor the job queue
	go func(email, subject, body string) {
		err := s.mailService.Send([]string{email}, subject, body)
		if err != nil {
			s.logger.Error("Failed to send email verification code", zap.Error(err))
		}
	}(
		code.CodeableId.String,
		"Verifikasi Email", fmt.Sprintf("Kode verifikasi email Anda adalah %s, berlaku selama 30 menit.",
			code.Value),
	)

	if _, err = s.activityService.Record(ctx, &activityDto.CreateActivity{
		CauserId:   &_dto.Email,
		CauserType: typeutil.Ptr(activityEnum.EmailCauserType),
		Action:     activityEnum.CodesCreateUserEmailVerification,
		Subject:    code,
	}); err != nil {
		return nil, err
	}

	return code, nil
}

func (s *codeUsecase) CreateUserResetPassword(ctx context.Context, _dto *dto.CreateUserResetPassword) (*entity.Code, error) {
	var err error

	code := &entity.Code{}
	code.Id = s.idService.Generate()
	code.Value = s.codeService.Generate()
	code.Type = enum.UserResetPassword
	code.CodeableId = null.StringFrom(_dto.Email)
	code.CodeableType = null.StringFrom("email")
	code.ExpiredAt = time.Now().Add(time.Minute * 30)

	err = s.codeRepository.Save(code)
	if err != nil {
		return nil, err
	}

	// TODO: move this to a job queue, and monitor the job queue
	go func(email, subject, body string) {
		err := s.mailService.Send([]string{email}, subject, body)
		if err != nil {
			s.logger.Error("Failed to send reset password code", zap.Error(err))
		}
	}(
		code.CodeableId.String,
		"Reset Password", fmt.Sprintf("Kode reset password Anda adalah %s, berlaku selama 30 menit.",
			code.Value),
	)

	if _, err = s.activityService.Record(ctx, &activityDto.CreateActivity{
		CauserId:   &_dto.Email,
		CauserType: typeutil.Ptr(activityEnum.EmailCauserType),
		Action:     activityEnum.CodesCreateUserResetPassword,
		Subject:    code,
	}); err != nil {
		return nil, err
	}

	return code, nil
}
