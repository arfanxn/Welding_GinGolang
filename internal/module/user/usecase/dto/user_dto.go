package dto

import (
	"time"

	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
)

type Register struct {
	Name                     string  `json:"name"`
	PhoneNumber              string  `json:"phone_number"`
	Email                    string  `json:"email"`
	Password                 string  `json:"password"`
	InvitationCode           *string `json:"invitation_code"`
	EmploymentIdentityNumber *string `json:"employment_identity_number"`
}

type VerifyEmail struct {
	Email string `json:"email"`
	Code  string `json:"code"`
}

type ResetPassword struct {
	Email    string `json:"email"`
	Code     string `json:"code"`
	Password string `json:"password"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResult struct {
	User  *entity.User
	Token string
}

type SaveUser struct {
	Id                       *string    `json:"id"`
	Name                     *string    `json:"name"`
	PhoneNumber              *string    `json:"phone_number"`
	Email                    *string    `json:"email"`
	Password                 *string    `json:"password"`
	RoleIds                  []string   `json:"role_ids"`
	ActivatedAt              *time.Time `json:"activated_at"`
	DeactivatedAt            *time.Time `json:"deactivated_at"`
	EmploymentIdentityNumber *string    `json:"employment_identity_number"`
}

type UpdateUserMePassword struct {
	CurrentPassword string `json:"current_password"`
	Password        string `json:"password"`
}

type UpdateUserPassword struct {
	Id       string `json:"id"`
	Password string `json:"password"`
}

type ToggleActivation struct {
	Id string `json:"id"`
}

type DestroyUser struct {
	Id string `json:"id"`
}
