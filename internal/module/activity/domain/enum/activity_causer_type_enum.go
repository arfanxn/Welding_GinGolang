package enum

type CauserType = string

const (
	UserCauserType   CauserType = "user"
	EmailCauserType  CauserType = "email"
	GuestCauserType  CauserType = "guest"
	SystemCauserType CauserType = "system"
)

var CauserTypes = []CauserType{
	UserCauserType,
	EmailCauserType,
	GuestCauserType,
	SystemCauserType,
}
