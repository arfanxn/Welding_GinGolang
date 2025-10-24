package entity

import (
	"time"

	"github.com/guregu/null/v6"
)

type User struct {
	Id              string    `json:"id" gorm:"primarykey"`
	Name            string    `json:"name"`
	PhoneNumber     string    `json:"phone_number"`
	Email           string    `json:"email"`
	EmailVerifiedAt null.Time `json:"email_verified_at"`
	Password        string    `json:"-"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       null.Time `json:"updated_at"`
	DeletedAt       null.Time `json:"deleted_at"`

	// Relations
	Roles    []*Role   `json:"roles,omitempty" gorm:"many2many:role_user"`
	Employee *Employee `json:"employee,omitempty" gorm:"foreignKey:UserId;references:Id"`

	// Joins
	EmploymentIdentityNumber string `json:"employment_identity_number,omitempty"`
}

func NewUser() *User {
	return &User{}
}

func (u User) TableName() string {
	return "users"
}
