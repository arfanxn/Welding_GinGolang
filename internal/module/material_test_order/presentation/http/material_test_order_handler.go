package http

import (
	"errors"
	"net/http"

	"github.com/arfanxn/welding/internal/infrastructure/http/helper"
	"github.com/arfanxn/welding/internal/infrastructure/http/response"
	materialTestOrderPresenter "github.com/arfanxn/welding/internal/module/material_test_order/presentation/http/presenter"
	materialTestOrderRequest "github.com/arfanxn/welding/internal/module/material_test_order/presentation/http/request"
	materialTestOrderUsecase "github.com/arfanxn/welding/internal/module/material_test_order/usecase"
	materialTestOrderDto "github.com/arfanxn/welding/internal/module/material_test_order/usecase/dto"
	"github.com/arfanxn/welding/internal/module/shared/domain/errorx"
	"github.com/arfanxn/welding/pkg/httperror"
	"github.com/arfanxn/welding/pkg/query"
	"github.com/gin-gonic/gin"
	"github.com/samber/lo"
	"go.uber.org/fx"
)

type MaterialTestOrderHandler interface {
	Paginate(c *gin.Context)
	Show(c *gin.Context)
	Store(c *gin.Context)
	Update(c *gin.Context)
	Destroy(c *gin.Context)

	StoreMedia(c *gin.Context)
}

type materialTestOrderHandler struct {
	materialTestOrderUsecase   materialTestOrderUsecase.MaterialTestOrderUsecase
	materialTestOrderPresenter materialTestOrderPresenter.MaterialTestOrderPresenter
}

type NewMaterialTestOrderHandlerParams struct {
	fx.In

	MaterialTestOrderUsecase   materialTestOrderUsecase.MaterialTestOrderUsecase
	MaterialTestOrderPresenter materialTestOrderPresenter.MaterialTestOrderPresenter
}

func NewMaterialTestOrderHandler(params NewMaterialTestOrderHandlerParams) MaterialTestOrderHandler {
	return &materialTestOrderHandler{
		materialTestOrderUsecase:   params.MaterialTestOrderUsecase,
		materialTestOrderPresenter: params.MaterialTestOrderPresenter,
	}
}

func (h *materialTestOrderHandler) Paginate(c *gin.Context) {
	q := query.NewQuery()
	c.ShouldBind(q)

	op, err := h.materialTestOrderUsecase.Paginate(c.Request.Context(), q)
	if err != nil {
		panic(err)
	}

	pp, err := h.materialTestOrderPresenter.FromEntityOffsetPaginationToViewModelPagePagination(c.Request.Context(), op)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, response.NewBodyWithData(
		http.StatusOK,
		"Material test orders berhasil diambil",
		pp,
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

	materialTestOrderVm, err := h.materialTestOrderPresenter.FromEntityToViewModel(c, materialTestOrder)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, response.NewBodyWithData(
		http.StatusOK,
		"Material test order berhasil diambil",
		gin.H{"material_test_order": materialTestOrderVm},
	))
}

func (h *materialTestOrderHandler) Store(c *gin.Context) {
	req := materialTestOrderRequest.NewStoreMaterialTestOrder()
	helper.MustBindValidate(c, req)

	materialTestOrder, err := h.materialTestOrderUsecase.Store(c.Request.Context(), &materialTestOrderDto.SaveMaterialTestOrder{
		WorkCategoryId:       &req.WorkCategoryId,       // work category
		WorkPackageName:      &req.WorkPackageName,      // work package name
		ApplicantName:        &req.ApplicantName,        // applicant name
		ApplicantPhoneNumber: &req.ApplicantPhoneNumber, // applicant phone number
		ApplicantEmail:       &req.ApplicantEmail,       // applicant email
		ApplicantFullAddress: &req.ApplicantFullAddress, // applicant full address
		ApplicantNote:        req.ApplicantNote,         // applicant note
		RecipientName:        &req.RecipientName,        // recipient name
		TesterNote:           req.TesterNote,            // tester note
		Status:               &req.Status,               // status
		OwnerUserIds:         req.OwnerUserIds,          // owner user ids
		OrderedServices: lo.Map(req.OrderedServices, // order services
			func(orderService materialTestOrderRequest.StoreMaterialTestOrderService, _ int) materialTestOrderDto.SaveMaterialTestOrderService {
				return materialTestOrderDto.SaveMaterialTestOrderService{
					ServiceId:  orderService.ServiceId,
					SampleName: orderService.SampleName,
					Quantity:   orderService.Quantity,
				}
			}),
	})
	if err != nil {
		if errors.Is(err, errorx.ErrMaterialTestOrderAlreadyExists) {
			httperror.Panic(http.StatusConflict, "Material test order sudah ada", nil)
		}
		if errors.Is(err, errorx.ErrUserNotFound) {
			httperror.Panic(http.StatusNotFound, "User tidak ditemukan", nil)
		}
		if errors.Is(err, errorx.ErrMaterialTestWorkCategoryNotFound) {
			httperror.Panic(http.StatusNotFound, "Material test work category tidak ditemukan", nil)
		}
		if errors.Is(err, errorx.ErrMaterialTestServiceNotFound) {
			httperror.Panic(http.StatusNotFound, "Material test service tidak ditemukan", nil)
		}
		// TODO: there might be more error handling here
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

	var (
		err error
	)

	materialTestOrder, err := h.materialTestOrderUsecase.Update(c.Request.Context(), &materialTestOrderDto.SaveMaterialTestOrder{
		Id:                   &req.Id,
		WorkCategoryId:       req.WorkCategoryId,       // work category
		WorkPackageName:      req.WorkPackageName,      // work package name
		ApplicantName:        req.ApplicantName,        // applicant name
		ApplicantPhoneNumber: req.ApplicantPhoneNumber, // applicant phone number
		ApplicantEmail:       req.ApplicantEmail,       // applicant email
		ApplicantFullAddress: req.ApplicantFullAddress, // applicant full address
		ApplicantNote:        req.ApplicantNote,        // applicant note
		RecipientName:        req.RecipientName,        // recipient name
		TesterNote:           req.TesterNote,           // tester note
		Status:               req.Status,               // status
		OwnerUserIds:         req.OwnerUserIds,         // owner user ids
		OrderedServices: lo.Map(req.OrderedServices,
			func(orderService materialTestOrderRequest.UpdateMaterialTestOrderService, _ int) materialTestOrderDto.SaveMaterialTestOrderService {
				return materialTestOrderDto.SaveMaterialTestOrderService{
					ServiceId:  orderService.ServiceId,
					SampleName: orderService.SampleName,
					Quantity:   orderService.Quantity,
				}
			}),
	})
	if err != nil {
		if errors.Is(err, errorx.ErrMaterialTestOrderNotFound) {
			httperror.Panic(http.StatusNotFound, "Material test order tidak ditemukan", nil)
		}
		if errors.Is(err, errorx.ErrMaterialTestOrderAlreadyExists) {
			httperror.Panic(http.StatusConflict, "Material test order sudah ada", nil)
		}
		if errors.Is(err, errorx.ErrMaterialTestWorkCategoryNotFound) {
			httperror.Panic(http.StatusNotFound, "Material test work category tidak ditemukan", nil)
		}
		if errors.Is(err, errorx.ErrMaterialTestServiceNotFound) {
			httperror.Panic(http.StatusNotFound, "Material test service tidak ditemukan", nil)
		}
		// TODO: there might be more error handling here
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
		// TODO: there might be more error handling here
		panic(err)
	}

	c.JSON(http.StatusOK, response.NewBody(http.StatusOK, "Material test order berhasil dihapus"))
}

func (h *materialTestOrderHandler) StoreMedia(c *gin.Context) {
	req := materialTestOrderRequest.NewStoreMaterialTestOrderMedia()
	req.OrderId = c.Param("id")
	helper.MustBindValidate(c, req)

	_, err := h.materialTestOrderUsecase.StoreMedia(c.Request.Context(), &materialTestOrderDto.SaveMaterialTestOrderMedia{
		OrderId: &req.OrderId,
		Name:    &req.Name,
		File:    req.File,
	})
	if err != nil {
		if errors.Is(err, errorx.ErrMaterialTestOrderNotFound) {
			httperror.Panic(http.StatusNotFound, "Material test order tidak ditemukan", nil)
		}
		// TODO: there might be more error handling here
		panic(err)
	}

	c.JSON(http.StatusOK, response.NewBody(http.StatusOK, "Material test order medias berhasil disimpan"))
}
