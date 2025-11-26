package repository

import (
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	"github.com/arfanxn/welding/pkg/pagination"
	"github.com/arfanxn/welding/pkg/query"
)

type MaterialTestMachineRepository interface {
	All() ([]*entity.MaterialTestMachine, error)
	Get(q *query.Query) ([]*entity.MaterialTestMachine, error)
	Paginate(q *query.Query) (*pagination.OffsetPagination[*entity.MaterialTestMachine], error)
	First(q *query.Query) (*entity.MaterialTestMachine, error)
	Find(id string) (*entity.MaterialTestMachine, error)
	FindByName(name string) (*entity.MaterialTestMachine, error)
	FindByIds(ids []string) ([]*entity.MaterialTestMachine, error)
	Save(mtm *entity.MaterialTestMachine) error
	SaveMany(mtms []*entity.MaterialTestMachine) error
	Destroy(mtm *entity.MaterialTestMachine) error
}
