package request

import (
	"mime/multipart"

	"github.com/arfanxn/welding/internal/infrastructure/http/request"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

var _ request.Request = (*RefundMaterialTestOrder)(nil)

type RefundMaterialTestOrder struct {
	Id   string                `form:"id" json:"id"`
	File *multipart.FileHeader `form:"file" json:"file"`
}

func NewRefundMaterialTestOrder() *RefundMaterialTestOrder {
	return &RefundMaterialTestOrder{}
}

func (r *RefundMaterialTestOrder) Validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.Id,
			validation.Required.Error("Id tidak boleh kosong"),
			validation.Length(26, 26).Error("Id harus 26 karakter"),
		),
		validation.Field(&r.File,
			validation.Required.Error("Bukti pembayaran wajib disi"),
		),
	)
}
