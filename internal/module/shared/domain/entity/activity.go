package entity

import (
	"time"

	activityEnum "github.com/arfanxn/welding/internal/module/activity/domain/enum"
	"github.com/arfanxn/welding/pkg/types"
)

type Activity struct {
	Id              string                            `json:"id" gorm:"primaryKey"`
	CauserId        *string                           `json:"causer_id"`
	CauserType      *activityEnum.ActivityCauserType  `json:"causer_type"`
	CauserIpAddress *string                           `json:"causer_ip_address"`
	Action          activityEnum.ActivityAction       `json:"action"`
	SubjectId       *string                           `json:"subject_id"`
	SubjectType     *activityEnum.ActivitySubjectType `json:"subject_type"`
	Properties      types.JSONMap                     `json:"properties" gorm:"type:jsonb"`
	CreatedAt       time.Time                         `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt       *time.Time                        `json:"updated_at" gorm:"autoUpdateTime"`

	CauserUser *User `json:"causer_user,omitempty" gorm:"foreignKey:CauserId;references:Id;"`

	SubjectUser       *User       `json:"subject_user,omitempty" gorm:"foreignKey:SubjectId;references:Id;"`
	SubjectRole       *Role       `json:"subject_role,omitempty" gorm:"foreignKey:SubjectId;references:Id;"`
	SubjectPermission *Permission `json:"subject_permission,omitempty" gorm:"foreignKey:SubjectId;references:Id;"`
	SubjectCode       *Code       `json:"subject_code,omitempty" gorm:"foreignKey:SubjectId;references:Id;"`
}

func NewActivity() *Activity {
	return &Activity{}
}

func (a *Activity) TableName() string {
	return "activities"
}
