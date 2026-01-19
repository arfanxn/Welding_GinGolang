package entity

import (
	"time"

	"gorm.io/gorm"
)

type MaterialTestMachine struct {
	Id          string         `json:"id" gorm:"primarykey"`
	Name        string         `json:"name"`
	Description *string        `json:"description"`
	CreatedAt   time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   *time.Time     `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"autoDeleteTime"`
}

func NewMaterialTestMachine() *MaterialTestMachine {
	return &MaterialTestMachine{}
}

func (m *MaterialTestMachine) TableName() string {
	return "material_test_machines"
}
