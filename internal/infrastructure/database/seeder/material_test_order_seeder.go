package seeder

import (
	"fmt"
	"os"
	"time"

	"github.com/arfanxn/welding/internal/infrastructure/id"
	materialTestOrderEnum "github.com/arfanxn/welding/internal/module/material_test_order/domain/enum"
	materialTestOrderRepository "github.com/arfanxn/welding/internal/module/material_test_order/domain/repository"
	materialTestOrderServiceRepository "github.com/arfanxn/welding/internal/module/material_test_order_service/domain/repository"
	materialTestOrderServiceEvaluationRepository "github.com/arfanxn/welding/internal/module/material_test_order_service_evaluation/domain/repository"
	materialTestOrderUserEnum "github.com/arfanxn/welding/internal/module/material_test_order_user/domain/enum"
	materialTestOrderUserRepository "github.com/arfanxn/welding/internal/module/material_test_order_user/domain/repository"
	materialTestServiceRepository "github.com/arfanxn/welding/internal/module/material_test_service/domain/repository"
	materialTestWorkCategoryRepository "github.com/arfanxn/welding/internal/module/material_test_work_category/domain/repository"
	mediaEnum "github.com/arfanxn/welding/internal/module/media/domain/enum"
	mediaRepository "github.com/arfanxn/welding/internal/module/media/domain/repository"
	roleEnum "github.com/arfanxn/welding/internal/module/role/domain/enum"
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	userRepository "github.com/arfanxn/welding/internal/module/user/domain/repository"
	"github.com/arfanxn/welding/pkg/query"
	"github.com/arfanxn/welding/pkg/sliceutil"
	"github.com/bluele/factory-go/factory"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/samber/lo"
	"go.uber.org/fx"
)

var _ Seeder = (*MaterialTestOrderSeeder)(nil)

type MaterialTestOrderSeeder struct {
	idService                                    id.IdService
	materialTestOrderFactory                     *factory.Factory
	materialTestOrderRepository                  materialTestOrderRepository.MaterialTestOrderRepository
	materialTestOrderUserRepository              materialTestOrderUserRepository.MaterialTestOrderUserRepository
	materialTestServiceRepository                materialTestServiceRepository.MaterialTestServiceRepository
	materialTestWorkCategoryRepository           materialTestWorkCategoryRepository.MaterialTestWorkCategoryRepository
	materialTestOrderServiceFactory              *factory.Factory
	materialTestOrderServiceRepository           materialTestOrderServiceRepository.MaterialTestOrderServiceRepository
	materialTestOrderServiceEvaluationFactory    *factory.Factory
	materialTestOrderServiceEvaluationRepository materialTestOrderServiceEvaluationRepository.MaterialTestOrderServiceEvaluationRepository
	mediaFactory                                 *factory.Factory
	mediaRepository                              mediaRepository.MediaRepository
	userRepository                               userRepository.UserRepository
}

type NewMaterialTestOrderSeederParams struct {
	fx.In

	IdService                                    id.IdService
	MaterialTestOrderFactory                     *factory.Factory `name:"material_test_order_factory"`
	MaterialTestOrderRepository                  materialTestOrderRepository.MaterialTestOrderRepository
	MaterialTestOrderUserRepository              materialTestOrderUserRepository.MaterialTestOrderUserRepository
	MaterialTestServiceRepository                materialTestServiceRepository.MaterialTestServiceRepository
	MaterialTestWorkCategoryRepository           materialTestWorkCategoryRepository.MaterialTestWorkCategoryRepository
	MaterialTestOrderServiceFactory              *factory.Factory `name:"material_test_order_service_factory"`
	MaterialTestOrderServiceRepository           materialTestOrderServiceRepository.MaterialTestOrderServiceRepository
	MaterialTestOrderServiceEvaluationFactory    *factory.Factory `name:"material_test_order_service_evaluation_factory"`
	MaterialTestOrderServiceEvaluationRepository materialTestOrderServiceEvaluationRepository.MaterialTestOrderServiceEvaluationRepository
	MediaFactory                                 *factory.Factory `name:"media_factory"`
	MediaRepository                              mediaRepository.MediaRepository
	UserRepository                               userRepository.UserRepository
}

func NewMaterialTestOrderSeeder(params NewMaterialTestOrderSeederParams) Seeder {
	return &MaterialTestOrderSeeder{
		idService:                                    params.IdService,
		materialTestOrderFactory:                     params.MaterialTestOrderFactory,
		materialTestOrderRepository:                  params.MaterialTestOrderRepository,
		materialTestOrderUserRepository:              params.MaterialTestOrderUserRepository,
		materialTestServiceRepository:                params.MaterialTestServiceRepository,
		materialTestWorkCategoryRepository:           params.MaterialTestWorkCategoryRepository,
		materialTestOrderServiceFactory:              params.MaterialTestOrderServiceFactory,
		materialTestOrderServiceRepository:           params.MaterialTestOrderServiceRepository,
		materialTestOrderServiceEvaluationFactory:    params.MaterialTestOrderServiceEvaluationFactory,
		materialTestOrderServiceEvaluationRepository: params.MaterialTestOrderServiceEvaluationRepository,
		mediaFactory:                                 params.MediaFactory,
		mediaRepository:                              params.MediaRepository,
		userRepository:                               params.UserRepository,
	}
}

func (s *MaterialTestOrderSeeder) Seed() error {
	materialTestOrderRepository := s.materialTestOrderRepository
	materialTestOrderUserRepository := s.materialTestOrderUserRepository
	materialTestServiceRepository := s.materialTestServiceRepository
	materialTestWorkCategoryRepository := s.materialTestWorkCategoryRepository
	materialTestOrderServiceFactory := s.materialTestOrderServiceFactory
	materialTestOrderServiceRepository := s.materialTestOrderServiceRepository
	materialTestOrderServiceEvaluationFactory := s.materialTestOrderServiceEvaluationFactory
	materialTestOrderServiceEvaluationRepository := s.materialTestOrderServiceEvaluationRepository
	mediaFactory := s.mediaFactory
	mediaRepository := s.mediaRepository
	userRepository := s.userRepository

	medias := []*entity.Media{}
	mtOrderUsers := []*entity.MaterialTestOrderUser{}
	mtOrderServices := []*entity.MaterialTestOrderService{}
	mtOrderServiceEvaluations := []*entity.MaterialTestOrderServiceEvaluation{}
	mtOrders := []*entity.MaterialTestOrder{
		// Draft
		s.createDraft(nil),

		// Awaiting Review
		s.createAwaitingReview(nil),

		// Awaiting Payment
		s.createAwaitingPayment(nil),

		// Payment Submitted
		s.createPaymentSubmitted(nil),

		// Payment Rejected
		s.createPaymentRejected(nil),

		// Payment Approved
		s.createPaymentApproved(nil),

		// Testing
		s.createTesting(nil),

		// Completed
		s.createCompleted(nil),

		// Cancelled
		s.createCancelled(nil),

		// Rejected
		s.createRejected(nil),

		// Refunded
		s.createRefunded(nil),
	}

	customerUsers, err := userRepository.Get(query.NewQuery().Include("roles").Filter("roles.name", query.OperatorEqual, roleEnum.Customer))
	if err != nil {
		return err
	}

	employeeUsers, err := userRepository.Get(query.NewQuery().Include("roles").Filter("roles.name", query.OperatorNotEqual, roleEnum.Customer))
	if err != nil {
		return err
	}

	users := []*entity.User{}
	users = append(users, customerUsers...)
	users = append(users, employeeUsers...)

	mtServices, err := materialTestServiceRepository.Get(nil)
	if err != nil {
		return err
	}

	mtWorkCategories, err := materialTestWorkCategoryRepository.Get(nil)
	if err != nil {
		return err
	}

	for mtOrderIndex, mtOrder := range mtOrders {
		mtOrderServicesCount := gofakeit.IntRange(0, 5)

		if mtOrderServicesCount > 0 {
			shuffledMTServices := sliceutil.Shuffle(mtServices)

			// Loop through each service count to create material test order services
			for i := range mtOrderServicesCount {
				// Get the material test service for this iteration
				mtService := shuffledMTServices[i]

				// Create a new material test order service using the factory
				mtOrderService := materialTestOrderServiceFactory.MustCreate().(*entity.MaterialTestOrderService)

				// Set up the order service details
				mtOrderService.OrderId = mtOrder.Id                                                // Link to the parent order
				mtOrderService.ServiceId = mtService.Id                                            // Reference the service
				mtOrderService.Price = mtService.Price                                             // Set the service price
				mtOrderService.Quantity = gofakeit.IntRange(1, 100)                                // Random quantity between 1-100
				mtOrderService.LineTotal = mtOrderService.Price * float64(mtOrderService.Quantity) // Calculate line total

				// Add the order service to our collection
				mtOrderServices = append(mtOrderServices, mtOrderService)

				// Update the order's subtotal with this service's line total
				mtOrder.SubTotal += mtOrderService.LineTotal

				// Create an evaluation record for this order service
				mtOrderServiceEvaluation := materialTestOrderServiceEvaluationFactory.MustCreate().(*entity.MaterialTestOrderServiceEvaluation)
				mtOrderServiceEvaluation.OrderServiceId = mtOrderService.Id // Link to the order service

				// Add the evaluation to our collection
				mtOrderServiceEvaluations = append(mtOrderServiceEvaluations, mtOrderServiceEvaluation)
			}
		}

		// Append two MaterialTestOrderUser entries for each material test order:
		// 1. A creator (can be either employee or customer user)
		// 2. An owner (must be a customer user)
		mtOrderUsers = append(mtOrderUsers,
			&entity.MaterialTestOrderUser{
				OrderId:   mtOrder.Id,
				UserId:    users[gofakeit.IntRange(0, len(users)-1)].Id, // Can be any user (employee or customer)
				Type:      materialTestOrderUserEnum.TypeCreator,
				CreatedAt: mtOrder.CreatedAt,
			},
			&entity.MaterialTestOrderUser{
				OrderId:   mtOrder.Id,
				UserId:    customerUsers[gofakeit.IntRange(0, len(customerUsers)-1)].Id, // Must be a customer
				Type:      materialTestOrderUserEnum.TypeOwner,
				CreatedAt: mtOrder.CreatedAt,
			})

		mtWorkCategory := mtWorkCategories[gofakeit.IntRange(0, len(mtWorkCategories)-1)]
		mtOrder.WorkCategoryId = mtWorkCategory.Id

		mtOrderHasMedias := gofakeit.Bool()
		if mtOrderHasMedias {
			mtOrderMediasCount := gofakeit.IntRange(1, 5)

			for i := range mtOrderMediasCount {
				orderColumn := i + 1

				media := mediaFactory.MustCreate().(*entity.Media)
				media.ModelId = mtOrder.Id
				media.ModelType = mediaEnum.ModelTypeMaterialTestOrder
				media.CollectionName = mediaEnum.CollectionNameMaterialTestOrder
				media.OrderColumn = &orderColumn

				{
					// TODO: implement abstract filesystem
					_path := "./storage/medias/" + media.Id
					_filePath := _path + "/" + media.FileName

					jpegBytes := gofakeit.ImageJpeg(500, 500)
					os.MkdirAll(_path, os.ModePerm)
					if err := os.WriteFile(_filePath, jpegBytes, 0644); err != nil {
						return err
					}

					_fileInfo, err := os.Stat(_filePath)
					if err != nil {
						return err
					}

					media.Size = _fileInfo.Size()
				}

				medias = append(medias, media)
			}
		}

		mtOrder.Number = fmt.Sprintf("%03d", len(mtOrders)-mtOrderIndex)

		if mtOrder.SubTotal > 0 {
			mtOrder.Total = mtOrder.SubTotal + mtOrder.Tax - mtOrder.Discount
		} else {
			mtOrder.Tax = 0
			mtOrder.Discount = 0
			mtOrder.Total = 0
		}
	}

	if err := materialTestOrderRepository.SaveMany(mtOrders); err != nil {
		return err
	}

	mtOrderServiceChunks := lo.Chunk(mtOrderServices, 100)
	for _, mtOrderServiceChunk := range mtOrderServiceChunks {
		if err := materialTestOrderServiceRepository.SaveMany(mtOrderServiceChunk); err != nil {
			return err
		}
	}

	mtOrderServiceEvaluationChunks := lo.Chunk(mtOrderServiceEvaluations, 100)
	for _, mtOrderServiceEvaluationChunk := range mtOrderServiceEvaluationChunks {
		if err := materialTestOrderServiceEvaluationRepository.SaveMany(mtOrderServiceEvaluationChunk); err != nil {
			return err
		}
	}

	mediaChunks := lo.Chunk(medias, 100)
	for _, mediaChunk := range mediaChunks {
		if err := mediaRepository.SaveMany(mediaChunk); err != nil {
			return err
		}
	}

	mtOrderUserChunks := lo.Chunk(mtOrderUsers, 100)
	for _, mtOrderUserChunk := range mtOrderUserChunks {
		if err := materialTestOrderUserRepository.SaveMany(mtOrderUserChunk); err != nil {
			return err
		}
	}

	return nil
}

func (s *MaterialTestOrderSeeder) createDateTimeOptions() map[string]any {
	createdAt := gofakeit.DateRange(
		time.Now().AddDate(-2, 0, 0),
		time.Now().AddDate(-1, 0, 0),
	)
	rejectedAt := gofakeit.DateRange(
		createdAt,
		createdAt.AddDate(0, 0, 2),
	)
	cancelledAt := gofakeit.DateRange(
		createdAt,
		createdAt.AddDate(0, 0, 2),
	)
	paymentSubmittedAt := gofakeit.DateRange(
		createdAt.AddDate(0, 0, 7),
		createdAt.AddDate(0, 0, 14),
	)
	paymentRejectedAt := gofakeit.DateRange(
		paymentSubmittedAt,
		paymentSubmittedAt.AddDate(0, 0, 2),
	)
	paymentApprovedAt := gofakeit.DateRange(
		paymentRejectedAt,
		paymentRejectedAt.AddDate(0, 0, 2),
	)
	testingAt := gofakeit.DateRange(
		createdAt,
		createdAt.AddDate(0, 0, 2),
	)
	completedAt := gofakeit.DateRange(
		testingAt.AddDate(0, 0, 2),
		testingAt.AddDate(0, 0, 14),
	)
	refundedAt := gofakeit.DateRange(
		testingAt.AddDate(0, 0, 2),
		testingAt.AddDate(0, 0, 14),
	)

	updatedAt := gofakeit.DateRange(
		createdAt,
		time.Now(),
	)

	return map[string]any{
		"PaymentSubmittedAt": &paymentSubmittedAt,
		"PaymentRejectedAt":  &paymentRejectedAt,
		"PaymentApprovedAt":  &paymentApprovedAt,

		"TestingAt":   &testingAt,
		"CompletedAt": &completedAt,

		"CancelledAt": &cancelledAt,
		"RejectedAt":  &rejectedAt,
		"RefundedAt":  &refundedAt,

		"CreatedAt": createdAt,
		"UpdatedAt": &updatedAt,
	}
}

func (s *MaterialTestOrderSeeder) createDraft(options map[string]any) *entity.MaterialTestOrder {
	datetimeOptions := lo.PickByKeys(s.createDateTimeOptions(), []string{
		"CreatedAt",
		"UpdatedAt",
	})

	options = lo.Assign(datetimeOptions, map[string]any{
		"Status": materialTestOrderEnum.MaterialTestOrderStatusDraft,
	}, options)

	mtoMap, err := s.materialTestOrderFactory.CreateWithOption(options)
	if err != nil {
		return nil
	}
	return mtoMap.(*entity.MaterialTestOrder)
}

func (s *MaterialTestOrderSeeder) createAwaitingReview(options map[string]any) *entity.MaterialTestOrder {
	datetimeOptions := lo.PickByKeys(s.createDateTimeOptions(), []string{
		"CreatedAt",
		"UpdatedAt",
	})

	options = lo.Assign(datetimeOptions, map[string]any{
		"Status": materialTestOrderEnum.MaterialTestOrderStatusAwaitingReview,
	}, options)

	return s.materialTestOrderFactory.MustCreateWithOption(options).(*entity.MaterialTestOrder)
}

func (s *MaterialTestOrderSeeder) createAwaitingPayment(options map[string]any) *entity.MaterialTestOrder {
	datetimeOptions := lo.PickByKeys(s.createDateTimeOptions(), []string{
		"CreatedAt",
		"UpdatedAt",
	})

	options = lo.Assign(datetimeOptions, map[string]any{
		"Status": materialTestOrderEnum.MaterialTestOrderStatusAwaitingPayment,
	}, options)

	return s.materialTestOrderFactory.MustCreateWithOption(options).(*entity.MaterialTestOrder)
}

func (s *MaterialTestOrderSeeder) createPaymentSubmitted(options map[string]any) *entity.MaterialTestOrder {
	datetimeOptions := lo.PickByKeys(s.createDateTimeOptions(), []string{
		"PaymentSubmittedAt",
		"CreatedAt",
		"UpdatedAt",
	})

	options = lo.Assign(datetimeOptions, map[string]any{
		"Status": materialTestOrderEnum.MaterialTestOrderStatusPaymentSubmitted,
	}, options)

	return s.materialTestOrderFactory.MustCreateWithOption(options).(*entity.MaterialTestOrder)
}

func (s *MaterialTestOrderSeeder) createPaymentRejected(options map[string]any) *entity.MaterialTestOrder {
	datetimeOptions := lo.PickByKeys(s.createDateTimeOptions(), []string{
		"PaymentSubmittedAt",
		"PaymentRejectedAt",
		"CreatedAt",
		"UpdatedAt",
	})

	options = lo.Assign(datetimeOptions, map[string]any{
		"Status": materialTestOrderEnum.MaterialTestOrderStatusPaymentRejected,
	}, options)

	return s.materialTestOrderFactory.MustCreateWithOption(options).(*entity.MaterialTestOrder)
}

func (s *MaterialTestOrderSeeder) createPaymentApproved(options map[string]any) *entity.MaterialTestOrder {
	datetimeOptions := lo.PickByKeys(s.createDateTimeOptions(), []string{
		"PaymentSubmittedAt",
		"PaymentApprovedAt",
		"CreatedAt",
		"UpdatedAt",
	})

	options = lo.Assign(datetimeOptions, map[string]any{
		"Status": materialTestOrderEnum.MaterialTestOrderStatusPaymentApproved,
	}, options)

	return s.materialTestOrderFactory.MustCreateWithOption(options).(*entity.MaterialTestOrder)
}

func (s *MaterialTestOrderSeeder) createTesting(options map[string]any) *entity.MaterialTestOrder {
	datetimeOptions := lo.PickByKeys(s.createDateTimeOptions(), []string{
		"PaymentSubmittedAt",
		"PaymentApprovedAt",
		"TestingAt",
		"CreatedAt",
		"UpdatedAt",
	})

	options = lo.Assign(datetimeOptions, map[string]any{
		"Status": materialTestOrderEnum.MaterialTestOrderStatusTesting,
	}, options)

	return s.materialTestOrderFactory.MustCreateWithOption(options).(*entity.MaterialTestOrder)
}

func (s *MaterialTestOrderSeeder) createCompleted(options map[string]any) *entity.MaterialTestOrder {
	datetimeOptions := lo.PickByKeys(s.createDateTimeOptions(), []string{
		"PaymentSubmittedAt",
		"PaymentApprovedAt",
		"TestingAt",
		"CompletedAt",
		"CreatedAt",
		"UpdatedAt",
	})

	options = lo.Assign(datetimeOptions, map[string]any{
		"Status": materialTestOrderEnum.MaterialTestOrderStatusCompleted,
	}, options)

	return s.materialTestOrderFactory.MustCreateWithOption(options).(*entity.MaterialTestOrder)
}

func (s *MaterialTestOrderSeeder) createCancelled(options map[string]any) *entity.MaterialTestOrder {
	datetimeOptions := lo.PickByKeys(s.createDateTimeOptions(), []string{
		"CancelledAt",
		"CreatedAt",
		"UpdatedAt",
	})

	options = lo.Assign(datetimeOptions, map[string]any{
		"Status": materialTestOrderEnum.MaterialTestOrderStatusCancelled,
	}, options)

	return s.materialTestOrderFactory.MustCreateWithOption(options).(*entity.MaterialTestOrder)
}

func (s *MaterialTestOrderSeeder) createRejected(options map[string]any) *entity.MaterialTestOrder {
	datetimeOptions := lo.PickByKeys(s.createDateTimeOptions(), []string{
		"RejectedAt",
		"CreatedAt",
		"UpdatedAt",
	})

	options = lo.Assign(datetimeOptions, map[string]any{
		"Status": materialTestOrderEnum.MaterialTestOrderStatusRejected,
	}, options)

	return s.materialTestOrderFactory.MustCreateWithOption(options).(*entity.MaterialTestOrder)
}

func (s *MaterialTestOrderSeeder) createRefunded(options map[string]any) *entity.MaterialTestOrder {
	datetimeOptions := lo.PickByKeys(s.createDateTimeOptions(), []string{
		"PaymentSubmittedAt",
		"PaymentApprovedAt",
		"TestingAt",
		"RefundedAt",
		"CreatedAt",
		"UpdatedAt",
	})

	options = lo.Assign(datetimeOptions, map[string]any{
		"Status": materialTestOrderEnum.MaterialTestOrderStatusRefunded,
	}, options)

	return s.materialTestOrderFactory.MustCreateWithOption(options).(*entity.MaterialTestOrder)
}
