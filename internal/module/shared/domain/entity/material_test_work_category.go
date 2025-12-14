package entity

import (
	"time"

	"gorm.io/gorm"
)

// WorkCategory
// Terjemahan langsung dari "Jenis Pekerjaan" pada dokumen fisik.
// Digunakan untuk mengkategorisasi PEMOHON:
// - Pendidik
// - IKM
// - Industri
// - Other
// BUKAN kategori pekerjaan fisik.
type MaterialTestWorkCategory struct {
	Id          string         `json:"id" gorm:"primarykey"`
	Name        string         `json:"name"`
	Description *string        `json:"description"`
	CreatedAt   time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   *time.Time     `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"autoDeleteTime"`
}

func NewMaterialTestWorkCategory() *MaterialTestWorkCategory {
	return &MaterialTestWorkCategory{}
}

func (m *MaterialTestWorkCategory) TableName() string {
	return "material_test_work_categories"
}
