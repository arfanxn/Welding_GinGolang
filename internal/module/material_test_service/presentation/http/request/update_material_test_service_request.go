package request

import (
	"github.com/arfanxn/welding/internal/infrastructure/http/request"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

var _ request.Request = (*UpdateMaterialTestService)(nil)

type UpdateMaterialTestService struct {
	Id          string   `form:"id" json:"id"`
	MachineId   *string  `form:"machine_id" json:"machine_id"`
	MethodId    *string  `form:"method_id" json:"method_id"`
	ServiceType *string  `form:"service_type" json:"service_type"`
	ServiceCode *string  `form:"service_code" json:"service_code"`
	Unit        *string  `form:"unit" json:"unit"`
	Price       *float64 `form:"price" json:"price"`
}

func NewUpdateMaterialTestService() *UpdateMaterialTestService {
	return &UpdateMaterialTestService{}
}

func (s *UpdateMaterialTestService) Validate() error {
	return validation.ValidateStruct(s,
		validation.Field(&s.Id,
			validation.Required,
			validation.Length(26, 26).Error("Id harus 26 karakter"),
		),
		validation.Field(&s.MachineId,
			validation.Length(26, 26).Error("Machine Id harus 26 karakter"),
		),
		validation.Field(&s.MethodId,
			validation.Length(26, 26).Error("Method Id harus 26 karakter"),
		),
		validation.Field(&s.ServiceType,
			validation.Length(3, 50).Error("Service Type harus di antara 3 dan 50 karakter"),
		),
		validation.Field(&s.ServiceCode,
			validation.Length(3, 50).Error("Service Code harus di antara 3 dan 50 karakter"),
		),
		validation.Field(&s.Unit,
			validation.Length(3, 50).Error("Unit harus di antara 3 dan 50 karakter"),
		),
		validation.Field(&s.Price,
			validation.Min(0.00).Error("Price harus lebih dari 0"),
		),
	)
}
