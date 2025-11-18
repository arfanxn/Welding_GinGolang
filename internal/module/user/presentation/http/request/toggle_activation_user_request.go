package request

import validation "github.com/go-ozzo/ozzo-validation/v4"

type ToggleActivationUser struct {
	Id string `form:"id" json:"id"`
}

func (r *ToggleActivationUser) Validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.Id,
			validation.Required.Error("Id tidak boleh kosong"),
			validation.Length(26, 26).Error("Id harus 26 karakter"),
		),
	)
}
