package request

import (
	"github.com/arfanxn/welding/internal/infrastructure/http/request"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

var _ request.Request = (*UpdateAddressRequest)(nil)

type UpdateAddressRequest struct {
	Id          string `form:"id" json:"id"`
	FullAddress string `form:"full_address" json:"full_address"`
}

func NewUpdateAddressRequest() *UpdateAddressRequest {
	return &UpdateAddressRequest{}
}

func (s *UpdateAddressRequest) Validate() error {
	return validation.ValidateStruct(s,
		validation.Field(&s.Id,
			validation.Required,
			validation.Length(26, 26).Error("Id harus 26 karakter"),
		),
		validation.Field(&s.FullAddress,
			validation.Length(3, 512).Error("Alamat harus di antara 3 dan 512 karakter"),
		),
	)
}
