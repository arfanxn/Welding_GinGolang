package request

import validation "github.com/go-ozzo/ozzo-validation/v4"

type DestroyMaterialTestService struct {
	Id string `form:"id" json:"id"`
}

func NewDestroyMaterialTestService() *DestroyMaterialTestService {
	return &DestroyMaterialTestService{}
}

func (s *DestroyMaterialTestService) Validate() error {
	return validation.ValidateStruct(s,
		validation.Field(&s.Id,
			validation.Required,
			validation.Length(26, 26).Error("Id harus 26 karakter"),
		),
	)
}
