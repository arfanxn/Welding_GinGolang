package usecase

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/arfanxn/welding/internal/infrastructure/logger"
	"github.com/arfanxn/welding/internal/infrastructure/mail"
	"github.com/arfanxn/welding/internal/module/code/domain/enum"
	"github.com/arfanxn/welding/internal/module/code/domain/repository"
	"github.com/arfanxn/welding/internal/module/code/infrastructure/policy"
	"github.com/arfanxn/welding/internal/module/code/usecase/dto"
	roleRepository "github.com/arfanxn/welding/internal/module/role/domain/repository"
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	"github.com/arfanxn/welding/pkg/errorutil"
	"github.com/guregu/null/v6"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type CodeUsecase interface {
	CreateUserRegisterInvitation(ctx context.Context, _dto *dto.CreateUserRegisterInvitation) (*entity.Code, error)
	CreateUserEmailVerification(ctx context.Context, _dto *dto.CreateUserEmailVerification) (*entity.Code, error)
	CreateUserResetPassword(ctx context.Context, _dto *dto.CreateUserResetPassword) (*entity.Code, error)
}

type codeUsecase struct {
	logger         *logger.Logger
	codePolicy     policy.CodePolicy
	codeRepository repository.CodeRepository
	roleRepository roleRepository.RoleRepository
	mailService    mail.MailService
}

func NewCodeUsecase(
	logger *logger.Logger,
	codePolicy policy.CodePolicy,
	codeRepository repository.CodeRepository,
	roleRepository roleRepository.RoleRepository,
	mailService mail.MailService,
) CodeUsecase {
	return &codeUsecase{
		logger:         logger,
		codePolicy:     codePolicy,
		codeRepository: codeRepository,
		roleRepository: roleRepository,
		mailService:    mailService,
	}
}

func (s *codeUsecase) CreateUserRegisterInvitation(ctx context.Context, _dto *dto.CreateUserRegisterInvitation) (*entity.Code, error) {
	var err error

	role, err := s.roleRepository.Find(_dto.RoleId)
	if err != nil {
		// TODO: return custom error on repository instead of gorm's error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorutil.NewHttpError(http.StatusNotFound, "Role not found", nil)
		}
		return nil, err
	}

	code := &entity.Code{}
	code.Type = enum.UserRegisterInvitation
	code.SetMeta(map[string]any{
		"role_id": role.Id,
	})
	code.ExpiredAt = _dto.ExpiredAt

	err = s.codeRepository.Save(code)
	if err != nil {
		return nil, err
	}

	return code, nil
}

func (s *codeUsecase) CreateUserEmailVerification(ctx context.Context, _dto *dto.CreateUserEmailVerification) (*entity.Code, error) {
	var err error

	code := &entity.Code{}
	code.Type = enum.UserEmailVerification
	code.CodeableId = null.StringFrom(_dto.Email)
	code.CodeableType = null.StringFrom("email")
	code.SetMeta(nil)
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

	return code, nil
}

func (s *codeUsecase) CreateUserResetPassword(ctx context.Context, _dto *dto.CreateUserResetPassword) (*entity.Code, error) {
	var err error

	code := &entity.Code{}
	code.Type = enum.UserResetPassword
	code.CodeableId = null.StringFrom(_dto.Email)
	code.CodeableType = null.StringFrom("email")
	code.SetMeta(nil)
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

	return code, nil
}
