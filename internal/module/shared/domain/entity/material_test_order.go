package entity

import "time"

type MaterialTestOrder struct {
	// ---- IDENTIFICATION ----
	Id             string  `json:"id" gorm:"primaryKey"`
	Number         string  `json:"number"`
	WorkCategoryId string  `json:"work_category_id"`
	WorkPackageId  string  `json:"work_package_id"`
	CustomerId     string  `json:"customer_id"`
	CustomerNote   *string `json:"customer_note"`
	TesterNote     *string `json:"tester_note"`
	IssuedFor      string  `json:"issued_for"`

	// ---- FINANCIAL METADATA ----
	SubTotal float64 `json:"sub_total"`
	Tax      float64 `json:"tax"`
	Discount float64 `json:"discount"`
	Total    float64 `json:"total"`

	// ---- TIMELINE ----
	EnteredAt time.Time `json:"entered_at"`

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
	WorkCategory              MaterialTestWorkCategory    `json:"work_category" gorm:"foreignKey:WorkCategoryId;references:Id"`
	WorkPackage               MaterialTestWorkPackage     `json:"work_package" gorm:"foreignKey:WorkPackageId;references:Id"`
	Customer                  Customer                    `json:"customer" gorm:"foreignKey:CustomerId;references:Id"`
	MaterialTestOrderServices []*MaterialTestOrderService `json:"material_test_order_services,omitempty" gorm:"foreignKey:OrderId;references:Id"`
	Medias                    []*Media                    `json:"medias,omitempty" gorm:"foreignKey:ModelId;references:Id"`
}

func NewMaterialTestOrder() *MaterialTestOrder {
	return &MaterialTestOrder{}
}

func (MaterialTestOrder) TableName() string {
	return "material_test_orders"
}
