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

type NewMaterialTestOrderServiceEvaluationFactoryParams struct {
	fx.In

	IdService id.IdService
}

func NewMaterialTestOrderServiceEvaluationFactory(params NewMaterialTestOrderServiceEvaluationFactoryParams) *factory.Factory {
	idService := params.IdService

	return factory.NewFactory(&entity.MaterialTestOrderServiceEvaluation{}).
		Attr("Id", func(args factory.Args) (any, error) {
			return idService.Generate(), nil
		}).
		Attr("IsEquipmentAvailable", func(args factory.Args) (any, error) {
			isAvailable := gofakeit.Bool()
			return isAvailable, nil
		}).
		Attr("IsPersonnelAvailable", func(args factory.Args) (any, error) {
			isAvailable := gofakeit.Bool()
			return isAvailable, nil
		}).
		Attr("IsTimeAvailable", func(args factory.Args) (any, error) {
			isAvailable := gofakeit.Bool()
			return isAvailable, nil
		}).
		Attr("IsTestReady", func(args factory.Args) (any, error) {
			isAvailable := gofakeit.Bool()
			return isAvailable, nil
		}).
		Attr("IsSubcontractLabAvailable", func(args factory.Args) (any, error) {
			isAvailable := gofakeit.Bool()
			return isAvailable, nil
		}).
		Attr("CreatedAt", func(args factory.Args) (any, error) {
			return gofakeit.DateRange(time.Now().Add(-time.Hour*24*365), time.Now()), nil
		}).
		Attr("UpdatedAt", func(args factory.Args) (any, error) {
			createdAt := args.Instance().(*entity.MaterialTestOrderServiceEvaluation).CreatedAt

			isUpdated := gofakeit.Bool()
			if isUpdated {
				return typeutil.Ptr(gofakeit.DateRange(createdAt, time.Now())), nil
			}

			return nil, nil
		})
}
