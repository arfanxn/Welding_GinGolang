package entity

import (
	"time"

	"gorm.io/gorm"
)

type Address struct {
	Id          string         `json:"id" gorm:"primarykey"`
	FullAddress string         `json:"full_address"`
	CreatedAt   time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   *time.Time     `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"autoDeleteTime"`
}

func NewAddress() *Address {
	return &Address{}
}

func (a *Address) TableName() string {
	return "addresses"
}
