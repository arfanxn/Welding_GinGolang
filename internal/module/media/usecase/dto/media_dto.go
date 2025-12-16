package dto

import (
	"mime/multipart"

	"github.com/arfanxn/welding/pkg/types"
)

type CreateFromMultipartFile struct {
	ModelType            string                `json:"model_type"`
	ModelId              string                `json:"model_id"`
	Ulid                 *string               `json:"ulid"`
	CollectionName       string                `json:"collection_name"`
	Name                 string                `json:"name"`
	File                 *multipart.FileHeader `json:"file"`
	Disk                 string                `json:"disk"`
	ConversionsDisk      *string               `json:"conversions_disk"`
	Manipulations        types.JSONMap         `json:"manipulations"`
	CustomProperties     types.JSONMap         `json:"custom_properties"`
	GeneratedConversions types.JSONMap         `json:"generated_conversions"`
	ResponsiveImages     types.JSONMap         `json:"responsive_images"`
	OrderColumn          *int                  `json:"order_column"`
}

type CreateFromMultipartFiles = []CreateFromMultipartFile

type SaveMedia struct {
	ModelType            *string               `json:"model_type"`
	ModelId              *string               `json:"model_id"`
	Ulid                 *string               `json:"ulid"`
	CollectionName       *string               `json:"collection_name"`
	Name                 *string               `json:"name"`
	File                 *multipart.FileHeader `json:"file"`
	Disk                 *string               `json:"disk"`
	ConversionsDisk      *string               `json:"conversions_disk"`
	Manipulations        types.JSONMap         `json:"manipulations"`
	CustomProperties     types.JSONMap         `json:"custom_properties"`
	GeneratedConversions types.JSONMap         `json:"generated_conversions"`
	ResponsiveImages     types.JSONMap         `json:"responsive_images"`
	OrderColumn          *int                  `json:"order_column"`
}
