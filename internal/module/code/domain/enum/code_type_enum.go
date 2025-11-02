package enum

type CodeType string

const (
	UserRegisterInvitation CodeType = "user_register_invitation"
	UserEmailVerification  CodeType = "user_email_verification"
	UserResetPassword      CodeType = "user_reset_password"
)

func (p CodeType) String() string {
	return string(p)
}

var CodeTypes = []CodeType{
	UserRegisterInvitation,
	UserEmailVerification,
	UserResetPassword,
}
