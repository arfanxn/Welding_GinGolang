package http

import (
	"net/http"

	"github.com/arfanxn/welding/internal/infrastructure/http/helper"
	"github.com/arfanxn/welding/internal/infrastructure/http/response"
	"github.com/arfanxn/welding/internal/module/user/presentation/http/request"
	"github.com/arfanxn/welding/internal/module/user/usecase"
	"github.com/arfanxn/welding/internal/module/user/usecase/dto"
	"github.com/arfanxn/welding/pkg/pagination"
	"github.com/arfanxn/welding/pkg/query"
	"github.com/gin-gonic/gin"
	"github.com/guregu/null/v6"
)

type UserHandler interface {
	Login(c *gin.Context)
	Logout(c *gin.Context)
	Me(c *gin.Context)
	Find(c *gin.Context)
	Paginate(c *gin.Context)
	Store(c *gin.Context)
	Update(c *gin.Context)
	Destroy(c *gin.Context)
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
	var req request.LoginUser
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
	userId := c.MustGet("user_id").(string)

	q := query.NewQuery()
	q.FilterById(userId)
	c.ShouldBind(q)

	user, err := h.userUsecase.Find(q)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, response.NewBodyWithData(http.StatusOK, "User berhasil diambil", gin.H{"user": user}))
}

func (h *userHandler) Find(c *gin.Context) {
	q := query.NewQuery()
	q.FilterById(c.Param("id"))
	c.ShouldBind(q)

	user, err := h.userUsecase.Find(q)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, response.NewBodyWithData(
		http.StatusOK,
		"User berhasil diambil",
		gin.H{"user": user},
	))
}

func (h *userHandler) Paginate(c *gin.Context) {
	q := query.NewQuery()
	c.ShouldBind(q)

	op, err := h.userUsecase.Paginate(q)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, response.NewBodyWithData(
		http.StatusOK,
		"Users berhasil diambil",
		pagination.PPFromOP(op, helper.URLFromC(c)),
	))
}

func (h *userHandler) Store(c *gin.Context) {
	req := &request.StoreUser{}
	helper.MustBindValidate(c, req)

	user, err := h.userUsecase.Save(&dto.SaveUser{
		Name:        req.Name,
		PhoneNumber: req.PhoneNumber,
		Email:       req.Email,
		Password:    req.Password,
		RoleIds:     req.RoleIds,
	})
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusCreated, response.NewBodyWithData(
		http.StatusCreated,
		"User berhasil disimpan",
		gin.H{"user": user},
	))
}

func (h *userHandler) Update(c *gin.Context) {
	req := &request.UpdateUser{}
	req.Id = c.Param("id")
	helper.MustBindValidate(c, req)

	user, err := h.userUsecase.Save(&dto.SaveUser{
		Id:          null.StringFrom(req.Id),
		Name:        req.Name,
		PhoneNumber: req.PhoneNumber,
		Email:       req.Email,
		Password:    req.Password,
		RoleIds:     req.RoleIds,
	})
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, response.NewBodyWithData(
		http.StatusOK,
		"User berhasil diperbarui",
		gin.H{"user": user},
	))
}

func (h *userHandler) Destroy(c *gin.Context) {
	req := &request.DestroyUser{}
	req.Id = c.Param("id")
	helper.MustBindValidate(c, req)

	err := h.userUsecase.Destroy(&dto.DestroyUser{Id: req.Id})
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, response.NewBody(http.StatusOK, "User berhasil dihapus"))
}
