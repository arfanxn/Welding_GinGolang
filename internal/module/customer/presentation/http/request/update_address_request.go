package request

import (
	"github.com/arfanxn/welding/internal/infrastructure/http/request"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

var _ request.Request = (*UpdateCustomerRequest)(nil)

type UpdateCustomerRequest struct {
	Id          string `form:"id" json:"id"`
	AddressId   string `form:"address_id" json:"address_id"`
	Name        string `form:"name" json:"name"`
	PhoneNumber string `form:"phone_number" json:"phone_number"`
	Email       string `form:"email" json:"email"`
}

func NewUpdateCustomerRequest() *UpdateCustomerRequest {
	return &UpdateCustomerRequest{}
}

func (r *UpdateCustomerRequest) Validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.Id,
			validation.Required,
			validation.Length(26, 26).Error("Id harus 26 karakter"),
		),
		validation.Field(&r.AddressId,
			validation.Required,
			validation.Length(26, 26).Error("Address Id harus 26 karakter"),
		),
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
	)
}
