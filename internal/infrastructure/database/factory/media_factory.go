package factory

import (
	"time"

	id "github.com/arfanxn/welding/internal/infrastructure/id"
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	"github.com/arfanxn/welding/pkg/boolutil"
	"github.com/bluele/factory-go/factory"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/iancoleman/strcase"
)

func NewMediaFactory(
	idService id.IdService,
) *factory.Factory {
	return factory.NewFactory(&entity.Media{}).
		Attr("Id", func(args factory.Args) (any, error) {
			return idService.Generate(), nil
		}).
		Attr("Name", func(args factory.Args) (any, error) {
			name := strcase.ToCamel(gofakeit.Word())
			return name, nil
		}).
		Attr("FileName", func(args factory.Args) (any, error) {
			name := args.Instance().(*entity.Media).Name
			fileName := strcase.ToSnake(name) + ".jpeg"
			return fileName, nil
		}).
		Attr("MimeType", func(args factory.Args) (any, error) {
			mimeType := "image/jpeg"
			return &mimeType, nil
		}).
		Attr("Disk", func(args factory.Args) (any, error) {
			return "local", nil
		}).
		Attr("CreatedAt", func(args factory.Args) (any, error) {
			createdAt := gofakeit.DateRange(time.Now().AddDate(-1, 0, 0), time.Now())
			return createdAt, nil
		}).
		Attr("UpdatedAt", func(args factory.Args) (any, error) {
			createdAt := args.Instance().(*entity.Media).CreatedAt
			updatedAt := gofakeit.DateRange(createdAt, time.Now())

			isUpdated := gofakeit.Bool()

			return boolutil.Ternary(isUpdated, &updatedAt, nil), nil
		})
}
