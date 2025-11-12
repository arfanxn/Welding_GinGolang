package entity

import (
	"time"

	"github.com/arfanxn/welding/pkg/boolutil"
	"github.com/gookit/goutil"
	"github.com/guregu/null/v6"
	"github.com/oklog/ulid/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
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
	EmploymentIdentityNumber null.String `json:"employment_identity_number,omitempty" gorm:"->"`
}

func NewUser() *User {
	return &User{}
}

func (u User) TableName() string {
	return "users"
}

func (u *User) BeforeSave(tx *gorm.DB) error {
	if goutil.IsEmpty(u.Id) {
		u.Id = ulid.Make().String()
	}
	return nil
}

func (u *User) SetPassword(password string) error {
	hashedPasswordBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPasswordBytes)
	return nil
}

func (u *User) SetEmploymentIdentityNumber(ein null.String) {
	u.Employee = boolutil.Ternary(ein.Valid,
		&Employee{
			UserId:                   u.Id,
			EmploymentIdentityNumber: ein.String,
		},
		nil,
	)
	u.EmploymentIdentityNumber = ein
}

func (u *User) MarkActivated() {
	u.ActivatedAt = null.TimeFrom(time.Now())
	u.DeactivatedAt = null.TimeFromPtr(nil)
}

func (u *User) MarkDeactivated() {
	u.ActivatedAt = null.TimeFromPtr(nil)
	u.DeactivatedAt = null.TimeFrom(time.Now())
}

func (u User) IsEmailVerified() bool {
	return u.EmailVerifiedAt.Valid
}

func (u User) IsActive() bool {
	return u.ActivatedAt.Valid && !u.DeactivatedAt.Valid
}
