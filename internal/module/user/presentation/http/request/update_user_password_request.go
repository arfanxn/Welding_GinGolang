package request

import validation "github.com/go-ozzo/ozzo-validation/v4"

type UpdateUserPassword struct {
	Id       string `form:"id" json:"id"`
	Password string `form:"password" json:"password"`
}

func (r *UpdateUserPassword) Validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.Id,
			validation.Required.Error("Id wajib diisi"),
		),
		validation.Field(&r.Password,
			validation.Required.Error("Kata sandi wajib diisi"),
			validation.Length(8, 255).Error("Panjang kata sandi minimal 8 karakter"),
		),
	)
}
