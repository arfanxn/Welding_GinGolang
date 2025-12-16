package factory

import (
	"fmt"
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
		Attr("SampleNumber", func(args factory.Args) (any, error) {
			now := time.Now()
			sampledAt := gofakeit.DateRange(now, now.AddDate(0, 0, 7))

			serviceCodes := []string{"UTK", "CNC", "CAL"}
			serviceCode := serviceCodes[gofakeit.IntRange(0, len(serviceCodes)-1)]
			sampleNumber := fmt.Sprintf("%02d/%d.%.1f/%s/%03d", // eg: 01/2025.1/UTK/001
				sampledAt.Day(),
				sampledAt.Year(),
				float64(sampledAt.Month())/1,
				serviceCode,
				gofakeit.Number(1, 999),
			)

			return sampleNumber, nil
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
