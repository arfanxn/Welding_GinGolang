package factory

import (
	"time"

	id "github.com/arfanxn/welding/internal/infrastructure/id"
	materialTestOrderEnum "github.com/arfanxn/welding/internal/module/material_test_order/domain/enum"
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	"github.com/arfanxn/welding/pkg/boolutil"
	"github.com/bluele/factory-go/factory"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/samber/lo"
)

func NewMaterialTestOrderFactory(
	idService id.IdService,
) *factory.Factory {
	statuses := materialTestOrderEnum.MaterialTestOrderStatuses

	return factory.NewFactory(&entity.MaterialTestOrder{}).
		Attr("Id", func(args factory.Args) (any, error) {
			return idService.Generate(), nil
		}).
		Attr("Number", func(args factory.Args) (any, error) {
			number := lo.Substring(idService.Generate(), 17, 26)
			return number, nil
		}).
		Attr("CustomerNote", func(args factory.Args) (any, error) {
			customerNote := gofakeit.Sentence()
			withCustomerNote := gofakeit.Bool()
			return boolutil.Ternary(withCustomerNote, &customerNote, nil), nil
		}).
		Attr("TesterNote", func(args factory.Args) (any, error) {
			testerNote := gofakeit.Sentence()
			withTesterNote := gofakeit.Bool()
			return boolutil.Ternary(withTesterNote, &testerNote, nil), nil
		}).
		Attr("IssuedFor", func(args factory.Args) (any, error) {
			return gofakeit.Name(), nil
		}).
		Attr("Tax", func(args factory.Args) (any, error) {
			withTax := gofakeit.Bool()
			tax := gofakeit.Float64Range(1000.00, 100000.00)
			return boolutil.Ternary(withTax, tax, 0.00), nil
		}).
		Attr("Discount", func(args factory.Args) (any, error) {
			withDiscount := gofakeit.Bool()
			discount := gofakeit.Float64Range(1000.00, 100000.00)
			return boolutil.Ternary(withDiscount, discount, 0.00), nil
		}).
		Attr("EnteredAt", func(args factory.Args) (any, error) {
			enteredAt := gofakeit.DateRange(
				time.Now().AddDate(-1, 0, 0),
				time.Now(),
			)
			return enteredAt, nil
		}).
		Attr("Status", func(args factory.Args) (any, error) {
			status := statuses[gofakeit.IntRange(0, len(statuses)-1)]
			return status, nil
		}).
		/*
			Attr("PaymentSubmittedAt", func(args factory.Args) (any, error) {
			}).
			Attr("PaymentRejectedAt", func(args factory.Args) (any, error) {
			}).
			Attr("PaymentApprovedAt", func(args factory.Args) (any, error) {
			}).
			Attr("TestingAt", func(args factory.Args) (any, error) {
			}).
			Attr("CompletedAt", func(args factory.Args) (any, error) {
			}).
			Attr("CancelledAt", func(args factory.Args) (any, error) {
			}).
			Attr("RejectedAt", func(args factory.Args) (any, error) {
			}).
			Attr("RefundedAt", func(args factory.Args) (any, error) {
			}).
		*/
		Attr("CreatedAt", func(args factory.Args) (any, error) {
			enteredAt := args.Instance().(*entity.MaterialTestOrder).EnteredAt
			createdAt := gofakeit.DateRange(
				enteredAt.AddDate(0, -6, 0),
				enteredAt,
			)
			return createdAt, nil
		}).
		Attr("UpdatedAt", func(args factory.Args) (any, error) {
			createdAt := args.Instance().(*entity.Address).CreatedAt
			updatedAt := gofakeit.DateRange(createdAt, time.Now())
			return &updatedAt, nil
		})
}
