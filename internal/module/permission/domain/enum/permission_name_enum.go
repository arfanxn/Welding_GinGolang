package enum

type PermissionName string

const (
	UserRead       PermissionName = "user.read"
	UserCreate     PermissionName = "user.create"
	UserUpdate     PermissionName = "user.update"
	UserDelete     PermissionName = "user.delete"
	RoleRead       PermissionName = "role.read"
	RoleCreate     PermissionName = "role.create"
	RoleUpdate     PermissionName = "role.update"
	RoleDelete     PermissionName = "role.delete"
	PermissionRead PermissionName = "permission.read"
)

func (p PermissionName) String() string {
	return string(p)
}

var PermissionNames = []PermissionName{
	UserRead,
	UserCreate,
	UserUpdate,
	UserDelete,
	RoleRead,
	RoleCreate,
	RoleUpdate,
	RoleDelete,
	PermissionRead,
}
