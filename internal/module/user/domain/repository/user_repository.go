package repository

import "github.com/arfanxn/welding/internal/module/user/domain/entity"

type UserRepository interface {
	Find(id string) (*entity.User, error)
	FindByEmail(email string) (*entity.User, error)
	Store(user *entity.User) error
}
