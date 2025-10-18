package repository

import (
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
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

func (r *GormUserRepository) Save(user *entity.User) error {
	return r.db.Save(user).Error
}

func (r *GormUserRepository) SaveMany(users []*entity.User) error {
	return r.db.CreateInBatches(users, 100).Error
}
