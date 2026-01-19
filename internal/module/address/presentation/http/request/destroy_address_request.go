package request

import (
	"github.com/arfanxn/welding/internal/infrastructure/http/request"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

var _ request.Request = (*DestroyAddressRequest)(nil)

type DestroyAddressRequest struct {
	Id string `form:"id" json:"id"`
}

func NewDestroyAddressRequest() *DestroyAddressRequest {
	return &DestroyAddressRequest{}
}

func (s *DestroyAddressRequest) Validate() error {
	return validation.ValidateStruct(s,
		validation.Field(&s.Id,
			validation.Required,
			validation.Length(26, 26).Error("Id harus 26 karakter"),
		),
	)
}
