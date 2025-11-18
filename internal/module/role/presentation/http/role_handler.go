package http

import (
	"errors"
	"net/http"

	"github.com/arfanxn/welding/internal/infrastructure/http/helper"
	"github.com/arfanxn/welding/internal/infrastructure/http/response"
	"github.com/arfanxn/welding/internal/module/role/domain/enum"
	roleRequest "github.com/arfanxn/welding/internal/module/role/presentation/http/request"
	"github.com/arfanxn/welding/internal/module/role/usecase"
	roleDto "github.com/arfanxn/welding/internal/module/role/usecase/dto"
	"github.com/arfanxn/welding/internal/module/shared/domain/errorx"
	"github.com/arfanxn/welding/pkg/httperror"
	"github.com/arfanxn/welding/pkg/pagination"
	"github.com/arfanxn/welding/pkg/query"
	"github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
)

type RoleHandler interface {
	Paginate(c *gin.Context)
	Show(c *gin.Context)
	Store(c *gin.Context)
	Update(c *gin.Context)
	SetDefault(c *gin.Context)
	Destroy(c *gin.Context)
}

type roleHandler struct {
	roleUsecase usecase.RoleUsecase
}

func NewRoleHandler(roleUsecase usecase.RoleUsecase) RoleHandler {
	return &roleHandler{
		roleUsecase: roleUsecase,
	}
}

func (h *roleHandler) Paginate(c *gin.Context) {
	q := query.NewQuery()
	c.ShouldBind(q)

	spew.Dump(q)

	paginationDto, err := h.roleUsecase.Paginate(c.Request.Context(), q)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, response.NewBodyWithData(
		http.StatusOK,
		"Roles berhasil diambil",
		pagination.PPFromOP(paginationDto, helper.URLFromC(c)),
	))
}

func (h *roleHandler) Show(c *gin.Context) {
	q := query.NewQuery()
	q.FilterById(c.Param("id"))
	c.ShouldBind(q)

	role, err := h.roleUsecase.Show(c.Request.Context(), q)
	if err != nil {
		if errors.Is(err, errorx.ErrRoleNotFound) {
			c.JSON(http.StatusNotFound, response.NewBody(http.StatusNotFound, "Role tidak ditemukan"))
			return
		}
		panic(err)
	}

	c.JSON(http.StatusOK, response.NewBodyWithData(
		http.StatusOK,
		"Role berhasil diambil",
		gin.H{"role": role},
	))
}

func (h *roleHandler) Store(c *gin.Context) {
	req := roleRequest.NewStoreRole()
	helper.MustBindValidate(c, req)

	roleName := enum.RoleName(req.Name)

	role, err := h.roleUsecase.Store(c.Request.Context(), &roleDto.SaveRole{
		Name:          &roleName,
		PermissionIds: req.PermissionIds,
	})
	if err != nil {
		if errors.Is(err, errorx.ErrRoleAlreadyExists) {
			httperror.Panic(http.StatusConflict, "Role sudah ada", nil)
		}
		if errors.Is(err, errorx.ErrRoleSuperAdminStoreForbidden) {
			httperror.Panic(http.StatusForbidden, "Role super admin tidak dapat dibuat", nil)
		}
		if errors.Is(err, errorx.ErrPermissionsNotFound) {
			httperror.Panic(http.StatusNotFound, "Satu atau lebih permission tidak ditemukan", nil)
		}
		panic(err)
	}

	c.JSON(http.StatusCreated, response.NewBodyWithData(
		http.StatusCreated,
		"Role berhasil disimpan",
		gin.H{"role": role},
	))
}

func (h *roleHandler) Update(c *gin.Context) {
	req := roleRequest.NewUpdateRole()
	req.Id = c.Param("id")
	helper.MustBindValidate(c, req)

	roleName := enum.RoleName(req.Name)

	role, err := h.roleUsecase.Update(c.Request.Context(), &roleDto.SaveRole{
		Id:            &req.Id,
		Name:          &roleName,
		PermissionIds: req.PermissionIds,
	})
	if err != nil {
		if errors.Is(err, errorx.ErrRoleNotFound) {
			httperror.Panic(http.StatusNotFound, "Role tidak ditemukan", nil)
		}
		if errors.Is(err, errorx.ErrRoleAlreadyExists) {
			httperror.Panic(http.StatusConflict, "Role sudah ada", nil)
		}
		if errors.Is(err, errorx.ErrRoleSuperAdminUpdateForbidden) {
			httperror.Panic(http.StatusForbidden, "Role super admin tidak dapat diubah", nil)
		}
		if errors.Is(err, errorx.ErrPermissionsNotFound) {
			httperror.Panic(http.StatusNotFound, "Satu atau lebih permission tidak ditemukan", nil)
		}
		panic(err)
	}

	c.JSON(http.StatusOK, response.NewBodyWithData(
		http.StatusOK,
		"Role berhasil diperbarui",
		gin.H{"role": role},
	))
}

func (h *roleHandler) SetDefault(c *gin.Context) {
	req := roleRequest.NewSetDefaultRole()
	req.Id = c.Param("id")
	helper.MustBindValidate(c, req)

	role, err := h.roleUsecase.SetDefault(c.Request.Context(), &roleDto.SetDefaultRole{Id: req.Id})
	if err != nil {
		if errors.Is(err, errorx.ErrRoleNotFound) {
			httperror.Panic(http.StatusNotFound, "Role tidak ditemukan", nil)
		}
		if errors.Is(err, errorx.ErrRoleAlreadyDefault) {
			httperror.Panic(http.StatusConflict, "Role sudah default", nil)
		}
		if errors.Is(err, errorx.ErrRoleSuperAdminSetDefaultForbidden) {
			httperror.Panic(http.StatusForbidden, "Role super admin tidak dapat diset default", nil)
		}
		panic(err)
	}

	c.JSON(http.StatusOK, response.NewBodyWithData(
		http.StatusOK,
		"Role berhasil diset default",
		gin.H{"role": role},
	))
}

func (h *roleHandler) Destroy(c *gin.Context) {
	req := roleRequest.NewDestroyRole()
	req.Id = c.Param("id")
	helper.MustBindValidate(c, req)

	err := h.roleUsecase.Destroy(c.Request.Context(), &roleDto.DestroyRole{Id: req.Id})
	if err != nil {
		if errors.Is(err, errorx.ErrRoleNotFound) {
			httperror.Panic(http.StatusNotFound, "Role tidak ditemukan", nil)
		}
		if errors.Is(err, errorx.ErrRoleDefaultDestroyForbidden) {
			httperror.Panic(http.StatusForbidden, "Role default tidak dapat dihapus", nil)
		}
		if errors.Is(err, errorx.ErrRoleSuperAdminDestroyForbidden) {
			httperror.Panic(http.StatusForbidden, "Role super admin tidak dapat dihapus", nil)
		}
		panic(err)
	}

	c.JSON(http.StatusOK, response.NewBody(http.StatusOK, "Role berhasil dihapus"))
}
