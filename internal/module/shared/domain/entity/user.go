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
	ActivatedAt     null.Time `json:"activated_at"`
	DeactivatedAt   null.Time `json:"deactivated_at"`
	CreatedAt       time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt       null.Time `json:"updated_at" gorm:"autoUpdateTime"`

	// Relations
	Roles    []*Role   `json:"roles,omitempty" gorm:"many2many:role_user"`
	Employee *Employee `json:"employee,omitempty" gorm:"foreignKey:UserId;references:Id"`

	// Joins
	EmploymentIdentityNumber *null.String `json:"employment_identity_number,omitempty" gorm:"->"`
}

func NewUser() *User {
	return &User{}
}

func (u User) TableName() string {
	return "users"
}

func (u *User) MarkActivated() {
	u.ActivatedAt = null.TimeFrom(time.Now())
	u.DeactivatedAt = null.TimeFromPtr(nil)
}

func (u *User) MarkDeactivated() {
	u.ActivatedAt = null.TimeFromPtr(nil)
	u.DeactivatedAt = null.TimeFrom(time.Now())
}

func (u *User) MarkEmailVerified() {
	u.EmailVerifiedAt = null.TimeFrom(time.Now())
}

func (u User) IsEmailVerified() bool {
	return u.EmailVerifiedAt.Valid
}

func (u User) IsActive() bool {
	return u.ActivatedAt.Valid && !u.DeactivatedAt.Valid
}
