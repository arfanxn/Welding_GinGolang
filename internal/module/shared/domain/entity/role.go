package entity

import (
	"time"

	"github.com/arfanxn/welding/internal/module/role/domain/enum"
	"github.com/guregu/null/v6"
)

type Role struct {
	Id        string        `json:"id" gorm:"primarykey;not null;unique;type:varchar(26);index"`
	Name      enum.RoleName `json:"name" gorm:"unique;not null;type:varchar(50);index"`
	IsDefault bool          `json:"is_default" gorm:"default:false"`
	CreatedAt time.Time     `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt null.Time     `json:"updated_at" gorm:"autoUpdateTime"`

	Users       []*User       `json:"users,omitempty" gorm:"many2many:role_user"`
	Permissions []*Permission `json:"permissions,omitempty" gorm:"many2many:permission_role"`
}

func NewRole() *Role {
	return &Role{}
}

func (u *Role) TableName() string {
	return "roles"
}

func (r *Role) IsSuperAdmin() bool {
	return r.Name == enum.SuperAdmin
}

func (r *Role) IsUpdateable() bool {
	return !r.IsSuperAdmin()
}

func (r *Role) IsSaveable() bool {
	return r.IsUpdateable()
}

func (r *Role) IsDestroyable() bool {
	return !r.IsSuperAdmin() && !r.IsDefault
}
