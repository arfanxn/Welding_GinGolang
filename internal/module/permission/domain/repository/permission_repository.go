package repository

import (
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	"github.com/arfanxn/welding/internal/module/shared/usecase/dto"
)

type PermissionRepository interface {
	All() ([]*entity.Permission, error)
	Get(queryDto *dto.Query) ([]*entity.Permission, error)
	Paginate(queryDto *dto.Query) (*dto.Pagination[*entity.Permission], error)
	Find(id string) (*entity.Permission, error)
	FindByName(name string) (*entity.Permission, error)
	Save(permission *entity.Permission) error
	SaveMany(permissions []*entity.Permission) error
}
