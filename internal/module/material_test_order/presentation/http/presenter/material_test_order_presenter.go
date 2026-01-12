package presenter

import (
	"context"
	"fmt"
	"net/url"

	"github.com/arfanxn/welding/internal/infrastructure/config"
	"github.com/arfanxn/welding/internal/module/material_test_order/presentation/http/viewmodel"
	mediaViewmodel "github.com/arfanxn/welding/internal/module/media/presentation/http/viewmodel"
	"github.com/arfanxn/welding/internal/module/shared/contextkey"
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	"github.com/arfanxn/welding/pkg/pagination"
	"go.uber.org/fx"
)

// MaterialTestOrderPresenter defines the interface for converting MaterialTestOrder domain entities to view models.
// It handles the transformation of material test order data between the domain layer and the presentation layer.
type MaterialTestOrderPresenter interface {
	// FromEntityToViewModel converts a single material test order entity to its corresponding view model.
	// It takes a context and a material test order entity, and returns the view model representation
	// along with any error that occurred during conversion.
	FromEntityToViewModel(ctx context.Context, mto *entity.MaterialTestOrder) (*viewmodel.MaterialTestOrderViewModel, error)

	// FromEntitiesToViewModels converts a slice of material test order entities to their corresponding view models.
	// It processes multiple material test orders in bulk and returns a slice of view models.
	// If any error occurs during conversion, it returns the error immediately.
	FromEntitiesToViewModels(ctx context.Context, mtos []*entity.MaterialTestOrder) ([]*viewmodel.MaterialTestOrderViewModel, error)

	// FromEntityOffsetPaginationToViewModelOffsetPagination converts a paginated result of material test order entities
	// to a paginated result of view models using offset-based pagination.
	FromEntityOffsetPaginationToViewModelOffsetPagination(
		ctx context.Context,
		op *pagination.OffsetPagination[*entity.MaterialTestOrder],
	) (*pagination.OffsetPagination[*viewmodel.MaterialTestOrderViewModel], error)

	// FromEntityOffsetPaginationToViewModelPagePagination converts a paginated result of material test order entities
	// to a page-based pagination result of view models.
	// This is useful for web interfaces that display material test orders with page numbers.
	FromEntityOffsetPaginationToViewModelPagePagination(
		ctx context.Context,
		op *pagination.OffsetPagination[*entity.MaterialTestOrder],
	) (*pagination.PagePagination[*viewmodel.MaterialTestOrderViewModel], error)
}

type materialTestOrderPresenter struct {
	config *config.Config
}

type NewMaterialTestOrderPresenterParams struct {
	fx.In

	Config *config.Config
}

func NewMaterialTestOrderPresenter(params NewMaterialTestOrderPresenterParams) MaterialTestOrderPresenter {
	return &materialTestOrderPresenter{
		config: params.Config,
	}
}

func (p *materialTestOrderPresenter) FromEntityToViewModel(ctx context.Context, mto *entity.MaterialTestOrder) (*viewmodel.MaterialTestOrderViewModel, error) {
	mtoVM := &viewmodel.MaterialTestOrderViewModel{}

	// -- IDENTIFICATION --
	mtoVM.Id = mto.Id
	mtoVM.Number = mto.Number
	mtoVM.WorkCategoryId = mto.WorkCategoryId

	// -- WORK PACKAGE & APPLICANT DETAILS --
	mtoVM.WorkPackageName = mto.WorkPackageName
	mtoVM.ApplicantName = mto.ApplicantName
	mtoVM.ApplicantPhoneNumber = mto.ApplicantPhoneNumber
	mtoVM.ApplicantEmail = mto.ApplicantEmail
	mtoVM.ApplicantFullAddress = mto.ApplicantFullAddress
	mtoVM.ApplicantNote = mto.ApplicantNote
	mtoVM.RecipientName = mto.RecipientName
	mtoVM.TesterNote = mto.TesterNote

	// -- FINANCIAL METADATA --
	mtoVM.SubTotal = mto.SubTotal
	mtoVM.Tax = mto.Tax
	mtoVM.Discount = mto.Discount
	mtoVM.Total = mto.Total

	// -- WORKFLOW STATUS --
	mtoVM.Status = mto.Status

	// -- TIMESTAMP MILESTONES --
	// Review lifecycle
	mtoVM.SubmittedAt = mto.SubmittedAt
	mtoVM.ApprovedAt = mto.ApprovedAt
	mtoVM.RejectedAt = mto.RejectedAt
	mtoVM.CancelledAt = mto.CancelledAt

	// Payment lifecycle
	mtoVM.PaymentSubmittedAt = mto.PaymentSubmittedAt
	mtoVM.PaymentApprovedAt = mto.PaymentApprovedAt
	mtoVM.PaymentRejectedAt = mto.PaymentRejectedAt

	// Testing lifecycle
	mtoVM.TestingAt = mto.TestingAt
	mtoVM.RefundedAt = mto.RefundedAt
	mtoVM.CompletedAt = mto.CompletedAt

	// -- SYSTEM METADATA --
	mtoVM.CreatedAt = mto.CreatedAt
	mtoVM.UpdatedAt = mto.UpdatedAt

	// -- RELATIONS --
	if mto.WorkCategory.Id != "" {
		mtoVM.WorkCategory = mto.WorkCategory
	}

	if len(mto.OrderUsers) > 0 {
		mtoVM.OrderUsers = mto.OrderUsers
	}

	if len(mto.OrderedServices) > 0 {
		mtoVM.OrderedServices = mto.OrderedServices
	}

	if len(mto.Medias) > 0 {
		mtoVM.Medias = []*mediaViewmodel.Media{}
		for _, media := range mto.Medias {
			mtoVM.Medias = append(mtoVM.Medias, &mediaViewmodel.Media{
				Id:             media.Id,
				ModelType:      media.ModelType,
				ModelId:        media.ModelId,
				Ulid:           media.Ulid,
				CollectionName: media.CollectionName,
				Name:           media.Name,
				FileName:       media.FileName,
				// TODO: implement abstarct filesystem
				FileUrl:              fmt.Sprintf("%s:%s/storage/medias/%s/%s", p.config.AppHost, p.config.AppPort, media.Id, media.FileName),
				MimeType:             media.MimeType,
				Disk:                 media.Disk,
				ConversionsDisk:      media.ConversionsDisk,
				Size:                 media.Size,
				Manipulations:        media.Manipulations,
				CustomProperties:     media.CustomProperties,
				GeneratedConversions: media.GeneratedConversions,
				ResponsiveImages:     media.ResponsiveImages,
				OrderColumn:          media.OrderColumn,
				CreatedAt:            media.CreatedAt,
				UpdatedAt:            media.UpdatedAt,
			})
		}
	}

	return mtoVM, nil
}

func (p *materialTestOrderPresenter) FromEntitiesToViewModels(ctx context.Context, mtos []*entity.MaterialTestOrder) ([]*viewmodel.MaterialTestOrderViewModel, error) {
	var viewModels []*viewmodel.MaterialTestOrderViewModel
	for _, mto := range mtos {
		viewModel, err := p.FromEntityToViewModel(ctx, mto)
		if err != nil {
			return nil, err
		}
		viewModels = append(viewModels, viewModel)
	}

	return viewModels, nil
}

func (p *materialTestOrderPresenter) FromEntityOffsetPaginationToViewModelOffsetPagination(
	ctx context.Context,
	op *pagination.OffsetPagination[*entity.MaterialTestOrder],
) (*pagination.OffsetPagination[*viewmodel.MaterialTestOrderViewModel], error) {
	viewModels, err := p.FromEntitiesToViewModels(ctx, op.Items)
	if err != nil {
		return nil, err
	}

	return pagination.NewOffsetPagination(op.Offset, op.Limit, op.TotalItems, viewModels), nil
}

func (p *materialTestOrderPresenter) FromEntityOffsetPaginationToViewModelPagePagination(
	ctx context.Context,
	op *pagination.OffsetPagination[*entity.MaterialTestOrder],
) (*pagination.PagePagination[*viewmodel.MaterialTestOrderViewModel], error) {
	convertedOP, err := p.FromEntityOffsetPaginationToViewModelOffsetPagination(ctx, op)
	if err != nil {
		return nil, err
	}

	url, _ := ctx.Value(contextkey.RequestURLKey).(url.URL)

	return pagination.FromOPToPP(convertedOP, url), nil
}
