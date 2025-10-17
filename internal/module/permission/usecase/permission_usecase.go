package usecase

import (
	"github.com/arfanxn/welding/internal/module/permission/domain/repository"
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	"github.com/arfanxn/welding/internal/module/shared/usecase/dto"
)

type PermissionUsecase interface {
	Paginate(queryDto *dto.Query) (*dto.Pagination[*entity.Permission], error)
}

type permissionUsecase struct {
	permissionRepository repository.PermissionRepository
}

func NewPermissionUsecase(permissionRepository repository.PermissionRepository) PermissionUsecase {
	return &permissionUsecase{
		permissionRepository: permissionRepository,
	}
}

func (u *permissionUsecase) Paginate(queryDto *dto.Query) (*dto.Pagination[*entity.Permission], error) {
	return u.permissionRepository.Paginate(queryDto)
}
