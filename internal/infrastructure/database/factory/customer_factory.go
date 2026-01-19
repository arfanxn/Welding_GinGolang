package factory

import (
	"time"

	id "github.com/arfanxn/welding/internal/infrastructure/id"
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	"github.com/arfanxn/welding/pkg/boolutil"
	"github.com/bluele/factory-go/factory"
	"github.com/brianvoe/gofakeit/v7"
)

func NewCustomerFactory(
	idService id.IdService,
) *factory.Factory {
	return factory.NewFactory(&entity.Customer{}).
		Attr("Id", func(args factory.Args) (any, error) {
			return idService.Generate(), nil
		}).
		Attr("Name", func(args factory.Args) (any, error) {
			isIndividual := gofakeit.Bool()
			return boolutil.Ternary(isIndividual, gofakeit.Name(), gofakeit.Company()), nil
		}).
		Attr("PhoneNumber", func(args factory.Args) (any, error) {
			return gofakeit.Phone(), nil
		}).
		Attr("Email", func(args factory.Args) (any, error) {
			return gofakeit.Email(), nil
		}).
		Attr("CreatedAt", func(args factory.Args) (any, error) {
			createdAt := gofakeit.DateRange(time.Now().AddDate(-1, 0, 0), time.Now())
			return createdAt, nil
		}).
		Attr("UpdatedAt", func(args factory.Args) (any, error) {
			createdAt := args.Instance().(*entity.Customer).CreatedAt
			updatedAt := gofakeit.DateRange(createdAt, time.Now())

			isUpdated := gofakeit.Bool()

			return boolutil.Ternary(isUpdated, &updatedAt, nil), nil
		})
}
