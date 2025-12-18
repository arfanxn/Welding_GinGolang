package repository

import (
	"errors"

	"github.com/arfanxn/welding/internal/infrastructure/database/helper"
	"github.com/arfanxn/welding/internal/module/material_test_order/domain/repository"
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	"github.com/arfanxn/welding/internal/module/shared/domain/errorx"
	"github.com/arfanxn/welding/pkg/pagination"
	"github.com/arfanxn/welding/pkg/query"
	"gorm.io/gorm"
)

type GormMaterialTestOrderRepository struct {
	db *gorm.DB
}

func NewGormMaterialTestOrderRepository(db *gorm.DB) repository.MaterialTestOrderRepository {
	return &GormMaterialTestOrderRepository{
		db: db,
	}
}

func (r *GormMaterialTestOrderRepository) query(db *gorm.DB, q *query.Query) (*gorm.DB, error) {
	mtoTableName := entity.NewMaterialTestOrder().TableName()

	if q != nil {
		if q.GetInclude("work_category") != nil {
			db = db.Preload("WorkCategory")
		}

		if q.GetInclude("order_users.user") != nil {
			db = db.Preload("OrderUsers.User")
		}

		if q.GetInclude("ordered_services") != nil {
			db = db.Preload("OrderedServices")
		}

		if q.GetInclude("ordered_services.evaluation") != nil {
			db = db.Preload("OrderedServices.Evaluation")
		}

		if q.GetInclude("medias") != nil {
			db = db.Preload("Medias")
		}

		if f := q.GetFilterById(); f != nil {
			db = db.Where(mtoTableName+".id = ?", f.Value)
		}

		if search := q.GetSearch(); search != nil {
			s := "%" + *search + "%"
			db = db.Where("CAST("+mtoTableName+".number AS TEXT) ILIKE ?", s).
				Or(mtoTableName+".work_package_name ILIKE ?", s).
				Or(mtoTableName+".applicant_name ILIKE ?", s).
				Or(mtoTableName+".applicant_phone_number ILIKE ?", s).
				Or(mtoTableName+".applicant_email ILIKE ?", s).
				Or(mtoTableName+".applicant_full_address ILIKE ?", s).
				Or(mtoTableName+".recipient_name ILIKE ?", s).
				Or(mtoTableName+".status ILIKE ?", s)
		}

		if number := q.GetFilter("number", query.OperatorEqual); number != nil {
			db = db.Where(mtoTableName+".number = ?", number.Value)
		}

		if workCategoryId := q.GetFilter("work_category_id", query.OperatorEqual); workCategoryId != nil {
			db = db.Where(mtoTableName+".work_category_id = ?", workCategoryId.Value)
		}

		if workPackageName := q.GetFilter("work_package_name", query.OperatorEqual); workPackageName != nil {
			db = db.Where(mtoTableName+".work_package_name = ?", workPackageName.Value)
		}

		if applicantName := q.GetFilter("applicant_name", query.OperatorEqual); applicantName != nil {
			db = db.Where(mtoTableName+".applicant_name = ?", applicantName.Value)
		}

		if applicantPhoneNumber := q.GetFilter("applicant_phone_number", query.OperatorEqual); applicantPhoneNumber != nil {
			db = db.Where(mtoTableName+".applicant_phone_number = ?", applicantPhoneNumber.Value)
		}

		if applicantEmail := q.GetFilter("applicant_email", query.OperatorEqual); applicantEmail != nil {
			db = db.Where(mtoTableName+".applicant_email = ?", applicantEmail.Value)
		}

		if recipientName := q.GetFilter("recipient_name", query.OperatorEqual); recipientName != nil {
			db = db.Where(mtoTableName+".recipient_name = ?", recipientName.Value)
		}

		if status := q.GetFilter("status", query.OperatorEqual); status != nil {
			db = db.Where(mtoTableName+".status = ?", status.Value)
		}

		if createdAtBetween := q.GetFilter("created_at", query.OperatorBetween); createdAtBetween != nil {
			db = db.Where(mtoTableName+".created_at BETWEEN ? AND ?", createdAtBetween.Values[0], createdAtBetween.Values[1])
		}

		if sort := q.GetSort("created_at"); sort != nil {
			db = db.Order(mtoTableName + ".created_at " + sort.Order)
		}
	}

	return db, nil
}

func (r *GormMaterialTestOrderRepository) Get(q *query.Query) ([]*entity.MaterialTestOrder, error) {
	var mtos []*entity.MaterialTestOrder

	db, err := r.query(r.db, q)
	if err != nil {
		return nil, err
	}

	if err := db.Find(&mtos).Error; err != nil {
		return nil, err
	}

	return mtos, nil
}

func (r *GormMaterialTestOrderRepository) Paginate(q *query.Query) (*pagination.OffsetPagination[*entity.MaterialTestOrder], error) {
	db := r.db.Model(&entity.MaterialTestOrder{})

	db, err := r.query(db, q)
	if err != nil {
		return nil, err
	}

	op, err := helper.GormDBPaginateWithQuery[*entity.MaterialTestOrder](db, q)
	if err != nil {
		return nil, err
	}
	return op, nil
}

func (r *GormMaterialTestOrderRepository) First(q *query.Query) (*entity.MaterialTestOrder, error) {
	var mto *entity.MaterialTestOrder

	db, err := r.query(r.db, q)
	if err != nil {
		return nil, err
	}

	if err := db.First(&mto).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorx.ErrMaterialTestOrderNotFound
		}
		return nil, err
	}

	return mto, nil
}

func (r *GormMaterialTestOrderRepository) Find(id string, q *query.Query) (*entity.MaterialTestOrder, error) {
	var mto entity.MaterialTestOrder

	db, err := r.query(r.db, q)
	if err != nil {
		return nil, err
	}

	if err := db.Where("id = ?", id).First(&mto).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorx.ErrMaterialTestOrderNotFound
		}
		return nil, err
	}
	return &mto, nil
}

func (r *GormMaterialTestOrderRepository) Save(mto *entity.MaterialTestOrder) error {
	err := r.db.Omit("WorkCategory", "OrderedServices", "OrderUsers", "Medias").Save(mto).Error
	if err != nil {
		if helper.IsPostgresDuplicateKeyError(err) {
			return errorx.ErrMaterialTestOrderAlreadyExists
		}
		return err
	}
	return nil
}

func (r *GormMaterialTestOrderRepository) SaveMany(mtos []*entity.MaterialTestOrder) error {
	return r.db.CreateInBatches(mtos, 100).Error
}

func (r *GormMaterialTestOrderRepository) Destroy(mto *entity.MaterialTestOrder) error {
	return r.db.Delete(mto).Error
}
