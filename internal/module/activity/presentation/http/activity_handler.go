package http

import (
	"errors"
	"net/http"

	"github.com/arfanxn/welding/internal/infrastructure/http/response"
	activityEnum "github.com/arfanxn/welding/internal/module/activity/domain/enum"
	activityPresenter "github.com/arfanxn/welding/internal/module/activity/presentation/http/presenter"
	"github.com/arfanxn/welding/internal/module/activity/usecase"
	"github.com/arfanxn/welding/internal/module/shared/domain/errorx"
	"github.com/arfanxn/welding/pkg/httperror"
	"github.com/arfanxn/welding/pkg/query"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type ActivityHandler interface {
	EnumValues(c *gin.Context)
	Index(c *gin.Context)
	Show(c *gin.Context)
}

type activityHandler struct {
	activityUsecase   usecase.ActivityUsecase
	activityPresenter activityPresenter.ActivityPresenter
}

type NewActivityHandlerParams struct {
	fx.In

	ActivityUsecase   usecase.ActivityUsecase
	ActivityPresenter activityPresenter.ActivityPresenter
}

func NewActivityHandler(params NewActivityHandlerParams) ActivityHandler {
	return &activityHandler{
		activityUsecase:   params.ActivityUsecase,
		activityPresenter: params.ActivityPresenter,
	}
}

func (h *activityHandler) EnumValues(c *gin.Context) {
	c.JSON(http.StatusOK, response.NewBodyWithData(
		http.StatusOK,
		"Activity enums berhasil diambil",
		gin.H{
			"actions":       activityEnum.ActivityActions,
			"causer_types":  activityEnum.ActivityCauserTypes,
			"subject_types": activityEnum.ActivitySubjectTypes,
		},
	))
}

func (h *activityHandler) Index(c *gin.Context) {
	q := query.NewQuery()
	c.ShouldBind(q)

	op, err := h.activityUsecase.Paginate(c.Request.Context(), q)
	if err != nil {
		if errors.Is(err, errorx.ErrActivityInvalidAction) {
			httperror.Panic(http.StatusBadRequest, "Invalid activity action", nil)
		}
		if errors.Is(err, errorx.ErrActivityInvalidCauserType) {
			httperror.Panic(http.StatusBadRequest, "Invalid activity causer type", nil)
		}
		if errors.Is(err, errorx.ErrActivityInvalidSubjectType) {
			httperror.Panic(http.StatusBadRequest, "Invalid activity subject type", nil)
		}
		panic(err)
	}

	pp, err := h.activityPresenter.FromEntityOffsetPaginationToViewModelPagePagination(c.Request.Context(), op)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, response.NewBodyWithData(
		http.StatusOK,
		"Activities berhasil diambil",
		pp,
	))
}

func (h *activityHandler) Show(c *gin.Context) {
	q := query.NewQuery()
	q.FilterById(c.Param("id"))
	c.ShouldBind(q)

	activity, err := h.activityUsecase.Show(c.Request.Context(), q)
	if err != nil {
		if errors.Is(err, errorx.ErrActivityNotFound) {
			httperror.Panic(http.StatusNotFound, "Activity tidak ditemukan", nil)
		}
		panic(err)
	}

	activityVm, err := h.activityPresenter.FromEntityToViewModel(c, activity)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, response.NewBodyWithData(
		http.StatusOK,
		"Activity berhasil diambil",
		activityVm,
	))
}
