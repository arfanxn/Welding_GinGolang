package request

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type ResetPassword struct {
	Email                string `form:"email" json:"email"`
	Code                 string `form:"code" json:"code"`
	Password             string `form:"password" json:"password"`
	PasswordConfirmation string `form:"password_confirmation" json:"password_confirmation"`
}

func (r *ResetPassword) Validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.Email,
			validation.Required.Error("Email wajib diisi"),
			validation.Length(3, 50).Error("Panjang email harus antara 3-50 karakter"),
			is.Email.Error("Format email tidak valid"),
		),
		validation.Field(&r.Code,
			validation.Required.Error("Kode reset wajib diisi"),
			validation.Length(6, 6).Error("Panjang kode reset harus 6 karakter"),
		),
		validation.Field(&r.Password,
			validation.Required.Error("Kata sandi wajib diisi"),
			validation.Length(8, 255).Error("Panjang kata sandi minimal 8 karakter"),
		),
		validation.Field(&r.PasswordConfirmation,
			validation.Required.Error("Konfirmasi kata sandi wajib diisi"),
			validation.By(func(value any) error {
				if value.(string) != r.Password {
					return validation.NewError("password_mismatch", "Kata sandi tidak cocok")
				}
				return nil
			}),
		),
	)
}
