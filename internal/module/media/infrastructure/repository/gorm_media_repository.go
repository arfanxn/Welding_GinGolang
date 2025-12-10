package repository

import (
	"errors"

	"github.com/arfanxn/welding/internal/infrastructure/database/helper"
	mediaRepository "github.com/arfanxn/welding/internal/module/media/domain/repository"
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	"github.com/arfanxn/welding/internal/module/shared/domain/errorx"
	"github.com/arfanxn/welding/pkg/pagination"
	"github.com/arfanxn/welding/pkg/query"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

type gormMediaRepository struct {
	db *gorm.DB
}

type NewGormMediaRepositoryParams struct {
	fx.In

	DB *gorm.DB
}

func NewGormMediaRepository(params NewGormMediaRepositoryParams) mediaRepository.MediaRepository {
	return &gormMediaRepository{
		db: params.DB,
	}
}

func (r *gormMediaRepository) query(db *gorm.DB, q *query.Query) (*gorm.DB, error) {
	mediaTableName := entity.NewMedia().TableName()

	if q != nil {
		if f := q.GetFilterById(); f != nil {
			db = db.Where(mediaTableName+".id = ?", f.Value)
		}

		if search := q.GetSearch(); search != nil {
			s := "%" + *search + "%"
			db = db.Where(mediaTableName+".name ILIKE ?", s).
				Or(mediaTableName+".file_name ILIKE ?", s).
				Or(mediaTableName+".collection_name ILIKE ?", s)
		}

		if modelType := q.GetFilter("model_type", query.OperatorEqual); modelType != nil {
			db = db.Where(mediaTableName+".model_type = ?", modelType.Value)
		}

		if modelId := q.GetFilter("model_id", query.OperatorEqual); modelId != nil {
			db = db.Where(mediaTableName+".model_id = ?", modelId.Value)
		}

		if collectionName := q.GetFilter("collection_name", query.OperatorEqual); collectionName != nil {
			db = db.Where(mediaTableName+".collection_name = ?", collectionName.Value)
		}

		if name := q.GetFilter("name", query.OperatorEqual); name != nil {
			db = db.Where(mediaTableName+".name = ?", name.Value)
		}

		if fileName := q.GetFilter("file_name", query.OperatorEqual); fileName != nil {
			db = db.Where(mediaTableName+".file_name = ?", fileName.Value)
		}

		if mimeType := q.GetFilter("mime_type", query.OperatorEqual); mimeType != nil {
			db = db.Where(mediaTableName+".mime_type = ?", mimeType.Value)
		}

		if disk := q.GetFilter("disk", query.OperatorEqual); disk != nil {
			db = db.Where(mediaTableName+".disk = ?", disk.Value)
		}

		if createdAtBetween := q.GetFilter("created_at", query.OperatorBetween); createdAtBetween != nil {
			db = db.Where(mediaTableName+".created_at BETWEEN ? AND ?", createdAtBetween.Values[0], createdAtBetween.Values[1])
		}

		if sort := q.GetSort("created_at"); sort != nil {
			db = db.Order(mediaTableName + ".created_at " + sort.Order)
		}

		if sort := q.GetSort("order_column"); sort != nil {
			db = db.Order(mediaTableName + ".order_column " + sort.Order)
		} else {
			db = db.Order(mediaTableName + ".order_column ASC")
		}
	}

	return db, nil
}

func (r *gormMediaRepository) Get(q *query.Query) ([]*entity.Media, error) {
	var mtos []*entity.Media

	db, err := r.query(r.db, q)
	if err != nil {
		return nil, err
	}

	if err := db.Find(&mtos).Error; err != nil {
		return nil, err
	}

	return mtos, nil
}

func (r *gormMediaRepository) Paginate(q *query.Query) (*pagination.OffsetPagination[*entity.Media], error) {
	db := r.db.Model(&entity.Media{})

	db, err := r.query(db, q)
	if err != nil {
		return nil, err
	}

	op, err := helper.GormDBPaginateWithQuery[*entity.Media](db, q)
	if err != nil {
		return nil, err
	}
	return op, nil
}

func (r *gormMediaRepository) First(q *query.Query) (*entity.Media, error) {
	var mto *entity.Media

	db, err := r.query(r.db, q)
	if err != nil {
		return nil, err
	}

	if err := db.First(&mto).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorx.ErrMediaNotFound
		}
		return nil, err
	}

	return mto, nil
}

func (r *gormMediaRepository) Find(id string, q *query.Query) (*entity.Media, error) {
	var mto entity.Media

	db, err := r.query(r.db, q)
	if err != nil {
		return nil, err
	}

	if err := db.Where("id = ?", id).First(&mto).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorx.ErrMediaNotFound
		}
		return nil, err
	}
	return &mto, nil
}

func (r *gormMediaRepository) Save(mto *entity.Media) error {
	return r.db.Save(mto).Error
}

func (r *gormMediaRepository) SaveMany(mtos []*entity.Media) error {
	return r.db.CreateInBatches(mtos, 100).Error
}

func (r *gormMediaRepository) Destroy(mto *entity.Media) error {
	return r.db.Delete(mto).Error
}
