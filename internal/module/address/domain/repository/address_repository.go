package repository

import (
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	"github.com/arfanxn/welding/pkg/pagination"
	"github.com/arfanxn/welding/pkg/query"
)

type AddressRepository interface {
	Get(query *query.Query) ([]*entity.Address, error)
	Paginate(query *query.Query) (*pagination.OffsetPagination[*entity.Address], error)
	First(query *query.Query) (*entity.Address, error)
	Find(id string, q *query.Query) (*entity.Address, error)
	Save(user *entity.Address) error
	SaveMany(users []*entity.Address) error
	Destroy(user *entity.Address) error
}
