package errorx

type Errorx interface {
	Error() string
}

type errorx struct {
	s string
}

func New(s string) Errorx {
	return &errorx{s: s}
}

func (e *errorx) Error() string {
	return e.s
}

var (
	// ========================================
	// User Errors
	// ========================================

	// ErrUserNotFound is returned when a user is not found in the system
	ErrUserNotFound Errorx = New("user not found")

	// ErrUserAlreadyExists is returned when attempting to create a user that already exists
	ErrUserAlreadyExists Errorx = New("user already exists")

	// ErrUserEmailAlreadyVerified is returned when trying to verify an already verified email
	ErrUserEmailAlreadyVerified Errorx = New("user email already verified")

	// ErrUserPasswordIncorrect is returned when the provided password is incorrect
	ErrUserPasswordIncorrect Errorx = New("user password incorrect")

	// ErrUserSuperAdminUpdateForbidden is returned when attempting to update a super admin user
	ErrUserSuperAdminUpdateForbidden Errorx = New("user super admin update forbidden")

	// ErrSuperAdminRoleChangeForbidden is returned when attempting to change a super admin's role
	ErrUserSuperAdminRoleChangeForbidden Errorx = New("user super admin role change forbidden")

	// ErrUserSuperAdminAssignmentForbidden is returned when attempting to assign super admin role to a user
	ErrUserSuperAdminAssignmentForbidden Errorx = New("user super admin assignment forbidden")

	// ========================================
	// Role Errors
	// ========================================

	// ErrRoleNotFound is returned when a specific role is not found
	ErrRoleNotFound Errorx = New("role not found")

	// ErrRolesNotFound is returned when one or more requested roles are not found
	ErrRolesNotFound Errorx = New("roles not found")

	// ErrRoleAlreadyExists is returned when attempting to create a role that already exists
	ErrRoleAlreadyExists Errorx = New("role already exists")

	// ErrRoleAlreadyDefault is returned when attempting to set a role as default that is already default
	ErrRoleAlreadyDefault Errorx = New("role already default")

	// ErrRoleDefaultNotConfigured is returned when the system default role is not configured
	ErrRoleDefaultNotConfigured Errorx = New("role default not configured")

	// ErrRoleDefaultDestroyForbidden is returned when attempting to destroy a default role
	ErrRoleDefaultDestroyForbidden Errorx = New("role default destroy forbidden")

	// ErrRoleSuperAdminStoreForbidden is returned when attempting to store a super admin role
	ErrRoleSuperAdminStoreForbidden Errorx = New("role super admin store forbidden")

	// ErrRoleSuperAdminUpdateForbidden is returned when attempting to update a super admin role
	ErrRoleSuperAdminUpdateForbidden Errorx = New("role super admin update forbidden")

	// ErrRoleSuperAdminSetDefaultForbidden is returned when attempting to set a super admin role as default
	ErrRoleSuperAdminSetDefaultForbidden Errorx = New("role super admin set default forbidden")

	// ErrRoleSuperAdminDestroyForbidden is returned when attempting to destroy a super admin role
	ErrRoleSuperAdminDestroyForbidden Errorx = New("role super admin destroy forbidden")

	// ========================================
	// Permission Errors
	// ========================================

	// ErrPermissionNotFound is returned when a specific permission is not found
	ErrPermissionNotFound Errorx = New("permission not found")

	// ErrPermissionsNotFound is returned when one or more requested permissions are not found
	ErrPermissionsNotFound Errorx = New("permissions not found")

	// ErrPermissionAlreadyExists is returned when attempting to create a permission that already exists
	ErrPermissionAlreadyExists Errorx = New("permission already exists")

	// ========================================
	// Permission Role Errors
	// ========================================

	// ErrPermissionRoleNotFound is returned when a specific permission role is not found
	ErrPermissionRoleNotFound Errorx = New("permission role not found")

	// ErrPermissionRoleAlreadyExists is returned when attempting to create a permission role that already exists
	ErrPermissionRoleAlreadyExists Errorx = New("permission role already exists")

	// ========================================
	// Code Errors
	// ========================================

	// ErrCodeNotFound is returned when a verification code is not found
	ErrCodeNotFound Errorx = New("code not found")

	// ErrCodeAlreadyExists is returned when attempting to create a code that already exists
	ErrCodeAlreadyExists Errorx = New("code already exists")

	// ErrCodeAlreadyUsed is returned when attempting to use a code that has already been used
	ErrCodeAlreadyUsed Errorx = New("code already used")

	// ErrCodeExpired is returned when attempting to use an expired verification code
	ErrCodeExpired Errorx = New("code expired")

	// ========================================
	// Employee Errors
	// ========================================

	// ErrEmployeeNotFound is returned when an employee record is not found
	ErrEmployeeNotFound Errorx = New("employee not found")

	// ========================================
	// Activity Errors
	// ========================================

	// ErrActivityNotFound is returned when an activity record is not found
	ErrActivityNotFound Errorx = New("activity not found")

	// ErrActivityInvalidAction is returned when an invalid action is provided
	ErrActivityInvalidAction Errorx = New("activity invalid action")

	// ErrActivityInvalidCauserType is returned when an invalid causer type is provided
	ErrActivityInvalidCauserType Errorx = New("activity invalid causer type")

	// ErrActivityInvalidSubjectType is returned when an invalid subject type is provided
	ErrActivityInvalidSubjectType Errorx = New("activity invalid subject type")

	// ========================================
	// Material Test Method Errors
	// ========================================

	// ErrMaterialTestMethodNotFound is returned when a material test method record is not found
	ErrMaterialTestMethodNotFound Errorx = New("material test method not found")

	// ErrMaterialTestMethodAlreadyExists is returned when attempting to create a material test method that already exists
	ErrMaterialTestMethodAlreadyExists Errorx = New("material test method already exists")

	// ! Deprecated
	// ErrMaterialTestMethodInUseDestroyForbidden is returned when attempting to destroy a material test method that has material test service, or it is in use
	// ErrMaterialTestMethodInUseDestroyForbidden Errorx = New("material test method in use destroy forbidden")

	// ========================================
	// Material Test Machine Errors
	// ========================================

	// ErrMaterialTestMachineNotFound is returned when a material test machine record is not found
	ErrMaterialTestMachineNotFound Errorx = New("material test machine not found")

	// ErrMaterialTestMachineAlreadyExists is returned when attempting to create a material test machine that already exists
	ErrMaterialTestMachineAlreadyExists Errorx = New("material test machine already exists")

	// ! Deprecated
	// ErrMaterialTestMachineInUseDestroyForbidden is returned when attempting to destroy a material test machine that has material test service, or it is in use
	// ErrMaterialTestMachineInUseDestroyForbidden Errorx = New("material test machine in use destroy forbidden")

	// ========================================
	// Material Test Service Errors
	// ========================================

	// ErrMaterialTestServiceNotFound is returned when a material test service record is not found
	ErrMaterialTestServiceNotFound Errorx = New("material test service not found")

	// ErrMaterialTestServiceAlreadyExists is returned when attempting to create a material test service that already exists
	ErrMaterialTestServiceAlreadyExists Errorx = New("material test service already exists")

	// ! Deprecated
	// ErrMaterialTestServiceInUseDestroyForbidden is returned when attempting to destroy a material test service that has material test service order, or it is in use
	// ErrMaterialTestServiceInUseDestroyForbidden Errorx = New("material test service in use destroy forbidden")

	// ========================================
	// Address Errors
	// ========================================

	// ErrAddressNotFound is returned when an address record is not found
	ErrAddressNotFound Errorx = New("address not found")

	// ErrAddressAlreadyExists is returned when attempting to create an address that already exists
	ErrAddressAlreadyExists Errorx = New("address already exists")

	// ========================================
	// Customer Errors
	// ========================================

	// ErrCustomerNotFound is returned when an customer record is not found
	ErrCustomerNotFound Errorx = New("customer not found")

	// ErrCustomerAlreadyExists is returned when attempting to create an customer that already exists
	ErrCustomerAlreadyExists Errorx = New("customer already exists")

	// ! Deprecated
	// ErrCustomerInUseDestroyForbidden is returned when attempting to destroy a customer that has material test service order, or it is in use
	// ErrCustomerInUseDestroyForbidden Errorx = New("customer in use destroy forbidden")

	// ========================================
	// Material Test Work Category Errors
	// ========================================

	// ErrMaterialTestWorkCategoryNotFound is returned when a material test work category record is not found
	ErrMaterialTestWorkCategoryNotFound Errorx = New("material test work category not found")

	// ErrMaterialTestWorkCategoryAlreadyExists is returned when attempting to create a material test work category that already exists
	ErrMaterialTestWorkCategoryAlreadyExists Errorx = New("material test work category already exists")

	// ========================================
	// Material Test Work Package Errors
	// ========================================

	// ErrMaterialTestWorkPackageNotFound is returned when a material test work package record is not found
	ErrMaterialTestWorkPackageNotFound Errorx = New("material test work package not found")

	// ErrMaterialTestWorkPackageAlreadyExists is returned when attempting to create a material test work package that already exists
	ErrMaterialTestWorkPackageAlreadyExists Errorx = New("material test work package already exists")

	// ========================================
	// Material Test Order Errors
	// ========================================

	// ErrMaterialTestOrderNotFound is returned when a material test order record is not found
	ErrMaterialTestOrderNotFound Errorx = New("material test order not found")

	// ErrMaterialTestOrderAlreadyExists is returned when attempting to create a material test order that already exists
	ErrMaterialTestOrderAlreadyExists Errorx = New("material test order already exists")

	// ========================================
	// Material Test Order Service Errors
	// ========================================

	// ErrMaterialTestOrderServiceNotFound is returned when a material test order service record is not found
	ErrMaterialTestOrderServiceNotFound Errorx = New("material test order service not found")

	// ErrMaterialTestOrderServiceAlreadyExists is returned when attempting to create a material test order service that already exists
	ErrMaterialTestOrderServiceAlreadyExists Errorx = New("material test order service already exists")

	// ========================================
	// Material Test Order Service Evaluation Errors
	// ========================================

	// ErrMaterialTestOrderServiceEvaluationNotFound is returned when a material test service order record is not found
	ErrMaterialTestOrderServiceEvaluationNotFound Errorx = New("material test order service evaluation not found")

	// ErrMaterialTestOrderServiceEvaluationAlreadyExists is returned when attempting to create a material test service order that already exists
	ErrMaterialTestOrderServiceEvaluationAlreadyExists Errorx = New("material test order service evaluation already exists")

	// ========================================
	// Media Errors
	// ========================================

	// ErrMediaNotFound is returned when a media record is not found
	ErrMediaNotFound Errorx = New("media not found")

	// ErrMediaAlreadyExists is returned when attempting to create a media that already exists
	ErrMediaAlreadyExists Errorx = New("media already exists")
)
