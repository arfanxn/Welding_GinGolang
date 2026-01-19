package request

import validation "github.com/go-ozzo/ozzo-validation/v4"

type DestroyMaterialTestWorkPackage struct {
	Id string `form:"id" json:"id"`
}

func NewDestroyMaterialTestWorkPackage() *DestroyMaterialTestWorkPackage {
	return &DestroyMaterialTestWorkPackage{}
}

func (s *DestroyMaterialTestWorkPackage) Validate() error {
	return validation.ValidateStruct(s,
		validation.Field(&s.Id,
			validation.Required,
			validation.Length(26, 26).Error("Id harus 26 karakter"),
		),
	)
}
