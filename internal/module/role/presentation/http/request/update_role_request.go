package request

import (
	"github.com/arfanxn/welding/internal/infrastructure/http/request"
	"github.com/creasty/defaults"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

var _ request.Request = (*UpdateRole)(nil)

type UpdateRole struct {
	Id            string   `form:"id" json:"id"`
	Name          string   `form:"name" json:"name"`
	PermissionIds []string `form:"permission_id" json:"permissions" default:"[]"`
}

func NewUpdateRole() *UpdateRole {
	defaults.Set(&UpdateRole{})
	return &UpdateRole{}
}

func (s *UpdateRole) Validate() error {
	return validation.ValidateStruct(s,
		validation.Field(&s.Id,
			validation.Required.Error("Id tidak boleh kosong"),
			validation.Length(26, 26).Error("Id harus 26 karakter"),
		),
		validation.Field(&s.Name,
			validation.Required.Error("Nama tidak boleh kosong"),
			validation.Length(3, 50).Error("Nama harus di antara 3 dan 50 karakter"),
		),
		validation.Field(&s.PermissionIds,
			validation.Each(
				is.Alphanumeric.Error("Permission id hanya huruf dan angka yang diperbolehkan"),
				validation.Length(26, 26).Error("Permission id harus 26 karakter"),
			),
		),
	)
}
