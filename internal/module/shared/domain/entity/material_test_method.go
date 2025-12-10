package entity

import (
	"time"
)

type MaterialTestMethod struct {
	Id          string     `json:"id" gorm:"primarykey"`
	Name        string     `json:"name"`
	Description *string    `json:"description"`
	CreatedAt   time.Time  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   *time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt   *time.Time `json:"deleted_at" gorm:"autoDeleteTime"`
}

func NewMaterialTestMethod() *MaterialTestMethod {
	return &MaterialTestMethod{}
}

func (m *MaterialTestMethod) TableName() string {
	return "material_test_methods"
}
