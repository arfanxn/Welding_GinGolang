package usecase

import (
	"github.com/arfanxn/welding/internal/module/permission/domain/repository"
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	"github.com/arfanxn/welding/pkg/pagination"
	"github.com/arfanxn/welding/pkg/query"
)

type PermissionUsecase interface {
	Paginate(q *query.Query) (*pagination.OffsetPagination[*entity.Permission], error)
}

type permissionUsecase struct {
	permissionRepository repository.PermissionRepository
}

func NewPermissionUsecase(permissionRepository repository.PermissionRepository) PermissionUsecase {
	return &permissionUsecase{
		permissionRepository: permissionRepository,
	}
}

func (u *permissionUsecase) Paginate(q *query.Query) (*pagination.OffsetPagination[*entity.Permission], error) {
	return u.permissionRepository.Paginate(q)
}
