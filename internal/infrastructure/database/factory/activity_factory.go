package factory

import (
	"time"

	id "github.com/arfanxn/welding/internal/infrastructure/id"
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	"github.com/bluele/factory-go/factory"
	"github.com/brianvoe/gofakeit/v7"
)

func NewActivityFactory(
	idService id.IdService,
) *factory.Factory {
	return factory.NewFactory(&entity.Activity{}).
		Attr("Id", func(args factory.Args) (any, error) {
			return idService.Generate(), nil
		}).
		Attr("CauserIpAddress", func(args factory.Args) (any, error) {
			ipAddr := gofakeit.IPv4Address()
			return &ipAddr, nil
		}).
		Attr("CreatedAt", func(args factory.Args) (any, error) {
			return gofakeit.DateRange(time.Now().Add(-time.Hour*24*365), time.Now()), nil
		}).
		Attr("UpdatedAt", func(args factory.Args) (any, error) {
			return nil, nil
		})

}
