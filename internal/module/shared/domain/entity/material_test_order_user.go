package entity

import (
	"time"

	materialTestOrderUserEnum "github.com/arfanxn/welding/internal/module/material_test_order_user/domain/enum"
	"github.com/guregu/null/v6"
)

type MaterialTestOrderUser struct {
	OrderId   string                                              `json:"role_id" gorm:"primaryKey"`
	UserId    string                                              `json:"user_id" gorm:"primaryKey"`
	Type      materialTestOrderUserEnum.MaterialTestOrderUserType `json:"type"`
	CreatedAt time.Time                                           `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt null.Time                                           `json:"updated_at" gorm:"autoUpdateTime"`

	Order *MaterialTestOrder `json:"order,omitempty" gorm:"foreignKey:OrderId;references:Id"`
	User  *User              `json:"user,omitempty" gorm:"foreignKey:UserId;references:Id"`
}

func NewMaterialTestOrderUser() *MaterialTestOrderUser {
	return &MaterialTestOrderUser{}
}

// TableName specifies the table name for the OrderUser model
func (MaterialTestOrderUser) TableName() string {
	return "material_test_order_user"
}
