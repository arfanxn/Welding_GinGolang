package request

import (
	"github.com/arfanxn/welding/internal/infrastructure/http/request"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

var _ request.Request = (*UpdateMaterialTestMachine)(nil)

type UpdateMaterialTestMachine struct {
	Id          string  `form:"id" json:"id"`
	Name        *string `form:"name" json:"name"`
	Description *string `form:"description" json:"description"`
}

func NewUpdateMaterialTestMachine() *UpdateMaterialTestMachine {
	return &UpdateMaterialTestMachine{}
}

func (s *UpdateMaterialTestMachine) Validate() error {
	return validation.ValidateStruct(s,
		validation.Field(&s.Id,
			validation.Required,
			validation.Length(26, 26).Error("Id harus 26 karakter"),
		),
		validation.Field(&s.Name,
			validation.Length(3, 50).Error("Nama harus di antara 3 dan 50 karakter"),
		),
		validation.Field(&s.Description,
			validation.Length(0, 255).Error("Deskripsi harus di antara 0 dan 255 karakter"),
		),
	)
}
