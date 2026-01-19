package factory

import (
	"fmt"
	"time"

	id "github.com/arfanxn/welding/internal/infrastructure/id"
	materialTestOrderEnum "github.com/arfanxn/welding/internal/module/material_test_order/domain/enum"
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	"github.com/arfanxn/welding/pkg/boolutil"
	"github.com/bluele/factory-go/factory"
	"github.com/brianvoe/gofakeit/v7"
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
			number := fmt.Sprintf("%03d", gofakeit.IntRange(1, 999))
			return number, nil
		}).
		Attr("WorkPackageName", func(args factory.Args) (any, error) {
			return gofakeit.Sentence(), nil
		}).
	Attr("ApplicantName", func(args factory.Args) (any, error) {
		isIndividual := gofakeit.Bool()
		return boolutil.Ternary(isIndividual, gofakeit.Name(), gofakeit.Company()), nil
	}).
	Attr("ApplicantCompanyName", func(args factory.Args) (any, error) {
		withCompany := gofakeit.Bool()
		company := gofakeit.Company()
		return boolutil.Ternary(withCompany, company, ""), nil
	}).
	Attr("ApplicantPhoneNumber", func(args factory.Args) (any, error) {
		return gofakeit.Phone(), nil
	}).
		Attr("ApplicantEmail", func(args factory.Args) (any, error) {
			return gofakeit.Email(), nil
		}).
		Attr("ApplicantFullAddress", func(args factory.Args) (any, error) {
			addrInfo := gofakeit.Address()
			fullAddr := addrInfo.Address
			return fullAddr, nil
		}).
		Attr("ApplicantNote", func(args factory.Args) (any, error) {
			applicantNote := gofakeit.Sentence()
			withApplicantNote := gofakeit.Bool()
			return boolutil.Ternary(withApplicantNote, &applicantNote, nil), nil
		}).
		Attr("TesterNote", func(args factory.Args) (any, error) {
			testerNote := gofakeit.Sentence()
			withTesterNote := gofakeit.Bool()
			return boolutil.Ternary(withTesterNote, &testerNote, nil), nil
		}).
	Attr("RecipientName", func(args factory.Args) (any, error) {
		isIndividual := gofakeit.Bool()
		return boolutil.Ternary(isIndividual, gofakeit.Name(), gofakeit.Company()), nil
	}).
	Attr("RecipientFullAddress", func(args factory.Args) (any, error) {
		addrInfo := gofakeit.Address()
		return addrInfo.Address, nil
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
			createdAt := gofakeit.DateRange(
				time.Now().AddDate(0, -6, 0),
				time.Now(),
			)
			return createdAt, nil
		}).
		Attr("UpdatedAt", func(args factory.Args) (any, error) {
			createdAt := args.Instance().(*entity.Address).CreatedAt
			updatedAt := gofakeit.DateRange(createdAt, time.Now())
			return &updatedAt, nil
		})
}
