package repository

import (
	"github.com/arfanxn/welding/internal/module/code/domain/enum"
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	"github.com/arfanxn/welding/pkg/query"
)

type CodeRepository interface {
	Find(id string, q *query.Query) (*entity.Code, error)
	FindByValue(value string, q *query.Query) (*entity.Code, error)
	FindByType(_type enum.CodeType, q *query.Query) (*entity.Code, error)
	FindByTypeAndValue(_type enum.CodeType, value string, q *query.Query) (*entity.Code, error)
	FindByCodeableAndTypeAndValue(codeableId string, codeableType string, _type enum.CodeType, value string, q *query.Query) (*entity.Code, error)
	Save(code *entity.Code) error
	SaveMany(codes []*entity.Code) error
	Destroy(code *entity.Code) error
}
