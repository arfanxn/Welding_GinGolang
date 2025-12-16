package repository

import (
	"github.com/arfanxn/welding/internal/module/material_test_order_user/domain/repository"
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	"gorm.io/gorm"

	materialTestOrderUserEnum "github.com/arfanxn/welding/internal/module/material_test_order_user/domain/enum"
)

var _ repository.MaterialTestOrderUserRepository = (*GormMaterialTestOrderUserRepository)(nil)

type GormMaterialTestOrderUserRepository struct {
	db *gorm.DB
}

func NewGormMaterialTestOrderUserRepository(db *gorm.DB) repository.MaterialTestOrderUserRepository {
	return &GormMaterialTestOrderUserRepository{
		db: db,
	}
}

func (r *GormMaterialTestOrderUserRepository) Save(materialTestOrderUser *entity.MaterialTestOrderUser) error {
	return r.db.Save(materialTestOrderUser).Error
}

func (r *GormMaterialTestOrderUserRepository) SaveMany(materialTestOrderUsers []*entity.MaterialTestOrderUser) error {
	return r.db.CreateInBatches(materialTestOrderUsers, 100).Error
}

func (r *GormMaterialTestOrderUserRepository) DestroyByUserId(userId string) error {
	return r.db.Delete(&entity.MaterialTestOrderUser{}, "user_id = ?", userId).Error
}

func (r *GormMaterialTestOrderUserRepository) Destroy(materialTestOrderUser *entity.MaterialTestOrderUser) error {
	return r.db.Delete(materialTestOrderUser).Error
}

// DestroyOwnerByOrderId deletes all owner type MaterialTestOrderUser records associated with the given order ID.
// It returns an error if the operation fails, or nil if the operation is successful.
// Parameters:
//   - orderId: The ID of the material test order whose owner records should be deleted
//
// Returns:
//   - error: An error if the deletion fails, nil otherwise
func (r *GormMaterialTestOrderUserRepository) DestroyOwnerByOrderId(orderId string) error {
	return r.db.Delete(&entity.MaterialTestOrderUser{}, "order_id = ? AND type = ?", orderId, materialTestOrderUserEnum.TypeOwner).Error
}
