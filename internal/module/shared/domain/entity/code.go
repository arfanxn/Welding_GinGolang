package entity

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/arfanxn/welding/internal/module/code/domain/enum"
	"github.com/arfanxn/welding/pkg/numberutil"
	"github.com/gookit/goutil"
	"github.com/guregu/null/v6"
	"github.com/oklog/ulid/v2"
	"gorm.io/datatypes"
	"gorm.io/gorm"
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

func (c *Code) BeforeSave(tx *gorm.DB) error {

	if goutil.IsZero(c.Id) {
		c.Id = ulid.Make().String()
	}

	if goutil.IsZero(c.Value) {
		c.Value = strconv.Itoa(numberutil.Random(100000, 999999))
	}

	return nil
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

func (c *Code) IsUsed() bool {
	return c.UsedAt.Valid
}

func (c *Code) IsExpired() bool {
	return c.ExpiredAt.Before(time.Now())
}
