package repository

import (
	"github.com/arfanxn/welding/internal/module/user/domain/entity"
	"github.com/arfanxn/welding/internal/module/user/domain/repository"
	"gorm.io/gorm"
)

var _ repository.UserRepository = (*GormUserRepository)(nil)

type GormUserRepository struct {
	db *gorm.DB
}

func NewGormUserRepository(db *gorm.DB) repository.UserRepository {
	return &GormUserRepository{
		db: db,
	}
}

func (r *GormUserRepository) Find(id string) (*entity.User, error) {
	var user entity.User
	if err := r.db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *GormUserRepository) FindByEmail(email string) (*entity.User, error) {
	var user entity.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *GormUserRepository) Store(user *entity.User) error {
	return r.db.Create(user).Error
}
