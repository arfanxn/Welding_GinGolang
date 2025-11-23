package enum

type ActivitySubjectType = string

const (
	UserSubjectType       ActivitySubjectType = "user"
	PermissionSubjectType ActivitySubjectType = "permission"
	RoleSubjectType       ActivitySubjectType = "role"
	CodeSubjectType       ActivitySubjectType = "code"
)

var ActivitySubjectTypes = []ActivitySubjectType{
	UserSubjectType,
	PermissionSubjectType,
	RoleSubjectType,
	CodeSubjectType,
}
