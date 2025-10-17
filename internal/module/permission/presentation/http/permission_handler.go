package http

import (
	"net/http"

	"github.com/arfanxn/welding/internal/infrastructure/http/helper"
	"github.com/arfanxn/welding/internal/infrastructure/http/request"
	"github.com/arfanxn/welding/internal/infrastructure/http/response"
	"github.com/arfanxn/welding/internal/module/permission/usecase"
	"github.com/gin-gonic/gin"
)

type PermissionHandler interface {
	Paginate(c *gin.Context)
}

type permissionHandler struct {
	permissionUsecase usecase.PermissionUsecase
}

func NewPermissionHandler(permissionUsecase usecase.PermissionUsecase) PermissionHandler {
	return &permissionHandler{
		permissionUsecase: permissionUsecase,
	}
}

func (h *permissionHandler) Paginate(c *gin.Context) {
	req := request.NewQuery()
	helper.MustBindValidate(c, req)

	paginationDto, err := h.permissionUsecase.Paginate(req.MustToQueryDTO())
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, response.NewBodyWithData(
		http.StatusOK,
		"Permissions berhasil diambil",
		response.NewPaginationFromContextPaginationDTO(c, paginationDto),
	))
}
