package request

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type StoreUser struct {
	Name                     string   `form:"name" json:"name"`
	PhoneNumber              string   `form:"phone_number" json:"phone_number"`
	Email                    string   `form:"email" json:"email"`
	Password                 string   `form:"password" json:"password"`
	RoleIds                  []string `form:"role_id" json:"role_ids"`
	EmploymentIdentityNumber *string  `form:"employment_identity_number" json:"employment_identity_number"`
}

func (r *StoreUser) Validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.Name,
			validation.Required.Error("Nama wajib diisi"),
			validation.Length(3, 64).Error("Panjang nama harus antara 3-64 karakter"),
		),
		validation.Field(&r.PhoneNumber,
			validation.Required.Error("Nomor telepon wajib diisi"),
			validation.Length(10, 15).Error("Panjang nomor telepon harus antara 10-15 karakter"),
		),
		validation.Field(&r.Email,
			validation.Required.Error("Email wajib diisi"),
			validation.Length(3, 50).Error("Panjang email harus antara 3-50 karakter"),
			is.Email.Error("Format email tidak valid"),
		),
		validation.Field(&r.Password,
			validation.Required.Error("Kata sandi wajib diisi"),
			validation.Length(8, 255).Error("Panjang kata sandi minimal 8 karakter"),
		),
		validation.Field(&r.RoleIds,
			validation.Each(
				is.Alphanumeric.Error("Role id hanya huruf dan angka yang diperbolehkan"),
				validation.Length(26, 26).Error("Role id harus 26 karakter"),
			),
		),
		validation.Field(&r.EmploymentIdentityNumber,
			validation.Length(10, 50).Error("Panjang NIP harus antara 10-50 karakter"),
		),
	)
}
