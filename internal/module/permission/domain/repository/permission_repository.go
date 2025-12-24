package repository

import (
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	"github.com/arfanxn/welding/pkg/pagination"
	"github.com/arfanxn/welding/pkg/query"
)

type PermissionRepository interface {
	All() ([]*entity.Permission, error)
	Get(q *query.Query) ([]*entity.Permission, error)
	Paginate(q *query.Query) (*pagination.OffsetPagination[*entity.Permission], error)
	First(q *query.Query) (*entity.Permission, error)
	Find(id string, q *query.Query) (*entity.Permission, error)
	FindByName(name string, q *query.Query) (*entity.Permission, error)
	FindByIds(ids []string, q *query.Query) ([]*entity.Permission, error)
	FindByUserId(userId string, q *query.Query) ([]*entity.Permission, error)
	Save(permission *entity.Permission) error
	SaveMany(permissions []*entity.Permission) error
}
