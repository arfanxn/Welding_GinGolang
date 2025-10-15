package entity

import (
	"time"

	"github.com/guregu/null/v6"
)

type Role struct {
	Id        string    `gorm:"primarykey"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt null.Time `json:"updated_at"`
	DeletedAt null.Time `json:"deleted_at"`

	Users       []*User       `gorm:"many2many:role_user;references:id;joinReferences:role_id;foreignReferences:user_id"`
	Permissions []*Permission `gorm:"many2many:permission_role;references:id;joinReferences:role_id;foreignReferences:permission_id"`
}

func NewRole() *Role {
	return &Role{}
}
