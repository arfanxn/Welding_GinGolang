package entity

import (
	"time"

	"github.com/guregu/null/v6"
)

type PermissionRole struct {
	PermissionId string    `json:"permission_id" gorm:"primaryKey"`
	RoleId       string    `json:"role_id" gorm:"primaryKey"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    null.Time `json:"updated_at"`

	Permission *Permission `gorm:"foreignKey:PermissionId;references:id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Role       *Role       `gorm:"foreignKey:RoleId;references:id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

// TableName specifies the table name for the PermissionRole model
func (PermissionRole) TableName() string {
	return "permission_role"
}
