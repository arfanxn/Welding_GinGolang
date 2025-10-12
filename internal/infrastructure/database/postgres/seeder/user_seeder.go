package seeder

import (
	"time"

	"github.com/arfanxn/welding/internal/module/user/domain/entity"
	"github.com/guregu/null/v6"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var _ Seeder = (*UserSeeder)(nil)

type UserSeeder struct {
	DB *gorm.DB
}

func NewUserSeeder(db *gorm.DB) *UserSeeder {
	return &UserSeeder{DB: db}
}

func (s *UserSeeder) Seed() error {
	passwordStr := "11112222"
	hashedPasswordBytes, err := bcrypt.GenerateFromPassword([]byte(passwordStr), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	hashedPasswordStr := string(hashedPasswordBytes)

	users := []entity.User{
		{
			Id:              "01K7CNB24108TVE4TWJE01MGGD",
			Name:            "Admin",
			Email:           "admin@gmail.com",
			EmailVerifiedAt: null.TimeFrom(time.Now()),
			Password:        hashedPasswordStr,
			CreatedAt:       time.Now(),
			UpdatedAt:       null.Time{},
			DeletedAt:       null.Time{},
		},
	}

	err = s.DB.CreateInBatches(users, 100).Error
	if err != nil {
		return err
	}

	return nil
}
