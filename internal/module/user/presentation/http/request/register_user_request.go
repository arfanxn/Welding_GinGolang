package request

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	is "github.com/go-ozzo/ozzo-validation/v4/is"
)

type RegisterUserRequest struct {
	Name                 string `json:"name"`
	Email                string `json:"email"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"password_confirmation"`
}

func (r *RegisterUserRequest) Validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.Name,
			validation.Required.Error("Nama wajib diisi"),
			validation.Length(2, 64).Error("Panjang nama harus antara 2-64 karakter"),
		),
		validation.Field(&r.Email,
			validation.Required.Error("Email wajib diisi"),
			validation.Length(3, 64).Error("Panjang email harus antara 3-64 karakter"),
			is.EmailFormat.Error("Format email tidak valid"),
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
