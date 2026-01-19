package repository

import (
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	"github.com/arfanxn/welding/pkg/pagination"
	"github.com/arfanxn/welding/pkg/query"
)

type MaterialTestOrderServiceEvaluationRepository interface {
	Get(q *query.Query) ([]*entity.MaterialTestOrderServiceEvaluation, error)
	Paginate(q *query.Query) (*pagination.OffsetPagination[*entity.MaterialTestOrderServiceEvaluation], error)
	First(q *query.Query) (*entity.MaterialTestOrderServiceEvaluation, error)
	Find(id string, q *query.Query) (*entity.MaterialTestOrderServiceEvaluation, error)
	Save(mts *entity.MaterialTestOrderServiceEvaluation) error
	SaveMany(mtss []*entity.MaterialTestOrderServiceEvaluation) error
	Destroy(mts *entity.MaterialTestOrderServiceEvaluation) error
}
