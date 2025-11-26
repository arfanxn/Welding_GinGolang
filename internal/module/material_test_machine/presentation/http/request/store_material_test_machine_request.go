package request

import (
	"github.com/arfanxn/welding/internal/infrastructure/http/request"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

var _ request.Request = (*StoreMaterialTestMachine)(nil)

type StoreMaterialTestMachine struct {
	Name        string  `form:"name" json:"name"`
	Description *string `form:"description" json:"description"`
}

func NewStoreMaterialTestMachine() *StoreMaterialTestMachine {
	return &StoreMaterialTestMachine{}
}

func (s *StoreMaterialTestMachine) Validate() error {
	return validation.ValidateStruct(s,
		validation.Field(&s.Name,
			validation.Required.Error("Nama tidak boleh kosong"),
			validation.Length(3, 50).Error("Nama harus di antara 3 dan 50 karakter"),
		),
		validation.Field(&s.Description,
			validation.Length(0, 255).Error("Deskripsi harus di antara 0 dan 255 karakter"),
		),
	)
}
