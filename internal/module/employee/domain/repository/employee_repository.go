package repository

import "github.com/arfanxn/welding/internal/module/shared/domain/entity"

type EmployeeRepository interface {
	Save(employee *entity.Employee) error
	SaveMany(employees []*entity.Employee) error
}
