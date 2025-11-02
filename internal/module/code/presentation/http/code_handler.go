package http

import (
	"net/http"
	"time"

	"github.com/arfanxn/welding/internal/infrastructure/http/helper"
	"github.com/arfanxn/welding/internal/infrastructure/http/response"
	"github.com/arfanxn/welding/internal/module/code/presentation/http/request"
	"github.com/arfanxn/welding/internal/module/code/usecase"
	"github.com/arfanxn/welding/internal/module/code/usecase/dto"
	"github.com/arfanxn/welding/pkg/errorutil"
	"github.com/gin-gonic/gin"
)

type CodeHandler interface {
	CreateUserRegisterInvitation(c *gin.Context)
	CreateUserEmailVerification(c *gin.Context)
	CreateUserResetPassword(c *gin.Context)
}

type codeHandler struct {
	codeUsecase usecase.CodeUsecase
}

func NewCodeHandler(codeUsecase usecase.CodeUsecase) CodeHandler {
	return &codeHandler{
		codeUsecase: codeUsecase,
	}
}

func (h *codeHandler) CreateUserRegisterInvitation(c *gin.Context) {
	var req request.CreateUserRegisterInvitation
	helper.MustBindValidate(c, &req)

	code, err := h.codeUsecase.CreateUserRegisterInvitation(c.Request.Context(), &dto.CreateUserRegisterInvitation{
		RoleId:    req.RoleId,
		ExpiredAt: errorutil.Must(time.Parse(time.DateTime, req.ExpiredAt)),
	})
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusCreated, response.NewBodyWithData(
		http.StatusCreated,
		"Kode undangan registrasi berhasil dibuat",
		gin.H{"code": code},
	))
}

func (h *codeHandler) CreateUserEmailVerification(c *gin.Context) {
	var req request.CreateUserEmailVerification
	helper.MustBindValidate(c, &req)

	_, err := h.codeUsecase.CreateUserEmailVerification(
		c.Request.Context(),
		&dto.CreateUserEmailVerification{Email: req.Email},
	)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusCreated, response.NewBody(
		http.StatusCreated,
		"Kode verifikasi email berhasil dibuat dan dikirim ke email",
	))
}

func (h *codeHandler) CreateUserResetPassword(c *gin.Context) {
	var req request.CreateUserResetPassword
	helper.MustBindValidate(c, &req)

	_, err := h.codeUsecase.CreateUserResetPassword(
		c.Request.Context(),
		&dto.CreateUserResetPassword{Email: req.Email},
	)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusCreated, response.NewBody(
		http.StatusCreated,
		"Kode reset password berhasil dibuat dan dikirim ke email",
	))
}
