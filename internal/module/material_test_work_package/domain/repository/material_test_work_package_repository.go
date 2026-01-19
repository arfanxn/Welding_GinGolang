package repository

import (
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	"github.com/arfanxn/welding/pkg/pagination"
	"github.com/arfanxn/welding/pkg/query"
)

type MaterialTestWorkPackageRepository interface {
	Get(*query.Query) ([]*entity.MaterialTestWorkPackage, error)
	Paginate(*query.Query) (*pagination.OffsetPagination[*entity.MaterialTestWorkPackage], error)
	First(*query.Query) (*entity.MaterialTestWorkPackage, error)
	Find(id string, q *query.Query) (*entity.MaterialTestWorkPackage, error)
	Save(materialTestWorkPackage *entity.MaterialTestWorkPackage) error
	SaveMany(materialTestWorkPackages []*entity.MaterialTestWorkPackage) error
	Destroy(materialTestWorkPackage *entity.MaterialTestWorkPackage) error
}
