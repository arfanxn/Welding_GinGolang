package request

// import (
// 	"fmt"

// 	"github.com/arfanxn/welding/internal/infrastructure/http/request"
// 	materialTestOrderEnum "github.com/arfanxn/welding/internal/module/material_test_order/domain/enum"
// 	validation "github.com/go-ozzo/ozzo-validation/v4"
// )

// var _ request.Request = (*UpdateMaterialTestOrderStatus)(nil)

// type UpdateMaterialTestOrderStatus struct {
// 	Id     string `form:"id" json:"id"`
// 	Status string `form:"status" json:"status"`
// }

// func NewUpdateMaterialTestOrderStatus() *UpdateMaterialTestOrderStatus {
// 	return &UpdateMaterialTestOrderStatus{}
// }

// func (r *UpdateMaterialTestOrderStatus) Validate() error {
// 	// Why payment_submitted and refunded arent not in the list?
// 	// Because payment_submitted and refunded have different flow

// 	allowedStatuses := []materialTestOrderEnum.MaterialTestOrderStatus{
// 		materialTestOrderEnum.MaterialTestOrderStatusRejected,
// 		materialTestOrderEnum.MaterialTestOrderStatusAwaitingPayment,
// 		materialTestOrderEnum.MaterialTestOrderStatusCancelled,
// 		materialTestOrderEnum.MaterialTestOrderStatusPaymentRejected,
// 		materialTestOrderEnum.MaterialTestOrderStatusPaymentApproved,
// 		materialTestOrderEnum.MaterialTestOrderStatusTesting,
// 		materialTestOrderEnum.MaterialTestOrderStatusCompleted,
// 	}

// 	return validation.ValidateStruct(r,
// 		validation.Field(&r.Id,
// 			validation.Required.Error("Id wajib disi"),
// 			validation.Length(26, 26).Error("Id harus 26 karakter"),
// 		),
// 		validation.Field(&r.Status,
// 			validation.Required.Error("Status wajib disi"),
// 			validation.In(
// 				materialTestOrderEnum.MaterialTestOrderStatusDraft, materialTestOrderEnum.MaterialTestOrderStatusAwaitingReview).
// 				Error(fmt.Sprintf("Status tidak valid. allowed: '%s', '%s'",
// 					materialTestOrderEnum.MaterialTestOrderStatusDraft,
// 					materialTestOrderEnum.MaterialTestOrderStatusAwaitingReview),
// 				),
// 		),
// 	)
// }
