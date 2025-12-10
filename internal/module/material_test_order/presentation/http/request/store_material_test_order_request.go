package request

import (
	"github.com/arfanxn/welding/internal/infrastructure/http/request"
	"github.com/creasty/defaults"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

var _ request.Request = (*StoreMaterialTestOrder)(nil)

type StoreMaterialTestOrder struct {
	Name string `form:"name" json:"name"`
}

func NewStoreMaterialTestOrder() *StoreMaterialTestOrder {
	defaults.Set(&StoreMaterialTestOrder{})
	return &StoreMaterialTestOrder{}
}

func (s *StoreMaterialTestOrder) Validate() error {
	return validation.ValidateStruct(s,
		validation.Field(&s.Name,
			validation.Required.Error("Nama tidak boleh kosong"),
			validation.Length(3, 50).Error("Nama harus di antara 3 dan 50 karakter"),
		),
	)
}
