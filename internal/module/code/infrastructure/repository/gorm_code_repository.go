package repository

import (
	"errors"

	"github.com/arfanxn/welding/internal/infrastructure/database/helper"
	"github.com/arfanxn/welding/internal/module/code/domain/enum"
	"github.com/arfanxn/welding/internal/module/code/domain/repository"
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	"github.com/arfanxn/welding/internal/module/shared/domain/errorx"
	"gorm.io/gorm"
)

var _ repository.CodeRepository = (*GormCodeRepository)(nil)

type GormCodeRepository struct {
	db *gorm.DB
}

func NewGormCodeRepository(db *gorm.DB) repository.CodeRepository {
	return &GormCodeRepository{
		db: db,
	}
}

func (r *GormCodeRepository) Find(id string) (*entity.Code, error) {
	var code entity.Code

	if err := r.db.Where("id = ?", id).First(&code).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorx.ErrCodeNotFound
		}
		return nil, err
	}
	return &code, nil
}

func (r *GormCodeRepository) FindByValue(value string) (*entity.Code, error) {
	var code entity.Code
	if err := r.db.Where("value = ?", value).First(&code).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorx.ErrCodeNotFound
		}
		return nil, err
	}
	return &code, nil
}

func (r *GormCodeRepository) FindByTypeAndValue(_type enum.CodeType, value string) (*entity.Code, error) {
	var code entity.Code
	if err := r.db.Where("type = ? AND value = ?", _type, value).First(&code).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorx.ErrCodeNotFound
		}
		return nil, err
	}
	return &code, nil
}

func (r *GormCodeRepository) FindByCodeableAndTypeAndValue(codeableId string, codeableType string, _type enum.CodeType, value string) (*entity.Code, error) {
	var code entity.Code
	if err := r.db.Where(
		"codeable_id = ? AND codeable_type = ? AND type = ? AND value = ?",
		codeableId, codeableType, _type, value,
	).First(&code).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorx.ErrCodeNotFound
		}
		return nil, err
	}
	return &code, nil
}

func (r *GormCodeRepository) Save(code *entity.Code) error {
	err := r.db.Save(code).Error
	if err != nil {
		if helper.IsPostgresDuplicateKeyError(err) {
			return errorx.ErrCodeAlreadyExists
		}
		return err
	}
	return nil
}

func (r *GormCodeRepository) SaveMany(codes []*entity.Code) error {
	return r.db.CreateInBatches(codes, 100).Error
}

func (r *GormCodeRepository) Destroy(code *entity.Code) error {
	return r.db.Delete(code).Error
}
