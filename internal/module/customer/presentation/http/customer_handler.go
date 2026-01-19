package http

import (
	"errors"
	"net/http"

	"github.com/arfanxn/welding/internal/infrastructure/http/helper"
	"github.com/arfanxn/welding/internal/infrastructure/http/response"
	customerRequest "github.com/arfanxn/welding/internal/module/customer/presentation/http/request"
	customerUsecase "github.com/arfanxn/welding/internal/module/customer/usecase"
	"github.com/arfanxn/welding/internal/module/customer/usecase/dto"
	"github.com/arfanxn/welding/internal/module/shared/domain/errorx"
	"github.com/arfanxn/welding/pkg/httperror"
	"github.com/arfanxn/welding/pkg/pagination"
	"github.com/arfanxn/welding/pkg/query"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type CustomerHandler interface {
	Paginate(c *gin.Context)
	Show(c *gin.Context)
	Store(c *gin.Context)
	Update(c *gin.Context)
	Destroy(c *gin.Context)
}

type customerHandler struct {
	customerUsecase customerUsecase.CustomerUsecase
}

type NewCustomerHandlerParams struct {
	fx.In

	CustomerUsecase customerUsecase.CustomerUsecase
}

func NewCustomerHandler(params NewCustomerHandlerParams) CustomerHandler {
	return &customerHandler{
		customerUsecase: params.CustomerUsecase,
	}
}

func (h *customerHandler) Paginate(c *gin.Context) {
	q := query.NewQuery()
	c.ShouldBind(q)

	op, err := h.customerUsecase.Paginate(c.Request.Context(), q)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, response.NewBodyWithData(
		http.StatusOK,
		"Customers berhasil diambil",
		pagination.FromOPToPP(op, helper.URLFromC(c)),
	))
}

func (h *customerHandler) Show(c *gin.Context) {
	q := query.NewQuery()
	q.FilterById(c.Param("id"))
	c.ShouldBind(q)

	customer, err := h.customerUsecase.Show(c.Request.Context(), q)
	if err != nil {
		if errors.Is(err, errorx.ErrCustomerNotFound) {
			httperror.Panic(http.StatusNotFound, "Customer tidak ditemukan", nil)
		}
		panic(err)
	}

	c.JSON(http.StatusOK, response.NewBodyWithData(
		http.StatusOK,
		"Customer berhasil diambil",
		gin.H{"customer": customer},
	))
}

func (h *customerHandler) Store(c *gin.Context) {
	req := customerRequest.NewStoreCustomerRequest()
	helper.MustBindValidate(c, req)

	customer, err := h.customerUsecase.Store(c.Request.Context(), &dto.SaveCustomer{
		AddressId:   &req.AddressId,
		Name:        &req.Name,
		PhoneNumber: &req.PhoneNumber,
		Email:       &req.Email,
	})
	if err != nil {
		if errors.Is(err, errorx.ErrCustomerAlreadyExists) {
			httperror.Panic(http.StatusConflict, "Customer sudah ada", nil)
		}
		if errors.Is(err, errorx.ErrAddressNotFound) {
			httperror.Panic(http.StatusNotFound, "Address tidak ditemukan", nil)
		}
		panic(err)
	}

	c.JSON(http.StatusCreated, response.NewBodyWithData(
		http.StatusCreated,
		"Customer berhasil disimpan",
		gin.H{"customer": customer},
	))
}

func (h *customerHandler) Update(c *gin.Context) {
	req := customerRequest.NewUpdateCustomerRequest()
	req.Id = c.Param("id")
	helper.MustBindValidate(c, req)

	customer, err := h.customerUsecase.Update(c.Request.Context(), &dto.SaveCustomer{
		Id:          &req.Id,
		AddressId:   &req.AddressId,
		Name:        &req.Name,
		PhoneNumber: &req.PhoneNumber,
		Email:       &req.Email,
	})
	if err != nil {
		if errors.Is(err, errorx.ErrCustomerNotFound) {
			httperror.Panic(http.StatusNotFound, "Customer tidak ditemukan", nil)
		}
		if errors.Is(err, errorx.ErrMaterialTestWorkPackageAlreadyExists) {
			httperror.Panic(http.StatusConflict, "Customer sudah ada", nil)
		}
		if errors.Is(err, errorx.ErrAddressNotFound) {
			httperror.Panic(http.StatusNotFound, "Address tidak ditemukan", nil)
		}
		panic(err)
	}

	c.JSON(http.StatusOK, response.NewBodyWithData(
		http.StatusOK,
		"Customer berhasil diperbarui",
		gin.H{"customer": customer},
	))
}

func (h *customerHandler) Destroy(c *gin.Context) {
	req := customerRequest.NewDestroyCustomerRequest()
	req.Id = c.Param("id")
	helper.MustBindValidate(c, req)

	err := h.customerUsecase.Destroy(c.Request.Context(), &dto.DestroyCustomer{Id: req.Id})
	if err != nil {
		if errors.Is(err, errorx.ErrCustomerNotFound) {
			httperror.Panic(http.StatusNotFound, "Customer tidak ditemukan", nil)
		}
		panic(err)
	}

	c.JSON(http.StatusOK, response.NewBody(http.StatusOK, "Customer berhasil dihapus"))
}
