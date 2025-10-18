package http

import (
	"net/http"

	"github.com/arfanxn/welding/internal/infrastructure/http/helper"
	"github.com/arfanxn/welding/internal/infrastructure/http/response"
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	"github.com/arfanxn/welding/internal/module/user/presentation/http/request"
	"github.com/arfanxn/welding/internal/module/user/usecase"
	"github.com/arfanxn/welding/internal/module/user/usecase/dto"
	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	Login(c *gin.Context)
	Logout(c *gin.Context)
	Me(c *gin.Context)
}

type userHandler struct {
	userUsecase usecase.UserUsecase
}

func NewUserHandler(userUsecase usecase.UserUsecase) UserHandler {
	return &userHandler{
		userUsecase: userUsecase,
	}
}

func (h *userHandler) Login(c *gin.Context) {
	var req request.LoginUserRequest
	helper.MustBindValidate(c, &req)

	loginResult, err := h.userUsecase.Login(&dto.Login{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, response.NewBodyWithData(
		http.StatusOK,
		"Login berhasil",
		gin.H{"user": loginResult.User, "token": loginResult.Token},
	))
}

func (h *userHandler) Logout(c *gin.Context) {
	c.JSON(http.StatusOK, response.NewBody(http.StatusOK, "Logout berhasil"))
}

func (h *userHandler) Me(c *gin.Context) {
	user := c.MustGet("user").(*entity.User)
	c.JSON(http.StatusOK, response.NewBodyWithData(http.StatusOK, "User berhasil diambil", gin.H{"user": user}))
}
