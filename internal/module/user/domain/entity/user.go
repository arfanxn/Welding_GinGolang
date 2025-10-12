package entity

import (
	"time"

	"github.com/guregu/null/v6"
)

type User struct {
	Id              string    `gorm:"primarykey"`
	Name            string    `json:"name"`
	Email           string    `json:"email"`
	EmailVerifiedAt null.Time `json:"email_verified_at"`
	Password        string    `json:"password,omitempty"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       null.Time `json:"updated_at"`
	DeletedAt       null.Time `json:"deleted_at"`
}

func NewUser() *User {
	return &User{}
}
