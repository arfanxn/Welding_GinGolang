package enum

type PermissionName = string

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

	ActivitiesIndex PermissionName = "activities.index"
	ActivitiesShow  PermissionName = "activities.show"

	AddressesIndex   PermissionName = "addresses.index"
	AddressesShow    PermissionName = "addresses.show"
	AddressesStore   PermissionName = "addresses.store"
	AddressesUpdate  PermissionName = "addresses.update"
	AddressesDestroy PermissionName = "addresses.destroy"

	CustomersIndex   PermissionName = "customers.index"
	CustomersShow    PermissionName = "customers.show"
	CustomersStore   PermissionName = "customers.store"
	CustomersUpdate  PermissionName = "customers.update"
	CustomersDestroy PermissionName = "customers.destroy"

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

	MaterialTestWorkCategoriesIndex   PermissionName = "material_test_work_categories.index"
	MaterialTestWorkCategoriesShow    PermissionName = "material_test_work_categories.show"
	MaterialTestWorkCategoriesStore   PermissionName = "material_test_work_categories.store"
	MaterialTestWorkCategoriesUpdate  PermissionName = "material_test_work_categories.update"
	MaterialTestWorkCategoriesDestroy PermissionName = "material_test_work_categories.destroy"

	MaterialTestWorkPackagesIndex   PermissionName = "material_test_work_packages.index"
	MaterialTestWorkPackagesShow    PermissionName = "material_test_work_packages.show"
	MaterialTestWorkPackagesStore   PermissionName = "material_test_work_packages.store"
	MaterialTestWorkPackagesUpdate  PermissionName = "material_test_work_packages.update"
	MaterialTestWorkPackagesDestroy PermissionName = "material_test_work_packages.destroy"

	MaterialTestOrdersIndex  PermissionName = "material_test_orders.index"
	MaterialTestOrdersShow   PermissionName = "material_test_orders.show"
	MaterialTestOrdersStore  PermissionName = "material_test_orders.store"
	MaterialTestOrdersUpdate PermissionName = "material_test_orders.update"

	AnalyticsIndex PermissionName = "analytics.index"
)

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

	ActivitiesIndex,
	ActivitiesShow,

	AddressesIndex,
	AddressesShow,
	AddressesStore,
	AddressesUpdate,
	AddressesDestroy,

	CustomersIndex,
	CustomersShow,
	CustomersStore,
	CustomersUpdate,
	CustomersDestroy,

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

	MaterialTestWorkCategoriesIndex,
	MaterialTestWorkCategoriesShow,
	MaterialTestWorkCategoriesStore,
	MaterialTestWorkCategoriesUpdate,
	MaterialTestWorkCategoriesDestroy,

	MaterialTestWorkPackagesIndex,
	MaterialTestWorkPackagesShow,
	MaterialTestWorkPackagesStore,
	MaterialTestWorkPackagesUpdate,
	MaterialTestWorkPackagesDestroy,

	MaterialTestOrdersIndex,
	MaterialTestOrdersShow,
	MaterialTestOrdersStore,
	MaterialTestOrdersUpdate,

	AnalyticsIndex,
}
