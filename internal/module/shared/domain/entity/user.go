package entity

import (
	"time"

	"github.com/guregu/null/v6"
)

type User struct {
	Id              string    `json:"id" gorm:"primarykey"`
	Name            string    `json:"name"`
	Email           string    `json:"email"`
	EmailVerifiedAt null.Time `json:"email_verified_at"`
	Password        string    `json:"-"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       null.Time `json:"updated_at"`
	DeletedAt       null.Time `json:"deleted_at"`

	Roles []*Role `json:"roles" gorm:"many2many:role_user"`
}

func NewUser() *User {
	return &User{}
}
