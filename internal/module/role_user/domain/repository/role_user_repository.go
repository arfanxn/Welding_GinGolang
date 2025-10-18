package repository

import "github.com/arfanxn/welding/internal/module/shared/domain/entity"

type RoleUserRepository interface {
	Save(role *entity.RoleUser) error
	SaveMany(roles []*entity.RoleUser) error
}
