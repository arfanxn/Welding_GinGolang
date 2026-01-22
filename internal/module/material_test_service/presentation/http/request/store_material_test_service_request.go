package request

import (
	"strings"

	"github.com/arfanxn/welding/internal/infrastructure/http/request"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

var _ request.Request = (*StoreMaterialTestService)(nil)

type StoreMaterialTestService struct {
	MachineId   *string `form:"machine_id" json:"machine_id"`
	MethodId    *string `form:"method_id" json:"method_id"`
	TestName    string  `form:"test_name" json:"test_name"`
	ServiceType string  `form:"service_type" json:"service_type"`
	ServiceCode string  `form:"service_code" json:"service_code"`
	Unit        string  `form:"unit" json:"unit"`
	Price       float64 `form:"price" json:"price"`
}

func NewStoreMaterialTestService() *StoreMaterialTestService {
	return &StoreMaterialTestService{}
}

func (s *StoreMaterialTestService) Validate() error {
	return validation.ValidateStruct(s,
		validation.Field(&s.MachineId,
			validation.When(hasValue(s.MachineId),
				validation.Length(26, 26).Error("Machine Id harus 26 karakter"),
			),
		),
		validation.Field(&s.MethodId,
			validation.When(hasValue(s.MethodId),
				validation.Length(26, 26).Error("Method Id harus 26 karakter"),
			),
		),
		validation.Field(&s.TestName,
			validation.Required.Error("Test Name tidak boleh kosong"),
			validation.Length(3, 255).Error("Test Name harus di antara 3 dan 255 karakter"),
		),
		validation.Field(&s.ServiceType,
			validation.Required.Error("Service Type tidak boleh kosong"),
			validation.Length(3, 50).Error("Service Type harus di antara 3 dan 50 karakter"),
		),
		validation.Field(&s.ServiceCode,
			validation.Required.Error("Service Code tidak boleh kosong"),
			validation.Length(3, 50).Error("Service Code harus di antara 3 dan 50 karakter"),
		),
		validation.Field(&s.Unit,
			validation.Required.Error("Unit tidak boleh kosong"),
			validation.Length(3, 50).Error("Unit harus di antara 3 dan 50 karakter"),
		),
		validation.Field(&s.Price,
			validation.Required.Error("Price tidak boleh kosong"),
			validation.Min(0.00).Error("Price harus lebih dari 0"),
		),
	)
}

func hasValue(value *string) bool {
	if value == nil {
		return false
	}
	return strings.TrimSpace(*value) != ""
}
