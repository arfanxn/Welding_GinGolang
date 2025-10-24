package dto

import (
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	"github.com/guregu/null/v6"
)

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResult struct {
	User  *entity.User
	Token string
}

type SaveUser struct {
	Id          null.String `json:"id"`
	Name        string      `json:"name"`
	PhoneNumber string      `json:"phone_number"`
	Email       string      `json:"email"`
	Password    string      `json:"password"`
	RoleIds     []string    `json:"role_ids"`
}

type DestroyUser struct {
	Id string `json:"id"`
}
