package http

import (
	"net/http"

	"github.com/arfanxn/welding/internal/infrastructure/http/response"
	analyticPresenter "github.com/arfanxn/welding/internal/module/analytic/presentation/http/presenter"
	analyticUsecase "github.com/arfanxn/welding/internal/module/analytic/usecase"
	"github.com/arfanxn/welding/pkg/query"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type AnalyticHandler interface {
	Summary(c *gin.Context)
	OrderTrends(c *gin.Context)
}

type analyticHandler struct {
	analyticUsecase  analyticUsecase.AnalyticUsecase
	summaryPresenter analyticPresenter.SummaryPresenter
}

type NewAnalyticHandlerParams struct {
	fx.In

	AnalyticUsecase  analyticUsecase.AnalyticUsecase
	SummaryPresenter analyticPresenter.SummaryPresenter
}

func NewAnalyticHandler(params NewAnalyticHandlerParams) AnalyticHandler {
	return &analyticHandler{
		analyticUsecase:  params.AnalyticUsecase,
		summaryPresenter: params.SummaryPresenter,
	}
}

func (h *analyticHandler) Summary(c *gin.Context) {
	q := query.NewQuery()
	c.ShouldBind(q)

	summary, err := h.analyticUsecase.GetSummary(q)
	if err != nil {
		panic(err)
	}

	summaryViewModel, err := h.summaryPresenter.FromEntityToViewModel(c.Request.Context(), summary)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, response.NewBodyWithData(http.StatusOK, "Summary berhasil diambil", gin.H{"summary": summaryViewModel}))
}

func (h *analyticHandler) OrderTrends(c *gin.Context) {
	q := query.NewQuery()
	c.ShouldBind(q)

	orderTrends, err := h.analyticUsecase.GetOrderTrends(q)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, response.NewBodyWithData(http.StatusOK, "Order trends berhasil diambil", gin.H{"order_trends": orderTrends}))
}
