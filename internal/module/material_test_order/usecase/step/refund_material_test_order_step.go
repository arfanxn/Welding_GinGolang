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

type RefundMaterialTestOrderStep interface {
	Handle(ctx context.Context, dto *dto.RefundMaterialTestOrder) (*entity.MaterialTestOrder, error)
}

type refundPaymentMaterialTestOrderStep struct {
	mtoRepository mtoRepository.MaterialTestOrderRepository
	mediaService  mediaService.MediaService
}

type NewRefundMaterialTestOrderStepParams struct {
	fx.In

	MaterialTestOrderRepository mtoRepository.MaterialTestOrderRepository
	MediaService                mediaService.MediaService
}

func NewRefundMaterialTestOrderStep(params NewRefundMaterialTestOrderStepParams) RefundMaterialTestOrderStep {
	return &refundPaymentMaterialTestOrderStep{
		mtoRepository: params.MaterialTestOrderRepository,
		mediaService:  params.MediaService,
	}
}

func (s *refundPaymentMaterialTestOrderStep) Handle(ctx context.Context, _dto *dto.RefundMaterialTestOrder) (
	mto *entity.MaterialTestOrder, err error,
) {
	q := query.NewQuery().Include("medias")

	mto, err = s.mtoRepository.Find(_dto.Id, q)
	if err != nil {
		return
	}

	mto.Status = mtoEnum.MaterialTestOrderStatusRefunded
	mto.RefundedAt = typeutil.Ptr(time.Now())
	mto.CompletedAt = nil

	s.mediaService.DestroyByQuery(
		ctx,
		query.NewQuery().
			Filter("model_type", query.OperatorEqual, mediaEnum.ModelTypeMaterialTestOrder).
			Filter("model_id", query.OperatorEqual, _dto.Id).
			Filter("collection_name", query.OperatorEqual, mediaEnum.CollectionNameMaterialTestOrderRefundProof),
	)

	s.mediaService.CreateFromMultipartFiles(ctx, mediaDto.CreateFromMultipartFiles{
		{
			ModelType:      mediaEnum.ModelTypeMaterialTestOrder,
			ModelId:        _dto.Id,
			CollectionName: mediaEnum.CollectionNameMaterialTestOrderRefundProof,
			Name:           "Bukti pengembalian (refund)",
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
