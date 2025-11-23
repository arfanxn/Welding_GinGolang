package repository

import (
	"errors"

	"github.com/arfanxn/welding/internal/infrastructure/database/helper"
	activityEnum "github.com/arfanxn/welding/internal/module/activity/domain/enum"
	"github.com/arfanxn/welding/internal/module/activity/domain/repository"
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	"github.com/arfanxn/welding/internal/module/shared/domain/errorx"
	"github.com/arfanxn/welding/pkg/pagination"
	"github.com/arfanxn/welding/pkg/query"
	"github.com/gookit/goutil"
	"gorm.io/gorm"
)

type gormActivityRepository struct {
	db *gorm.DB
}

func NewGormActivityRepository(db *gorm.DB) repository.ActivityRepository {
	return &gormActivityRepository{db: db}
}

func (r *gormActivityRepository) query(db *gorm.DB, q *query.Query) (*gorm.DB, error) {
	activityTableName := entity.NewActivity().TableName()

	if q != nil {
		if f := q.GetFilterById(); f != nil {
			db = db.Where(activityTableName+".id = ?", f.Value)
		}

		if search := q.GetSearch(); search != nil {
			db = db.Where(activityTableName+".description ILIKE ?", "%"+*search+"%")
		}

		if action := q.GetFilter("action", query.OperatorEqual); action != nil {
			if !goutil.Contains(activityEnum.ActivityActions, activityEnum.ActivityAction(action.Value)) {
				return nil, errorx.ErrActivityInvalidAction
			}
			db = db.Where(activityTableName+".action = ?", action.Value)
		}

		if causerType := q.GetFilter("causer_type", query.OperatorEqual); causerType != nil {
			if !goutil.Contains(activityEnum.ActivityCauserTypes, activityEnum.ActivityCauserType(causerType.Value)) {
				return nil, errorx.ErrActivityInvalidCauserType
			}
			db = db.Where(activityTableName+".causer_type = ?", causerType.Value)
		}

		if causerId := q.GetFilter("causer_id", query.OperatorEqual); causerId != nil {
			db = db.Where(activityTableName+".causer_id = ?", causerId.Value)
		}

		if causerIpAddress := q.GetFilter("causer_ip_address", query.OperatorEqual); causerIpAddress != nil {
			db = db.Where(activityTableName+".causer_ip_address = ?", causerIpAddress.Value)
		}

		if subjectType := q.GetFilter("subject_type", query.OperatorEqual); subjectType != nil {
			if !goutil.Contains(activityEnum.ActivitySubjectTypes, activityEnum.ActivitySubjectType(subjectType.Value)) {
				return nil, errorx.ErrActivityInvalidSubjectType
			}
			db = db.Where(activityTableName+".subject_type = ?", subjectType.Value)
		}

		if subjectId := q.GetFilter("subject_id", query.OperatorEqual); subjectId != nil {
			db = db.Where(activityTableName+".subject_id = ?", subjectId.Value)
		}

		if createdAtBetween := q.GetFilter("created_at", query.OperatorBetween); createdAtBetween != nil {
			db = db.Where(activityTableName+".created_at BETWEEN ? AND ?", createdAtBetween.Values[0], createdAtBetween.Values[1])
		}

		if q.GetInclude("causer") != nil {
			db = db.Preload("CauserUser")
		}

		if q.GetInclude("subject") != nil {
			db = db.
				Preload("SubjectUser").
				Preload("SubjectRole").
				Preload("SubjectPermission").
				Preload("SubjectCode")
		}

		if sort := q.GetSort("action"); sort != nil {
			db = db.Order(activityTableName + ".action" + sort.Order)
		}

		if sort := q.GetSort("created_at"); sort != nil {
			db = db.Order(activityTableName + ".created_at " + sort.Order)
		}
	}

	return db, nil
}

func (r *gormActivityRepository) Paginate(q *query.Query) (*pagination.OffsetPagination[*entity.Activity], error) {
	db := r.db.Model(&entity.Activity{})

	db, err := r.query(db, q)
	if err != nil {
		return nil, err
	}

	op, err := helper.GormDBPaginateWithQuery[*entity.Activity](db, q)
	if err != nil {
		return nil, err
	}
	return op, nil
}

func (r *gormActivityRepository) First(q *query.Query) (*entity.Activity, error) {
	var activity *entity.Activity

	db, err := r.query(r.db, q)
	if err != nil {
		return nil, err
	}

	if err := db.First(&activity).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorx.ErrActivityNotFound
		}
		return nil, err
	}

	return activity, nil
}

func (r *gormActivityRepository) Find(id string) (*entity.Activity, error) {
	var activity entity.Activity
	if err := r.db.Where("id = ?", id).First(&activity).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorx.ErrActivityNotFound
		}
		return nil, err
	}
	return &activity, nil
}

func (r *gormActivityRepository) Save(activity *entity.Activity) error {
	return r.db.Save(activity).Error
}

func (r *gormActivityRepository) SaveMany(activities []*entity.Activity) error {
	return r.db.CreateInBatches(activities, 100).Error
}
