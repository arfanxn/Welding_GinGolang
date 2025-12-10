package request

import (
	"github.com/arfanxn/welding/internal/infrastructure/http/request"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

var _ request.Request = (*StoreAddressRequest)(nil)

type StoreAddressRequest struct {
	FullAddress string `form:"full_address" json:"full_address"`
}

func NewStoreAddressRequest() *StoreAddressRequest {
	return &StoreAddressRequest{}
}

func (s *StoreAddressRequest) Validate() error {
	return validation.ValidateStruct(s,
		validation.Field(&s.FullAddress,
			validation.Required.Error("Alamat tidak boleh kosong"),
			validation.Length(3, 512).Error("Alamat harus di antara 3 dan 512 karakter"),
		),
	)
}
