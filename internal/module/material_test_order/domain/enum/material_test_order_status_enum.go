package enum

// MaterialTestOrderStatus represents the workflow state of a material test order.
type MaterialTestOrderStatus = string

const (
	// draft — Order is created but not yet submitted for review.
	MaterialTestOrderStatusDraft MaterialTestOrderStatus = "draft"

	// awaiting_review — Order has been submitted and is waiting for admin validation.
	MaterialTestOrderStatusAwaitingReview MaterialTestOrderStatus = "awaiting_review"

	// awaiting_payment — Order has passed review; invoice is issued and waiting for customer payment.
	MaterialTestOrderStatusAwaitingPayment MaterialTestOrderStatus = "awaiting_payment"

	// payment_submitted — Customer has submitted payment proof; waiting for admin verification.
	MaterialTestOrderStatusPaymentSubmitted MaterialTestOrderStatus = "payment_submitted"

	// payment_rejected — Payment proof was reviewed but rejected; customer must re-upload or correct it.
	MaterialTestOrderStatusPaymentRejected MaterialTestOrderStatus = "payment_rejected"

	// payment_approved — Payment has been verified by admin; order is ready to proceed to testing.
	MaterialTestOrderStatusPaymentApproved MaterialTestOrderStatus = "payment_approved"

	// testing — Testing process has started; sample is actively being tested.
	MaterialTestOrderStatusTesting MaterialTestOrderStatus = "testing"

	// completed — Testing process has finished; results are ready or being prepared.
	MaterialTestOrderStatusCompleted MaterialTestOrderStatus = "completed"

	// cancelled — Order was cancelled by customer or admin; no further processing will occur.
	MaterialTestOrderStatusCancelled MaterialTestOrderStatus = "cancelled"

	// rejected — Order was rejected during review (not related to payment).
	MaterialTestOrderStatusRejected MaterialTestOrderStatus = "rejected"

	// refunded — Order cannot continue due to lab failure; customer payment has been refunded.
	MaterialTestOrderStatusRefunded MaterialTestOrderStatus = "refunded"
)

var MaterialTestOrderStatuses = []MaterialTestOrderStatus{
	MaterialTestOrderStatusDraft,            // 1
	MaterialTestOrderStatusAwaitingReview,   // 2
	MaterialTestOrderStatusAwaitingPayment,  // 3
	MaterialTestOrderStatusPaymentSubmitted, // 4
	MaterialTestOrderStatusPaymentRejected,  // 5
	MaterialTestOrderStatusPaymentApproved,  // 6
	MaterialTestOrderStatusTesting,          // 7
	MaterialTestOrderStatusCompleted,        // 8
	MaterialTestOrderStatusCancelled,        // 9
	MaterialTestOrderStatusRejected,         // 10
	MaterialTestOrderStatusRefunded,         // 11
}
