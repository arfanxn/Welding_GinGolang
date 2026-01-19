package repository

import (
	"errors"

	"github.com/arfanxn/welding/internal/infrastructure/database/helper"
	"github.com/arfanxn/welding/internal/module/material_test_order_service_evaluation/domain/repository"
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	"github.com/arfanxn/welding/internal/module/shared/domain/errorx"
	"github.com/arfanxn/welding/pkg/pagination"
	"github.com/arfanxn/welding/pkg/query"
	"gorm.io/gorm"
)

var _ repository.MaterialTestOrderServiceEvaluationRepository = (*GormMaterialTestOrderServiceEvaluationRepository)(nil)

type GormMaterialTestOrderServiceEvaluationRepository struct {
	db *gorm.DB
}

func NewGormMaterialTestOrderServiceEvaluationRepository(db *gorm.DB) repository.MaterialTestOrderServiceEvaluationRepository {
	return &GormMaterialTestOrderServiceEvaluationRepository{
		db: db,
	}
}

// query applies query filters and sorting to the database query based on the provided Query DTO.
// It supports searching by name (case-insensitive) and sorting by name in ascending or descending order.
// The modified *gorm.DB is returned with the applied scopes.
func (r *GormMaterialTestOrderServiceEvaluationRepository) query(db *gorm.DB, q *query.Query) *gorm.DB {
	mtOrderServiceEvaluationTableName := entity.NewMaterialTestOrderServiceEvaluation().TableName()

	if q != nil {
		if id := q.GetFilterById(); id != nil {
			db = db.Where(mtOrderServiceEvaluationTableName+".id = ?", id.Value)
		}

		if sort := q.GetSort("created_at"); sort != nil {
			db = db.Order(mtOrderServiceEvaluationTableName + ".created_at " + sort.Order)
		}
	}

	return db
}

func (r *GormMaterialTestOrderServiceEvaluationRepository) Get(q *query.Query) ([]*entity.MaterialTestOrderServiceEvaluation, error) {
	var materialTestOrderServiceEvaluations []*entity.MaterialTestOrderServiceEvaluation

	db := r.query(r.db, q)

	if err := db.Find(&materialTestOrderServiceEvaluations).Error; err != nil {
		return nil, err
	}

	return materialTestOrderServiceEvaluations, nil
}

func (r *GormMaterialTestOrderServiceEvaluationRepository) Paginate(q *query.Query) (*pagination.OffsetPagination[*entity.MaterialTestOrderServiceEvaluation], error) {
	db := r.db.Model(&entity.MaterialTestOrderServiceEvaluation{})

	db = r.query(db, q)

	pagination, err := helper.GormDBPaginateWithQuery[*entity.MaterialTestOrderServiceEvaluation](db, q)
	if err != nil {
		return nil, err
	}
	return pagination, nil
}

func (r *GormMaterialTestOrderServiceEvaluationRepository) First(q *query.Query) (*entity.MaterialTestOrderServiceEvaluation, error) {
	var materialTestOrderServiceEvaluation *entity.MaterialTestOrderServiceEvaluation

	db := r.query(r.db, q)

	if err := db.First(&materialTestOrderServiceEvaluation).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorx.ErrMaterialTestOrderServiceEvaluationNotFound
		}
		return nil, err
	}

	return materialTestOrderServiceEvaluation, nil
}

func (r *GormMaterialTestOrderServiceEvaluationRepository) Find(id string, q *query.Query) (*entity.MaterialTestOrderServiceEvaluation, error) {
	var materialTestOrderServiceEvaluation entity.MaterialTestOrderServiceEvaluation

	db := r.query(r.db, q)

	if err := db.Where("id = ?", id).First(&materialTestOrderServiceEvaluation).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorx.ErrMaterialTestOrderServiceEvaluationNotFound
		}
		return nil, err
	}
	return &materialTestOrderServiceEvaluation, nil
}

func (r *GormMaterialTestOrderServiceEvaluationRepository) Save(materialTestOrderServiceEvaluation *entity.MaterialTestOrderServiceEvaluation) error {
	err := r.db.Save(materialTestOrderServiceEvaluation).Error
	if err != nil {
		if helper.IsPostgresDuplicateKeyError(err) {
			return errorx.ErrMaterialTestOrderServiceEvaluationAlreadyExists
		}
		return err
	}
	return nil
}

func (r *GormMaterialTestOrderServiceEvaluationRepository) SaveMany(materialTestOrderServiceEvaluations []*entity.MaterialTestOrderServiceEvaluation) error {
	return r.db.CreateInBatches(materialTestOrderServiceEvaluations, 100).Error
}

func (r *GormMaterialTestOrderServiceEvaluationRepository) Destroy(materialTestOrderServiceEvaluation *entity.MaterialTestOrderServiceEvaluation) error {
	return r.db.Delete(materialTestOrderServiceEvaluation).Error
}
