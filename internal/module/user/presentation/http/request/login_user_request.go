package request

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	is "github.com/go-ozzo/ozzo-validation/v4/is"
)

type LoginUser struct {
	Email    string `form:"email" json:"email"`
	Password string `form:"password" json:"password"`
}

func (r *LoginUser) Validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.Email,
			validation.Required.Error("Email wajib diisi"),
			validation.Length(3, 64).Error("Panjang email harus antara 3-64 karakter"),
			is.EmailFormat.Error("Format email tidak valid"),
		),
		validation.Field(&r.Password,
			validation.Required.Error("Kata sandi wajib diisi"),
			validation.Length(8, 255).Error("Panjang kata sandi minimal 8 karakter"),
		),
	)
}
