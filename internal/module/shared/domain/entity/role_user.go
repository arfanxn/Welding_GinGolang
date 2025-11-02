package entity

import (
	"time"

	"github.com/guregu/null/v6"
)

type RoleUser struct {
	RoleId    string    `json:"role_id" gorm:"primaryKey"`
	UserId    string    `json:"user_id"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt null.Time `json:"updated_at" gorm:"autoUpdateTime"`

	Role *Role `gorm:"foreignKey:RoleId;references:id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	User *User `gorm:"foreignKey:UserId;references:id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

// TableName specifies the table name for the RoleUser model
func (RoleUser) TableName() string {
	return "role_user"
}
