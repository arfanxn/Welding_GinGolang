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
	// Database Errors
	// ========================================

	// ErrRecordNotFound is returned when a database record is not found
	ErrRecordNotFound Errorx = New("record not found")

	// ErrDuplicatedKey is returned when attempting to insert a duplicate key
	ErrDuplicatedKey Errorx = New("duplicated key")

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

	// ErrUserSuperAdminCannotBeModified is returned when attempting to modify a super admin user
	ErrUserSuperAdminCannotBeModified Errorx = New("user super admin cannot be modified")

	// ErrUserSuperAdminRoleCannotBeChanged is returned when attempting to change a super admin's role
	ErrUserSuperAdminRoleCannotBeChanged Errorx = New("user super admin role cannot be changed")

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

	// ErrRoleDefaultNotConfigured is returned when the system default role is not configured
	ErrRoleDefaultNotConfigured Errorx = New("role default not configured")

	// ========================================
	// Permission Errors
	// ========================================

	// ErrPermissionNotFound is returned when a specific permission is not found
	ErrPermissionNotFound Errorx = New("permission not found")

	// ErrPermissionAlreadyExists is returned when attempting to create a permission that already exists
	ErrPermissionAlreadyExists Errorx = New("permission already exists")

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
