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

	ActivitiesIndex PermissionName = "activities.index"
	ActivitiesShow  PermissionName = "activities.show"

	MaterialTestMethodsIndex   PermissionName = "material_test_methods.index"
	MaterialTestMethodsShow    PermissionName = "material_test_methods.show"
	MaterialTestMethodsStore   PermissionName = "material_test_methods.store"
	MaterialTestMethodsUpdate  PermissionName = "material_test_methods.update"
	MaterialTestMethodsDestroy PermissionName = "material_test_methods.destroy"

	MaterialTestMachinesIndex   PermissionName = "material_test_machines.index"
	MaterialTestMachinesShow    PermissionName = "material_test_machines.show"
	MaterialTestMachinesStore   PermissionName = "material_test_machines.store"
	MaterialTestMachinesUpdate  PermissionName = "material_test_machines.update"
	MaterialTestMachinesDestroy PermissionName = "material_test_machines.destroy"

	MaterialTestServicesIndex   PermissionName = "material_test_services.index"
	MaterialTestServicesShow    PermissionName = "material_test_services.show"
	MaterialTestServicesStore   PermissionName = "material_test_services.store"
	MaterialTestServicesUpdate  PermissionName = "material_test_services.update"
	MaterialTestServicesDestroy PermissionName = "material_test_services.destroy"
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

	ActivitiesIndex,
	ActivitiesShow,

	MaterialTestMethodsIndex,
	MaterialTestMethodsShow,
	MaterialTestMethodsStore,
	MaterialTestMethodsUpdate,
	MaterialTestMethodsDestroy,

	MaterialTestMachinesIndex,
	MaterialTestMachinesShow,
	MaterialTestMachinesStore,
	MaterialTestMachinesUpdate,
	MaterialTestMachinesDestroy,

	MaterialTestServicesIndex,
	MaterialTestServicesShow,
	MaterialTestServicesStore,
	MaterialTestServicesUpdate,
	MaterialTestServicesDestroy,
}
