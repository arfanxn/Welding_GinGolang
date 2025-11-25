package viewmodel

import (
	"time"

	activityEnum "github.com/arfanxn/welding/internal/module/activity/domain/enum"
	"github.com/arfanxn/welding/pkg/types"
)

type ActivityViewModel struct {
	Id              string                      `json:"id"`
	CauserId        *string                     `json:"causer_id"`
	CauserType      *activityEnum.CauserType    `json:"causer_type"`
	Causer          any                         `json:"causer,omitempty"`
	CauserIpAddress *string                     `json:"causer_ip_address"`
	Action          activityEnum.ActivityAction `json:"action"`
	Description     string                      `json:"description"`
	SubjectId       *string                     `json:"subject_id"`
	SubjectType     *activityEnum.SubjectType   `json:"subject_type"`
	Subject         any                         `json:"subject,omitempty"`
	Properties      types.JSONMap               `json:"properties"`
	CreatedAt       time.Time                   `json:"created_at"`
	UpdatedAt       *time.Time                  `json:"updated_at"`
}
