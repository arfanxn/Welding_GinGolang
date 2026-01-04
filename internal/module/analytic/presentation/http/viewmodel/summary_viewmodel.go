package viewmodel

type SummaryViewModel struct {
	// User counts
	UserCount                 int64 `json:"user_count"`
	ActiveUserCount           int64 `json:"active_user_count"`
	DeactiveUserCount         int64 `json:"deactive_user_count"`
	CustomerUserCount         int64 `json:"customer_user_count"`
	ActiveCustomerUserCount   int64 `json:"active_customer_user_count"`
	DeactiveCustomerUserCount int64 `json:"deactive_customer_user_count"`
	EmployeeUserCount         int64 `json:"employee_user_count"`
	ActiveEmployeeUserCount   int64 `json:"active_employee_user_count"`
	DeactiveEmployeeUserCount int64 `json:"deactive_employee_user_count"`

	// Order counts
	OrderCount                 int64 `json:"order_count"`
	DraftOrderCount            int64 `json:"draft_order_count"`
	AwaitingReviewOrderCount   int64 `json:"awaiting_review_order_count"`
	RejectedOrderCount         int64 `json:"rejected_order_count"`
	AwaitingPaymentOrderCount  int64 `json:"awaiting_payment_order_count"`
	CancelledOrderCount        int64 `json:"cancelled_order_count"`
	PaymentSubmittedOrderCount int64 `json:"payment_submitted_order_count"`
	PaymentRejectedOrderCount  int64 `json:"payment_rejected_order_count"`
	PaymentApprovedOrderCount  int64 `json:"payment_approved_order_count"`
	TestingOrderCount          int64 `json:"testing_order_count"`
	RefundedOrderCount         int64 `json:"refunded_order_count"`
	CompletedOrderCount        int64 `json:"completed_order_count"`

	// Order sum
	CompletedOrderSum float64 `json:"completed_order_sum"` // Total sum of completed orders

	// Service count
	ServiceCount int64 `json:"service_count"`
}
