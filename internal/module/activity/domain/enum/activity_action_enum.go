package enum

type ActivityAction = string

const (
	// Users
	UsersRegister         ActivityAction = "users.register"
	UsersResetPassword    ActivityAction = "users.reset_password"
	UsersVerifyEmail      ActivityAction = "users.verify_email"
	UsersLogin            ActivityAction = "users.login"
	UsersLogout           ActivityAction = "users.logout"
	UsersMe               ActivityAction = "users.me"
	UsersUpdateMeProfile  ActivityAction = "users.update_me_profile"
	UsersUpdateMePassword ActivityAction = "users.update_me_password"
	UsersIndex            ActivityAction = "users.index"
	UsersShow             ActivityAction = "users.show"
	UsersStore            ActivityAction = "users.store"
	UsersUpdate           ActivityAction = "users.update"
	UsersToggleActivation ActivityAction = "users.toggle_activation"
	UsersDestroy          ActivityAction = "users.destroy"

	// Permissions
	PermissionsIndex ActivityAction = "permissions.index"

	// Roles
	RolesIndex      ActivityAction = "roles.index"
	RolesShow       ActivityAction = "roles.show"
	RolesStore      ActivityAction = "roles.store"
	RolesUpdate     ActivityAction = "roles.update"
	RolesSetDefault ActivityAction = "roles.set_default"
	RolesDestroy    ActivityAction = "roles.destroy"

	// Codes
	CodesCreateUserRegisterInvitation ActivityAction = "codes.create_user_register_invitation"
	CodesCreateUserEmailVerification  ActivityAction = "codes.create_user_email_verification"
	CodesCreateUserResetPassword      ActivityAction = "codes.create_user_reset_password"

	// Activities
	ActivitiesIndex ActivityAction = "activities.index"
	ActivitiesShow  ActivityAction = "activities.show"

	// Address
	AddressesIndex   ActivityAction = "addresses.index"
	AddressesShow    ActivityAction = "addresses.show"
	AddressesStore   ActivityAction = "addresses.store"
	AddressesUpdate  ActivityAction = "addresses.update"
	AddressesDestroy ActivityAction = "addresses.destroy"

	// Customers
	CustomersIndex   ActivityAction = "customers.index"
	CustomersShow    ActivityAction = "customers.show"
	CustomersStore   ActivityAction = "customers.store"
	CustomersUpdate  ActivityAction = "customers.update"
	CustomersDestroy ActivityAction = "customers.destroy"

	// Material test methods
	MaterialTestMethodsIndex   ActivityAction = "material_test_methods.index"
	MaterialTestMethodsShow    ActivityAction = "material_test_methods.show"
	MaterialTestMethodsStore   ActivityAction = "material_test_methods.store"
	MaterialTestMethodsUpdate  ActivityAction = "material_test_methods.update"
	MaterialTestMethodsDestroy ActivityAction = "material_test_methods.destroy"

	// Material test machines
	MaterialTestMachinesIndex   ActivityAction = "material_test_machines.index"
	MaterialTestMachinesShow    ActivityAction = "material_test_machines.show"
	MaterialTestMachinesStore   ActivityAction = "material_test_machines.store"
	MaterialTestMachinesUpdate  ActivityAction = "material_test_machines.update"
	MaterialTestMachinesDestroy ActivityAction = "material_test_machines.destroy"

	// Material test services
	MaterialTestServicesIndex   ActivityAction = "material_test_services.index"
	MaterialTestServicesShow    ActivityAction = "material_test_services.show"
	MaterialTestServicesStore   ActivityAction = "material_test_services.store"
	MaterialTestServicesUpdate  ActivityAction = "material_test_services.update"
	MaterialTestServicesDestroy ActivityAction = "material_test_services.destroy"

	// Material work categories
	MaterialTestWorkCategoriesIndex   ActivityAction = "material_test_work_categories.index"
	MaterialTestWorkCategoriesShow    ActivityAction = "material_test_work_categories.show"
	MaterialTestWorkCategoriesStore   ActivityAction = "material_test_work_categories.store"
	MaterialTestWorkCategoriesUpdate  ActivityAction = "material_test_work_categories.update"
	MaterialTestWorkCategoriesDestroy ActivityAction = "material_test_work_categories.destroy"

	// Material work packages
	MaterialTestWorkPackagesIndex   ActivityAction = "material_test_work_packages.index"
	MaterialTestWorkPackagesShow    ActivityAction = "material_test_work_packages.show"
	MaterialTestWorkPackagesStore   ActivityAction = "material_test_work_packages.store"
	MaterialTestWorkPackagesUpdate  ActivityAction = "material_test_work_packages.update"
	MaterialTestWorkPackagesDestroy ActivityAction = "material_test_work_packages.destroy"

	// Material test orders
	MaterialTestOrdersIndex   ActivityAction = "material_test_orders.index"
	MaterialTestOrdersShow    ActivityAction = "material_test_orders.show"
	MaterialTestOrdersStore   ActivityAction = "material_test_orders.store"
	MaterialTestOrdersUpdate  ActivityAction = "material_test_orders.update"
	MaterialTestOrdersDestroy ActivityAction = "material_test_orders.destroy"
)

var ActivityActions = []ActivityAction{
	// Users
	UsersRegister,
	UsersResetPassword,
	UsersVerifyEmail,
	UsersLogin,
	UsersLogout,
	UsersMe,
	UsersUpdateMeProfile,
	UsersUpdateMePassword,
	UsersIndex,
	UsersShow,
	UsersStore,
	UsersUpdate,
	UsersToggleActivation,
	UsersDestroy,

	// Permissions
	PermissionsIndex,

	// Roles
	RolesIndex,
	RolesShow,
	RolesStore,
	RolesUpdate,
	RolesSetDefault,
	RolesDestroy,

	// Codes
	CodesCreateUserRegisterInvitation,
	CodesCreateUserEmailVerification,
	CodesCreateUserResetPassword,

	// Activities
	ActivitiesIndex,
	ActivitiesShow,

	// Addresses
	AddressesIndex,
	AddressesShow,
	AddressesStore,
	AddressesUpdate,
	AddressesDestroy,

	// Customers
	CustomersIndex,
	CustomersShow,
	CustomersStore,
	CustomersUpdate,
	CustomersDestroy,

	// Material test methods
	MaterialTestMethodsIndex,
	MaterialTestMethodsShow,
	MaterialTestMethodsStore,
	MaterialTestMethodsUpdate,
	MaterialTestMethodsDestroy,

	// Material test machines
	MaterialTestMachinesIndex,
	MaterialTestMachinesShow,
	MaterialTestMachinesStore,
	MaterialTestMachinesUpdate,
	MaterialTestMachinesDestroy,

	// Material test services
	MaterialTestServicesIndex,
	MaterialTestServicesShow,
	MaterialTestServicesStore,
	MaterialTestServicesUpdate,
	MaterialTestServicesDestroy,

	// Material test work categories
	MaterialTestWorkCategoriesIndex,
	MaterialTestWorkCategoriesShow,
	MaterialTestWorkCategoriesStore,
	MaterialTestWorkCategoriesUpdate,
	MaterialTestWorkCategoriesDestroy,

	// Material test work packages
	MaterialTestWorkPackagesIndex,
	MaterialTestWorkPackagesShow,
	MaterialTestWorkPackagesStore,
	MaterialTestWorkPackagesUpdate,
	MaterialTestWorkPackagesDestroy,

	// Material test orders
	MaterialTestOrdersIndex,
	MaterialTestOrdersShow,
	MaterialTestOrdersStore,
	MaterialTestOrdersUpdate,
	MaterialTestOrdersDestroy,
}
