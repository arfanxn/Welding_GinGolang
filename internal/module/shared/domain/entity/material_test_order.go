package entity

import (
	"time"

	mtoEnum "github.com/arfanxn/welding/internal/module/material_test_order/domain/enum"
)

type MaterialTestOrder struct {
	// ---- IDENTIFICATION ----
	Id             string `json:"id" gorm:"primaryKey"`
	Number         string `json:"number"`
	WorkCategoryId string `json:"work_category_id"`

	// ---- WORK PACKAGE & APPLICANT DETAILS ----
	WorkPackageName      string  `json:"work_package_name"`
	ApplicantName        string  `json:"applicant_name"`
	ApplicantPhoneNumber string  `json:"applicant_phone_number"`
	ApplicantEmail       string  `json:"applicant_email"`
	ApplicantFullAddress string  `json:"applicant_full_address"`
	ApplicantNote        *string `json:"applicant_note"`
	RecipientName        string  `json:"recipient_name"`
	TesterNote           *string `json:"tester_note"`

	// ---- FINANCIAL METADATA ----
	SubTotal float64 `json:"sub_total"`
	Tax      float64 `json:"tax"`
	Discount float64 `json:"discount"`
	Total    float64 `json:"total"`

	// ---- WORKFLOW STATUS ----
	Status string `json:"status"`

	// ---- TIMESTAMP MILESTONES ----
	// Payment lifecycle
	PaymentSubmittedAt *time.Time `json:"payment_submitted_at"`
	PaymentRejectedAt  *time.Time `json:"payment_rejected_at"`
	PaymentApprovedAt  *time.Time `json:"payment_approved_at"`

	// Testing lifecycle
	TestingAt   *time.Time `json:"testing_at"`
	CompletedAt *time.Time `json:"completed_at"`

	// End-of-flow outcomes
	CancelledAt *time.Time `json:"cancelled_at"`
	RejectedAt  *time.Time `json:"rejected_at"`
	RefundedAt  *time.Time `json:"refunded_at"`

	// ---- SYSTEM METADATA ----
	CreatedAt time.Time  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	// ---- RELATIONS ----
	OrderUsers      []*MaterialTestOrderUser    `json:"order_users,omitempty" gorm:"foreignKey:OrderId;references:Id"`
	WorkCategory    MaterialTestWorkCategory    `json:"work_category" gorm:"foreignKey:WorkCategoryId;references:Id"`
	OrderedServices []*MaterialTestOrderService `json:"ordered_services,omitempty" gorm:"foreignKey:OrderId;references:Id"`
	Medias          []*Media                    `json:"medias,omitempty" gorm:"foreignKey:ModelId;references:Id"`
}

func NewMaterialTestOrder() *MaterialTestOrder {
	return &MaterialTestOrder{}
}

func (MaterialTestOrder) TableName() string {
	return "material_test_orders"
}

func (mto *MaterialTestOrder) IsDraft() bool {
	return mto.Status == mtoEnum.MaterialTestOrderStatusDraft
}

func (mto *MaterialTestOrder) IsAwaitingReview() bool {
	return mto.Status == mtoEnum.MaterialTestOrderStatusAwaitingReview
}
