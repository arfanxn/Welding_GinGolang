package request

import (
	"github.com/arfanxn/welding/internal/infrastructure/http/request"
	"github.com/creasty/defaults"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

var _ request.Request = (*UpdateMaterialTestOrder)(nil)

type UpdateMaterialTestOrder struct {
	Id string `form:"id" json:"id"`
}

func NewUpdateMaterialTestOrder() *UpdateMaterialTestOrder {
	defaults.Set(&UpdateMaterialTestOrder{})
	return &UpdateMaterialTestOrder{}
}

func (s *UpdateMaterialTestOrder) Validate() error {
	return validation.ValidateStruct(s,
		validation.Field(&s.Id,
			validation.Required.Error("Id tidak boleh kosong"),
			validation.Length(26, 26).Error("Id harus 26 karakter"),
		),
	)
}
