package repository

import (
	permissionEnum "github.com/arfanxn/welding/internal/module/permission/domain/enum"
	roleEnum "github.com/arfanxn/welding/internal/module/role/domain/enum"
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	"github.com/arfanxn/welding/pkg/pagination"
	"github.com/arfanxn/welding/pkg/query"
)

type UserRepository interface {
	Get(query *query.Query) ([]*entity.User, error)
	Paginate(query *query.Query) (*pagination.OffsetPagination[*entity.User], error)
	First(query *query.Query) (*entity.User, error)
	Find(id string) (*entity.User, error)
	FindByEmail(email string) (*entity.User, error)
	HasPermissionNames(user *entity.User, permissionNames []permissionEnum.PermissionName) (bool, error)
	HasRoleNames(user *entity.User, roleNames []roleEnum.RoleName) (bool, error)
	ToggleActivation(user *entity.User) (*entity.User, error)
	Save(user *entity.User) error
	SaveMany(users []*entity.User) error
	Destroy(user *entity.User) error
}
