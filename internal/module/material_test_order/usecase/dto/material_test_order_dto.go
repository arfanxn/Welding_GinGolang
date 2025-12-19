package dto

import (
	"mime/multipart"

	materialTestOrderEnum "github.com/arfanxn/welding/internal/module/material_test_order/domain/enum"
)

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

type DestroyMaterialTestOrder struct {
	Id string `json:"id"`
}

type SaveMaterialTestOrderService struct {
	ServiceId  string `json:"service_id"`
	SampleName string `json:"sample_name"`
	Quantity   int    `json:"quantity"`
}

type SaveMaterialTestOrderMedia struct {
	OrderId *string               `json:"order_id"`
	MediaId *string               `json:"media_id"`
	Name    *string               `json:"name"`
	File    *multipart.FileHeader `json:"file"`
}

type DestroyMaterialTestOrderMedia struct {
	MediaId string `json:"media_id"`
}
