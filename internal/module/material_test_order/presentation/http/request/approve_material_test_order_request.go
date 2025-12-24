package request

import (
	"github.com/arfanxn/welding/internal/infrastructure/http/request"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

var _ request.Request = (*ApproveMaterialTestOrder)(nil)

type ApproveMaterialTestOrder struct {
	Id string `form:"id" json:"id"`
}

func NewApproveMaterialTestOrder() *ApproveMaterialTestOrder {
	return &ApproveMaterialTestOrder{}
}

func (r *ApproveMaterialTestOrder) Validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.Id,
			validation.Required.Error("Id tidak boleh kosong"),
			validation.Length(26, 26).Error("Id harus 26 karakter"),
		),
	)
}
