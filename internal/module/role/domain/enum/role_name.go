package enum

type RoleName = string

const (
	SuperAdmin           RoleName = "super_admin"
	CustomerServiceAdmin RoleName = "customer_service_admin"
	Customer             RoleName = "customer"
)

var RoleNames = []RoleName{
	SuperAdmin,
	CustomerServiceAdmin,
	Customer,
}

const DefaultRoleName = Customer
