package domain

import (
	"errors"

	"github.com/arfanxn/welding/internal/module/employee/domain/repository"
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	"github.com/arfanxn/welding/internal/module/shared/domain/errorx"
	"gorm.io/gorm"
)

var _ repository.EmployeeRepository = (*GormEmployeeRepository)(nil)

type GormEmployeeRepository struct {
	db *gorm.DB
}

func NewGormEmployeeRepository(db *gorm.DB) repository.EmployeeRepository {
	return &GormEmployeeRepository{
		db: db,
	}
}

func (r *GormEmployeeRepository) FindByUserId(userId string) (*entity.Employee, error) {
	var employee entity.Employee
	if err := r.db.Where("user_id = ?", userId).First(&employee).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorx.ErrEmployeeNotFound
		}
		return nil, err
	}
	return &employee, nil
}

func (r *GormEmployeeRepository) Save(employee *entity.Employee) error {
	return r.db.Save(employee).Error
}

func (r *GormEmployeeRepository) SaveMany(employees []*entity.Employee) error {
	return r.db.CreateInBatches(employees, 100).Error
}

func (r *GormEmployeeRepository) DestroyByUserId(userId string) error {
	return r.db.Where("user_id = ?", userId).Delete(&entity.Employee{}).Error
}
