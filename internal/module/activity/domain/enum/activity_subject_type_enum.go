package enum

type SubjectType = string

const (
	UserSubjectType               SubjectType = "user"
	PermissionSubjectType         SubjectType = "permission"
	RoleSubjectType               SubjectType = "role"
	CodeSubjectType               SubjectType = "code"
	MaterialTestMethodSubjectType SubjectType = "material_test_method"
	ActivitySubjectType           SubjectType = "activity"
)

var SubjectTypes = []SubjectType{
	UserSubjectType,
	PermissionSubjectType,
	RoleSubjectType,
	CodeSubjectType,
	MaterialTestMethodSubjectType,
	ActivitySubjectType,
}
