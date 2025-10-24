package repository

import (
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	"github.com/arfanxn/welding/pkg/pagination"
	"github.com/arfanxn/welding/pkg/query"
)

type UserRepository interface {
	Get(query *query.Query) ([]*entity.User, error)
	Paginate(query *query.Query) (*pagination.OffsetPagination[*entity.User], error)
	Find(id string) (*entity.User, error)
	FindByEmail(email string) (*entity.User, error)
	Save(user *entity.User) error
	SaveMany(users []*entity.User) error
	Destroy(user *entity.User) error
}
