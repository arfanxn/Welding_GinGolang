package enum

type ActivityAction = string

const (
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

	PermissionsIndex ActivityAction = "permissions.index"

	RolesIndex      ActivityAction = "roles.index"
	RolesShow       ActivityAction = "roles.show"
	RolesStore      ActivityAction = "roles.store"
	RolesUpdate     ActivityAction = "roles.update"
	RolesSetDefault ActivityAction = "roles.set_default"
	RolesDestroy    ActivityAction = "roles.destroy"

	CodesCreateUserRegisterInvitation ActivityAction = "codes.create_user_register_invitation"
	CodesCreateUserEmailVerification  ActivityAction = "codes.create_user_email_verification"
	CodesCreateUserResetPassword      ActivityAction = "codes.create_user_reset_password"

	ActivitiesIndex ActivityAction = "activities.index"
	ActivitiesShow  ActivityAction = "activities.show"

	MaterialTestMethodsIndex   ActivityAction = "material_test_methods.index"
	MaterialTestMethodsShow    ActivityAction = "material_test_methods.show"
	MaterialTestMethodsStore   ActivityAction = "material_test_methods.store"
	MaterialTestMethodsUpdate  ActivityAction = "material_test_methods.update"
	MaterialTestMethodsDestroy ActivityAction = "material_test_methods.destroy"

	MaterialTestMachinesIndex   ActivityAction = "material_test_machines.index"
	MaterialTestMachinesShow    ActivityAction = "material_test_machines.show"
	MaterialTestMachinesStore   ActivityAction = "material_test_machines.store"
	MaterialTestMachinesUpdate  ActivityAction = "material_test_machines.update"
	MaterialTestMachinesDestroy ActivityAction = "material_test_machines.destroy"
)

var ActivityActions = []ActivityAction{
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

	PermissionsIndex,

	RolesIndex,
	RolesShow,
	RolesStore,
	RolesUpdate,
	RolesSetDefault,
	RolesDestroy,

	CodesCreateUserRegisterInvitation,
	CodesCreateUserEmailVerification,
	CodesCreateUserResetPassword,

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
}
