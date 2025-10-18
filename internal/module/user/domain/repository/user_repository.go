package repository

import "github.com/arfanxn/welding/internal/module/shared/domain/entity"

type UserRepository interface {
	Find(id string) (*entity.User, error)
	FindByEmail(email string) (*entity.User, error)
	Save(user *entity.User) error
	SaveMany(users []*entity.User) error
}
