package entity

import (
	"time"

	"github.com/guregu/null/v6"
)

type Permission struct {
	Id        string    `json:"id" gorm:"primarykey"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt null.Time `json:"updated_at"`
	DeletedAt null.Time `json:"deleted_at"`

	Roles []*Role `json:"roles" gorm:"many2many:permission_role;references:id;joinReferences:permission_id;foreignReferences:role_id"`
}

func NewPermission() *Permission {
	return &Permission{}
}
