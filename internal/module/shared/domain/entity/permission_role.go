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
}

// TableName specifies the table name for the PermissionRole model
func (PermissionRole) TableName() string {
	return "permission_role"
}
