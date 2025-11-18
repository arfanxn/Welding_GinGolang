package entity

import (
	"encoding/json"
	"time"

	"github.com/arfanxn/welding/internal/module/code/domain/enum"
	"github.com/guregu/null/v6"
	"gorm.io/datatypes"
)

type Code struct {
	Id           string         `json:"id" gorm:"primaryKey"`
	CodeableId   null.String    `json:"codeable_id,omitzero"`
	CodeableType null.String    `json:"codeable_type,omitzero"`
	Type         enum.CodeType  `json:"type" gorm:"type:code_type_enum"`
	Value        string         `json:"value"`
	Meta         datatypes.JSON `json:"meta" gorm:"type:jsonb"`
	UsedAt       null.Time      `json:"used_at"`
	ExpiredAt    time.Time      `json:"expired_at"`
	CreatedAt    time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    null.Time      `json:"updated_at" gorm:"autoUpdateTime"`
}

// TableName specifies the table name for the Code model
func (Code) TableName() string {
	return "codes"
}

func (c *Code) GetMeta() (map[string]any, error) {
	var meta map[string]any
	err := json.Unmarshal(c.Meta, &meta)
	if err != nil {
		return nil, err
	}
	return meta, nil
}

func (c *Code) SetMeta(meta map[string]any) error {
	if meta == nil {
		c.Meta = nil
		return nil
	}

	jsonBytes, err := json.Marshal(meta)
	if err != nil {
		return err
	}
	c.Meta = datatypes.JSON(jsonBytes)
	return nil
}

func (c *Code) MarkUsed() {
	c.UsedAt = null.TimeFrom(time.Now())
}

func (c *Code) IsUsed() bool {
	return c.UsedAt.Valid
}

func (c *Code) IsExpired() bool {
	return c.ExpiredAt.Before(time.Now())
}
