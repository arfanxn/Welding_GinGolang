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
	"github.com/arfanxn/welding/pkg/boolutil"
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

	// Status relateds
	Approve(c *gin.Context)
	Reject(c *gin.Context)
	Cancel(c *gin.Context)
	SubmitPayment(c *gin.Context)
	ApprovePayment(c *gin.Context)
	RejectPayment(c *gin.Context)
	Test(c *gin.Context)
	Refund(c *gin.Context)
	Complete(c *gin.Context)

	// Service evaluation related
	UpdateOrderServiceEvaluation(c *gin.Context)

	// Media relateds
	StoreMedia(c *gin.Context)
	UpdateMedia(c *gin.Context)
	DestroyMedia(c *gin.Context)
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
		ApplicantCompanyName: req.ApplicantCompanyName,  // applicant company name
		ApplicantPhoneNumber: &req.ApplicantPhoneNumber, // applicant phone number
		ApplicantEmail:       &req.ApplicantEmail,       // applicant email
		ApplicantFullAddress: &req.ApplicantFullAddress, // applicant full address
		ApplicantNote:        req.ApplicantNote,         // applicant note
		RecipientName:        &req.RecipientName,        // recipient name
		RecipientFullAddress: req.RecipientFullAddress,  // recipient full address
		TesterNote:           req.TesterNote,            // tester note
		Status:               &req.Status,               // status
		OwnerUserIds:         req.OwnerUserIds,          // owner user ids
		OrderedServices: boolutil.Ternary(req.OrderedServices != nil, lo.Map(req.OrderedServices, // order services
			func(orderService materialTestOrderRequest.StoreMaterialTestOrderService, _ int) materialTestOrderDto.SaveMaterialTestOrderService {
				return materialTestOrderDto.SaveMaterialTestOrderService{
					ServiceId:  orderService.ServiceId,
					SampleName: orderService.SampleName,
					Quantity:   orderService.Quantity,
				}
			}), nil),
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
		panic(err)
	}

	materialTestOrderVm, err := h.materialTestOrderPresenter.FromEntityToViewModel(c, materialTestOrder)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusCreated, response.NewBodyWithData(
		http.StatusCreated,
		"Material test order berhasil disimpan",
		gin.H{"material_test_order": materialTestOrderVm},
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
		ApplicantCompanyName: req.ApplicantCompanyName, // applicant company name
		ApplicantPhoneNumber: req.ApplicantPhoneNumber, // applicant phone number
		ApplicantEmail:       req.ApplicantEmail,       // applicant email
		ApplicantFullAddress: req.ApplicantFullAddress, // applicant full address
		ApplicantNote:        req.ApplicantNote,        // applicant note
		RecipientName:        req.RecipientName,        // recipient name
		RecipientFullAddress: req.RecipientFullAddress, // recipient full address
		TesterNote:           req.TesterNote,           // tester note
		Status:               req.Status,               // status
		OwnerUserIds:         req.OwnerUserIds,         // owner user ids
		OrderedServices: boolutil.Ternary(req.OrderedServices != nil, lo.Map(req.OrderedServices,
			func(orderService materialTestOrderRequest.UpdateMaterialTestOrderService, _ int) materialTestOrderDto.SaveMaterialTestOrderService {
				return materialTestOrderDto.SaveMaterialTestOrderService{
					ServiceId:  orderService.ServiceId,
					SampleName: orderService.SampleName,
					Quantity:   orderService.Quantity,
				}
			}), nil),
	})
	if err != nil {
		if errors.Is(err, errorx.ErrMaterialTestOrderNotFound) {
			httperror.Panic(http.StatusNotFound, "Material test order tidak ditemukan", nil)
		}
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
		if errors.Is(err, errorx.ErrMaterialTestOrderUpdateForbidden) {
			httperror.Panic(http.StatusForbidden, "User tidak memiliki izin untuk memperbarui material test order ini", nil)
		}
		if errors.Is(err, errorx.ErrMaterialTestOrderStatusNotUpdateableUpdateForbidden) {
			httperror.Panic(http.StatusForbidden, "Material test order tidak dalam status yang dapat diperbarui", nil)
		}
		panic(err)
	}

	materialTestOrderVm, err := h.materialTestOrderPresenter.FromEntityToViewModel(c, materialTestOrder)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, response.NewBodyWithData(
		http.StatusOK,
		"Material test order berhasil diperbarui",
		gin.H{"material_test_order": materialTestOrderVm},
	))
}

/* ==============================
		Status relateds
============================== */

func (h *materialTestOrderHandler) Approve(c *gin.Context) {
	req := materialTestOrderRequest.NewApproveMaterialTestOrder()
	req.Id = c.Param("id")
	helper.MustBindValidate(c, req)

	materialTestOrder, err := h.materialTestOrderUsecase.Approve(c.Request.Context(), &materialTestOrderDto.ApproveMaterialTestOrder{
		Id: req.Id,
	})
	if err != nil {
		if errors.Is(err, errorx.ErrMaterialTestOrderNotFound) {
			httperror.Panic(http.StatusNotFound, "Material test order tidak ditemukan", nil)
		}
		if errors.Is(err, errorx.ErrMaterialTestOrderStatusNotAwaitingReviewApproveForbidden) {
			httperror.Panic(http.StatusForbidden, "Material test order tidak dalam status yang dapat disetujui (approved / awaiting_payment)", nil)
		}
		panic(err)
	}

	materialTestOrderVm, err := h.materialTestOrderPresenter.FromEntityToViewModel(c, materialTestOrder)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, response.NewBodyWithData(
		http.StatusOK,
		"Material test order berhasil disetujui (approved)",
		gin.H{"material_test_order": materialTestOrderVm},
	))
}

func (h *materialTestOrderHandler) Reject(c *gin.Context) {
	req := materialTestOrderRequest.NewApproveMaterialTestOrder()
	req.Id = c.Param("id")
	helper.MustBindValidate(c, req)

	materialTestOrder, err := h.materialTestOrderUsecase.Reject(c.Request.Context(), &materialTestOrderDto.RejectMaterialTestOrder{
		Id: req.Id,
	})
	if err != nil {
		if errors.Is(err, errorx.ErrMaterialTestOrderNotFound) {
			httperror.Panic(http.StatusNotFound, "Material test order tidak ditemukan", nil)
		}
		if errors.Is(err, errorx.ErrMaterialTestOrderStatusNotAwaitingReviewRejectForbidden) {
			httperror.Panic(http.StatusForbidden, "Material test order tidak dalam status yang dapat ditolak (rejected)", nil)
		}
		panic(err)
	}

	materialTestOrderVm, err := h.materialTestOrderPresenter.FromEntityToViewModel(c, materialTestOrder)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, response.NewBodyWithData(
		http.StatusOK,
		"Material test order berhasil ditolak (rejected)",
		gin.H{"material_test_order": materialTestOrderVm},
	))
}

func (h *materialTestOrderHandler) Cancel(c *gin.Context) {
	req := materialTestOrderRequest.NewCancelMaterialTestOrder()
	req.Id = c.Param("id")
	helper.MustBindValidate(c, req)

	materialTestOrder, err := h.materialTestOrderUsecase.Cancel(c.Request.Context(), &materialTestOrderDto.CancelMaterialTestOrder{
		Id: req.Id,
	})
	if err != nil {
		if errors.Is(err, errorx.ErrMaterialTestOrderNotFound) {
			httperror.Panic(http.StatusNotFound, "Material test order tidak ditemukan", nil)
		}
		if errors.Is(err, errorx.ErrMaterialTestOrderStatusNotCancellableCancelForbidden) {
			httperror.Panic(http.StatusForbidden, "Material test order tidak dalam status yang dapat dibatalkan (cancelled)", nil)
		}
		if errors.Is(err, errorx.ErrMaterialTestOrderUpdateForbidden) {
			httperror.Panic(http.StatusForbidden, "User tidak memiliki izin untuk membatalkan material test order ini", nil)
		}

		panic(err)
	}

	materialTestOrderVm, err := h.materialTestOrderPresenter.FromEntityToViewModel(c, materialTestOrder)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, response.NewBodyWithData(
		http.StatusOK,
		"Material test order berhasil dibatalkan (cancelled)",
		gin.H{"material_test_order": materialTestOrderVm},
	))
}

func (h *materialTestOrderHandler) SubmitPayment(c *gin.Context) {
	req := materialTestOrderRequest.NewSubmitPaymentMaterialTestOrder()
	req.Id = c.Param("id")
	helper.MustBindValidate(c, req)

	materialTestOrder, err := h.materialTestOrderUsecase.SubmitPayment(c.Request.Context(), &materialTestOrderDto.SubmitPaymentMaterialTestOrder{
		Id:   req.Id,
		File: req.File,
	})
	if err != nil {
		if errors.Is(err, errorx.ErrMaterialTestOrderNotFound) {
			httperror.Panic(http.StatusNotFound, "Material test order tidak ditemukan", nil)
		}
		if errors.Is(err, errorx.ErrMaterialTestOrderStatusNotPaymentSubmittableSubmitPaymentForbidden) {
			httperror.Panic(http.StatusForbidden, "Material test order tidak dalam status yang dapat mengirim bukti pembayaran (payment_submitted)", nil)
		}
		if errors.Is(err, errorx.ErrMaterialTestOrderUpdateForbidden) {
			httperror.Panic(http.StatusForbidden, "User tidak memiliki izin untuk mengirim bukti pembayaran material test order ini", nil)
		}
		panic(err)
	}

	materialTestOrderVm, err := h.materialTestOrderPresenter.FromEntityToViewModel(c, materialTestOrder)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, response.NewBodyWithData(
		http.StatusOK,
		"Material test order bukti pembayaran dikirim (payment submitted)",
		gin.H{"material_test_order": materialTestOrderVm},
	))
}

func (h *materialTestOrderHandler) ApprovePayment(c *gin.Context) {
	req := materialTestOrderRequest.NewApprovePaymentMaterialTestOrder()
	req.Id = c.Param("id")
	helper.MustBindValidate(c, req)

	materialTestOrder, err := h.materialTestOrderUsecase.ApprovePayment(c.Request.Context(), &materialTestOrderDto.ApprovePaymentMaterialTestOrder{
		Id: req.Id,
	})
	if err != nil {
		if errors.Is(err, errorx.ErrMaterialTestOrderNotFound) {
			httperror.Panic(http.StatusNotFound, "Material test order tidak ditemukan", nil)
		}
		if errors.Is(err, errorx.ErrMaterialTestOrderStatusNotPaymentSubmittedApprovePaymentForbidden) {
			httperror.Panic(http.StatusForbidden, "Material test order tidak dalam status yang dapat disetujui pembayaran (payment_approved)", nil)
		}
		panic(err)
	}

	materialTestOrderVm, err := h.materialTestOrderPresenter.FromEntityToViewModel(c, materialTestOrder)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, response.NewBodyWithData(
		http.StatusOK,
		"Material test order pembayaran berhasil disetujui (payment approved)",
		gin.H{"material_test_order": materialTestOrderVm},
	))
}

func (h *materialTestOrderHandler) RejectPayment(c *gin.Context) {
	req := materialTestOrderRequest.NewRejectPaymentMaterialTestOrder()
	req.Id = c.Param("id")
	helper.MustBindValidate(c, req)

	materialTestOrder, err := h.materialTestOrderUsecase.RejectPayment(c.Request.Context(), &materialTestOrderDto.RejectPaymentMaterialTestOrder{
		Id: req.Id,
	})
	if err != nil {
		if errors.Is(err, errorx.ErrMaterialTestOrderNotFound) {
			httperror.Panic(http.StatusNotFound, "Material test order tidak ditemukan", nil)
		}
		if errors.Is(err, errorx.ErrMaterialTestOrderStatusNotPaymentSubmittedRejectPaymentForbidden) {
			httperror.Panic(http.StatusForbidden, "Material test order tidak dalam status yang dapat ditolak pembayaran (payment_rejected)", nil)
		}
		panic(err)
	}

	materialTestOrderVm, err := h.materialTestOrderPresenter.FromEntityToViewModel(c, materialTestOrder)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, response.NewBodyWithData(
		http.StatusOK,
		"Material test order pembayaran berhasil ditolak (payment rejected)",
		gin.H{"material_test_order": materialTestOrderVm},
	))
}

func (h *materialTestOrderHandler) Test(c *gin.Context) {
	req := materialTestOrderRequest.NewTestMaterialTestOrder()
	req.Id = c.Param("id")
	helper.MustBindValidate(c, req)

	materialTestOrder, err := h.materialTestOrderUsecase.Test(c.Request.Context(), &materialTestOrderDto.TestMaterialTestOrder{
		Id: req.Id,
	})
	if err != nil {
		if errors.Is(err, errorx.ErrMaterialTestOrderNotFound) {
			httperror.Panic(http.StatusNotFound, "Material test order tidak ditemukan", nil)
		}
		if errors.Is(err, errorx.ErrMaterialTestOrderStatusNotPaymentApprovedTestForbidden) {
			httperror.Panic(http.StatusForbidden, "Material test order tidak dalam status yang dapat dilanjutkan ke tahap uji (test / testing)", nil)
		}
		panic(err)
	}

	materialTestOrderVm, err := h.materialTestOrderPresenter.FromEntityToViewModel(c, materialTestOrder)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, response.NewBodyWithData(
		http.StatusOK,
		"Material test order dilanjutkan ke tahap uji (testing)",
		gin.H{"material_test_order": materialTestOrderVm},
	))
}

func (h *materialTestOrderHandler) Refund(c *gin.Context) {
	req := materialTestOrderRequest.NewRefundMaterialTestOrder()
	req.Id = c.Param("id")
	helper.MustBindValidate(c, req)

	materialTestOrder, err := h.materialTestOrderUsecase.Refund(c.Request.Context(), &materialTestOrderDto.RefundMaterialTestOrder{
		Id:   req.Id,
		File: req.File,
	})
	if err != nil {
		if errors.Is(err, errorx.ErrMaterialTestOrderNotFound) {
			httperror.Panic(http.StatusNotFound, "Material test order tidak ditemukan", nil)
		}
		if errors.Is(err, errorx.ErrMaterialTestOrderStatusNotRefundableRefundForbidden) {
			httperror.Panic(http.StatusForbidden, "Material test order tidak dalam status yang dapat dikembalikan (refunded)", nil)
		}
		panic(err)
	}

	materialTestOrderVm, err := h.materialTestOrderPresenter.FromEntityToViewModel(c, materialTestOrder)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, response.NewBodyWithData(
		http.StatusOK,
		"Material test order pembayaran berhasil dikembalikan (refunded)",
		gin.H{"material_test_order": materialTestOrderVm},
	))
}

func (h *materialTestOrderHandler) Complete(c *gin.Context) {
	req := materialTestOrderRequest.NewCompleteMaterialTestOrder()
	req.Id = c.Param("id")
	helper.MustBindValidate(c, req)

	materialTestOrder, err := h.materialTestOrderUsecase.Complete(c.Request.Context(), &materialTestOrderDto.CompleteMaterialTestOrder{
		Id: req.Id,
	})
	if err != nil {
		if errors.Is(err, errorx.ErrMaterialTestOrderNotFound) {
			httperror.Panic(http.StatusNotFound, "Material test order tidak ditemukan", nil)
		}
		if errors.Is(err, errorx.ErrMaterialTestOrderStatusNotTestingCompleteForbidden) {
			httperror.Panic(http.StatusForbidden, "Material test order tidak dalam status yang dapat diselesaikan (completed)", nil)
		}
		panic(err)
	}

	materialTestOrderVm, err := h.materialTestOrderPresenter.FromEntityToViewModel(c, materialTestOrder)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, response.NewBodyWithData(
		http.StatusOK,
		"Material test order berhasil diselesaikan (completed)",
		gin.H{"material_test_order": materialTestOrderVm},
	))
}

/* ==============================
	Service evaluation related
============================== */

func (h *materialTestOrderHandler) UpdateOrderServiceEvaluation(c *gin.Context) {
	req := materialTestOrderRequest.NewUpdateMaterialTestOrderServiceEvaluation()
	req.Id = c.Param("id")
	req.OrderServiceEvaluationId = c.Param("order_service_evaluation_id")
	helper.MustBindValidate(c, req)

	materialTestOrder, err := h.materialTestOrderUsecase.UpdateOrderServiceEvaluation(c.Request.Context(), &materialTestOrderDto.UpdateMaterialTestOrderServiceEvaluation{
		Id:                       req.Id,
		OrderServiceEvaluationId: req.OrderServiceEvaluationId,

		IsEquipmentAvailable:      req.IsEquipmentAvailable,
		IsPersonnelAvailable:      req.IsPersonnelAvailable,
		IsTimeAvailable:           req.IsTimeAvailable,
		IsTestReady:               req.IsTestReady,
		IsSubcontractLabAvailable: req.IsSubcontractLabAvailable,
	})
	if err != nil {
		if errors.Is(err, errorx.ErrMaterialTestOrderNotFound) {
			httperror.Panic(http.StatusNotFound, "Material test order tidak ditemukan", nil)
		}
		if errors.Is(err, errorx.ErrMaterialTestOrderServiceEvaluationNotFound) {
			httperror.Panic(http.StatusNotFound, "Material test order service evaluation tidak ditemukan", nil)
		}
		if errors.Is(err, errorx.ErrMaterialTestOrderStatusNotAwaitingReviewUpdateServiceEvaluationForbidden) {
			httperror.Panic(http.StatusForbidden, "Material test order tidak dalam status yang dapat diperbarui evaluasi layanan", nil)
		}
		panic(err)
	}

	materialTestOrderVm, err := h.materialTestOrderPresenter.FromEntityToViewModel(c, materialTestOrder)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusCreated, response.NewBodyWithData(
		http.StatusCreated,
		"Material test order service evaluation berhasil diperbarui",
		gin.H{"material_test_order": materialTestOrderVm},
	))
}

/* ==============================
		Media relateds
============================== */

func (h *materialTestOrderHandler) StoreMedia(c *gin.Context) {
	req := materialTestOrderRequest.NewStoreMaterialTestOrderMedia()
	req.Id = c.Param("id")
	helper.MustBindValidate(c, req)

	materialTestOrder, err := h.materialTestOrderUsecase.StoreMedia(c.Request.Context(), &materialTestOrderDto.SaveMaterialTestOrderMedia{
		Id:   &req.Id,
		Name: &req.Name,
		File: req.File,
	})
	if err != nil {
		if errors.Is(err, errorx.ErrMaterialTestOrderNotFound) {
			httperror.Panic(http.StatusNotFound, "Material test order tidak ditemukan", nil)
		}
		if errors.Is(err, errorx.ErrMaterialTestOrderUpdateForbidden) {
			httperror.Panic(http.StatusForbidden, "User tidak memiliki izin untuk memperbarui material test order ini", nil)
		}
		panic(err)
	}

	materialTestOrderVm, err := h.materialTestOrderPresenter.FromEntityToViewModel(c, materialTestOrder)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusCreated, response.NewBodyWithData(
		http.StatusCreated,
		"Material test order media berhasil disimpan",
		gin.H{"material_test_order": materialTestOrderVm},
	))
}

func (h *materialTestOrderHandler) UpdateMedia(c *gin.Context) {
	req := materialTestOrderRequest.NewUpdateMaterialTestOrderMedia()
	req.Id = c.Param("id")
	req.MediaId = c.Param("media_id")
	helper.MustBindValidate(c, req)

	materialTestOrder, err := h.materialTestOrderUsecase.UpdateMedia(c.Request.Context(), &materialTestOrderDto.SaveMaterialTestOrderMedia{
		Id:      &req.Id,
		MediaId: &req.MediaId,
		Name:    &req.Name,
	})
	if err != nil {
		if errors.Is(err, errorx.ErrMaterialTestOrderNotFound) {
			httperror.Panic(http.StatusNotFound, "Material test order tidak ditemukan", nil)
		}
		if errors.Is(err, errorx.ErrMediaNotFound) {
			httperror.Panic(http.StatusNotFound, "Media tidak ditemukan", nil)
		}
		if errors.Is(err, errorx.ErrMaterialTestOrderUpdateForbidden) {
			httperror.Panic(http.StatusForbidden, "User tidak memiliki izin untuk memperbarui material test order ini", nil)
		}
		if errors.Is(err, errorx.ErrMaterialTestOrderMediaPaymentProofUpdateForbidden) {
			httperror.Panic(http.StatusForbidden, "Material test order media payment proof tidak dapat diperbarui", nil)
		}
		if errors.Is(err, errorx.ErrMaterialTestOrderMediaRefundProofUpdateForbidden) {
			httperror.Panic(http.StatusForbidden, "Material test order media refund proof tidak dapat diperbarui", nil)
		}
		panic(err)
	}

	materialTestOrderVm, err := h.materialTestOrderPresenter.FromEntityToViewModel(c, materialTestOrder)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, response.NewBodyWithData(
		http.StatusOK,
		"Material test order media berhasil diperbarui",
		gin.H{"material_test_order": materialTestOrderVm},
	))
}

func (h *materialTestOrderHandler) DestroyMedia(c *gin.Context) {
	req := materialTestOrderRequest.NewDestroyMaterialTestOrderMedia()
	req.Id = c.Param("id")
	req.MediaId = c.Param("media_id")
	helper.MustBindValidate(c, req)

	materialTestOrder, err := h.materialTestOrderUsecase.DestroyMedia(c.Request.Context(), &materialTestOrderDto.DestroyMaterialTestOrderMedia{
		Id:      req.Id,
		MediaId: req.MediaId,
	})
	if err != nil {
		if errors.Is(err, errorx.ErrMaterialTestOrderNotFound) {
			httperror.Panic(http.StatusNotFound, "Material test order tidak ditemukan", nil)
		}
		if errors.Is(err, errorx.ErrMediaNotFound) {
			httperror.Panic(http.StatusNotFound, "Media tidak ditemukan", nil)
		}
		if errors.Is(err, errorx.ErrMaterialTestOrderUpdateForbidden) {
			httperror.Panic(http.StatusForbidden, "User tidak memiliki izin untuk memperbarui material test order ini", nil)
		}
		if errors.Is(err, errorx.ErrMaterialTestOrderMediaPaymentProofDestroyForbidden) {
			httperror.Panic(http.StatusForbidden, "Material test order media payment proof tidak dapat dihapus", nil)
		}
		if errors.Is(err, errorx.ErrMaterialTestOrderMediaRefundProofDestroyForbidden) {
			httperror.Panic(http.StatusForbidden, "Material test order media refund proof tidak dapat dihapus", nil)
		}
		panic(err)
	}

	materialTestOrderVm, err := h.materialTestOrderPresenter.FromEntityToViewModel(c, materialTestOrder)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, response.NewBodyWithData(
		http.StatusOK,
		"Material test order media berhasil dihapus",
		gin.H{"material_test_order": materialTestOrderVm},
	))
}
