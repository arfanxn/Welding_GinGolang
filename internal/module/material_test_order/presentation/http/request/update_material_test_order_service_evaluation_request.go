package request

import (
	"github.com/arfanxn/welding/internal/infrastructure/http/request"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

var _ request.Request = (*UpdateMaterialTestOrderServiceEvaluation)(nil)

type UpdateMaterialTestOrderServiceEvaluation struct {
	Id                       string `form:"id" json:"id"`
	OrderServiceEvaluationId string `form:"order_service_evaluation_id" json:"order_service_evaluation_id"`

	IsEquipmentAvailable      bool `form:"is_equipment_available" json:"is_equipment_available"`
	IsPersonnelAvailable      bool `form:"is_personnel_available" json:"is_personnel_available"`
	IsTimeAvailable           bool `form:"is_time_available" json:"is_time_available"`
	IsTestReady               bool `form:"is_test_ready" json:"is_test_ready"`
	IsSubcontractLabAvailable bool `form:"is_subcontract_lab_available" json:"is_subcontract_lab_available"`
}

func NewUpdateMaterialTestOrderServiceEvaluation() *UpdateMaterialTestOrderServiceEvaluation {
	return &UpdateMaterialTestOrderServiceEvaluation{}
}

func (r *UpdateMaterialTestOrderServiceEvaluation) Validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.Id,
			validation.Required.Error("Id tidak boleh kosong"),
			validation.Length(26, 26).Error("Id harus 26 karakter"),
		),
		validation.Field(&r.OrderServiceEvaluationId,
			validation.Required.Error("Order service evaluation id tidak boleh kosong"),
			validation.Length(26, 26).Error("Order service evaluation id harus 26 karakter"),
		),
		validation.Field(&r.IsEquipmentAvailable),
		validation.Field(&r.IsPersonnelAvailable),
		validation.Field(&r.IsTimeAvailable),
		validation.Field(&r.IsTestReady),
		validation.Field(&r.IsSubcontractLabAvailable),
	)
}
