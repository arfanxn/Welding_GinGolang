package entity

import (
	"time"

	"github.com/guregu/null/v6"
)

type Role struct {
	Id        string    `json:"id" gorm:"primarykey;not null;unique;type:varchar(26);index"`
	Name      string    `json:"name" gorm:"unique;not null;type:varchar(50);index"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt null.Time `json:"updated_at" gorm:"autoUpdateTime"`

	Users       []*User       `json:"users,omitempty" gorm:"many2many:role_user"`
	Permissions []*Permission `json:"permissions,omitempty" gorm:"many2many:permission_role"`
}

func NewRole() *Role {
	return &Role{}
}
