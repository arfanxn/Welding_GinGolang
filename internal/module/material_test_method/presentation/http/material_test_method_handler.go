package http

import (
	"errors"
	"net/http"

	"github.com/arfanxn/welding/internal/infrastructure/http/helper"
	"github.com/arfanxn/welding/internal/infrastructure/http/response"
	mtmRequest "github.com/arfanxn/welding/internal/module/material_test_method/presentation/http/request"
	"github.com/arfanxn/welding/internal/module/material_test_method/usecase"
	"github.com/arfanxn/welding/internal/module/material_test_method/usecase/dto"
	"github.com/arfanxn/welding/internal/module/shared/domain/errorx"
	"github.com/arfanxn/welding/pkg/httperror"
	"github.com/arfanxn/welding/pkg/pagination"
	"github.com/arfanxn/welding/pkg/query"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type MaterialTestMethodHandler interface {
	Paginate(c *gin.Context)
	Show(c *gin.Context)
	Store(c *gin.Context)
	Update(c *gin.Context)
	Destroy(c *gin.Context)
}

type mtmHandler struct {
	mtmUsecase usecase.MaterialTestMethodUsecase
}

type NewMaterialTestMethodHandlerParams struct {
	fx.In

	MtmUsecase usecase.MaterialTestMethodUsecase
}

func NewMaterialTestMethodHandler(params NewMaterialTestMethodHandlerParams) MaterialTestMethodHandler {
	return &mtmHandler{
		mtmUsecase: params.MtmUsecase,
	}
}

func (h *mtmHandler) Paginate(c *gin.Context) {
	q := query.NewQuery()
	c.ShouldBind(q)

	op, err := h.mtmUsecase.Paginate(c.Request.Context(), q)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, response.NewBodyWithData(
		http.StatusOK,
		"Material test methods berhasil diambil",
		pagination.FromOPToPP(op, helper.URLFromC(c)),
	))
}

func (h *mtmHandler) Show(c *gin.Context) {
	q := query.NewQuery()
	q.FilterById(c.Param("id"))
	c.ShouldBind(q)

	mtm, err := h.mtmUsecase.Show(c.Request.Context(), q)
	if err != nil {
		if errors.Is(err, errorx.ErrMaterialTestMethodNotFound) {
			httperror.Panic(http.StatusNotFound, "Material test method tidak ditemukan", nil)
		}
		panic(err)
	}

	c.JSON(http.StatusOK, response.NewBodyWithData(
		http.StatusOK,
		"Material test method berhasil diambil",
		gin.H{"material_test_method": mtm},
	))
}

func (h *mtmHandler) Store(c *gin.Context) {
	req := mtmRequest.NewStoreMaterialTestMethod()
	helper.MustBindValidate(c, req)

	mtm, err := h.mtmUsecase.Store(c.Request.Context(), &dto.SaveMaterialTestMethod{
		Name:        &req.Name,
		Description: req.Description,
	})
	if err != nil {
		if errors.Is(err, errorx.ErrMaterialTestMethodAlreadyExists) {
			httperror.Panic(http.StatusConflict, "Material test method sudah ada", nil)
		}
		panic(err)
	}

	c.JSON(http.StatusCreated, response.NewBodyWithData(
		http.StatusCreated,
		"Material test method berhasil disimpan",
		gin.H{"material_test_method": mtm},
	))
}

func (h *mtmHandler) Update(c *gin.Context) {
	req := mtmRequest.NewUpdateMaterialTestMethod()
	req.Id = c.Param("id")
	helper.MustBindValidate(c, req)

	mtm, err := h.mtmUsecase.Update(c.Request.Context(), &dto.SaveMaterialTestMethod{
		Id:          &req.Id,
		Name:        req.Name,
		Description: req.Description,
	})
	if err != nil {
		if errors.Is(err, errorx.ErrMaterialTestMethodNotFound) {
			httperror.Panic(http.StatusNotFound, "Material test method tidak ditemukan", nil)
		}
		if errors.Is(err, errorx.ErrMaterialTestMethodAlreadyExists) {
			httperror.Panic(http.StatusConflict, "Material test method sudah ada", nil)
		}
		panic(err)
	}

	c.JSON(http.StatusOK, response.NewBodyWithData(
		http.StatusOK,
		"Material test method berhasil diperbarui",
		gin.H{"material_test_method": mtm},
	))
}

func (h *mtmHandler) Destroy(c *gin.Context) {
	req := mtmRequest.NewDestroyMaterialTestMethod()
	req.Id = c.Param("id")
	helper.MustBindValidate(c, req)

	err := h.mtmUsecase.Destroy(c.Request.Context(), &dto.DestroyMaterialTestMethod{Id: req.Id})
	if err != nil {
		if errors.Is(err, errorx.ErrMaterialTestMethodNotFound) {
			httperror.Panic(http.StatusNotFound, "Material test method tidak ditemukan", nil)
		}
		panic(err)
	}

	c.JSON(http.StatusOK, response.NewBody(http.StatusOK, "Material test method berhasil dihapus"))
}
