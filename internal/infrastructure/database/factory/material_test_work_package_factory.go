package factory

import (
	"time"

	id "github.com/arfanxn/welding/internal/infrastructure/id"
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	"github.com/arfanxn/welding/pkg/boolutil"
	"github.com/bluele/factory-go/factory"
	"github.com/brianvoe/gofakeit/v7"
	"go.uber.org/fx"
)

type NewMaterialTestWorkPackageFactoryParams struct {
	fx.In

	IdService id.IdService
}

func NewMaterialTestWorkPackageFactory(params NewMaterialTestWorkPackageFactoryParams) *factory.Factory {
	idService := params.IdService

	return factory.NewFactory(&entity.MaterialTestWorkPackage{}).
		Attr("Id", func(args factory.Args) (any, error) {
			return idService.Generate(), nil
		}).
		Attr("Name", func(args factory.Args) (any, error) {
			name := gofakeit.Word()
			return name, nil
		}).
		Attr("Description", func(args factory.Args) (any, error) {
			description := gofakeit.Sentence()
			return &description, nil
		}).
		Attr("CreatedAt", func(args factory.Args) (any, error) {
			createdAt := gofakeit.DateRange(time.Now().Add(-time.Hour*24*365), time.Now())
			return createdAt, nil
		}).
		Attr("UpdatedAt", func(args factory.Args) (any, error) {
			createdAt := args.Instance().(*entity.MaterialTestWorkPackage).CreatedAt
			updatedAt := gofakeit.DateRange(createdAt, time.Now())
			isUpdated := gofakeit.Bool()
			return boolutil.Ternary(isUpdated, &updatedAt, nil), nil
		})
}
