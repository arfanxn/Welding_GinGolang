package http

import (
	"errors"
	"net/http"

	"github.com/arfanxn/welding/internal/infrastructure/http/helper"
	"github.com/arfanxn/welding/internal/infrastructure/http/response"
	materialTestWorkCategoryRequest "github.com/arfanxn/welding/internal/module/material_test_work_category/presentation/http/request"
	"github.com/arfanxn/welding/internal/module/material_test_work_category/usecase"
	"github.com/arfanxn/welding/internal/module/material_test_work_category/usecase/dto"
	"github.com/arfanxn/welding/internal/module/shared/domain/errorx"
	"github.com/arfanxn/welding/pkg/httperror"
	"github.com/arfanxn/welding/pkg/pagination"
	"github.com/arfanxn/welding/pkg/query"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type MaterialTestWorkCategoryHandler interface {
	Paginate(c *gin.Context)
	Show(c *gin.Context)
	Store(c *gin.Context)
	Update(c *gin.Context)
	Destroy(c *gin.Context)
}

type materialTestWorkCategoryHandler struct {
	materialTestWorkCategoryUsecase usecase.MaterialTestWorkCategoryUsecase
}

type NewMaterialTestWorkCategoryHandlerParams struct {
	fx.In

	MaterialTestWorkCategoryUsecase usecase.MaterialTestWorkCategoryUsecase
}

func NewMaterialTestWorkCategoryHandler(params NewMaterialTestWorkCategoryHandlerParams) MaterialTestWorkCategoryHandler {
	return &materialTestWorkCategoryHandler{
		materialTestWorkCategoryUsecase: params.MaterialTestWorkCategoryUsecase,
	}
}

func (h *materialTestWorkCategoryHandler) Paginate(c *gin.Context) {
	q := query.NewQuery()
	c.ShouldBind(q)

	op, err := h.materialTestWorkCategoryUsecase.Paginate(c.Request.Context(), q)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, response.NewBodyWithData(
		http.StatusOK,
		"Material test work categories berhasil diambil",
		pagination.FromOPToPP(op, helper.URLFromC(c)),
	))
}

func (h *materialTestWorkCategoryHandler) Show(c *gin.Context) {
	q := query.NewQuery()
	q.FilterById(c.Param("id"))
	c.ShouldBind(q)

	mtm, err := h.materialTestWorkCategoryUsecase.Show(c.Request.Context(), q)
	if err != nil {
		if errors.Is(err, errorx.ErrMaterialTestWorkCategoryNotFound) {
			httperror.Panic(http.StatusNotFound, "Material test work category tidak ditemukan", nil)
		}
		panic(err)
	}

	c.JSON(http.StatusOK, response.NewBodyWithData(
		http.StatusOK,
		"Material test work category berhasil diambil",
		gin.H{"material_test_work_category": mtm},
	))
}

func (h *materialTestWorkCategoryHandler) Store(c *gin.Context) {
	req := materialTestWorkCategoryRequest.NewStoreMaterialTestWorkCategory()
	helper.MustBindValidate(c, req)

	mtm, err := h.materialTestWorkCategoryUsecase.Store(c.Request.Context(), &dto.SaveMaterialTestWorkCategory{
		Name:        &req.Name,
		Description: req.Description,
	})
	if err != nil {
		if errors.Is(err, errorx.ErrMaterialTestWorkCategoryAlreadyExists) {
			httperror.Panic(http.StatusConflict, "Material test work category sudah ada", nil)
		}
		panic(err)
	}

	c.JSON(http.StatusCreated, response.NewBodyWithData(
		http.StatusCreated,
		"Material test work category berhasil disimpan",
		gin.H{"material_test_work_category": mtm},
	))
}

func (h *materialTestWorkCategoryHandler) Update(c *gin.Context) {
	req := materialTestWorkCategoryRequest.NewUpdateMaterialTestWorkCategory()
	req.Id = c.Param("id")
	helper.MustBindValidate(c, req)

	mtm, err := h.materialTestWorkCategoryUsecase.Update(c.Request.Context(), &dto.SaveMaterialTestWorkCategory{
		Id:          &req.Id,
		Name:        req.Name,
		Description: req.Description,
	})
	if err != nil {
		if errors.Is(err, errorx.ErrMaterialTestWorkCategoryNotFound) {
			httperror.Panic(http.StatusNotFound, "Material test work category tidak ditemukan", nil)
		}
		if errors.Is(err, errorx.ErrMaterialTestWorkCategoryAlreadyExists) {
			httperror.Panic(http.StatusConflict, "Material test work category sudah ada", nil)
		}
		panic(err)
	}

	c.JSON(http.StatusOK, response.NewBodyWithData(
		http.StatusOK,
		"Material test work category berhasil diperbarui",
		gin.H{"material_test_work_category": mtm},
	))
}

func (h *materialTestWorkCategoryHandler) Destroy(c *gin.Context) {
	req := materialTestWorkCategoryRequest.NewDestroyMaterialTestWorkCategory()
	req.Id = c.Param("id")
	helper.MustBindValidate(c, req)

	err := h.materialTestWorkCategoryUsecase.Destroy(c.Request.Context(), &dto.DestroyMaterialTestWorkCategory{Id: req.Id})
	if err != nil {
		if errors.Is(err, errorx.ErrMaterialTestWorkCategoryNotFound) {
			httperror.Panic(http.StatusNotFound, "Material test work category tidak ditemukan", nil)
		}
		panic(err)
	}

	c.JSON(http.StatusOK, response.NewBody(http.StatusOK, "Material test work category berhasil dihapus"))
}
