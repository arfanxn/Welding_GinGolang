package enum

type PermissionName string

const (
	UsersIndex   PermissionName = "users.index"
	UsersShow    PermissionName = "users.show"
	UsersStore   PermissionName = "users.store"
	UsersUpdate  PermissionName = "users.update"
	UsersDestroy PermissionName = "users.destroy"

	RolesIndex   PermissionName = "roles.index"
	RolesShow    PermissionName = "roles.show"
	RolesStore   PermissionName = "roles.store"
	RolesUpdate  PermissionName = "roles.update"
	RolesDestroy PermissionName = "roles.destroy"

	PermissionsIndex PermissionName = "permissions.index"
	PermissionsShow  PermissionName = "permissions.show"
)

func (p PermissionName) String() string {
	return string(p)
}

var PermissionNames = []PermissionName{
	UsersIndex,
	UsersShow,
	UsersStore,
	UsersUpdate,
	UsersDestroy,

	RolesIndex,
	RolesShow,
	RolesStore,
	RolesUpdate,
	RolesDestroy,

	PermissionsIndex,
	PermissionsShow,
}
