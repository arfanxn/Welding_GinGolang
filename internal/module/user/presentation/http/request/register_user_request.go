package request

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	is "github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/gookit/goutil"
)

type RegisterUser struct {
	Name                     string `form:"name" json:"name"`
	PhoneNumber              string `form:"phone_number" json:"phone_number"`
	Email                    string `form:"email" json:"email"`
	Password                 string `form:"password" json:"password"`
	PasswordConfirmation     string `form:"password_confirmation" json:"password_confirmation"`
	InvitationCode           string `form:"invitation_code" json:"invitation_code"`
	EmploymentIdentityNumber string `form:"employment_identity_number" json:"employment_identity_number"`
}

func (r *RegisterUser) Validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.Name,
			validation.Required.Error("Nama wajib diisi"),
			validation.Length(2, 64).Error("Panjang nama harus antara 2-64 karakter"),
		),
		validation.Field(&r.PhoneNumber,
			validation.Required.Error("Nomor telepon wajib diisi"),
			validation.Length(10, 15).Error("Panjang nomor telepon harus antara 10-15 karakter"),
		),
		validation.Field(&r.Email,
			validation.Required.Error("Email wajib diisi"),
			validation.Length(3, 50).Error("Panjang email harus antara 3-50 karakter"),
			is.EmailFormat.Error("Format email tidak valid"),
		),
		validation.Field(&r.Password,
			validation.Required.Error("Kata sandi wajib diisi"),
			validation.Length(8, 255).Error("Panjang kata sandi minimal 8 karakter"),
		),
		validation.Field(&r.PasswordConfirmation,
			validation.Required.Error("Konfirmasi kata sandi wajib diisi"),
			validation.By(func(value any) error {
				if value.(string) != r.Password {
					return validation.NewError("password_mismatch", "Kata sandi tidak cocok")
				}
				return nil
			}),
		),
		validation.Field(&r.InvitationCode,
			validation.Length(6, 6).Error("Panjang kode undangan harus 6 karakter"),
		),
		validation.Field(&r.EmploymentIdentityNumber,
			validation.When(
				!goutil.IsZero(r.InvitationCode),
				validation.Required.Error("NIP wajib diisi jika menggunakan kode undangan"),
			),
			validation.Length(10, 50).Error("Panjang NIP harus antara 10-50 karakter"),
		),
	)
}
