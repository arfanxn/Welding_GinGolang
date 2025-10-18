package repository

import (
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	"github.com/arfanxn/welding/internal/module/shared/usecase/dto"
)

type RoleRepository interface {
	All() ([]*entity.Role, error)
	Get(*dto.Query) ([]*entity.Role, error)
	Paginate(*dto.Query) (*dto.Pagination[*entity.Role], error)
	Find(id string) (*entity.Role, error)
	FindByName(name string) (*entity.Role, error)
	Save(role *entity.Role) error
	SaveMany(roles []*entity.Role) error
	Destroy(role *entity.Role) error
}
