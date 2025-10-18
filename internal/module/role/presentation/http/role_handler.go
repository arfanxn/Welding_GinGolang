package http

import (
	"net/http"

	"github.com/arfanxn/welding/internal/infrastructure/http/helper"
	"github.com/arfanxn/welding/internal/infrastructure/http/request"
	"github.com/arfanxn/welding/internal/infrastructure/http/response"
	roleRequest "github.com/arfanxn/welding/internal/module/role/presentation/http/request"
	"github.com/arfanxn/welding/internal/module/role/usecase"
	roleDto "github.com/arfanxn/welding/internal/module/role/usecase/dto"
	"github.com/gin-gonic/gin"
)

type RoleHandler interface {
	Paginate(c *gin.Context)
	Find(c *gin.Context)
	Store(c *gin.Context)
	Update(c *gin.Context)
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
	req := request.NewQuery()
	helper.MustBindValidate(c, req)

	paginationDto, err := h.roleUsecase.Paginate(req.MustToQueryDTO())
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, response.NewBodyWithData(
		http.StatusOK,
		"Roles berhasil diambil",
		response.NewPaginationFromContextPaginationDTO(c, paginationDto),
	))
}

func (h *roleHandler) Store(c *gin.Context) {
	req := roleRequest.NewStoreRole()
	helper.MustBindValidate(c, req)

	role, err := h.roleUsecase.Save(&roleDto.SaveRole{
		Name:          req.Name,
		PermissionIds: req.PermissionIds,
	})
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, response.NewBodyWithData(
		http.StatusOK,
		"Role berhasil disimpan",
		gin.H{"role": role},
	))
}

func (h *roleHandler) Find(c *gin.Context) {
	req := request.NewQuery()
	req.AppendFilter("id == " + c.Param("id"))
	helper.MustBindValidate(c, req)

	role, err := h.roleUsecase.Find(req.MustToQueryDTO())
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, response.NewBodyWithData(
		http.StatusOK,
		"Role berhasil diambil",
		gin.H{"role": role},
	))
}

func (h *roleHandler) Update(c *gin.Context) {
	req := roleRequest.NewUpdateRole()
	req.Id = c.Param("id")
	helper.MustBindValidate(c, req)

	role, err := h.roleUsecase.Save(&roleDto.SaveRole{
		Id:            req.Id,
		Name:          req.Name,
		PermissionIds: req.PermissionIds,
	})
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, response.NewBodyWithData(
		http.StatusOK,
		"Role berhasil diperbarui",
		gin.H{"role": role},
	))
}

func (h *roleHandler) Destroy(c *gin.Context) {
	req := roleRequest.NewDestroyRole()
	req.Id = c.Param("id")
	helper.MustBindValidate(c, req)

	err := h.roleUsecase.Destroy(&roleDto.DestroyRole{Id: req.Id})
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, response.NewBody(http.StatusOK, "Role berhasil dihapus"))
}
