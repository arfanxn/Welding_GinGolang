package request

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type CreateUserResetPassword struct {
	Email string `form:"email" json:"email"`
}

func (r *CreateUserResetPassword) Validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.Email,
			validation.Required.Error("Email wajib diisi"),
			validation.Length(3, 50).Error("Panjang email harus antara 3-50 karakter"),
			is.EmailFormat.Error("Format email tidak valid"),
		),
	)
}
