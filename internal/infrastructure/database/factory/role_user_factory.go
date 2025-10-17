package factory

import (
	"time"

	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	"github.com/bluele/factory-go/factory"
	"github.com/guregu/null/v6"
)

var RoleUserFactory = factory.NewFactory(&entity.RoleUser{}).
	Attr("CreatedAt", func(args factory.Args) (any, error) {
		return time.Now(), nil
	}).
	Attr("UpdatedAt", func(args factory.Args) (any, error) {
		return null.TimeFromPtr(nil), nil
	})
