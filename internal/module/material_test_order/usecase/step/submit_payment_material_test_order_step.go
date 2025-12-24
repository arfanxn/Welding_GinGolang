package step

import (
	"context"
	"time"

	mtoEnum "github.com/arfanxn/welding/internal/module/material_test_order/domain/enum"
	mtoRepository "github.com/arfanxn/welding/internal/module/material_test_order/domain/repository"
	"github.com/arfanxn/welding/internal/module/material_test_order/usecase/dto"
	mediaEnum "github.com/arfanxn/welding/internal/module/media/domain/enum"
	mediaDto "github.com/arfanxn/welding/internal/module/media/usecase/dto"
	mediaService "github.com/arfanxn/welding/internal/module/media/usecase/service"
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	"github.com/arfanxn/welding/pkg/query"
	"github.com/arfanxn/welding/pkg/typeutil"
	"go.uber.org/fx"
)

type SubmitPaymentMaterialTestOrderStep interface {
	Handle(ctx context.Context, dto *dto.SubmitPaymentMaterialTestOrder) (*entity.MaterialTestOrder, error)
}

type submitPaymentMaterialTestOrderStep struct {
	mtoRepository mtoRepository.MaterialTestOrderRepository
	mediaService  mediaService.MediaService
}

type NewSubmitPaymentMaterialTestOrderStepParams struct {
	fx.In

	MaterialTestOrderRepository mtoRepository.MaterialTestOrderRepository
	MediaService                mediaService.MediaService
}

func NewSubmitPaymentMaterialTestOrderStep(params NewSubmitPaymentMaterialTestOrderStepParams) SubmitPaymentMaterialTestOrderStep {
	return &submitPaymentMaterialTestOrderStep{
		mtoRepository: params.MaterialTestOrderRepository,
		mediaService:  params.MediaService,
	}
}

func (s *submitPaymentMaterialTestOrderStep) Handle(ctx context.Context, _dto *dto.SubmitPaymentMaterialTestOrder) (
	mto *entity.MaterialTestOrder, err error,
) {
	q := query.NewQuery().Include("medias")

	mto, err = s.mtoRepository.Find(_dto.Id, q)
	if err != nil {
		return
	}

	mto.Status = mtoEnum.MaterialTestOrderStatusPaymentSubmitted
	mto.PaymentSubmittedAt = typeutil.Ptr(time.Now())

	s.mediaService.DestroyByQuery(
		ctx,
		query.NewQuery().
			Filter("model_type", query.OperatorEqual, mediaEnum.ModelTypeMaterialTestOrder).
			Filter("model_id", query.OperatorEqual, _dto.Id).
			Filter("collection_name", query.OperatorEqual, mediaEnum.CollectionNameMaterialTestOrderPaymentProof),
	)

	s.mediaService.CreateFromMultipartFiles(ctx, mediaDto.CreateFromMultipartFiles{
		{
			ModelType:      mediaEnum.ModelTypeMaterialTestOrder,
			ModelId:        _dto.Id,
			CollectionName: mediaEnum.CollectionNameMaterialTestOrderPaymentProof,
			Name:           "Bukti pembayaran (payment)",
			File:           _dto.File,
		},
	})

	if err = s.mtoRepository.Save(mto); err != nil {
		return
	}

	mto, err = s.mtoRepository.Find(_dto.Id, q)
	if err != nil {
		return
	}

	return
}
