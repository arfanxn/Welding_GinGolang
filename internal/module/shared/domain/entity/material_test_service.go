package entity

import (
	"time"
)

type MaterialTestService struct {
	Id          string     `json:"id" gorm:"primarykey"`
	MachineId   string     `json:"machine_id"`
	MethodId    string     `json:"method_id"`
	ServiceType string     `json:"service_type"`
	ServiceCode string     `json:"service_code"`
	Unit        string     `json:"unit"`
	Price       float64    `json:"price"`
	CreatedAt   time.Time  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   *time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	Machine *MaterialTestMachine `gorm:"foreignKey:MachineId;references:Id;constraint:OnDelete:CASCADE"`
	Method  *MaterialTestMethod  `gorm:"foreignKey:MethodId;references:Id;constraint:OnDelete:CASCADE"`
}

func NewMaterialTestService() *MaterialTestService {
	return &MaterialTestService{}
}

func (m *MaterialTestService) TableName() string {
	return "material_test_services"
}
