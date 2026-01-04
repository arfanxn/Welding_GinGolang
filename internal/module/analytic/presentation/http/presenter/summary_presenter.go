package presenter

import (
	"context"

	analyticEntity "github.com/arfanxn/welding/internal/module/analytic/domain/entity"
	analyticViewModel "github.com/arfanxn/welding/internal/module/analytic/presentation/http/viewmodel"
	"go.uber.org/fx"
)

type SummaryPresenter interface {
	FromEntityToViewModel(ctx context.Context, summary *analyticEntity.Summary) (*analyticViewModel.SummaryViewModel, error)
}

type summaryPresenter struct {
}

type NewSummaryPresenterParams struct {
	fx.In
}

func NewSummaryPresenter(params NewSummaryPresenterParams) SummaryPresenter {
	return &summaryPresenter{}
}

func (p *summaryPresenter) FromEntityToViewModel(ctx context.Context, summary *analyticEntity.Summary) (summaryViewModel *analyticViewModel.SummaryViewModel, err error) {
	summaryViewModel = &analyticViewModel.SummaryViewModel{}

	// User counts
	summaryViewModel.UserCount = summary.UserCount
	summaryViewModel.ActiveUserCount = summary.ActiveUserCount
	summaryViewModel.DeactiveUserCount = summary.DeactiveUserCount
	summaryViewModel.CustomerUserCount = summary.CustomerUserCount
	summaryViewModel.ActiveCustomerUserCount = summary.ActiveCustomerUserCount
	summaryViewModel.DeactiveCustomerUserCount = summary.DeactiveCustomerUserCount
	summaryViewModel.EmployeeUserCount = summary.EmployeeUserCount
	summaryViewModel.ActiveEmployeeUserCount = summary.ActiveEmployeeUserCount
	summaryViewModel.DeactiveEmployeeUserCount = summary.DeactiveEmployeeUserCount

	// Order counts
	summaryViewModel.OrderCount = summary.OrderCount
	summaryViewModel.DraftOrderCount = summary.DraftOrderCount
	summaryViewModel.AwaitingReviewOrderCount = summary.AwaitingReviewOrderCount
	summaryViewModel.RejectedOrderCount = summary.RejectedOrderCount
	summaryViewModel.AwaitingPaymentOrderCount = summary.AwaitingPaymentOrderCount
	summaryViewModel.CancelledOrderCount = summary.CancelledOrderCount
	summaryViewModel.PaymentSubmittedOrderCount = summary.PaymentSubmittedOrderCount
	summaryViewModel.PaymentRejectedOrderCount = summary.PaymentRejectedOrderCount
	summaryViewModel.PaymentApprovedOrderCount = summary.PaymentApprovedOrderCount
	summaryViewModel.TestingOrderCount = summary.TestingOrderCount
	summaryViewModel.RefundedOrderCount = summary.RefundedOrderCount
	summaryViewModel.CompletedOrderCount = summary.CompletedOrderCount

	// Order sum
	summaryViewModel.CompletedOrderSum = summary.CompletedOrderSum // Total sum of completed orders

	// Service count
	summaryViewModel.ServiceCount = summary.ServiceCount

	return summaryViewModel, nil
}
