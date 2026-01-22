package http

import (
	"errors"
	"net/http"
	"strings"

	"github.com/arfanxn/welding/internal/infrastructure/http/helper"
	"github.com/arfanxn/welding/internal/infrastructure/http/response"
	mtsRequest "github.com/arfanxn/welding/internal/module/material_test_service/presentation/http/request"
	"github.com/arfanxn/welding/internal/module/material_test_service/usecase"
	"github.com/arfanxn/welding/internal/module/material_test_service/usecase/dto"
	"github.com/arfanxn/welding/internal/module/shared/domain/errorx"
	"github.com/arfanxn/welding/pkg/httperror"
	"github.com/arfanxn/welding/pkg/pagination"
	"github.com/arfanxn/welding/pkg/query"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type MaterialTestServiceHandler interface {
	Paginate(c *gin.Context)
	Show(c *gin.Context)
	Store(c *gin.Context)
	Update(c *gin.Context)
	Destroy(c *gin.Context)
}

type mtsHandler struct {
	mtsUsecase usecase.MaterialTestServiceUsecase
}

type NewMaterialTestServiceHandlerParams struct {
	fx.In

	MaterialTestServiceUsecase usecase.MaterialTestServiceUsecase
}

func NewMaterialTestServiceHandler(params NewMaterialTestServiceHandlerParams) MaterialTestServiceHandler {
	return &mtsHandler{
		mtsUsecase: params.MaterialTestServiceUsecase,
	}
}

func (h *mtsHandler) Paginate(c *gin.Context) {
	q := query.NewQuery()
	c.ShouldBind(q)

	op, err := h.mtsUsecase.Paginate(c.Request.Context(), q)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, response.NewBodyWithData(
		http.StatusOK,
		"Material test services berhasil diambil",
		pagination.FromOPToPP(op, helper.URLFromC(c)),
	))
}

func (h *mtsHandler) Show(c *gin.Context) {
	q := query.NewQuery()
	q.FilterById(c.Param("id"))
	c.ShouldBind(q)

	mts, err := h.mtsUsecase.Show(c.Request.Context(), q)
	if err != nil {
		if errors.Is(err, errorx.ErrMaterialTestServiceNotFound) {
			httperror.Panic(http.StatusNotFound, "Material test service tidak ditemukan", nil)
		}
		panic(err)
	}

	c.JSON(http.StatusOK, response.NewBodyWithData(
		http.StatusOK,
		"Material test service berhasil diambil",
		gin.H{"material_test_service": mts},
	))
}

func (h *mtsHandler) Store(c *gin.Context) {
	req := mtsRequest.NewStoreMaterialTestService()
	helper.MustBindValidate(c, req)

	machineId := normalizeOptionalString(req.MachineId)
	methodId := normalizeOptionalString(req.MethodId)

	mts, err := h.mtsUsecase.Store(c.Request.Context(), &dto.SaveMaterialTestService{
		MachineId:   machineId,
		MethodId:    methodId,
		TestName:    &req.TestName,
		ServiceType: &req.ServiceType,
		ServiceCode: &req.ServiceCode,
		Unit:        &req.Unit,
		Price:       &req.Price,
	})
	if err != nil {
		if errors.Is(err, errorx.ErrMaterialTestServiceAlreadyExists) {
			httperror.Panic(http.StatusConflict, "Material test service sudah ada", nil)
		}
		if errors.Is(err, errorx.ErrMaterialTestMachineNotFound) {
			httperror.Panic(http.StatusNotFound, "Machine tidak ditemukan", nil)
		}
		if errors.Is(err, errorx.ErrMaterialTestMethodNotFound) {
			httperror.Panic(http.StatusNotFound, "Method tidak ditemukan", nil)
		}
		panic(err)
	}

	c.JSON(http.StatusCreated, response.NewBodyWithData(
		http.StatusCreated,
		"Material test service berhasil disimpan",
		gin.H{"material_test_service": mts},
	))
}

func (h *mtsHandler) Update(c *gin.Context) {
	req := mtsRequest.NewUpdateMaterialTestService()
	req.Id = c.Param("id")
	helper.MustBindValidate(c, req)

	machineId := normalizeOptionalString(req.MachineId)
	methodId := normalizeOptionalString(req.MethodId)

	mts, err := h.mtsUsecase.Update(c.Request.Context(), &dto.SaveMaterialTestService{
		Id:          &req.Id,
		MachineId:   machineId,
		MethodId:    methodId,
		TestName:    req.TestName,
		ServiceType: req.ServiceType,
		ServiceCode: req.ServiceCode,
		Unit:        req.Unit,
		Price:       req.Price,
	})
	if err != nil {
		if errors.Is(err, errorx.ErrMaterialTestServiceNotFound) {
			httperror.Panic(http.StatusNotFound, "Material test service tidak ditemukan", nil)
		}
		if errors.Is(err, errorx.ErrMaterialTestServiceAlreadyExists) {
			httperror.Panic(http.StatusConflict, "Material test service sudah ada", nil)
		}
		if errors.Is(err, errorx.ErrMaterialTestMachineNotFound) {
			httperror.Panic(http.StatusNotFound, "Machine tidak ditemukan", nil)
		}
		if errors.Is(err, errorx.ErrMaterialTestMethodNotFound) {
			httperror.Panic(http.StatusNotFound, "Method tidak ditemukan", nil)
		}
		panic(err)
	}

	c.JSON(http.StatusOK, response.NewBodyWithData(
		http.StatusOK,
		"Material test service berhasil diperbarui",
		gin.H{"material_test_service": mts},
	))
}

func (h *mtsHandler) Destroy(c *gin.Context) {
	req := mtsRequest.NewDestroyMaterialTestService()
	req.Id = c.Param("id")
	helper.MustBindValidate(c, req)

	err := h.mtsUsecase.Destroy(c.Request.Context(), &dto.DestroyMaterialTestService{Id: req.Id})
	if err != nil {
		if errors.Is(err, errorx.ErrMaterialTestServiceNotFound) {
			httperror.Panic(http.StatusNotFound, "Material test service tidak ditemukan", nil)
		}
		panic(err)
	}

	c.JSON(http.StatusOK, response.NewBody(http.StatusOK, "Material test service berhasil dihapus"))
}

func normalizeOptionalString(value *string) *string {
	if value == nil {
		return nil
	}
	trimmed := strings.TrimSpace(*value)
	if trimmed == "" {
		return nil
	}
	return &trimmed
}
