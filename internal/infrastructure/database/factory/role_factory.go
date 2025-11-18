package factory

import (
	"time"

	id "github.com/arfanxn/welding/internal/infrastructure/id"
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	"github.com/bluele/factory-go/factory"
	"github.com/guregu/null/v6"
)

func NewRoleFactory(
	idService id.IdService,
) *factory.Factory {
	return factory.NewFactory(&entity.Role{}).
		Attr("Id", func(args factory.Args) (any, error) {
			return idService.Generate(), nil
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

}
