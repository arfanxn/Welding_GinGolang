package request

import (
	"github.com/arfanxn/welding/internal/infrastructure/http/request"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

var _ request.Request = (*RejectPaymentMaterialTestOrder)(nil)

type RejectPaymentMaterialTestOrder struct {
	Id string `form:"id" json:"id"`
}

func NewRejectPaymentMaterialTestOrder() *RejectPaymentMaterialTestOrder {
	return &RejectPaymentMaterialTestOrder{}
}

func (r *RejectPaymentMaterialTestOrder) Validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.Id,
			validation.Required.Error("Id tidak boleh kosong"),
			validation.Length(26, 26).Error("Id harus 26 karakter"),
		),
	)
}
