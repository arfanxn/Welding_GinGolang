package request

import (
	"github.com/arfanxn/welding/internal/infrastructure/http/request"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

var _ request.Request = (*SetDefaultRole)(nil)

type SetDefaultRole struct {
	Id string `form:"id" json:"id"`
}

func NewSetDefaultRole() *SetDefaultRole {
	return &SetDefaultRole{}
}

func (s *SetDefaultRole) Validate() error {
	return validation.ValidateStruct(s,
		validation.Field(&s.Id,
			validation.Required.Error("Id tidak boleh kosong"),
			validation.Length(26, 26).Error("Id harus 26 karakter"),
		),
	)
}
