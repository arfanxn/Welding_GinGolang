package factory

import (
	"time"

	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	"github.com/bluele/factory-go/factory"
	"github.com/guregu/null/v6"
	"github.com/oklog/ulid/v2"
)

var RoleFactory = factory.NewFactory(&entity.Role{}).
	Attr("Id", func(args factory.Args) (any, error) {
		return ulid.Make().String(), nil
	}).
	Attr("IsDefault", func(args factory.Args) (any, error) {
		return false, nil
	}).
	Attr("CreatedAt", func(args factory.Args) (any, error) {
		return time.Now(), nil
	}).
	Attr("UpdatedAt", func(args factory.Args) (any, error) {
		return null.TimeFromPtr(nil), nil
	})
