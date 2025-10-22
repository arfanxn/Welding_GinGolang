package domain

import (
	"github.com/arfanxn/welding/internal/module/employee/domain/repository"
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
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

func (r *GormEmployeeRepository) Save(employee *entity.Employee) error {
	return r.db.Save(employee).Error
}

func (r *GormEmployeeRepository) SaveMany(employees []*entity.Employee) error {
	return r.db.CreateInBatches(employees, 100).Error
}
