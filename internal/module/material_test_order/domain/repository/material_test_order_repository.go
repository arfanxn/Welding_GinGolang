package repository

import (
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	"github.com/arfanxn/welding/pkg/pagination"
	"github.com/arfanxn/welding/pkg/query"
)

type MaterialTestOrderRepository interface {
	Get(q *query.Query) ([]*entity.MaterialTestOrder, error)
	Paginate(q *query.Query) (*pagination.OffsetPagination[*entity.MaterialTestOrder], error)
	First(q *query.Query) (*entity.MaterialTestOrder, error)
	Find(id string, q *query.Query) (*entity.MaterialTestOrder, error)
	FindLatestThisYear(q *query.Query) (*entity.MaterialTestOrder, error)
	Save(order *entity.MaterialTestOrder) error
	SaveMany(orders []*entity.MaterialTestOrder) error
	Destroy(order *entity.MaterialTestOrder) error
}
