package dto

import "github.com/arfanxn/welding/internal/module/shared/domain/entity"

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResult struct {
	User  *entity.User
	Token string
}
