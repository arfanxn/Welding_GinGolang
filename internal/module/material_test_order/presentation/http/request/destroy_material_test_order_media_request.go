package request

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type DestroyMaterialTestOrderMedia struct {
	Id      string `form:"id" json:"id"`
	MediaId string `form:"media_id" json:"media_id"`
}

func NewDestroyMaterialTestOrderMedia() *DestroyMaterialTestOrderMedia {
	return &DestroyMaterialTestOrderMedia{}
}

func (r *DestroyMaterialTestOrderMedia) Validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.Id,
			validation.Required.Error("Id tidak boleh kosong"),
			validation.Length(26, 26).Error("Id harus 26 karakter"),
		),
		validation.Field(&r.MediaId,
			validation.Required.Error("Media id tidak boleh kosong"),
			validation.Length(26, 26).Error("Media id harus 26 karakter"),
		),
	)
}
