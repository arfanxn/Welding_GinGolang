package dto

import (
	activityEnum "github.com/arfanxn/welding/internal/module/activity/domain/enum"
	"github.com/arfanxn/welding/pkg/types"
)

type CreateActivity struct {
	CauserId   *string
	CauserType *activityEnum.ActivityCauserType
	Causer     any

	CauserIpAddress *string

	Action activityEnum.ActivityAction

	SubjectId   *string
	SubjectType *activityEnum.ActivitySubjectType
	Subject     any

	Properties types.JSONMap
}
