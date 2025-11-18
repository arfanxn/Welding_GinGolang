package http

import (
	"errors"
	"net/http"
	"time"

	"github.com/arfanxn/welding/internal/infrastructure/http/helper"
	"github.com/arfanxn/welding/internal/infrastructure/http/response"
	roleEnum "github.com/arfanxn/welding/internal/module/role/domain/enum"
	"github.com/arfanxn/welding/internal/module/shared/contextkey"
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	"github.com/arfanxn/welding/internal/module/shared/domain/errorx"
	"github.com/arfanxn/welding/internal/module/user/presentation/http/request"
	"github.com/arfanxn/welding/internal/module/user/usecase"
	"github.com/arfanxn/welding/internal/module/user/usecase/dto"
	"github.com/arfanxn/welding/pkg/boolutil"
	"github.com/arfanxn/welding/pkg/httperror"
	"github.com/arfanxn/welding/pkg/pagination"
	"github.com/arfanxn/welding/pkg/query"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserHandler interface {
	Register(c *gin.Context)
	VerifyEmail(c *gin.Context)   // Verify email
	ResetPassword(c *gin.Context) // Reset password
	Login(c *gin.Context)
	Logout(c *gin.Context)
	Me(c *gin.Context)
	Show(c *gin.Context)
	Paginate(c *gin.Context)
	Store(c *gin.Context)
	Update(c *gin.Context)
	UpdateMeProfile(c *gin.Context)
	// ! Deprecated
	// UpdatePassword(c *gin.Context)
	UpdateMePassword(c *gin.Context)
	ToggleActivation(c *gin.Context)
	Destroy(c *gin.Context)
}

type userHandler struct {
	db          *gorm.DB
	userUsecase usecase.UserUsecase
}

func NewUserHandler(db *gorm.DB, userUsecase usecase.UserUsecase) UserHandler {
	return &userHandler{
		db:          db,
		userUsecase: userUsecase,
	}
}

func (h *userHandler) Register(c *gin.Context) {
	var req request.RegisterUser
	helper.MustBindValidate(c, &req)

	user, err := h.userUsecase.Register(c.Request.Context(), &dto.Register{
		Name:                     req.Name,
		PhoneNumber:              req.PhoneNumber,
		Email:                    req.Email,
		Password:                 req.Password,
		InvitationCode:           req.InvitationCode,
		EmploymentIdentityNumber: req.EmploymentIdentityNumber,
	})
	if err != nil {
		if errors.Is(err, errorx.ErrCodeNotFound) {
			httperror.Panic(http.StatusBadRequest, "Kode undangan tidak ditemukan", nil)
		}
		if errors.Is(err, errorx.ErrCodeAlreadyUsed) {
			httperror.Panic(http.StatusBadRequest, "Kode undangan sudah digunakan", nil)
		}
		if errors.Is(err, errorx.ErrCodeExpired) {
			httperror.Panic(http.StatusBadRequest, "Kode undangan sudah expired", nil)
		}
		if errors.Is(err, errorx.ErrRoleDefaultNotConfigured) {
			httperror.Panic(http.StatusBadRequest, "Role default belum dikonfigurasi", nil)
		}
		if errors.Is(err, errorx.ErrUserAlreadyExists) {
			httperror.Panic(http.StatusConflict, "User sudah ada", nil)
		}
		panic(err)
	}

	c.JSON(http.StatusCreated, response.NewBodyWithData(
		http.StatusCreated,
		"Registrasi berhasil",
		gin.H{"user": user},
	))
}

func (h *userHandler) VerifyEmail(c *gin.Context) {
	var req request.VerifyEmail
	helper.MustBindValidate(c, &req)

	user, err := h.userUsecase.VerifyEmail(c.Request.Context(), &dto.VerifyEmail{
		Email: req.Email,
		Code:  req.Code,
	})
	if err != nil {
		if errors.Is(err, errorx.ErrCodeNotFound) {
			httperror.Panic(http.StatusBadRequest, "Kode verifikasi salah", nil)
		}
		if errors.Is(err, errorx.ErrCodeAlreadyUsed) {
			httperror.Panic(http.StatusBadRequest, "Kode verifikasi sudah digunakan", nil)
		}
		if errors.Is(err, errorx.ErrCodeExpired) {
			httperror.Panic(http.StatusBadRequest, "Kode verifikasi sudah kadaluarsa", nil)
		}
		if errors.Is(err, errorx.ErrUserNotFound) {
			httperror.Panic(http.StatusNotFound, "User tidak ditemukan", nil)
		}
		if errors.Is(err, errorx.ErrUserEmailAlreadyVerified) {
			httperror.Panic(http.StatusBadRequest, "Email sudah diverifikasi", nil)
		}
		panic(err)
	}

	c.JSON(http.StatusOK, response.NewBodyWithData(
		http.StatusOK,
		"Email berhasil diverifikasi",
		gin.H{"user": user}),
	)
}

func (h *userHandler) ResetPassword(c *gin.Context) {
	var req request.ResetPassword
	helper.MustBindValidate(c, &req)

	user, err := h.userUsecase.ResetPassword(c.Request.Context(), &dto.ResetPassword{
		Email:    req.Email,
		Code:     req.Code,
		Password: req.Password,
	})
	if err != nil {
		if errors.Is(err, errorx.ErrCodeNotFound) {
			httperror.Panic(http.StatusBadRequest, "Kode reset password salah", nil)
		}
		if errors.Is(err, errorx.ErrCodeAlreadyUsed) {
			httperror.Panic(http.StatusBadRequest, "Kode reset password sudah digunakan", nil)
		}
		if errors.Is(err, errorx.ErrCodeExpired) {
			httperror.Panic(http.StatusBadRequest, "Kode reset password sudah kadaluarsa", nil)
		}
		if errors.Is(err, errorx.ErrUserNotFound) {
			httperror.Panic(http.StatusNotFound, "User tidak ditemukan", nil)
		}
		panic(err)
	}

	c.JSON(http.StatusOK, response.NewBodyWithData(
		http.StatusOK,
		"Password berhasil direset",
		gin.H{"user": user}),
	)
}

func (h *userHandler) Login(c *gin.Context) {
	var req request.LoginUser
	helper.MustBindValidate(c, &req)

	loginResult, err := h.userUsecase.Login(c.Request.Context(), &dto.Login{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		if errors.Is(err, errorx.ErrUserPasswordIncorrect) {
			httperror.Panic(http.StatusUnauthorized, "Email atau password salah", nil)
		}
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
	userId := c.MustGet(contextkey.UserIdKey).(string)

	q := query.NewQuery()
	q.FilterById(userId)
	c.ShouldBind(q)

	user, err := h.userUsecase.Show(c.Request.Context(), q)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, response.NewBodyWithData(http.StatusOK, "User berhasil diambil", gin.H{"user": user}))
}

func (h *userHandler) Show(c *gin.Context) {
	q := query.NewQuery()
	q.FilterById(c.Param("id"))
	c.ShouldBind(q)

	user, err := h.userUsecase.Show(c.Request.Context(), q)
	if err != nil {
		if errors.Is(err, errorx.ErrUserNotFound) {
			httperror.Panic(http.StatusNotFound, "User tidak ditemukan", nil)
		}
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

	op, err := h.userUsecase.Paginate(c.Request.Context(), q)
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

	activatedAt := time.Now()

	user, err := h.userUsecase.Store(c.Request.Context(), &dto.SaveUser{
		Name:                     &req.Name,
		PhoneNumber:              &req.PhoneNumber,
		Email:                    &req.Email,
		Password:                 &req.Password,
		ActivatedAt:              &activatedAt,
		RoleIds:                  req.RoleIds,
		EmploymentIdentityNumber: req.EmploymentIdentityNumber,
	})
	if err != nil {
		if errors.Is(err, errorx.ErrRolesNotFound) {
			httperror.Panic(http.StatusNotFound, "Satu atau lebih role tidak ditemukan", nil)
		}
		if errors.Is(err, errorx.ErrUserSuperAdminAssignmentForbidden) {
			httperror.Panic(http.StatusForbidden, "User tidak dapat diberi role "+string(roleEnum.SuperAdmin), nil)
		}
		if errors.Is(err, errorx.ErrUserAlreadyExists) {
			httperror.Panic(http.StatusConflict, "User sudah ada", nil)
		}
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

	user, err := h.userUsecase.Update(c.Request.Context(), &dto.SaveUser{
		Id:                       &req.Id,
		Name:                     req.Name,
		PhoneNumber:              req.PhoneNumber,
		Email:                    req.Email,
		Password:                 req.Password,
		RoleIds:                  req.RoleIds,
		EmploymentIdentityNumber: req.EmploymentIdentityNumber,
	})
	if err != nil {
		if errors.Is(err, errorx.ErrUserNotFound) {
			httperror.Panic(http.StatusNotFound, "User tidak ditemukan", nil)
		}
		if errors.Is(err, errorx.ErrRolesNotFound) {
			httperror.Panic(http.StatusNotFound, "Satu atau lebih role tidak ditemukan", nil)
		}
		if errors.Is(err, errorx.ErrUserSuperAdminUpdateForbidden) {
			httperror.Panic(http.StatusForbidden, "User super admin tidak dapat diubah", nil)
		}
		if errors.Is(err, errorx.ErrUserSuperAdminRoleChangeForbidden) {
			httperror.Panic(http.StatusForbidden, "User super admin tidak dapat diubah role", nil)
		}
		if errors.Is(err, errorx.ErrUserSuperAdminAssignmentForbidden) {
			httperror.Panic(http.StatusForbidden, "User tidak dapat diberi role "+string(roleEnum.SuperAdmin), nil)
		}
		if errors.Is(err, errorx.ErrUserAlreadyExists) {
			httperror.Panic(http.StatusConflict, "User sudah ada", nil)
		}
		panic(err)
	}

	c.JSON(http.StatusOK, response.NewBodyWithData(
		http.StatusOK,
		"User berhasil diperbarui",
		gin.H{"user": user},
	))
}

func (h *userHandler) UpdateMeProfile(c *gin.Context) {
	req := &request.UpdateUserMeProfile{}
	helper.MustBindValidate(c, req)

	user := c.MustGet(contextkey.UserKey).(*entity.User)

	user, err := h.userUsecase.Update(c.Request.Context(), &dto.SaveUser{
		Id:                       &user.Id,
		Name:                     req.Name,
		PhoneNumber:              req.PhoneNumber,
		Email:                    req.Email,
		EmploymentIdentityNumber: req.EmploymentIdentityNumber,
	})
	if err != nil {
		if errors.Is(err, errorx.ErrUserAlreadyExists) {
			httperror.Panic(http.StatusConflict, "User sudah ada", nil)
		}
		panic(err)
	}

	c.JSON(http.StatusOK, response.NewBodyWithData(
		http.StatusOK,
		"User berhasil diperbarui",
		gin.H{"user": user},
	))
}

/*
! Deprecated
func (h *userHandler) UpdatePassword(c *gin.Context) {
	req := &request.UpdateUserPassword{}
	req.Id = c.Param("id")
	helper.MustBindValidate(c, req)

	user, err := h.userUsecase.Save(c.Request.Context(), &dto.SaveUser{
		Id:       &req.Id,
		Password: &req.Password,
	})
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, response.NewBodyWithData(
		http.StatusOK,
		"Password berhasil diperbarui",
		gin.H{"user": user},
	))
}
*/

func (h *userHandler) UpdateMePassword(c *gin.Context) {
	// Bind and validate the request payload
	req := &request.UpdateUserMePassword{}
	helper.MustBindValidate(c, req)

	// Update user's password
	user, err := h.userUsecase.UpdateMePassword(c.Request.Context(), &dto.UpdateUserMePassword{
		CurrentPassword: req.CurrentPassword,
		Password:        req.Password,
	})
	if err != nil {
		if errors.Is(err, errorx.ErrUserPasswordIncorrect) {
			httperror.Panic(http.StatusBadRequest, "Password saat ini tidak sesuai", nil)
		}
		panic(err)
	}

	// Return success response with updated user data
	c.JSON(http.StatusOK, response.NewBodyWithData(
		http.StatusOK,
		"Password berhasil diperbarui",
		gin.H{"user": user},
	))
}

func (h *userHandler) ToggleActivation(c *gin.Context) {
	req := &request.ToggleActivationUser{}
	req.Id = c.Param("id")
	helper.MustBindValidate(c, req)

	user, err := h.userUsecase.ToggleActivation(c.Request.Context(), &dto.ToggleActivation{Id: req.Id})
	if err != nil {
		if errors.Is(err, errorx.ErrUserNotFound) {
			httperror.Panic(http.StatusNotFound, "User tidak ditemukan", nil)
		}
		if errors.Is(err, errorx.ErrUserSuperAdminUpdateForbidden) {
			httperror.Panic(http.StatusForbidden, "User dengan role "+string(roleEnum.SuperAdmin)+" tidak dapat dinonaktifkan", nil)
		}
		panic(err)
	}

	message := boolutil.Ternary(user.ActivatedAt.Valid, "User berhasil diaktifkan", "User berhasil dinonaktifkan")

	c.JSON(http.StatusOK, response.NewBodyWithData(
		http.StatusOK,
		message,
		gin.H{"user": user},
	))
}

func (h *userHandler) Destroy(c *gin.Context) {
	req := &request.DestroyUser{}
	req.Id = c.Param("id")
	helper.MustBindValidate(c, req)

	err := h.userUsecase.Destroy(c.Request.Context(), &dto.DestroyUser{Id: req.Id})
	if err != nil {
		if errors.Is(err, errorx.ErrUserNotFound) {
			httperror.Panic(http.StatusNotFound, "User tidak ditemukan", nil)
		}
		if errors.Is(err, errorx.ErrUserSuperAdminUpdateForbidden) {
			httperror.Panic(http.StatusForbidden, "User dengan role "+string(roleEnum.SuperAdmin)+" tidak dapat dihapus", nil)
		}
		panic(err)
	}

	c.JSON(http.StatusOK, response.NewBody(http.StatusOK, "User berhasil dihapus"))
}
