package request

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type CreateUserRegisterInvitation struct {
	RoleId    string `form:"role_id" json:"role_id"`
	ExpiredAt string `form:"expired_at" json:"expired_at"`
}

func (r *CreateUserRegisterInvitation) Validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.RoleId,
			validation.Required.Error("Role Id wajib diisi"),
			validation.Length(26, 26).Error("Role Id harus 26 karakter"),
		),
		validation.Field(&r.ExpiredAt,
			validation.Required.Error("Expired in wajib diisi"),
			validation.Date(time.DateTime).Error("Format tanggal tidak valid. Gunakan format: YYYY-MM-DD HH:MM:SS"),
			validation.Date(time.DateTime).Min(time.Now().AddDate(0, 0, 1)).Error("Expired date harus lebih dari hari ini"),
		),
	)
}
