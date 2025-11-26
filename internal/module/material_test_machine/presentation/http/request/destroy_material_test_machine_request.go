package request

import validation "github.com/go-ozzo/ozzo-validation/v4"

type DestroyMaterialTestMachine struct {
	Id string `form:"id" json:"id"`
}

func NewDestroyMaterialTestMachine() *DestroyMaterialTestMachine {
	return &DestroyMaterialTestMachine{}
}

func (s *DestroyMaterialTestMachine) Validate() error {
	return validation.ValidateStruct(s,
		validation.Field(&s.Id,
			validation.Required,
			validation.Length(26, 26).Error("Id harus 26 karakter"),
		),
	)
}
