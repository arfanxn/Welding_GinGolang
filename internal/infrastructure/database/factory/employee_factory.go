package factory

import (
	"time"

	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	"github.com/bluele/factory-go/factory"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/gookit/goutil"
	"github.com/guregu/null/v6"
)

var EmployeeFactory = factory.NewFactory(&entity.Employee{}).
	Attr("EmploymentIdentityNumber", func(args factory.Args) (any, error) {
		return goutil.ToString(gofakeit.IntRange(100000000000000000, 900000000000000000))
	}).
	Attr("CreatedAt", func(args factory.Args) (any, error) {
		return gofakeit.DateRange(time.Now().Add(-time.Hour*24*365), time.Now()), nil
	}).
	Attr("UpdatedAt", func(args factory.Args) (any, error) {
		createdAt := args.Instance().(*entity.Employee).CreatedAt

		isUpdated := gofakeit.Bool()
		if isUpdated {
			return null.
				TimeFrom(gofakeit.DateRange(createdAt, time.Now())), nil
		}

		return null.TimeFromPtr(nil), nil
	})
