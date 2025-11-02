package repository

import (
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	"github.com/arfanxn/welding/pkg/pagination"
	"github.com/arfanxn/welding/pkg/query"
)

type RoleRepository interface {
	All() ([]*entity.Role, error)
	Get(*query.Query) ([]*entity.Role, error)
	Paginate(*query.Query) (*pagination.OffsetPagination[*entity.Role], error)
	Find(id string) (*entity.Role, error)
	FindDefault() (*entity.Role, error)
	FindByIds(ids []string) ([]*entity.Role, error)
	FindByName(name string) (*entity.Role, error)
	Save(role *entity.Role) error
	SetDefault(role *entity.Role) error
	SaveMany(roles []*entity.Role) error
	Destroy(role *entity.Role) error
}
