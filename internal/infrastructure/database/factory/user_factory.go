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
	Attr("PhoneNumber", func(args factory.Args) (any, error) {
		return gofakeit.Phone(), nil
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
	// ActivatedAt represents when the user was activated.
	// - 50% chance to generate an activation timestamp
	// - If activated, the timestamp will be between user creation time and now
	// - If not activated, returns null
	Attr("ActivatedAt", func(args factory.Args) (any, error) {
		user := args.Instance().(*entity.User)
		isActivated := gofakeit.Bool()
		if isActivated {
			activatedAt := gofakeit.DateRange(user.CreatedAt, time.Now())
			return null.TimeFrom(activatedAt), nil
		}
		return null.TimeFromPtr(nil), nil
	}).

	// DeactivatedAt represents when the user was deactivated.
	// - Only set if the user was never activated (ActivatedAt is zero)
	// - If deactivated, the timestamp will be between user creation time and now
	// - Returns null if the user was activated (non-zero ActivatedAt)
	Attr("DeactivatedAt", func(args factory.Args) (any, error) {
		user := args.Instance().(*entity.User)
		if user.ActivatedAt.IsZero() {
			deactivatedAt := gofakeit.DateRange(user.CreatedAt, time.Now())
			return null.TimeFrom(deactivatedAt), nil
		}
		return null.TimeFromPtr(nil), nil
	}).
	Attr("UpdatedAt", func(args factory.Args) (any, error) {
		createdAt := args.Instance().(*entity.User).CreatedAt

		isUpdated := gofakeit.Bool()
		if isUpdated {
			return null.
				TimeFrom(gofakeit.DateRange(createdAt, time.Now())), nil
		}

		return null.TimeFromPtr(nil), nil
	})
