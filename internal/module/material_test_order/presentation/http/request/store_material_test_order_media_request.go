package request

import (
	"mime/multipart"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type StoreMaterialTestOrderMedia struct {
	OrderId string                `form:"order_id" json:"order_id"`
	Name    string                `form:"name" json:"name"`
	File    *multipart.FileHeader `form:"file" json:"file"`
}

func NewStoreMaterialTestOrderMedia() *StoreMaterialTestOrderMedia {
	return &StoreMaterialTestOrderMedia{}
}

func (r *StoreMaterialTestOrderMedia) Validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.OrderId,
			validation.Required.Error("Id tidak boleh kosong"),
			validation.Length(26, 26).Error("Id harus 26 karakter"),
		),
		validation.Field(&r.Name,
			validation.Required.Error("Nama media wajib disi"),
			validation.Length(3, 255).Error("Nama media harus di antara 3 dan 255 karakter"),
		),
	)
}
