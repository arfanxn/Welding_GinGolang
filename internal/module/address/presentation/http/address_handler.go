package http

import (
	"errors"
	"net/http"

	"github.com/arfanxn/welding/internal/infrastructure/http/helper"
	"github.com/arfanxn/welding/internal/infrastructure/http/response"
	addressRequest "github.com/arfanxn/welding/internal/module/address/presentation/http/request"
	addressUsecase "github.com/arfanxn/welding/internal/module/address/usecase"
	"github.com/arfanxn/welding/internal/module/address/usecase/dto"
	"github.com/arfanxn/welding/internal/module/shared/domain/errorx"
	"github.com/arfanxn/welding/pkg/httperror"
	"github.com/arfanxn/welding/pkg/pagination"
	"github.com/arfanxn/welding/pkg/query"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type AddressHandler interface {
	Paginate(c *gin.Context)
	Show(c *gin.Context)
	Store(c *gin.Context)
	Update(c *gin.Context)
	Destroy(c *gin.Context)
}

type addressHandler struct {
	addressUsecase addressUsecase.AddressUsecase
}

type NewAddressHandlerParams struct {
	fx.In

	AddressUsecase addressUsecase.AddressUsecase
}

func NewAddressHandler(params NewAddressHandlerParams) AddressHandler {
	return &addressHandler{
		addressUsecase: params.AddressUsecase,
	}
}

func (h *addressHandler) Paginate(c *gin.Context) {
	q := query.NewQuery()
	c.ShouldBind(q)

	op, err := h.addressUsecase.Paginate(c.Request.Context(), q)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, response.NewBodyWithData(
		http.StatusOK,
		"Addresses berhasil diambil",
		pagination.FromOPToPP(op, helper.URLFromC(c)),
	))
}

func (h *addressHandler) Show(c *gin.Context) {
	q := query.NewQuery()
	q.FilterById(c.Param("id"))
	c.ShouldBind(q)

	mtm, err := h.addressUsecase.Show(c.Request.Context(), q)
	if err != nil {
		if errors.Is(err, errorx.ErrAddressNotFound) {
			httperror.Panic(http.StatusNotFound, "Address tidak ditemukan", nil)
		}
		panic(err)
	}

	c.JSON(http.StatusOK, response.NewBodyWithData(
		http.StatusOK,
		"Address berhasil diambil",
		gin.H{"address": mtm},
	))
}

func (h *addressHandler) Store(c *gin.Context) {
	req := addressRequest.NewStoreAddressRequest()
	helper.MustBindValidate(c, req)

	mtm, err := h.addressUsecase.Store(c.Request.Context(), &dto.SaveAddress{
		FullAddress: &req.FullAddress,
	})
	if err != nil {
		if errors.Is(err, errorx.ErrAddressAlreadyExists) {
			httperror.Panic(http.StatusConflict, "Address sudah ada", nil)
		}
		panic(err)
	}

	c.JSON(http.StatusCreated, response.NewBodyWithData(
		http.StatusCreated,
		"Address berhasil disimpan",
		gin.H{"address": mtm},
	))
}

func (h *addressHandler) Update(c *gin.Context) {
	req := addressRequest.NewUpdateAddressRequest()
	req.Id = c.Param("id")
	helper.MustBindValidate(c, req)

	mtm, err := h.addressUsecase.Update(c.Request.Context(), &dto.SaveAddress{
		Id:          &req.Id,
		FullAddress: &req.FullAddress,
	})
	if err != nil {
		if errors.Is(err, errorx.ErrAddressNotFound) {
			httperror.Panic(http.StatusNotFound, "Address tidak ditemukan", nil)
		}
		if errors.Is(err, errorx.ErrAddressAlreadyExists) {
			httperror.Panic(http.StatusConflict, "Address sudah ada", nil)
		}
		panic(err)
	}

	c.JSON(http.StatusOK, response.NewBodyWithData(
		http.StatusOK,
		"Address berhasil diperbarui",
		gin.H{"address": mtm},
	))
}

func (h *addressHandler) Destroy(c *gin.Context) {
	req := addressRequest.NewDestroyAddressRequest()
	req.Id = c.Param("id")
	helper.MustBindValidate(c, req)

	err := h.addressUsecase.Destroy(c.Request.Context(), &dto.DestroyAddress{Id: req.Id})
	if err != nil {
		if errors.Is(err, errorx.ErrAddressNotFound) {
			httperror.Panic(http.StatusNotFound, "Address tidak ditemukan", nil)
		}
		panic(err)
	}

	c.JSON(http.StatusOK, response.NewBody(http.StatusOK, "Address berhasil dihapus"))
}
