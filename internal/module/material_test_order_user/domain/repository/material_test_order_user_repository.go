package repository

import "github.com/arfanxn/welding/internal/module/shared/domain/entity"

type MaterialTestOrderUserRepository interface {
	Save(orderUser *entity.MaterialTestOrderUser) error
	SaveMany(orderUsers []*entity.MaterialTestOrderUser) error
	DestroyByUserId(userId string) error
	Destroy(orderUser *entity.MaterialTestOrderUser) error

	// DestroyOwnerByOrderId deletes all owner type MaterialTestOrderUser records associated with the given order ID.
	// It returns an error if the operation fails, or nil if the operation is successful.
	// Parameters:
	//   - orderId: The ID of the material test order whose owner records should be deleted
	//
	// Returns:
	//   - error: An error if the deletion fails, nil otherwise
	DestroyOwnerByOrderId(orderId string) error
}
