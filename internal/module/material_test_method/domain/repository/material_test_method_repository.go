package repository

import (
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	"github.com/arfanxn/welding/pkg/pagination"
	"github.com/arfanxn/welding/pkg/query"
)

type MaterialTestMethodRepository interface {
	All() ([]*entity.MaterialTestMethod, error)
	Get(q *query.Query) ([]*entity.MaterialTestMethod, error)
	Paginate(q *query.Query) (*pagination.OffsetPagination[*entity.MaterialTestMethod], error)
	First(q *query.Query) (*entity.MaterialTestMethod, error)
	Find(id string, q *query.Query) (*entity.MaterialTestMethod, error)
	FindByName(name string, q *query.Query) (*entity.MaterialTestMethod, error)
	FindByIds(ids []string, q *query.Query) ([]*entity.MaterialTestMethod, error)
	Save(mtm *entity.MaterialTestMethod) error
	SaveMany(mtms []*entity.MaterialTestMethod) error
	Destroy(mtm *entity.MaterialTestMethod) error
}
