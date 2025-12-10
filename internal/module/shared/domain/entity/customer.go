package entity

import (
	"time"

	"gorm.io/gorm"
)

type Customer struct {
	Id          string         `json:"id" gorm:"primarykey"`
	AddressId   string         `json:"address_id"`
	Name        string         `json:"name"`
	PhoneNumber string         `json:"phone_number"`
	Email       string         `json:"email"`
	CreatedAt   time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   *time.Time     `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"autoDeleteTime"`

	Address *Address `json:"address,omitempty" gorm:"foreignKey:AddressId;references:Id"`
}

func NewCustomer() *Customer {
	return &Customer{}
}

func (c *Customer) TableName() string {
	return "customers"
}
