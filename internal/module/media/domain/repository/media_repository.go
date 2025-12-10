package repository

import (
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	"github.com/arfanxn/welding/pkg/pagination"
	"github.com/arfanxn/welding/pkg/query"
)

type MediaRepository interface {
	Get(q *query.Query) ([]*entity.Media, error)
	Paginate(q *query.Query) (*pagination.OffsetPagination[*entity.Media], error)
	First(q *query.Query) (*entity.Media, error)
	Find(id string, q *query.Query) (*entity.Media, error)
	Save(media *entity.Media) error
	SaveMany(medias []*entity.Media) error
	Destroy(media *entity.Media) error
}
