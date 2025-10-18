package request

import validation "github.com/go-ozzo/ozzo-validation/v4"

type DestroyRole struct {
	Id string `form:"id" json:"id"`
}

func NewDestroyRole() *DestroyRole {
	return &DestroyRole{}
}

func (s *DestroyRole) Validate() error {
	return validation.ValidateStruct(s,
		validation.Field(&s.Id,
			validation.Required,
			validation.Length(26, 26).Error("Id harus 26 karakter"),
		),
	)
}
