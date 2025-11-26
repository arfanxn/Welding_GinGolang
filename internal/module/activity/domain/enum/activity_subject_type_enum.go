package enum

type SubjectType = string

const (
	UserSubjectType                SubjectType = "user"
	PermissionSubjectType          SubjectType = "permission"
	RoleSubjectType                SubjectType = "role"
	CodeSubjectType                SubjectType = "code"
	MaterialTestMethodSubjectType  SubjectType = "material_test_method"
	MaterialTestMachineSubjectType SubjectType = "material_test_machine"
	MaterialTestServiceSubjectType SubjectType = "material_test_service"
	ActivitySubjectType            SubjectType = "activity"
)

var SubjectTypes = []SubjectType{
	UserSubjectType,
	PermissionSubjectType,
	RoleSubjectType,
	CodeSubjectType,
	MaterialTestMethodSubjectType,
	MaterialTestMachineSubjectType,
	MaterialTestServiceSubjectType,
	ActivitySubjectType,
}
