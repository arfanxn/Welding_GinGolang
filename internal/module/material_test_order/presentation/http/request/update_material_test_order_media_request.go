package request

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type UpdateMaterialTestOrderMedia struct {
	Id      string `form:"id" json:"id"`
	MediaId string `form:"media_id" json:"media_id"`
	Name    string `form:"name" json:"name"`
}

func NewUpdateMaterialTestOrderMedia() *UpdateMaterialTestOrderMedia {
	return &UpdateMaterialTestOrderMedia{}
}

func (r *UpdateMaterialTestOrderMedia) Validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.Id,
			validation.Required.Error("Id tidak boleh kosong"),
			validation.Length(26, 26).Error("Id harus 26 karakter"),
		),
		validation.Field(&r.MediaId,
			validation.Required.Error("Media id tidak boleh kosong"),
			validation.Length(26, 26).Error("Media id harus 26 karakter"),
		),
		validation.Field(&r.Name,
			validation.Required.Error("Nama media wajib disi"),
			validation.Length(3, 255).Error("Nama media harus di antara 3 dan 255 karakter"),
		),
	)
}
