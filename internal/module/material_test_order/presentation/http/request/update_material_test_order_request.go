package request

import (
	"fmt"

	"github.com/arfanxn/welding/internal/infrastructure/http/request"
	materialTestOrderEnum "github.com/arfanxn/welding/internal/module/material_test_order/domain/enum"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

var _ request.Request = (*UpdateMaterialTestOrder)(nil)
var _ request.Request = (*UpdateMaterialTestOrderService)(nil)

type UpdateMaterialTestOrder struct {
	Id string `form:"id" json:"id"`

	WorkCategoryId *string `form:"work_category_id" json:"work_category_id"`

	WorkPackageName      *string `form:"work_package_name" json:"work_package_name"`
	ApplicantName        *string `form:"applicant_name" json:"applicant_name"`
	ApplicantPhoneNumber *string `form:"applicant_phone_number" json:"applicant_phone_number"`
	ApplicantEmail       *string `form:"applicant_email" json:"applicant_email"`
	ApplicantFullAddress *string `form:"applicant_full_address" json:"applicant_full_address"`
	ApplicantNote        *string `form:"applicant_note" json:"applicant_note"`
	RecipientName        *string `form:"recipient_name" json:"recipient_name"`
	TesterNote           *string `form:"tester_note" json:"tester_note"`

	Status *string `form:"status" json:"status"`

	OwnerUserIds []string `form:"owner_user_ids" json:"owner_user_ids"`

	OrderedServices []UpdateMaterialTestOrderService `form:"ordered_services" json:"ordered_services"`
}

func NewUpdateMaterialTestOrder() *UpdateMaterialTestOrder {
	return &UpdateMaterialTestOrder{}
}

func (r *UpdateMaterialTestOrder) Validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.Id,
			validation.Required.Error("Id tidak boleh kosong"),
			validation.Length(26, 26).Error("Id harus 26 karakter"),
		),
		validation.Field(&r.WorkCategoryId,
			validation.Length(26, 26).Error("Id kategori pekerjaan harus 26 karakter"),
		),
		validation.Field(&r.WorkPackageName,
			validation.Length(3, 255).Error("Nama paket pekerjaan harus di antara 3 dan 255 karakter"),
		),
		validation.Field(&r.ApplicantName,
			validation.Length(3, 255).Error("Nama pemohon harus di antara 3 dan 255 karakter"),
		),
		validation.Field(&r.ApplicantPhoneNumber,
			validation.Length(10, 15).Error("Nomor telpon pemohon harus antara 10-15 karakter"),
		),
		validation.Field(&r.ApplicantEmail,
			validation.Length(3, 50).Error("Panjang email pemohon harus antara 3-50 karakter"),
			is.Email.Error("Format email pemohon tidak valid"),
		),
		validation.Field(&r.ApplicantFullAddress,
			validation.Length(3, 512).Error("Alamat pemohon harus di antara 3 dan 512 karakter"),
		),
		validation.Field(&r.ApplicantNote),
		validation.Field(&r.RecipientName,
			validation.Length(3, 255).Error("Nama penerima harus di antara 3 dan 255 karakter"),
		),
		validation.Field(&r.TesterNote),
		validation.Field(&r.Status,
			validation.In(
				materialTestOrderEnum.MaterialTestOrderStatusDraft, materialTestOrderEnum.MaterialTestOrderStatusAwaitingReview).
				Error(fmt.Sprintf("Status tidak valid. available: '%s', '%s'",
					materialTestOrderEnum.MaterialTestOrderStatusDraft,
					materialTestOrderEnum.MaterialTestOrderStatusAwaitingReview),
				),
		),

		validation.Field(&r.OwnerUserIds,
			validation.Each(
				validation.Length(26, 26).Error("Owner user id harus 26 karakter"),
			),
		),

		validation.Field(
			&r.OrderedServices,
			validation.Each(
				validation.By(func(value any) error {
					orderService := value.(UpdateMaterialTestOrderService)
					return orderService.Validate()
				}),
			),
		),
	)
}

type UpdateMaterialTestOrderService struct {
	ServiceId  string `form:"service_id" json:"service_id"`
	SampleName string `form:"sample_name" json:"sample_name"`
	Quantity   int    `form:"quantity" json:"quantity"`
}

func (r *UpdateMaterialTestOrderService) Validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.ServiceId,
			validation.Required.Error("Service id wajib disi"),
			validation.Length(26, 26).Error("Service id harus 26 karakter"),
		),
		validation.Field(&r.SampleName,
			validation.Required.Error("Sample name wajib disi"),
			validation.Length(3, 255).Error("Sample name harus di antara 3 dan 255 karakter"),
		),
		validation.Field(&r.Quantity,
			validation.Required.Error("Quantity wajib disi"),
			validation.Min(1).Error("Quantity minimal 1"),
		),
	)
}
