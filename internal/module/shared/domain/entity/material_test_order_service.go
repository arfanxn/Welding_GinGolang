package entity

import (
	"time"
)

type MaterialTestOrderService struct {
	Id           string     `json:"id" gorm:"primarykey"`
	OrderId      string     `json:"order_id"`
	ServiceId    string     `json:"service_id"`
	SampleNumber *string    `json:"sample_number"`
	SampleName   string     `json:"sample_name"`
	Price        float64    `json:"price"`
	Quantity     int        `json:"quantity"`
	LineTotal    float64    `json:"line_total"`
	CreatedAt    time.Time  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    *time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	Order      *MaterialTestOrder                  `json:"order,omitempty" gorm:"foreignKey:OrderId;references:Id;"`
	Service    *MaterialTestService                `json:"service,omitempty" gorm:"foreignKey:ServiceId;references:Id;"`
	Evaluation *MaterialTestOrderServiceEvaluation `json:"evaluation,omitempty" gorm:"foreignKey:OrderServiceId;references:Id;"`
}

func NewMaterialTestOrderService() *MaterialTestOrderService {
	return &MaterialTestOrderService{}
}

func (m *MaterialTestOrderService) TableName() string {
	return "material_test_order_services"
}
