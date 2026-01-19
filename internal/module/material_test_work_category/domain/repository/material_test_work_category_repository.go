package repository

import (
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	"github.com/arfanxn/welding/pkg/pagination"
	"github.com/arfanxn/welding/pkg/query"
)

type MaterialTestWorkCategoryRepository interface {
	Get(*query.Query) ([]*entity.MaterialTestWorkCategory, error)
	Paginate(*query.Query) (*pagination.OffsetPagination[*entity.MaterialTestWorkCategory], error)
	First(*query.Query) (*entity.MaterialTestWorkCategory, error)
	Find(id string, q *query.Query) (*entity.MaterialTestWorkCategory, error)
	Save(materialTestWorkCategory *entity.MaterialTestWorkCategory) error
	SaveMany(materialTestWorkCategories []*entity.MaterialTestWorkCategory) error
	Destroy(materialTestWorkCategory *entity.MaterialTestWorkCategory) error
}
