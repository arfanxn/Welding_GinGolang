package repository

import (
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	"github.com/arfanxn/welding/pkg/pagination"
	"github.com/arfanxn/welding/pkg/query"
)

type ActivityRepository interface {
	Paginate(q *query.Query) (*pagination.OffsetPagination[*entity.Activity], error)
	First(query *query.Query) (*entity.Activity, error)
	Find(id string) (*entity.Activity, error)
	Save(activity *entity.Activity) error
	SaveMany(activities []*entity.Activity) error
}
