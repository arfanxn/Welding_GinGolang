package dto

import "time"

type CreateUserRegisterInvitation struct {
	RoleId    string    `json:"role_id"`
	ExpiredAt time.Time `json:"expired_at"`
}

type CreateUserEmailVerification struct {
	Email string `json:"email"`
}

type CreateUserResetPassword struct {
	Email string `json:"email"`
}
