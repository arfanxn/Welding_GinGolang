package repository

import (
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	"github.com/arfanxn/welding/pkg/pagination"
	"github.com/arfanxn/welding/pkg/query"
)

type MaterialTestOrderServiceRepository interface {
	Get(q *query.Query) ([]*entity.MaterialTestOrderService, error)
	Paginate(q *query.Query) (*pagination.OffsetPagination[*entity.MaterialTestOrderService], error)
	First(q *query.Query) (*entity.MaterialTestOrderService, error)
	Find(id string, q *query.Query) (*entity.MaterialTestOrderService, error)
	FindLatestPerServiceByServiceIds(serviceIds []string, q *query.Query) (mtoss []*entity.MaterialTestOrderService, err error)
	Save(mtso *entity.MaterialTestOrderService) error
	SaveMany(mtsos []*entity.MaterialTestOrderService) error
	DestroyByOrderId(orderId string) error
	Destroy(mtso *entity.MaterialTestOrderService) error
}
