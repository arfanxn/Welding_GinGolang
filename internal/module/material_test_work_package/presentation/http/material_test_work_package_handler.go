package http

import (
	"errors"
	"net/http"

	"github.com/arfanxn/welding/internal/infrastructure/http/helper"
	"github.com/arfanxn/welding/internal/infrastructure/http/response"
	materialTestWorkPackageRequest "github.com/arfanxn/welding/internal/module/material_test_work_package/presentation/http/request"
	"github.com/arfanxn/welding/internal/module/material_test_work_package/usecase"
	"github.com/arfanxn/welding/internal/module/material_test_work_package/usecase/dto"
	"github.com/arfanxn/welding/internal/module/shared/domain/errorx"
	"github.com/arfanxn/welding/pkg/httperror"
	"github.com/arfanxn/welding/pkg/pagination"
	"github.com/arfanxn/welding/pkg/query"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type MaterialTestWorkPackageHandler interface {
	Paginate(c *gin.Context)
	Show(c *gin.Context)
	Store(c *gin.Context)
	Update(c *gin.Context)
	Destroy(c *gin.Context)
}

type materialTestWorkPackageHandler struct {
	materialTestWorkPackageUsecase usecase.MaterialTestWorkPackageUsecase
}

type NewMaterialTestWorkPackageHandlerParams struct {
	fx.In

	MaterialTestWorkPackageUsecase usecase.MaterialTestWorkPackageUsecase
}

func NewMaterialTestWorkPackageHandler(params NewMaterialTestWorkPackageHandlerParams) MaterialTestWorkPackageHandler {
	return &materialTestWorkPackageHandler{
		materialTestWorkPackageUsecase: params.MaterialTestWorkPackageUsecase,
	}
}

func (h *materialTestWorkPackageHandler) Paginate(c *gin.Context) {
	q := query.NewQuery()
	c.ShouldBind(q)

	op, err := h.materialTestWorkPackageUsecase.Paginate(c.Request.Context(), q)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, response.NewBodyWithData(
		http.StatusOK,
		"Material test work packages berhasil diambil",
		pagination.FromOPToPP(op, helper.URLFromC(c)),
	))
}

func (h *materialTestWorkPackageHandler) Show(c *gin.Context) {
	q := query.NewQuery()
	q.FilterById(c.Param("id"))
	c.ShouldBind(q)

	mtm, err := h.materialTestWorkPackageUsecase.Show(c.Request.Context(), q)
	if err != nil {
		if errors.Is(err, errorx.ErrMaterialTestWorkPackageNotFound) {
			httperror.Panic(http.StatusNotFound, "Material test work package tidak ditemukan", nil)
		}
		panic(err)
	}

	c.JSON(http.StatusOK, response.NewBodyWithData(
		http.StatusOK,
		"Material test work package berhasil diambil",
		gin.H{"material_test_work_package": mtm},
	))
}

func (h *materialTestWorkPackageHandler) Store(c *gin.Context) {
	req := materialTestWorkPackageRequest.NewStoreMaterialTestWorkPackage()
	helper.MustBindValidate(c, req)

	mtm, err := h.materialTestWorkPackageUsecase.Store(c.Request.Context(), &dto.SaveMaterialTestWorkPackage{
		Name:        &req.Name,
		Description: req.Description,
	})
	if err != nil {
		if errors.Is(err, errorx.ErrMaterialTestWorkPackageAlreadyExists) {
			httperror.Panic(http.StatusConflict, "Material test work package sudah ada", nil)
		}
		panic(err)
	}

	c.JSON(http.StatusCreated, response.NewBodyWithData(
		http.StatusCreated,
		"Material test work package berhasil disimpan",
		gin.H{"material_test_work_package": mtm},
	))
}

func (h *materialTestWorkPackageHandler) Update(c *gin.Context) {
	req := materialTestWorkPackageRequest.NewUpdateMaterialTestWorkPackage()
	req.Id = c.Param("id")
	helper.MustBindValidate(c, req)

	mtm, err := h.materialTestWorkPackageUsecase.Update(c.Request.Context(), &dto.SaveMaterialTestWorkPackage{
		Id:          &req.Id,
		Name:        req.Name,
		Description: req.Description,
	})
	if err != nil {
		if errors.Is(err, errorx.ErrMaterialTestWorkPackageNotFound) {
			httperror.Panic(http.StatusNotFound, "Material test work package tidak ditemukan", nil)
		}
		if errors.Is(err, errorx.ErrMaterialTestWorkPackageAlreadyExists) {
			httperror.Panic(http.StatusConflict, "Material test work package sudah ada", nil)
		}
		panic(err)
	}

	c.JSON(http.StatusOK, response.NewBodyWithData(
		http.StatusOK,
		"Material test work package berhasil diperbarui",
		gin.H{"material_test_work_package": mtm},
	))
}

func (h *materialTestWorkPackageHandler) Destroy(c *gin.Context) {
	req := materialTestWorkPackageRequest.NewDestroyMaterialTestWorkPackage()
	req.Id = c.Param("id")
	helper.MustBindValidate(c, req)

	err := h.materialTestWorkPackageUsecase.Destroy(c.Request.Context(), &dto.DestroyMaterialTestWorkPackage{Id: req.Id})
	if err != nil {
		if errors.Is(err, errorx.ErrMaterialTestWorkPackageNotFound) {
			httperror.Panic(http.StatusNotFound, "Material test work package tidak ditemukan", nil)
		}
		panic(err)
	}

	c.JSON(http.StatusOK, response.NewBody(http.StatusOK, "Material test work package berhasil dihapus"))
}
