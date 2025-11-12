package dto

import (
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	"github.com/guregu/null/v6"
)

type Register struct {
	Name                     string      `json:"name"`
	PhoneNumber              string      `json:"phone_number"`
	Email                    string      `json:"email"`
	Password                 string      `json:"password"`
	InvitationCode           null.String `json:"invitation_code"`
	EmploymentIdentityNumber null.String `json:"employment_identity_number"`
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
	Id                       null.String `json:"id"`
	Name                     string      `json:"name"`
	PhoneNumber              string      `json:"phone_number"`
	Email                    string      `json:"email"`
	Password                 string      `json:"password"`
	RoleIds                  []string    `json:"role_ids"`
	ActivatedAt              null.Time   `json:"activated_at"`
	DeactivatedAt            null.Time   `json:"deactivated_at"`
	EmploymentIdentityNumber null.String `json:"employment_identity_number"`
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
