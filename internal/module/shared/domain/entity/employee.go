package entity

import (
	"time"

	"github.com/guregu/null/v6"
)

type Employee struct {
	UserId                   string    `json:"user_id" gorm:"primarykey"`
	EmploymentIdentityNumber string    `json:"employment_identity_number"`
	CreatedAt                time.Time `json:"created_at"`
	UpdatedAt                null.Time `json:"updated_at"`

	User *User `json:"user" gorm:"foreignKey:UserId;references:Id"`
}

func NewEmployee() *Employee {
	return &Employee{}
}

func (e Employee) TableName() string {
	return "employees"
}
