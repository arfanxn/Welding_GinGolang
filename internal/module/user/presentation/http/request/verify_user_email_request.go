package request

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type VerifyEmail struct {
	Email string `form:"email" json:"email"`
	Code  string `form:"code" json:"code"`
}

func (r *VerifyEmail) Validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.Email,
			validation.Required.Error("Email wajib diisi"),
			validation.Length(3, 50).Error("Panjang email harus antara 3-50 karakter"),
			is.Email.Error("Format email tidak valid"),
		),
		validation.Field(&r.Code,
			validation.Required.Error("Kode verifikasi wajib diisi"),
			validation.Length(6, 6).Error("Panjang kode verifikasi harus 6 karakter"),
		),
	)
}
