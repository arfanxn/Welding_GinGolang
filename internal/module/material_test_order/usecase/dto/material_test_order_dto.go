package dto

import (
	"mime/multipart"

	materialTestOrderEnum "github.com/arfanxn/welding/internal/module/material_test_order/domain/enum"
)

// SaveMaterialTestOrder represents the data transfer object for creating or updating a material test order
type SaveMaterialTestOrder struct {
	Id             *string `json:"id"`
	WorkCategoryId *string `json:"work_category_id"`

	WorkPackageName *string `json:"work_package_name"`

	ApplicantName        *string `json:"applicant_name"`
	ApplicantPhoneNumber *string `json:"applicant_phone_number"`
	ApplicantEmail       *string `json:"applicant_email"`
	ApplicantFullAddress *string `json:"applicant_full_address"`
	ApplicantNote        *string `json:"applicant_note"`

	RecipientName *string `json:"recipient_name"`

	TesterNote *string `json:"tester_note"`

	Status *materialTestOrderEnum.MaterialTestOrderStatus `json:"status"`

	OwnerUserIds []string `json:"owner_user_ids"`

	OrderedServices []SaveMaterialTestOrderService `json:"ordered_services"`
}

// DestroyMaterialTestOrder represents the data transfer object for deleting a material test order
type DestroyMaterialTestOrder struct {
	Id string `json:"id"`
}

// SaveMaterialTestOrderService represents the data transfer object for a service included in a material test order
type SaveMaterialTestOrderService struct {
	ServiceId  string `json:"service_id"`
	SampleName string `json:"sample_name"`
	Quantity   int    `json:"quantity"`
}

type UpdateMaterialTestOrderServiceEvaluation struct {
	Id                       string `json:"id"`                          // the order id
	OrderServiceEvaluationId string `json:"order_service_evaluation_id"` // the ordered service's evaluation id

	IsEquipmentAvailable      bool `json:"is_equipment_available"`
	IsPersonnelAvailable      bool `json:"is_personnel_available"`
	IsTimeAvailable           bool `json:"is_time_available"`
	IsTestReady               bool `json:"is_test_ready"`
	IsSubcontractLabAvailable bool `json:"is_subcontract_lab_available"`
}

/* ==============================
	Status related DTOs
============================== */

// ApproveMaterialTestOrder represents the data transfer object for approving a material test order
type ApproveMaterialTestOrder struct {
	Id string `json:"id"`
}

// RejectMaterialTestOrder represents the data transfer object for rejecting a material test order
type RejectMaterialTestOrder struct {
	Id string `json:"id"`
}

// SubmitPaymentMaterialTestOrder represents the data transfer object for submitting payment for a material test order
type SubmitPaymentMaterialTestOrder struct {
	Id   string                `json:"id"`
	File *multipart.FileHeader `json:"file"` // payment proof
}

// ApprovePaymentMaterialTestOrder represents the data transfer object for approving a payment for a material test order
type ApprovePaymentMaterialTestOrder struct {
	Id string `json:"id"`
}

// RejectPaymentMaterialTestOrder represents the data transfer object for rejecting a payment for a material test order
type RejectPaymentMaterialTestOrder struct {
	Id string `json:"id"`
}

// CancelMaterialTestOrder represents the data transfer object for canceling a material test order
type CancelMaterialTestOrder struct {
	Id string `json:"id"`
}

// TestMaterialTestOrder represents the data transfer object for testing a material test order
type TestMaterialTestOrder struct {
	Id string `json:"id"`
}

// RefundMaterialTestOrder represents the data transfer object for processing a refund for a material test order
type RefundMaterialTestOrder struct {
	Id   string                `json:"id"`
	File *multipart.FileHeader `json:"file"` // refund proof
}

// CompleteMaterialTestOrder represents the data transfer object for marking a material test order as completed
type CompleteMaterialTestOrder struct {
	Id string `json:"id"`
}

/* ==============================
	Media related DTOs
============================== */

// SaveMaterialTestOrderMedia represents the data transfer object for adding media to a material test order
type SaveMaterialTestOrderMedia struct {
	Id      *string               `json:"id"`
	MediaId *string               `json:"media_id"`
	Name    *string               `json:"name"`
	File    *multipart.FileHeader `json:"file"`
}

// DestroyMaterialTestOrderMedia represents the data transfer object for removing media from a material test order
type DestroyMaterialTestOrderMedia struct {
	Id      string `json:"id"`
	MediaId string `json:"media_id"`
}
