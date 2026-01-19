package enum

type SubjectType = string

const (
	UserSubjectType                     SubjectType = "user"
	PermissionSubjectType               SubjectType = "permission"
	RoleSubjectType                     SubjectType = "role"
	CodeSubjectType                     SubjectType = "code"
	MaterialTestMethodSubjectType       SubjectType = "material_test_method"
	MaterialTestMachineSubjectType      SubjectType = "material_test_machine"
	MaterialTestServiceSubjectType      SubjectType = "material_test_service"
	MaterialTestWorkCategorySubjectType SubjectType = "material_test_work_category"
	MaterialTestWorkPackageSubjectType  SubjectType = "material_test_work_package"
	ActivitySubjectType                 SubjectType = "activity"
	AddressSubjectType                  SubjectType = "address"
	CustomerSubjectType                 SubjectType = "customer"
	MaterialTestOrderSubjectType        SubjectType = "material_test_order"
)

var SubjectTypes = []SubjectType{
	UserSubjectType,
	PermissionSubjectType,
	RoleSubjectType,
	CodeSubjectType,
	MaterialTestMethodSubjectType,
	MaterialTestMachineSubjectType,
	MaterialTestServiceSubjectType,
	MaterialTestWorkCategorySubjectType,
	MaterialTestWorkPackageSubjectType,
	ActivitySubjectType,
	AddressSubjectType,
	CustomerSubjectType,
	MaterialTestOrderSubjectType,
}
