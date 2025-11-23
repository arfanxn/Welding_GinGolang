package enum

type ActivityCauserType = string

const (
	UserCauserType   ActivityCauserType = "user"
	EmailCauserType  ActivityCauserType = "email"
	GuestCauserType  ActivityCauserType = "guest"
	SystemCauserType ActivityCauserType = "system"
)

var ActivityCauserTypes = []ActivityCauserType{
	UserCauserType,
	EmailCauserType,
	GuestCauserType,
	SystemCauserType,
}
