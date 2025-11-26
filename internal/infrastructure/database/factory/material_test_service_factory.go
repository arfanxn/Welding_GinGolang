package factory

import (
	"time"

	id "github.com/arfanxn/welding/internal/infrastructure/id"
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	"github.com/arfanxn/welding/pkg/typeutil"
	"github.com/bluele/factory-go/factory"
	"github.com/brianvoe/gofakeit/v7"
	"go.uber.org/fx"
)

type NewMaterialTestServiceFactoryParams struct {
	fx.In

	IdService id.IdService
}

func NewMaterialTestServiceFactory(params NewMaterialTestServiceFactoryParams) *factory.Factory {
	idService := params.IdService

	return factory.NewFactory(&entity.MaterialTestService{}).
		Attr("Id", func(args factory.Args) (any, error) {
			return idService.Generate(), nil
		}).
		Attr("ServiceType", func(args factory.Args) (any, error) {
			serviceTypes := []string{"Testing", "Machining"}
			serviceTypeIndex := gofakeit.IntRange(0, len(serviceTypes)-1)
			serviceType := serviceTypes[serviceTypeIndex]
			return serviceType, nil
		}).
		Attr("ServiceCode", func(args factory.Args) (any, error) {
			serviceCodes := []string{"UTK", "CNC"}
			serviceCodeIndex := gofakeit.IntRange(0, len(serviceCodes)-1)
			serviceCode := serviceCodes[serviceCodeIndex]
			return serviceCode, nil
		}).
		Attr("Unit", func(args factory.Args) (any, error) {
			units := []string{"sample", "jam"}
			unitIndex := gofakeit.IntRange(0, len(units)-1)
			unit := units[unitIndex]
			return unit, nil
		}).
		Attr("Price", func(args factory.Args) (any, error) {
			return gofakeit.Float64Range(10000, 99999), nil
		}).
		Attr("CreatedAt", func(args factory.Args) (any, error) {
			return gofakeit.DateRange(time.Now().Add(-time.Hour*24*365), time.Now()), nil
		}).
		Attr("UpdatedAt", func(args factory.Args) (any, error) {
			createdAt := args.Instance().(*entity.MaterialTestService).CreatedAt

			isUpdated := gofakeit.Bool()
			if isUpdated {
				return typeutil.Ptr(gofakeit.DateRange(createdAt, time.Now())), nil
			}

			return nil, nil
		})
}
