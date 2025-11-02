package repository

import (
	"github.com/arfanxn/welding/internal/module/code/domain/enum"
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
)

type CodeRepository interface {
	Find(id string) (*entity.Code, error)
	FindByValue(value string) (*entity.Code, error)
	FindByTypeAndValue(_type enum.CodeType, value string) (*entity.Code, error)
	FindByCodeableAndTypeAndValue(codeableId string, codeableType string, _type enum.CodeType, value string) (*entity.Code, error)
	Save(code *entity.Code) error
	SaveMany(codes []*entity.Code) error
	Destroy(code *entity.Code) error
}
