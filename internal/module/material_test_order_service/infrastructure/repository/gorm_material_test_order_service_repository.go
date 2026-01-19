package repository

import (
	"errors"

	"github.com/arfanxn/welding/internal/infrastructure/database/helper"
	"github.com/arfanxn/welding/internal/module/material_test_order_service/domain/repository"
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	"github.com/arfanxn/welding/internal/module/shared/domain/errorx"
	"github.com/arfanxn/welding/pkg/pagination"
	"github.com/arfanxn/welding/pkg/query"
	"gorm.io/gorm"
)

type GormMaterialTestOrderServiceRepository struct {
	db *gorm.DB
}

func NewGormMaterialTestOrderServiceRepository(db *gorm.DB) repository.MaterialTestOrderServiceRepository {
	return &GormMaterialTestOrderServiceRepository{
		db: db,
	}
}

func (r *GormMaterialTestOrderServiceRepository) query(db *gorm.DB, q *query.Query) *gorm.DB {
	mtsoTableName := entity.NewMaterialTestOrderService().TableName()

	if q != nil {
		if id := q.GetFilterById(); id != nil {
			db = db.Where(mtsoTableName+".id = ?", id.Value)
		}

		if search := q.GetSearch(); search != nil {
			db = db.Where(mtsoTableName+".name ILIKE ?", "%"+*search+"%")
		}

		if orderId := q.GetFilter("order_id", query.OperatorEqual); orderId != nil {
			db = db.Where(mtsoTableName+".order_id = ?", orderId.Value)
		}

		if name := q.GetFilter("name", query.OperatorEqual); name != nil {
			db = db.Where(mtsoTableName+".name = ?", name.Value)
		}

		if priceBetween := q.GetFilter("price", query.OperatorBetween); priceBetween != nil {
			db = db.Where(mtsoTableName+".price BETWEEN ? AND ?", priceBetween.Values[0], priceBetween.Values[1])
		}

		if quantityBetween := q.GetFilter("quantity", query.OperatorBetween); quantityBetween != nil {
			db = db.Where(mtsoTableName+".quantity BETWEEN ? AND ?", quantityBetween.Values[0], quantityBetween.Values[1])
		}

		if lineTotalBetween := q.GetFilter("line_total", query.OperatorBetween); lineTotalBetween != nil {
			db = db.Where(mtsoTableName+".line_total BETWEEN ? AND ?", lineTotalBetween.Values[0], lineTotalBetween.Values[1])
		}

		if sort := q.GetSort("name"); sort != nil {
			db = db.Order(mtsoTableName + ".name" + sort.Order)
		}

		if sort := q.GetSort("created_at"); sort != nil {
			db = db.Order(mtsoTableName + ".created_at " + sort.Order)
		}
	}

	return db
}

func (r *GormMaterialTestOrderServiceRepository) Get(q *query.Query) ([]*entity.MaterialTestOrderService, error) {
	var mtsos []*entity.MaterialTestOrderService

	db := r.query(r.db, q)

	if err := db.Find(&mtsos).Error; err != nil {
		return nil, err
	}

	return mtsos, nil
}

func (r *GormMaterialTestOrderServiceRepository) Paginate(q *query.Query) (*pagination.OffsetPagination[*entity.MaterialTestOrderService], error) {
	db := r.db.Model(&entity.MaterialTestOrderService{})

	db = r.query(db, q)

	paginator, err := helper.GormDBPaginateWithQuery[*entity.MaterialTestOrderService](db, q)
	if err != nil {
		return nil, err
	}
	return paginator, nil
}

func (r *GormMaterialTestOrderServiceRepository) First(q *query.Query) (*entity.MaterialTestOrderService, error) {
	var mtso *entity.MaterialTestOrderService

	db := r.query(r.db, q)

	if err := db.First(&mtso).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorx.ErrMaterialTestOrderServiceNotFound
		}
		return nil, err
	}

	return mtso, nil
}

func (r *GormMaterialTestOrderServiceRepository) Find(id string, q *query.Query) (*entity.MaterialTestOrderService, error) {
	var mtso entity.MaterialTestOrderService

	db := r.query(r.db, q)

	if err := db.Where("id = ?", id).First(&mtso).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorx.ErrMaterialTestOrderServiceNotFound
		}
		return nil, err
	}
	return &mtso, nil
}

func (r *GormMaterialTestOrderServiceRepository) FindByIds(ids []string, q *query.Query) (mtoss []*entity.MaterialTestOrderService, err error) {
	db := r.query(r.db, q)

	if err := db.Where("id IN (?)", ids).Find(&mtoss).Error; err != nil {
		return nil, err
	}
	if len(mtoss) != len(ids) {
		return nil, errorx.ErrMaterialTestOrderServiceNotFound
	}
	return
}

func (r *GormMaterialTestOrderServiceRepository) FindLatestThisYearPerServiceByServiceIds(
	serviceIds []string, q *query.Query,
) (mtoss []*entity.MaterialTestOrderService, err error) {
	// TODO: do something on `q *query.Query`

	if len(serviceIds) == 0 {
		return []*entity.MaterialTestOrderService{}, nil
	}

	table := entity.NewMaterialTestOrderService().TableName()

	subQuery := r.db.
		Table(table).
		Select(`
			*,
			ROW_NUMBER() OVER (
				PARTITION BY service_id
				ORDER BY RIGHT(sample_number, 3)::int DESC
			) AS rn
		`).
		Where("service_id IN ?", serviceIds).
		Where(`
			SPLIT_PART(SPLIT_PART(sample_number, '/', 2), '.', 1)::int
			= EXTRACT(YEAR FROM CURRENT_DATE)
		`)

	err = r.db.
		Table("(?) AS ranked", subQuery).
		Where("rn = 1").
		Find(&mtoss).Error

	return mtoss, err
}

func (r *GormMaterialTestOrderServiceRepository) Save(mtso *entity.MaterialTestOrderService) error {
	err := r.db.Save(mtso).Error
	if err != nil {
		if helper.IsPostgresDuplicateKeyError(err) {
			return errorx.ErrMaterialTestOrderServiceAlreadyExists
		}
		return err
	}

	return nil
}

func (r *GormMaterialTestOrderServiceRepository) SaveMany(mtsos []*entity.MaterialTestOrderService) error {
	return r.db.CreateInBatches(mtsos, 100).Error
}

func (r *GormMaterialTestOrderServiceRepository) DestroyByOrderId(orderId string) error {
	return r.db.Where("order_id = ?", orderId).Delete(&entity.MaterialTestOrderService{}).Error
}

// DestroyByOrderIdAndExceptIds deletes all MaterialTestOrderService records that match the given orderId
// except those whose IDs are in the provided ids slice.
//
// Parameters:
//   - orderId: The ID of the order to match against
//   - ids: Slice of order service IDs to exclude from deletion
//
// Returns:
//   - error: An error if the deletion fails, nil otherwise
func (r *GormMaterialTestOrderServiceRepository) DestroyByOrderIdAndExceptIds(orderId string, ids []string) error {
	return r.db.Where("order_id = ? AND id NOT IN ?", orderId, ids).
		Delete(&entity.MaterialTestOrderService{}).Error
}

func (r *GormMaterialTestOrderServiceRepository) Destroy(mtso *entity.MaterialTestOrderService) error {
	return r.db.Delete(mtso).Error
}
