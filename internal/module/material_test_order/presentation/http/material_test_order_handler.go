package http

import (
	"errors"
	"net/http"

	"github.com/arfanxn/welding/internal/infrastructure/http/helper"
	"github.com/arfanxn/welding/internal/infrastructure/http/response"
	materialTestOrderRequest "github.com/arfanxn/welding/internal/module/material_test_order/presentation/http/request"
	materialTestOrderUsecase "github.com/arfanxn/welding/internal/module/material_test_order/usecase"
	materialTestOrderDto "github.com/arfanxn/welding/internal/module/material_test_order/usecase/dto"
	"github.com/arfanxn/welding/internal/module/shared/domain/errorx"
	"github.com/arfanxn/welding/pkg/httperror"
	"github.com/arfanxn/welding/pkg/pagination"
	"github.com/arfanxn/welding/pkg/query"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type MaterialTestOrderHandler interface {
	Paginate(c *gin.Context)
	Show(c *gin.Context)
	Store(c *gin.Context)
	Update(c *gin.Context)
	Destroy(c *gin.Context)
}

type materialTestOrderHandler struct {
	materialTestOrderUsecase materialTestOrderUsecase.MaterialTestOrderUsecase
}

type NewMaterialTestOrderHandlerParams struct {
	fx.In

	MaterialTestOrderUsecase materialTestOrderUsecase.MaterialTestOrderUsecase
}

func NewMaterialTestOrderHandler(params NewMaterialTestOrderHandlerParams) MaterialTestOrderHandler {
	return &materialTestOrderHandler{
		materialTestOrderUsecase: params.MaterialTestOrderUsecase,
	}
}

func (h *materialTestOrderHandler) Paginate(c *gin.Context) {
	q := query.NewQuery()
	c.ShouldBind(q)

	op, err := h.materialTestOrderUsecase.Paginate(c.Request.Context(), q)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, response.NewBodyWithData(
		http.StatusOK,
		"Material test orders berhasil diambil",
		pagination.FromOPToPP(op, helper.URLFromC(c)),
	))
}

func (h *materialTestOrderHandler) Show(c *gin.Context) {
	q := query.NewQuery()
	q.FilterById(c.Param("id"))
	c.ShouldBind(q)

	materialTestOrder, err := h.materialTestOrderUsecase.Show(c.Request.Context(), q)
	if err != nil {
		if errors.Is(err, errorx.ErrMaterialTestOrderNotFound) {
			c.JSON(http.StatusNotFound, response.NewBody(http.StatusNotFound, "Material test order tidak ditemukan"))
			return
		}
		panic(err)
	}

	c.JSON(http.StatusOK, response.NewBodyWithData(
		http.StatusOK,
		"Material test order berhasil diambil",
		gin.H{"material_test_order": materialTestOrder},
	))
}

func (h *materialTestOrderHandler) Store(c *gin.Context) {
	req := materialTestOrderRequest.NewStoreMaterialTestOrder()
	helper.MustBindValidate(c, req)

	materialTestOrder, err := h.materialTestOrderUsecase.Store(c.Request.Context(), &materialTestOrderDto.SaveMaterialTestOrder{
		// TODO
	})
	if err != nil {
		if errors.Is(err, errorx.ErrMaterialTestOrderAlreadyExists) {
			httperror.Panic(http.StatusConflict, "Material test order sudah ada", nil)
		}
		// TODO
		panic(err)
	}

	c.JSON(http.StatusCreated, response.NewBodyWithData(
		http.StatusCreated,
		"Material test order berhasil disimpan",
		gin.H{"material_test_order": materialTestOrder},
	))
}

func (h *materialTestOrderHandler) Update(c *gin.Context) {
	req := materialTestOrderRequest.NewUpdateMaterialTestOrder()
	req.Id = c.Param("id")
	helper.MustBindValidate(c, req)

	materialTestOrder, err := h.materialTestOrderUsecase.Update(c.Request.Context(), &materialTestOrderDto.SaveMaterialTestOrder{
		Id: &req.Id,
		// TODO
	})
	if err != nil {
		if errors.Is(err, errorx.ErrMaterialTestOrderNotFound) {
			httperror.Panic(http.StatusNotFound, "Material test order tidak ditemukan", nil)
		}
		if errors.Is(err, errorx.ErrMaterialTestOrderAlreadyExists) {
			httperror.Panic(http.StatusConflict, "Material test order sudah ada", nil)
		}
		// TODO
		panic(err)
	}

	c.JSON(http.StatusOK, response.NewBodyWithData(
		http.StatusOK,
		"Material test order berhasil diperbarui",
		gin.H{"material_test_order": materialTestOrder},
	))
}

func (h *materialTestOrderHandler) Destroy(c *gin.Context) {
	req := materialTestOrderRequest.NewDestroyMaterialTestOrder()
	req.Id = c.Param("id")
	helper.MustBindValidate(c, req)

	err := h.materialTestOrderUsecase.Destroy(c.Request.Context(), &materialTestOrderDto.DestroyMaterialTestOrder{Id: req.Id})
	if err != nil {
		if errors.Is(err, errorx.ErrMaterialTestOrderNotFound) {
			httperror.Panic(http.StatusNotFound, "Material test order tidak ditemukan", nil)
		}
		// TODO
		panic(err)
	}

	c.JSON(http.StatusOK, response.NewBody(http.StatusOK, "Material test order berhasil dihapus"))
}
