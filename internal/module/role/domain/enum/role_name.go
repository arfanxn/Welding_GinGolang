package enum

type RoleName string

func (p RoleName) String() string {
	return string(p)
}

const (
	SuperAdmin           RoleName = "super_admin"
	Admin                RoleName = "admin"
	Head                 RoleName = "head"
	Manager              RoleName = "manager"
	Supervisor           RoleName = "supervisor"
	Engineer             RoleName = "engineer"
	Staff                RoleName = "staff"
	Operator             RoleName = "operator"
	CustomerServiceAdmin RoleName = "customer_service_admin"
	Customer             RoleName = "customer"
)

var RoleNames = []RoleName{
	SuperAdmin,
	Admin,
	Head,
	Manager,
	Supervisor,
	Engineer,
	Staff,
	Operator,
	CustomerServiceAdmin,
	Customer,
}

const DefaultRoleName = Customer
