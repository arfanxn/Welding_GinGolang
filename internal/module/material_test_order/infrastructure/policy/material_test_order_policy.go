package policy

import (
	"context"
	"errors"

	mtoEnum "github.com/arfanxn/welding/internal/module/material_test_order/domain/enum"
	mtoRepository "github.com/arfanxn/welding/internal/module/material_test_order/domain/repository"
	"github.com/arfanxn/welding/internal/module/material_test_order/usecase/dto"
	mtosRepository "github.com/arfanxn/welding/internal/module/material_test_order_service/domain/repository"
	mtoseRepository "github.com/arfanxn/welding/internal/module/material_test_order_service_evaluation/domain/repository"
	mtsRepository "github.com/arfanxn/welding/internal/module/material_test_service/domain/repository"
	mtwcRepository "github.com/arfanxn/welding/internal/module/material_test_work_category/domain/repository"
	mediaEnum "github.com/arfanxn/welding/internal/module/media/domain/enum"
	mediaRepository "github.com/arfanxn/welding/internal/module/media/domain/repository"
	permissionEnum "github.com/arfanxn/welding/internal/module/permission/domain/enum"
	permissionService "github.com/arfanxn/welding/internal/module/permission/usecase/service"
	"github.com/arfanxn/welding/internal/module/shared/contextkey"
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	"github.com/arfanxn/welding/internal/module/shared/domain/errorx"
	userRepository "github.com/arfanxn/welding/internal/module/user/domain/repository"
	"github.com/gookit/goutil"
	"github.com/samber/lo"
	"go.uber.org/fx"
)

type MaterialTestOrderPolicy interface {
	Store(ctx context.Context, _dto *dto.SaveMaterialTestOrder) error
	Update(ctx context.Context, _dto *dto.SaveMaterialTestOrder) error

	// Status relateds
	Approve(context.Context, *dto.ApproveMaterialTestOrder) error
	Reject(context.Context, *dto.RejectMaterialTestOrder) error
	Cancel(context.Context, *dto.CancelMaterialTestOrder) error
	SubmitPayment(context.Context, *dto.SubmitPaymentMaterialTestOrder) error
	ApprovePayment(context.Context, *dto.ApprovePaymentMaterialTestOrder) error
	RejectPayment(context.Context, *dto.RejectPaymentMaterialTestOrder) error
	Test(context.Context, *dto.TestMaterialTestOrder) error
	Refund(context.Context, *dto.RefundMaterialTestOrder) error
	Complete(context.Context, *dto.CompleteMaterialTestOrder) error

	UpdateOrderServiceEvaluation(ctx context.Context, _dto *dto.UpdateMaterialTestOrderServiceEvaluation) error

	StoreMedia(ctx context.Context, _dto *dto.SaveMaterialTestOrderMedia) error
	UpdateMedia(ctx context.Context, _dto *dto.SaveMaterialTestOrderMedia) error
	DestroyMedia(ctx context.Context, _dto *dto.DestroyMaterialTestOrderMedia) error
}

type materialTestOrderPolicy struct {
	permissionService permissionService.PermissionService
	mtoRepository     mtoRepository.MaterialTestOrderRepository
	mtsRepository     mtsRepository.MaterialTestServiceRepository
	mtosRepository    mtosRepository.MaterialTestOrderServiceRepository
	mtoseRepository   mtoseRepository.MaterialTestOrderServiceEvaluationRepository
	mtwcRepository    mtwcRepository.MaterialTestWorkCategoryRepository
	userRepository    userRepository.UserRepository
	mediaRepository   mediaRepository.MediaRepository
}

type NewMaterialTestOrderPolicyParams struct {
	fx.In

	PermissionService                            permissionService.PermissionService
	MaterialTestOrderRepository                  mtoRepository.MaterialTestOrderRepository
	MaterialTestOrderServiceRepository           mtosRepository.MaterialTestOrderServiceRepository
	MaterialTestOrderServiceEvaluationRepository mtoseRepository.MaterialTestOrderServiceEvaluationRepository
	MaterialTestServiceRepository                mtsRepository.MaterialTestServiceRepository
	MaterialTestWorkCategoryRepository           mtwcRepository.MaterialTestWorkCategoryRepository
	UserRepository                               userRepository.UserRepository
	MediaRepository                              mediaRepository.MediaRepository
}

func NewMaterialTestOrderPolicy(params NewMaterialTestOrderPolicyParams) MaterialTestOrderPolicy {
	return &materialTestOrderPolicy{
		permissionService: params.PermissionService,
		mtoRepository:     params.MaterialTestOrderRepository,
		mtosRepository:    params.MaterialTestOrderServiceRepository,
		mtoseRepository:   params.MaterialTestOrderServiceEvaluationRepository,
		mtsRepository:     params.MaterialTestServiceRepository,
		mtwcRepository:    params.MaterialTestWorkCategoryRepository,
		userRepository:    params.UserRepository,
		mediaRepository:   params.MediaRepository,
	}
}

func (p *materialTestOrderPolicy) Store(ctx context.Context, _dto *dto.SaveMaterialTestOrder) error {
	if err := p.validateWorkCategory(*_dto.WorkCategoryId); err != nil {
		return err
	}

	if _dto.OwnerUserIds != nil {
		if len(_dto.OwnerUserIds) > 0 && !goutil.IsEmptyReal(_dto.OwnerUserIds[0]) {
			if err := p.validateUsers(_dto.OwnerUserIds); err != nil {
				return err
			}
		}
	}

	if _dto.OrderedServices != nil {
		if len(_dto.OrderedServices) > 0 && !goutil.IsEmptyReal(_dto.OrderedServices[0]) {
			serviceIds := []string{}
			for _, orderedServiceDto := range _dto.OrderedServices {
				serviceIds = append(serviceIds, orderedServiceDto.ServiceId)
			}

			if err := p.validateServices(serviceIds); err != nil {
				return err
			}

		}
	}

	return nil
}

func (p *materialTestOrderPolicy) Update(ctx context.Context, _dto *dto.SaveMaterialTestOrder) error {
	mto, err := p.mtoRepository.Find(*_dto.Id, nil)
	if err != nil {
		return err
	}

	if mto.IsAwaitingReview() {
		_hasPermission, err := p.hasPermission(ctx, permissionEnum.MaterialTestOrdersUpdate)
		if err != nil {
			return err
		} else if !_hasPermission {
			return errorx.ErrMaterialTestOrderUpdateForbidden
		}
	} else if mto.IsDraft() || mto.IsRejected() {
		_hasPermission, err := p.hasPermissionOrOwnership(ctx, permissionEnum.MaterialTestOrdersUpdate, *_dto.Id)
		if err != nil {
			return err
		} else if !_hasPermission {
			return errorx.ErrMaterialTestOrderUpdateForbidden
		}
	} else {
		return errorx.ErrMaterialTestOrderStatusNotUpdateableUpdateForbidden
	}

	if !goutil.IsEmptyReal(_dto.WorkCategoryId) {
		if err := p.validateWorkCategory(*_dto.WorkCategoryId); err != nil {
			return err
		}
	}

	if _dto.OwnerUserIds != nil {
		if len(_dto.OwnerUserIds) > 0 && !goutil.IsEmptyReal(_dto.OwnerUserIds[0]) {
			if err := p.validateUsers(_dto.OwnerUserIds); err != nil {
				return err
			}
		}
	}

	if _dto.OrderedServices != nil {
		if len(_dto.OrderedServices) > 0 && !goutil.IsEmptyReal(_dto.OrderedServices[0]) {
			serviceIds := []string{}
			for _, orderedServiceDto := range _dto.OrderedServices {
				serviceIds = append(serviceIds, orderedServiceDto.ServiceId)

			}

			if err := p.validateServices(serviceIds); err != nil {
				return err
			}

		}
	}

	return nil
}

/* ==============================
	Status related policies
============================== */

func (p *materialTestOrderPolicy) Approve(ctx context.Context, _dto *dto.ApproveMaterialTestOrder) error {
	mto, err := p.mtoRepository.Find(_dto.Id, nil)
	if err != nil {
		return err
	}

	if !mto.IsAwaitingReview() {
		return errorx.ErrMaterialTestOrderStatusNotAwaitingReviewApproveForbidden
	}

	return nil
}

func (p *materialTestOrderPolicy) Reject(ctx context.Context, _dto *dto.RejectMaterialTestOrder) error {
	mto, err := p.mtoRepository.Find(_dto.Id, nil)
	if err != nil {
		return err
	}

	if !mto.IsAwaitingReview() {
		return errorx.ErrMaterialTestOrderStatusNotAwaitingReviewRejectForbidden
	}

	return nil
}

func (p *materialTestOrderPolicy) Cancel(ctx context.Context, _dto *dto.CancelMaterialTestOrder) error {
	mto, err := p.mtoRepository.Find(_dto.Id, nil)
	if err != nil {
		return err
	}

	cancellableStatuses := []mtoEnum.MaterialTestOrderStatus{
		mtoEnum.MaterialTestOrderStatusDraft,
		mtoEnum.MaterialTestOrderStatusAwaitingReview,
		mtoEnum.MaterialTestOrderStatusRejected,
		mtoEnum.MaterialTestOrderStatusAwaitingPayment,
		mtoEnum.MaterialTestOrderStatusPaymentRejected,
	}

	isCancellable := lo.Contains(cancellableStatuses, mto.Status)
	if !isCancellable {
		return errorx.ErrMaterialTestOrderStatusNotCancellableCancelForbidden
	}

	_hasPermissionOrOwnership, err := p.hasPermissionOrOwnership(ctx, permissionEnum.MaterialTestOrdersUpdate, mto.Id)
	if err != nil {
		return err
	} else if !_hasPermissionOrOwnership {
		return errorx.ErrMaterialTestOrderUpdateForbidden
	}

	return nil
}

func (p *materialTestOrderPolicy) SubmitPayment(ctx context.Context, _dto *dto.SubmitPaymentMaterialTestOrder) error {
	mto, err := p.mtoRepository.Find(_dto.Id, nil)
	if err != nil {
		return err
	}

	isPaymentSubmittable := mto.IsAwaitingPayment() || mto.IsPaymentSubmitted() || mto.IsPaymentRejected()
	if !isPaymentSubmittable {
		return errorx.ErrMaterialTestOrderStatusNotPaymentSubmittableSubmitPaymentForbidden
	}

	_hasPermissionOrOwnership, err := p.hasPermissionOrOwnership(ctx, permissionEnum.MaterialTestOrdersUpdate, mto.Id)
	if err != nil {
		return err
	} else if !_hasPermissionOrOwnership {
		return errorx.ErrMaterialTestOrderUpdateForbidden
	}

	return nil
}

func (p *materialTestOrderPolicy) ApprovePayment(ctx context.Context, _dto *dto.ApprovePaymentMaterialTestOrder) error {
	mto, err := p.mtoRepository.Find(_dto.Id, nil)
	if err != nil {
		return err
	}

	if !mto.IsPaymentSubmitted() {
		return errorx.ErrMaterialTestOrderStatusNotPaymentSubmittedApprovePaymentForbidden
	}

	return nil
}

func (p *materialTestOrderPolicy) RejectPayment(ctx context.Context, _dto *dto.RejectPaymentMaterialTestOrder) error {
	mto, err := p.mtoRepository.Find(_dto.Id, nil)
	if err != nil {
		return err
	}

	if !mto.IsPaymentSubmitted() {
		return errorx.ErrMaterialTestOrderStatusNotPaymentSubmittedRejectPaymentForbidden
	}

	return nil
}

func (p *materialTestOrderPolicy) Test(ctx context.Context, _dto *dto.TestMaterialTestOrder) error {
	mto, err := p.mtoRepository.Find(_dto.Id, nil)
	if err != nil {
		return err
	}

	if !mto.IsPaymentApproved() {
		return errorx.ErrMaterialTestOrderStatusNotPaymentApprovedTestForbidden
	}

	return nil
}

func (p *materialTestOrderPolicy) Refund(ctx context.Context, _dto *dto.RefundMaterialTestOrder) error {
	mto, err := p.mtoRepository.Find(_dto.Id, nil)
	if err != nil {
		return err
	}

	refundableStatuses := []mtoEnum.MaterialTestOrderStatus{
		mtoEnum.MaterialTestOrderStatusPaymentApproved,
		mtoEnum.MaterialTestOrderStatusTesting,
		mtoEnum.MaterialTestOrderStatusCompleted,
		mtoEnum.MaterialTestOrderStatusRefunded, // why? for reuploading the refund proof if mistaken
	}

	isRefundable := lo.Contains(refundableStatuses, mto.Status)
	if !isRefundable {
		return errorx.ErrMaterialTestOrderStatusNotRefundableRefundForbidden
	}

	return nil
}

func (p *materialTestOrderPolicy) Complete(ctx context.Context, _dto *dto.CompleteMaterialTestOrder) error {
	mto, err := p.mtoRepository.Find(_dto.Id, nil)
	if err != nil {
		return err
	}

	if !mto.IsTesting() {
		return errorx.ErrMaterialTestOrderStatusNotTestingCompleteForbidden
	}

	return nil
}

/* ==============================
	Service evaluation related policy
============================== */

func (p *materialTestOrderPolicy) UpdateOrderServiceEvaluation(ctx context.Context, _dto *dto.UpdateMaterialTestOrderServiceEvaluation) error {
	mto, err := p.mtoRepository.Find(_dto.Id, nil)
	if err != nil {
		return err
	}

	_, err = p.mtoseRepository.Find(_dto.OrderServiceEvaluationId, nil)
	if err != nil {
		return err
	}

	if !mto.IsAwaitingReview() {
		return errorx.ErrMaterialTestOrderStatusNotAwaitingReviewUpdateServiceEvaluationForbidden
	}

	return nil
}

/* ==============================
	Media related policies
============================== */

func (p *materialTestOrderPolicy) StoreMedia(ctx context.Context, _dto *dto.SaveMaterialTestOrderMedia) error {
	mto, err := p.mtoRepository.Find(*_dto.Id, nil)
	if err != nil {
		return err
	}

	_hasPermissionOrOwnership, err := p.hasPermissionOrOwnership(ctx, permissionEnum.MaterialTestOrdersUpdate, mto.Id)
	if err != nil {
		return err
	} else if !_hasPermissionOrOwnership {
		return errorx.ErrMaterialTestOrderUpdateForbidden
	}

	return nil
}

func (p *materialTestOrderPolicy) UpdateMedia(ctx context.Context, _dto *dto.SaveMaterialTestOrderMedia) error {
	mto, err := p.mtoRepository.Find(*_dto.Id, nil)
	if err != nil {
		return err
	}

	_hasPermissionOrOwnership, err := p.hasPermissionOrOwnership(ctx, permissionEnum.MaterialTestOrdersUpdate, mto.Id)
	if err != nil {
		return err
	} else if !_hasPermissionOrOwnership {
		return errorx.ErrMaterialTestOrderUpdateForbidden
	}

	media, err := p.mediaRepository.Find(*_dto.MediaId, nil)
	if err != nil {
		return err
	}

	switch media.CollectionName {
	case mediaEnum.CollectionNameMaterialTestOrderPaymentProof:
		return errorx.ErrMaterialTestOrderMediaPaymentProofUpdateForbidden
	case mediaEnum.CollectionNameMaterialTestOrderRefundProof:
		return errorx.ErrMaterialTestOrderMediaRefundProofUpdateForbidden
	}

	return nil
}

func (p *materialTestOrderPolicy) DestroyMedia(ctx context.Context, _dto *dto.DestroyMaterialTestOrderMedia) error {
	mto, err := p.mtoRepository.Find(_dto.Id, nil)
	if err != nil {
		return err
	}

	_hasPermissionOrOwnership, err := p.hasPermissionOrOwnership(ctx, permissionEnum.MaterialTestOrdersUpdate, mto.Id)
	if err != nil {
		return err
	} else if !_hasPermissionOrOwnership {
		return errorx.ErrMaterialTestOrderUpdateForbidden
	}

	media, err := p.mediaRepository.Find(_dto.MediaId, nil)
	if err != nil {
		return err
	}

	switch media.CollectionName {
	case mediaEnum.CollectionNameMaterialTestOrderPaymentProof:
		return errorx.ErrMaterialTestOrderMediaPaymentProofDestroyForbidden
	case mediaEnum.CollectionNameMaterialTestOrderRefundProof:
		return errorx.ErrMaterialTestOrderMediaRefundProofDestroyForbidden
	}

	return nil
}

/* ==============================
	Private Helper Methods
============================== */

/*
func (p *materialTestOrderPolicy) validateOrderServices(orderServiceIds []string) error {
	_, err := p.mtosRepository.FindByIds(orderServiceIds, nil)
	if err != nil {
		return err
	}
	return nil
}
*/

func (p *materialTestOrderPolicy) hasPermission(ctx context.Context, requiredPermName permissionEnum.PermissionName) (bool, error) {
	userPermissions := ctx.Value(contextkey.UserPermissionsKey).([]*entity.Permission)
	_hasPermission, err := p.permissionService.CheckByNames(userPermissions, requiredPermName)
	if err != nil {
		return false, err
	}
	return _hasPermission, nil
}

func (p *materialTestOrderPolicy) hasPermissionOrOwnership(ctx context.Context, requiredPermName permissionEnum.PermissionName, orderId string) (bool, error) {
	userId := ctx.Value(contextkey.UserIdKey).(string)

	_hasPermission, err := p.hasPermission(ctx, requiredPermName)
	if err != nil {
		return false, err
	} else if _hasPermission {
		return true, nil
	}

	_, err = p.mtoRepository.FindByIdAndUserId(orderId, userId, nil)
	if err != nil {
		if errors.Is(err, errorx.ErrMaterialTestOrderNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (p *materialTestOrderPolicy) validateServices(serviceIds []string) error {
	_, err := p.mtsRepository.FindByIds(serviceIds, nil)
	if err != nil {
		return err
	}
	return nil
}

func (p *materialTestOrderPolicy) validateWorkCategory(WorkCategoryId string) error {
	_, err := p.mtwcRepository.Find(WorkCategoryId, nil)
	if err != nil {
		return err
	}
	return nil
}

func (p *materialTestOrderPolicy) validateUsers(userIds []string) error {
	_, err := p.userRepository.FindByIds(userIds, nil)
	if err != nil {
		return err
	}
	return nil
}
