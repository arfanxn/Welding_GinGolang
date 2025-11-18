package entity

import (
	"time"

	"github.com/arfanxn/welding/internal/module/permission/domain/enum"
)

type Permission struct {
	Id        string              `json:"id" gorm:"primarykey"`
	Name      enum.PermissionName `json:"name"`
	CreatedAt time.Time           `json:"created_at" gorm:"autoCreateTime"`

	Roles []*Role `json:"roles,omitempty" gorm:"many2many:permission_role"`
}

func NewPermission() *Permission {
	return &Permission{}
}
