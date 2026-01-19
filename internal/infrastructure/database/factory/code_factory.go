package factory

import (
	// "time"

	"time"

	id "github.com/arfanxn/welding/internal/infrastructure/id"
	codeService "github.com/arfanxn/welding/internal/module/code/usecase/service"
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	"github.com/bluele/factory-go/factory"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/guregu/null/v6"
	// "github.com/brianvoe/gofakeit/v7"
	// "github.com/guregu/null/v6"
)

func NewCodeFactory(
	idService id.IdService,
	codeService codeService.CodeService,
) *factory.Factory {
	return factory.NewFactory(&entity.Code{}).
		Attr("Id", func(args factory.Args) (any, error) {
			return idService.Generate(), nil
		}).
		Attr("Value", func(args factory.Args) (any, error) {
			return codeService.Generate(), nil
		}).
		Attr("UsedAt", func(args factory.Args) (any, error) {
			isUsed := gofakeit.Bool()

			if isUsed {
				usedAt := gofakeit.DateRange(time.Now().Add(-time.Hour*24*365), time.Now())
				return null.TimeFrom(usedAt), nil
			}

			return null.TimeFromPtr(nil), nil
		}).
		Attr("ExpiredAt", func(args factory.Args) (any, error) {
			usedAt := args.Instance().(*entity.Code).UsedAt

			if usedAt.Valid {
				return gofakeit.DateRange(usedAt.Time, time.Now()), nil
			} else {
				return gofakeit.DateRange(time.Now().Add(-time.Hour*24*365), time.Now()), nil
			}

		}).
		Attr("CreatedAt", func(args factory.Args) (any, error) {
			expiredAt := args.Instance().(*entity.Code).ExpiredAt

			createdAt := expiredAt.Add(-time.Minute * 30)
			return createdAt, nil
		}).
		Attr("UpdatedAt", func(args factory.Args) (any, error) {
			return null.TimeFromPtr(nil), nil
		})

}
