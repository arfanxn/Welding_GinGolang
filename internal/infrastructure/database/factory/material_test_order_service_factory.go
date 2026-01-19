package factory

import (
	"time"

	id "github.com/arfanxn/welding/internal/infrastructure/id"
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	"github.com/arfanxn/welding/pkg/boolutil"
	"github.com/bluele/factory-go/factory"
	"github.com/brianvoe/gofakeit/v7"
)

func NewMaterialTestOrderServiceFactory(
	idService id.IdService,
) *factory.Factory {
	return factory.NewFactory(&entity.MaterialTestOrderService{}).
		Attr("Id", func(args factory.Args) (any, error) {
			return idService.Generate(), nil
		}).
		Attr("SampleName", func(args factory.Args) (any, error) {
			return gofakeit.ProductMaterial(), nil
		}).
		Attr("Quantity", func(args factory.Args) (any, error) {
			return gofakeit.Number(1, 100), nil
		}).
		Attr("CreatedAt", func(args factory.Args) (any, error) {
			return time.Now(), nil
		}).
		Attr("UpdatedAt", func(args factory.Args) (any, error) {
			createdAt := args.Instance().(*entity.MaterialTestOrderService).CreatedAt
			updatedAt := gofakeit.DateRange(createdAt, createdAt.AddDate(0, 0, 2))
			isUpdated := gofakeit.Bool()
			return boolutil.Ternary(isUpdated, &updatedAt, nil), nil
		})
}
