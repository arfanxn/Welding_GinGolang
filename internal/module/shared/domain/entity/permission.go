package entity

import (
	"time"

	"github.com/arfanxn/welding/internal/module/permission/domain/enum"
	"github.com/gookit/goutil"
	"github.com/oklog/ulid/v2"
	"gorm.io/gorm"
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

func (u *Permission) BeforeSave(tx *gorm.DB) error {
	if goutil.IsEmpty(u.Id) {
		u.Id = ulid.Make().String()
	}
	return nil
}
