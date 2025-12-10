package factory

import (
	"time"

	id "github.com/arfanxn/welding/internal/infrastructure/id"
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	"github.com/arfanxn/welding/pkg/boolutil"
	"github.com/bluele/factory-go/factory"
	"github.com/brianvoe/gofakeit/v7"
)

func NewAddressFactory(
	idService id.IdService,
) *factory.Factory {
	return factory.NewFactory(&entity.Address{}).
		Attr("Id", func(args factory.Args) (any, error) {
			return idService.Generate(), nil
		}).
		Attr("FullAddress", func(args factory.Args) (any, error) {
			addrInfo := gofakeit.Address()
			fullAddr := addrInfo.Address
			return fullAddr, nil
		}).
		Attr("CreatedAt", func(args factory.Args) (any, error) {
			return time.Now(), nil
		}).
		Attr("UpdatedAt", func(args factory.Args) (any, error) {
			createdAt := args.Instance().(*entity.Address).CreatedAt
			updatedAt := gofakeit.DateRange(createdAt, time.Now())

			isUpdated := gofakeit.Bool()

			return boolutil.Ternary(isUpdated, &updatedAt, nil), nil
		})

}
