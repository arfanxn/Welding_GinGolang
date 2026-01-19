package repository

import (
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	"github.com/arfanxn/welding/pkg/pagination"
	"github.com/arfanxn/welding/pkg/query"
)

type MaterialTestServiceRepository interface {
	All() ([]*entity.MaterialTestService, error)
	Get(q *query.Query) ([]*entity.MaterialTestService, error)
	Paginate(q *query.Query) (*pagination.OffsetPagination[*entity.MaterialTestService], error)
	First(q *query.Query) (*entity.MaterialTestService, error)
	Find(id string, q *query.Query) (*entity.MaterialTestService, error)
	FindByIds(ids []string, q *query.Query) ([]*entity.MaterialTestService, error)
	CountByMachineId(machineId string) (int64, error)
	CountByMethodId(methodId string) (int64, error)
	Save(mts *entity.MaterialTestService) error
	SaveMany(mtss []*entity.MaterialTestService) error
	Destroy(mts *entity.MaterialTestService) error
}
