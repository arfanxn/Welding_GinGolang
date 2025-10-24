package http

import (
	"net/http"

	"github.com/arfanxn/welding/internal/infrastructure/http/helper"
	"github.com/arfanxn/welding/internal/infrastructure/http/response"
	"github.com/arfanxn/welding/internal/module/permission/usecase"
	"github.com/arfanxn/welding/pkg/pagination"
	"github.com/arfanxn/welding/pkg/query"
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
	q := query.NewQuery()
	c.ShouldBind(q)

	op, err := h.permissionUsecase.Paginate(q)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, response.NewBodyWithData(
		http.StatusOK,
		"Permissions berhasil diambil",
		pagination.PPFromOP(op, helper.URLFromC(c)),
	))
}
