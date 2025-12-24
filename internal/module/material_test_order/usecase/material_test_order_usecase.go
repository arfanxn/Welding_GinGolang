package usecase

import (
	"context"

	activityEnum "github.com/arfanxn/welding/internal/module/activity/domain/enum"
	activityDto "github.com/arfanxn/welding/internal/module/activity/usecase/dto"
	activityService "github.com/arfanxn/welding/internal/module/activity/usecase/service"
	mtoRepository "github.com/arfanxn/welding/internal/module/material_test_order/domain/repository"
	mtoPolicy "github.com/arfanxn/welding/internal/module/material_test_order/infrastructure/policy"
	"github.com/arfanxn/welding/internal/module/material_test_order/usecase/dto"
	mtoStep "github.com/arfanxn/welding/internal/module/material_test_order/usecase/step"
	mediaService "github.com/arfanxn/welding/internal/module/media/usecase/service"
	"github.com/arfanxn/welding/internal/module/shared/contextkey"
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	"github.com/arfanxn/welding/pkg/pagination"
	"github.com/arfanxn/welding/pkg/query"
	"github.com/arfanxn/welding/pkg/typeutil"
	"go.uber.org/fx"
)

type MaterialTestOrderUsecase interface {
	Paginate(context.Context, *query.Query) (*pagination.OffsetPagination[*entity.MaterialTestOrder], error)
	Show(context.Context, *query.Query) (*entity.MaterialTestOrder, error)
	Store(context.Context, *dto.SaveMaterialTestOrder) (*entity.MaterialTestOrder, error)
	Update(context.Context, *dto.SaveMaterialTestOrder) (*entity.MaterialTestOrder, error)

	// Status relateds
	Approve(context.Context, *dto.ApproveMaterialTestOrder) (*entity.MaterialTestOrder, error)
	Reject(context.Context, *dto.RejectMaterialTestOrder) (*entity.MaterialTestOrder, error)
	Cancel(context.Context, *dto.CancelMaterialTestOrder) (*entity.MaterialTestOrder, error)
	SubmitPayment(context.Context, *dto.SubmitPaymentMaterialTestOrder) (*entity.MaterialTestOrder, error)
	ApprovePayment(context.Context, *dto.ApprovePaymentMaterialTestOrder) (*entity.MaterialTestOrder, error)
	RejectPayment(context.Context, *dto.RejectPaymentMaterialTestOrder) (*entity.MaterialTestOrder, error)
	Test(context.Context, *dto.TestMaterialTestOrder) (*entity.MaterialTestOrder, error)
	Refund(context.Context, *dto.RefundMaterialTestOrder) (*entity.MaterialTestOrder, error)
	Complete(context.Context, *dto.CompleteMaterialTestOrder) (*entity.MaterialTestOrder, error)

	// Service evaluation related
	UpdateOrderServiceEvaluation(context.Context, *dto.UpdateMaterialTestOrderServiceEvaluation) (*entity.MaterialTestOrder, error)

	// Media relateds
	StoreMedia(context.Context, *dto.SaveMaterialTestOrderMedia) (*entity.MaterialTestOrder, error)
	UpdateMedia(context.Context, *dto.SaveMaterialTestOrderMedia) (*entity.MaterialTestOrder, error)
	DestroyMedia(context.Context, *dto.DestroyMaterialTestOrderMedia) (*entity.MaterialTestOrder, error)
}

type materialTestOrderUsecase struct {
	// Services
	activityService activityService.ActivityService
	mediaService    mediaService.MediaService

	// Steps
	saveMtoStep                    mtoStep.SaveMaterialTestOrderStep
	approveMtoStep                 mtoStep.ApproveMaterialTestOrderStep
	rejectMtoStep                  mtoStep.RejectMaterialTestOrderStep
	cancelMtoStep                  mtoStep.CancelMaterialTestOrderStep
	submitPaymentMtoStep           mtoStep.SubmitPaymentMaterialTestOrderStep
	approvePaymentMtoStep          mtoStep.ApprovePaymentMaterialTestOrderStep
	rejectPaymentMtoStep           mtoStep.RejectPaymentMaterialTestOrderStep
	testMtoStep                    mtoStep.TestMaterialTestOrderStep
	refundMtoStep                  mtoStep.RefundMaterialTestOrderStep
	completeMtoStep                mtoStep.CompleteMaterialTestOrderStep
	updateMtoServiceEvaluationStep mtoStep.UpdateMaterialTestOrderServiceEvaluationStep
	storeMtoMediaStep              mtoStep.StoreMaterialTestOrderMediaStep
	updateMtoMediaStep             mtoStep.UpdateMaterialTestOrderMediaStep
	destroyMtoMediaStep            mtoStep.DestroyMaterialTestOrderMediaStep

	// Repository
	mtoRepository mtoRepository.MaterialTestOrderRepository

	// Policy
	mtoPolicy mtoPolicy.MaterialTestOrderPolicy
}

type NewMaterialTestOrderUsecaseParams struct {
	fx.In

	// Services
	ActivityService activityService.ActivityService
	MediaService    mediaService.MediaService

	// Steps
	SaveMaterialTestOrderStep                    mtoStep.SaveMaterialTestOrderStep
	ApproveMaterialTestOrderStep                 mtoStep.ApproveMaterialTestOrderStep
	RejectMaterialTestOrderStep                  mtoStep.RejectMaterialTestOrderStep
	CancelMaterialTestOrderStep                  mtoStep.CancelMaterialTestOrderStep
	SubmitPaymentMaterialTestOrderStep           mtoStep.SubmitPaymentMaterialTestOrderStep
	ApprovePaymentMaterialTestOrderStep          mtoStep.ApprovePaymentMaterialTestOrderStep
	RejectPaymentMaterialTestOrderStep           mtoStep.RejectPaymentMaterialTestOrderStep
	TestMaterialTestOrderStep                    mtoStep.TestMaterialTestOrderStep
	RefundMaterialTestOrderStep                  mtoStep.RefundMaterialTestOrderStep
	CompleteMaterialTestOrderStep                mtoStep.CompleteMaterialTestOrderStep
	UpdateMaterialTestOrderServiceEvaluationStep mtoStep.UpdateMaterialTestOrderServiceEvaluationStep
	StoreMaterialTestOrderMediaStep              mtoStep.StoreMaterialTestOrderMediaStep
	UpdateMaterialTestOrderMediaStep             mtoStep.UpdateMaterialTestOrderMediaStep
	DestroyMaterialTestOrderMediaStep            mtoStep.DestroyMaterialTestOrderMediaStep

	// Repository
	MaterialTestOrderRepository mtoRepository.MaterialTestOrderRepository

	// Policy
	MaterialTestOrderPolicy mtoPolicy.MaterialTestOrderPolicy
}

func NewMaterialTestOrderUsecase(params NewMaterialTestOrderUsecaseParams) MaterialTestOrderUsecase {
	return &materialTestOrderUsecase{
		// Services
		activityService: params.ActivityService,
		mediaService:    params.MediaService,

		// Steps
		saveMtoStep:                    params.SaveMaterialTestOrderStep,
		approveMtoStep:                 params.ApproveMaterialTestOrderStep,
		rejectMtoStep:                  params.RejectMaterialTestOrderStep,
		cancelMtoStep:                  params.CancelMaterialTestOrderStep,
		submitPaymentMtoStep:           params.SubmitPaymentMaterialTestOrderStep,
		approvePaymentMtoStep:          params.ApprovePaymentMaterialTestOrderStep,
		rejectPaymentMtoStep:           params.RejectPaymentMaterialTestOrderStep,
		testMtoStep:                    params.TestMaterialTestOrderStep,
		refundMtoStep:                  params.RefundMaterialTestOrderStep,
		completeMtoStep:                params.CompleteMaterialTestOrderStep,
		updateMtoServiceEvaluationStep: params.UpdateMaterialTestOrderServiceEvaluationStep,
		storeMtoMediaStep:              params.StoreMaterialTestOrderMediaStep,
		updateMtoMediaStep:             params.UpdateMaterialTestOrderMediaStep,
		destroyMtoMediaStep:            params.DestroyMaterialTestOrderMediaStep,

		// Repository
		mtoRepository: params.MaterialTestOrderRepository,

		// Policy
		mtoPolicy: params.MaterialTestOrderPolicy,
	}
}

func (u *materialTestOrderUsecase) Paginate(ctx context.Context, q *query.Query) (op *pagination.OffsetPagination[*entity.MaterialTestOrder], err error) {
	op, err = u.mtoRepository.Paginate(q)
	if err != nil {
		return nil, err
	}

	if _, err = u.activityService.Record(ctx, &activityDto.CreateActivity{
		Causer:      contextkey.GetUser(ctx),
		Action:      activityEnum.MaterialTestOrdersIndex,
		SubjectType: typeutil.Ptr(activityEnum.MaterialTestOrderSubjectType),
	}); err != nil {
		return nil, err
	}

	return
}

func (u *materialTestOrderUsecase) Show(ctx context.Context, q *query.Query) (mto *entity.MaterialTestOrder, err error) {
	mto, err = u.mtoRepository.First(q)
	if err != nil {
		return nil, err
	}

	if _, err = u.activityService.Record(ctx, &activityDto.CreateActivity{
		Causer:  contextkey.GetUser(ctx),
		Action:  activityEnum.MaterialTestOrdersShow,
		Subject: mto,
	}); err != nil {
		return nil, err
	}

	return
}

func (u *materialTestOrderUsecase) Store(ctx context.Context, _dto *dto.SaveMaterialTestOrder) (mto *entity.MaterialTestOrder, err error) {
	if err = u.mtoPolicy.Store(ctx, _dto); err != nil {
		return nil, err
	}

	if mto, err = u.saveMtoStep.Handle(ctx, _dto); err != nil {
		return nil, err
	}

	if _, err = u.activityService.Record(ctx, &activityDto.CreateActivity{
		Causer:  contextkey.GetUser(ctx),
		Action:  activityEnum.MaterialTestOrdersStore,
		Subject: mto,
	}); err != nil {
		return nil, err
	}

	return
}

func (u *materialTestOrderUsecase) Update(ctx context.Context, _dto *dto.SaveMaterialTestOrder) (mto *entity.MaterialTestOrder, err error) {
	if err = u.mtoPolicy.Update(ctx, _dto); err != nil {
		return nil, err
	}

	if mto, err = u.saveMtoStep.Handle(ctx, _dto); err != nil {
		return nil, err
	}

	if _, err = u.activityService.Record(ctx, &activityDto.CreateActivity{
		Causer:  contextkey.GetUser(ctx),
		Action:  activityEnum.MaterialTestOrdersUpdate,
		Subject: mto,
	}); err != nil {
		return nil, err
	}

	return
}

/* ==============================
		Status relateds
============================== */

func (u *materialTestOrderUsecase) Approve(ctx context.Context, _dto *dto.ApproveMaterialTestOrder) (mto *entity.MaterialTestOrder, err error) {
	if err = u.mtoPolicy.Approve(ctx, _dto); err != nil {
		return nil, err
	}

	if mto, err = u.approveMtoStep.Handle(ctx, _dto); err != nil {
		return nil, err
	}

	if _, err = u.activityService.Record(ctx, &activityDto.CreateActivity{
		Causer:  contextkey.GetUser(ctx),
		Action:  activityEnum.MaterialTestOrdersUpdate,
		Subject: mto,
	}); err != nil {
		return nil, err
	}

	return
}

func (u *materialTestOrderUsecase) Reject(ctx context.Context, _dto *dto.RejectMaterialTestOrder) (mto *entity.MaterialTestOrder, err error) {
	if err = u.mtoPolicy.Reject(ctx, _dto); err != nil {
		return nil, err
	}

	if mto, err = u.rejectMtoStep.Handle(ctx, _dto); err != nil {
		return nil, err
	}

	if _, err = u.activityService.Record(ctx, &activityDto.CreateActivity{
		Causer:  contextkey.GetUser(ctx),
		Action:  activityEnum.MaterialTestOrdersUpdate,
		Subject: mto,
	}); err != nil {
		return nil, err
	}

	return
}

func (u *materialTestOrderUsecase) Cancel(ctx context.Context, _dto *dto.CancelMaterialTestOrder) (mto *entity.MaterialTestOrder, err error) {
	if err = u.mtoPolicy.Cancel(ctx, _dto); err != nil {
		return nil, err
	}

	if mto, err = u.cancelMtoStep.Handle(ctx, _dto); err != nil {
		return nil, err
	}

	if _, err = u.activityService.Record(ctx, &activityDto.CreateActivity{
		Causer:  contextkey.GetUser(ctx),
		Action:  activityEnum.MaterialTestOrdersUpdate,
		Subject: mto,
	}); err != nil {
		return nil, err
	}

	return
}

func (u *materialTestOrderUsecase) SubmitPayment(ctx context.Context, _dto *dto.SubmitPaymentMaterialTestOrder) (mto *entity.MaterialTestOrder, err error) {
	if err = u.mtoPolicy.SubmitPayment(ctx, _dto); err != nil {
		return nil, err
	}

	if mto, err = u.submitPaymentMtoStep.Handle(ctx, _dto); err != nil {
		return nil, err
	}

	if _, err = u.activityService.Record(ctx, &activityDto.CreateActivity{
		Causer:  contextkey.GetUser(ctx),
		Action:  activityEnum.MaterialTestOrdersUpdate,
		Subject: mto,
	}); err != nil {
		return nil, err
	}

	return
}

func (u *materialTestOrderUsecase) ApprovePayment(ctx context.Context, _dto *dto.ApprovePaymentMaterialTestOrder) (mto *entity.MaterialTestOrder, err error) {
	if err = u.mtoPolicy.ApprovePayment(ctx, _dto); err != nil {
		return nil, err
	}

	if mto, err = u.approvePaymentMtoStep.Handle(ctx, _dto); err != nil {
		return nil, err
	}

	if _, err = u.activityService.Record(ctx, &activityDto.CreateActivity{
		Causer:  contextkey.GetUser(ctx),
		Action:  activityEnum.MaterialTestOrdersUpdate,
		Subject: mto,
	}); err != nil {
		return nil, err
	}

	return
}

func (u *materialTestOrderUsecase) RejectPayment(ctx context.Context, _dto *dto.RejectPaymentMaterialTestOrder) (mto *entity.MaterialTestOrder, err error) {
	if err = u.mtoPolicy.RejectPayment(ctx, _dto); err != nil {
		return nil, err
	}

	if mto, err = u.rejectPaymentMtoStep.Handle(ctx, _dto); err != nil {
		return nil, err
	}

	if _, err = u.activityService.Record(ctx, &activityDto.CreateActivity{
		Causer:  contextkey.GetUser(ctx),
		Action:  activityEnum.MaterialTestOrdersUpdate,
		Subject: mto,
	}); err != nil {
		return nil, err
	}

	return
}

func (u *materialTestOrderUsecase) Test(ctx context.Context, _dto *dto.TestMaterialTestOrder) (mto *entity.MaterialTestOrder, err error) {
	if err = u.mtoPolicy.Test(ctx, _dto); err != nil {
		return nil, err
	}

	if mto, err = u.testMtoStep.Handle(ctx, _dto); err != nil {
		return nil, err
	}

	if _, err = u.activityService.Record(ctx, &activityDto.CreateActivity{
		Causer:  contextkey.GetUser(ctx),
		Action:  activityEnum.MaterialTestOrdersUpdate,
		Subject: mto,
	}); err != nil {
		return nil, err
	}

	return
}

func (u *materialTestOrderUsecase) Refund(ctx context.Context, _dto *dto.RefundMaterialTestOrder) (mto *entity.MaterialTestOrder, err error) {
	if err = u.mtoPolicy.Refund(ctx, _dto); err != nil {
		return nil, err
	}

	if mto, err = u.refundMtoStep.Handle(ctx, _dto); err != nil {
		return nil, err
	}

	if _, err = u.activityService.Record(ctx, &activityDto.CreateActivity{
		Causer:  contextkey.GetUser(ctx),
		Action:  activityEnum.MaterialTestOrdersUpdate,
		Subject: mto,
	}); err != nil {
		return nil, err
	}

	return
}

func (u *materialTestOrderUsecase) Complete(ctx context.Context, _dto *dto.CompleteMaterialTestOrder) (mto *entity.MaterialTestOrder, err error) {
	if err = u.mtoPolicy.Complete(ctx, _dto); err != nil {
		return nil, err
	}

	if mto, err = u.completeMtoStep.Handle(ctx, _dto); err != nil {
		return nil, err
	}

	if _, err = u.activityService.Record(ctx, &activityDto.CreateActivity{
		Causer:  contextkey.GetUser(ctx),
		Action:  activityEnum.MaterialTestOrdersUpdate,
		Subject: mto,
	}); err != nil {
		return nil, err
	}

	return
}

/* ==============================
		Service evaluation related
============================== */

func (u *materialTestOrderUsecase) UpdateOrderServiceEvaluation(ctx context.Context, _dto *dto.UpdateMaterialTestOrderServiceEvaluation) (*entity.MaterialTestOrder, error) {
	if err := u.mtoPolicy.UpdateOrderServiceEvaluation(ctx, _dto); err != nil {
		return nil, err
	}

	mto, err := u.updateMtoServiceEvaluationStep.Handle(ctx, _dto)
	if err != nil {
		return nil, err
	}

	if _, err = u.activityService.Record(ctx, &activityDto.CreateActivity{
		Causer:  contextkey.GetUser(ctx),
		Action:  activityEnum.MaterialTestOrdersUpdate,
		Subject: mto,
	}); err != nil {
		return nil, err
	}

	return mto, nil
}

/* ==============================
		Media relateds
============================== */

func (u *materialTestOrderUsecase) StoreMedia(ctx context.Context, _dto *dto.SaveMaterialTestOrderMedia) (mto *entity.MaterialTestOrder, err error) {
	if err = u.mtoPolicy.StoreMedia(ctx, _dto); err != nil {
		return nil, err
	}

	mto, err = u.storeMtoMediaStep.Handle(ctx, _dto)
	if err != nil {
		return nil, err
	}

	if _, err = u.activityService.Record(ctx, &activityDto.CreateActivity{
		Causer:  contextkey.GetUser(ctx),
		Action:  activityEnum.MaterialTestOrdersUpdate,
		Subject: mto,
	}); err != nil {
		return nil, err
	}

	return
}

func (u *materialTestOrderUsecase) UpdateMedia(ctx context.Context, _dto *dto.SaveMaterialTestOrderMedia) (mto *entity.MaterialTestOrder, err error) {
	if err = u.mtoPolicy.UpdateMedia(ctx, _dto); err != nil {
		return nil, err
	}

	mto, err = u.updateMtoMediaStep.Handle(ctx, _dto)
	if err != nil {
		return nil, err
	}

	if _, err = u.activityService.Record(ctx, &activityDto.CreateActivity{
		Causer:  contextkey.GetUser(ctx),
		Action:  activityEnum.MaterialTestOrdersUpdate,
		Subject: mto,
	}); err != nil {
		return nil, err
	}

	return
}

func (u *materialTestOrderUsecase) DestroyMedia(ctx context.Context, _dto *dto.DestroyMaterialTestOrderMedia) (mto *entity.MaterialTestOrder, err error) {
	if err = u.mtoPolicy.DestroyMedia(ctx, _dto); err != nil {
		return nil, err
	}

	mto, err = u.destroyMtoMediaStep.Handle(ctx, _dto)
	if err != nil {
		return nil, err
	}

	if _, err = u.activityService.Record(ctx, &activityDto.CreateActivity{
		Causer:  contextkey.GetUser(ctx),
		Action:  activityEnum.MaterialTestOrdersUpdate,
		Subject: mto,
	}); err != nil {
		return nil, err
	}

	return mto, nil
}
