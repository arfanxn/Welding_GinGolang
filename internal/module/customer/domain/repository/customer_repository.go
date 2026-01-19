package repository

import (
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	"github.com/arfanxn/welding/pkg/pagination"
	"github.com/arfanxn/welding/pkg/query"
)

type CustomerRepository interface {
	Get(q *query.Query) ([]*entity.Customer, error)
	Paginate(q *query.Query) (*pagination.OffsetPagination[*entity.Customer], error)
	First(q *query.Query) (*entity.Customer, error)
	Find(id string, q *query.Query) (*entity.Customer, error)
	Save(customer *entity.Customer) error
	SaveMany(customers []*entity.Customer) error
	Destroy(customer *entity.Customer) error
}
