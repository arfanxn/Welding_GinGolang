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
)
