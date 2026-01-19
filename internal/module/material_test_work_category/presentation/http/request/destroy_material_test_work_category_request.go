package request

import validation "github.com/go-ozzo/ozzo-validation/v4"

type DestroyMaterialTestWorkCategory struct {
	Id string `form:"id" json:"id"`
}

func NewDestroyMaterialTestWorkCategory() *DestroyMaterialTestWorkCategory {
	return &DestroyMaterialTestWorkCategory{}
}

func (s *DestroyMaterialTestWorkCategory) Validate() error {
	return validation.ValidateStruct(s,
		validation.Field(&s.Id,
			validation.Required,
			validation.Length(26, 26).Error("Id harus 26 karakter"),
		),
	)
}
