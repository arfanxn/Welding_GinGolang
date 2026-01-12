package viewmodel

import (
	"time"

	mediaViewmodel "github.com/arfanxn/welding/internal/module/media/presentation/http/viewmodel"
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
)

type MaterialTestOrderViewModel struct {
	// ---- IDENTIFICATION ----
	Id             string `json:"id"`
	Number         string `json:"number"`
	WorkCategoryId string `json:"work_category_id"`

	// ---- WORK PACKAGE & APPLICANT DETAILS ----
	WorkPackageName      string  `json:"work_package_name"`
	ApplicantName        string  `json:"applicant_name"`
	ApplicantCompanyName string  `json:"applicant_company_name"`
	ApplicantPhoneNumber string  `json:"applicant_phone_number"`
	ApplicantEmail       string  `json:"applicant_email"`
	ApplicantFullAddress string  `json:"applicant_full_address"`
	ApplicantNote        *string `json:"applicant_note"`
	RecipientName        string  `json:"recipient_name"`
	RecipientFullAddress string  `json:"recipient_full_address"`
	TesterNote           *string `json:"tester_note"`

	// ---- FINANCIAL METADATA ----
	SubTotal float64 `json:"sub_total"`
	Tax      float64 `json:"tax"`
	Discount float64 `json:"discount"`
	Total    float64 `json:"total"`

	// ---- WORKFLOW STATUS ----
	Status string `json:"status"`

	// ---- TIMESTAMP MILESTONES ----
	// Review lifecycle
	SubmittedAt *time.Time `json:"submitted_at"`
	ApprovedAt  *time.Time `json:"approved_at"`
	RejectedAt  *time.Time `json:"rejected_at"`
	CancelledAt *time.Time `json:"cancelled_at"`

	// Payment lifecycle
	PaymentSubmittedAt *time.Time `json:"payment_submitted_at"`
	PaymentApprovedAt  *time.Time `json:"payment_approved_at"`
	PaymentRejectedAt  *time.Time `json:"payment_rejected_at"`

	// Testing lifecycle
	TestingAt   *time.Time `json:"testing_at"`
	RefundedAt  *time.Time `json:"refunded_at"`
	CompletedAt *time.Time `json:"completed_at"`

	// ---- SYSTEM METADATA ----
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`

	// ---- RELATIONS ----
	WorkCategory    entity.MaterialTestWorkCategory    `json:"work_category,omitzero"`
	OrderUsers      []*entity.MaterialTestOrderUser    `json:"order_users,omitempty"`
	OrderedServices []*entity.MaterialTestOrderService `json:"ordered_services,omitempty"`
	Medias          []*mediaViewmodel.Media            `json:"medias,omitempty"`
}
