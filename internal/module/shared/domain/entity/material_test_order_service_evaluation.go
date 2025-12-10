package entity

import (
	"time"
)

type MaterialTestOrderServiceEvaluation struct {
	Id             string `json:"id" gorm:"primarykey"`
	OrderServiceId string `json:"order_service_id"`

	IsEquipmentAvailable      bool `json:"is_equipment_available"`
	IsPersonnelAvailable      bool `json:"is_personnel_available"`
	IsTimeAvailable           bool `json:"is_time_available"`
	IsTestReady               bool `json:"is_test_ready"`
	IsSubcontractLabAvailable bool `json:"is_subcontract_lab_available"`

	CreatedAt time.Time  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	OrderService *MaterialTestOrderService `json:"order_service,omitempty" gorm:"foreignKey:OrderServiceId;references:Id;"`
}

func NewMaterialTestOrderServiceEvaluation() *MaterialTestOrderServiceEvaluation {
	return &MaterialTestOrderServiceEvaluation{}
}

func (m *MaterialTestOrderServiceEvaluation) TableName() string {
	return "material_test_order_service_evaluations"
}
