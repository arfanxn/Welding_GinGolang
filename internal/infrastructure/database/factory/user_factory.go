package factory

import (
	"time"

	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	"github.com/bluele/factory-go/factory"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/guregu/null/v6"
	"github.com/oklog/ulid/v2"
	"golang.org/x/crypto/bcrypt"
)

var UserFactory = factory.NewFactory(&entity.User{}).
	Attr("Id", func(args factory.Args) (any, error) {
		return ulid.Make().String(), nil
	}).
	Attr("Name", func(args factory.Args) (any, error) {
		return gofakeit.Name(), nil
	}).
	Attr("Email", func(args factory.Args) (any, error) {
		return gofakeit.Email(), nil
	}).
	Attr("EmailVerifiedAt", func(args factory.Args) (any, error) {
		isEmailVerified := gofakeit.Bool()
		if isEmailVerified {
			return null.TimeFrom(time.Now()), nil
		}
		return null.TimeFromPtr(nil), nil
	}).
	Attr("Password", func(args factory.Args) (any, error) {
		passwordStr := "11112222"
		hashedPasswordBytes, err := bcrypt.GenerateFromPassword([]byte(passwordStr), bcrypt.DefaultCost)
		if err != nil {
			return "", err
		}
		return string(hashedPasswordBytes), nil
	}).
	Attr("CreatedAt", func(args factory.Args) (any, error) {
		user := args.Instance().(*entity.User)

		end := user.EmailVerifiedAt.Time
		if end.IsZero() {
			end = time.Now()
		}

		return gofakeit.DateRange(time.Now().Add(-time.Hour*24*365), end), nil
	}).
	Attr("UpdatedAt", func(args factory.Args) (any, error) {
		createdAt := args.Instance().(*entity.User).CreatedAt

		isUpdated := gofakeit.Bool()
		if isUpdated {
			return null.
				TimeFrom(gofakeit.DateRange(createdAt, time.Now())), nil
		}

		return null.TimeFromPtr(nil), nil
	}).
	Attr("DeletedAt", func(args factory.Args) (any, error) {
		return null.TimeFromPtr(nil), nil
	})
