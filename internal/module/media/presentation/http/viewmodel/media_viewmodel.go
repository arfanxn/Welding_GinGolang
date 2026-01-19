package viewmodel

import (
	"time"

	"github.com/arfanxn/welding/pkg/types"
)

type Media struct {
	// Primary identifier for the media record (ULID/UUID/string).
	Id string `json:"id"`

	// The type of model this media belongs to (polymorphic).
	// Example: "MaterialTestOrder", "User", "Invoice".
	ModelType string `json:"model_type"`

	// The ID of the model this media is attached to.
	ModelId string `json:"model_id"`

	// Extra identifier (often ULID/string) for compatibility or external references.
	Ulid *string `json:"ulid"`

	// Logical group of media files. Similar to Laravel Media Library collections.
	// Example: "customer_attachments", "tester_attachments", "sample_photos".
	CollectionName string `json:"collection_name"`

	// Human-readable display name for this media.
	Name string `json:"name"`

	// Original file name as uploaded by the user (with extension).
	FileName string `json:"file_name"`

	// URL to access the file.
	FileUrl string `json:"file_url"`

	// MIME type of the file. Example: "image/jpeg", "application/pdf".
	MimeType *string `json:"mime_type"`

	// Disk/storage driver where this file is stored.
	// Example: "s3", "local", "gcs".
	Disk string `json:"disk"`

	// Disk/storage used for generated conversions (thumbnails, webp versions, etc.).
	// Normally same as Disk; optional if unused.
	ConversionsDisk *string `json:"conversions_disk"`

	// Size of the stored file in bytes.
	Size int64 `json:"size"`

	// Manipulation instructions saved as JSON.
	// Typically contains resizing/cropping configuration.
	Manipulations types.JSONMap `json:"manipulations"`

	// Flexible JSON bag for adding metadata like width/height/custom tags/user info.
	CustomProperties types.JSONMap `json:"custom_properties"`

	// Metadata for generated image conversions.
	// Example JSON: {"thumb": "path/to/thumb.jpg", "webp": "path/to/image.webp"}
	GeneratedConversions types.JSONMap `json:"generated_conversions"`

	// Metadata for responsive images (screensize-specific variants).
	ResponsiveImages types.JSONMap `json:"responsive_images"`

	// Allows manually sorting media items inside the same collection.
	OrderColumn *int `json:"order_column"`

	// Timestamp when the media record was created.
	CreatedAt time.Time `json:"created_at"`

	// Timestamp when the media record was last updated.
	UpdatedAt *time.Time `json:"updated_at"`
}
