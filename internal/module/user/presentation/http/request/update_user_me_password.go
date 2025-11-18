package request

import validation "github.com/go-ozzo/ozzo-validation/v4"

type UpdateUserMePassword struct {
	CurrentPassword      string `form:"current_password" json:"current_password"`
	Password             string `form:"password" json:"password"`
	PasswordConfirmation string `form:"password_confirmation" json:"password_confirmation"`
}

func (r *UpdateUserMePassword) Validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.CurrentPassword,
			validation.Required.Error("Kata sandi saat ini wajib diisi"),
			validation.Length(8, 255).Error("Panjang kata sandi saat ini minimal 8 karakter"),
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
