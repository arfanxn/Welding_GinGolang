package step

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/arfanxn/welding/internal/infrastructure/id"
	mtoEnum "github.com/arfanxn/welding/internal/module/material_test_order/domain/enum"
	mtoRepository "github.com/arfanxn/welding/internal/module/material_test_order/domain/repository"
	"github.com/arfanxn/welding/internal/module/material_test_order/usecase/dto"
	mtosRepository "github.com/arfanxn/welding/internal/module/material_test_order_service/domain/repository"
	mtoseRepository "github.com/arfanxn/welding/internal/module/material_test_order_service_evaluation/domain/repository"
	materialTestOrderUserEnum "github.com/arfanxn/welding/internal/module/material_test_order_user/domain/enum"
	mtouRepository "github.com/arfanxn/welding/internal/module/material_test_order_user/domain/repository"
	mtsRepository "github.com/arfanxn/welding/internal/module/material_test_service/domain/repository"
	"github.com/arfanxn/welding/internal/module/shared/contextkey"
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	"github.com/arfanxn/welding/internal/module/shared/domain/errorx"
	"github.com/arfanxn/welding/pkg/query"
	"github.com/arfanxn/welding/pkg/typeutil"
	"github.com/gookit/goutil"
	"github.com/samber/lo"
	"go.uber.org/fx"
)

type SaveMaterialTestOrderStep interface {
	Handle(ctx context.Context, dto *dto.SaveMaterialTestOrder) (*entity.MaterialTestOrder, error)
}

type saveMaterialTestOrderStep struct {
	idService       id.IdService
	mtoRepository   mtoRepository.MaterialTestOrderRepository
	mtouRepository  mtouRepository.MaterialTestOrderUserRepository
	mtsRepository   mtsRepository.MaterialTestServiceRepository
	mtosRepository  mtosRepository.MaterialTestOrderServiceRepository
	mtoseRepository mtoseRepository.MaterialTestOrderServiceEvaluationRepository
}

type NewSaveMaterialTestOrderStepParams struct {
	fx.In

	IdService                                    id.IdService
	MaterialTestOrderRepository                  mtoRepository.MaterialTestOrderRepository
	MaterialTestOrderUserRepository              mtouRepository.MaterialTestOrderUserRepository
	MaterialTestServiceRepository                mtsRepository.MaterialTestServiceRepository
	MaterialTestOrderServiceEvaluationRepository mtoseRepository.MaterialTestOrderServiceEvaluationRepository
	MaterialTestOrderServiceRepository           mtosRepository.MaterialTestOrderServiceRepository
}

func NewSaveMaterialTestOrderStep(params NewSaveMaterialTestOrderStepParams) SaveMaterialTestOrderStep {
	return &saveMaterialTestOrderStep{
		idService:       params.IdService,
		mtoRepository:   params.MaterialTestOrderRepository,
		mtouRepository:  params.MaterialTestOrderUserRepository,
		mtsRepository:   params.MaterialTestServiceRepository,
		mtosRepository:  params.MaterialTestOrderServiceRepository,
		mtoseRepository: params.MaterialTestOrderServiceEvaluationRepository,
	}
}

func (s *saveMaterialTestOrderStep) Handle(ctx context.Context, _dto *dto.SaveMaterialTestOrder) (
	*entity.MaterialTestOrder, error,
) {
	var (
		authUser = contextkey.GetUser(ctx)
		err      error
		q        = query.NewQuery()
		mto      *entity.MaterialTestOrder
		mtoId    string
		mtous    = []*entity.MaterialTestOrderUser{}
		mtoss    = []*entity.MaterialTestOrderService{}
		mtoses   = []*entity.MaterialTestOrderServiceEvaluation{}
	)

	q = q.Include("work_category").
		Include("ordered_services.evaluation").
		Include("ordered_services.service")

	// Retrieve existing material test order or create new one
	if !goutil.IsEmptyReal(_dto.Id) {
		// Update scenario: fetch existing mto
		mtoId = *_dto.Id
		q = q.FilterById(mtoId)
		mto, err = s.mtoRepository.First(q)
		if err != nil {
			return nil, err
		}
	} else {
		mtoId = s.idService.Generate()

		// Create scenario: initialize new material test order with generated ID
		latestMto, err := s.mtoRepository.FindLatestThisYear(nil)
		if err != nil {
			if errors.Is(err, errorx.ErrMaterialTestOrderNotFound) {
				// DO NOTHING
			} else {
				return nil, err
			}
		}
		mtoNumber := "001"
		if latestMto != nil {
			latestMtoNumberInt, err := strconv.Atoi(latestMto.Number)
			if err != nil {
				return nil, err
			}
			mtoNumber = fmt.Sprintf("%03d", latestMtoNumberInt+1)
		}

		mtoCreatedAt := time.Now()

		q = q.FilterById(mtoId)
		mto = &entity.MaterialTestOrder{
			Id:        mtoId,
			Number:    mtoNumber,
			CreatedAt: mtoCreatedAt,
		}

		// Add the authenticated user as the creator of this material test order
		// This creates an association between the order and the user with 'creator' type
		mtous = append(mtous, &entity.MaterialTestOrderUser{
			OrderId: mto.Id,
			UserId:  authUser.Id,
			Type:    materialTestOrderUserEnum.TypeCreator,
		})
	}

	if !goutil.IsEmptyReal(_dto.WorkCategoryId) {
		mto.WorkCategoryId = *_dto.WorkCategoryId
	}

	if !goutil.IsEmptyReal(_dto.WorkPackageName) {
		mto.WorkPackageName = *_dto.WorkPackageName
	}

	if !goutil.IsEmptyReal(_dto.ApplicantName) {
		mto.ApplicantName = *_dto.ApplicantName
	}

	if _dto.ApplicantCompanyName != nil {
		if !goutil.IsEmptyReal(_dto.ApplicantCompanyName) {
			mto.ApplicantCompanyName = *_dto.ApplicantCompanyName
		} else {
			mto.ApplicantCompanyName = ""
		}
	}

	if !goutil.IsEmptyReal(_dto.ApplicantPhoneNumber) {
		mto.ApplicantPhoneNumber = *_dto.ApplicantPhoneNumber
	}

	if !goutil.IsEmptyReal(_dto.ApplicantEmail) {
		mto.ApplicantEmail = *_dto.ApplicantEmail
	}

	if !goutil.IsEmptyReal(_dto.ApplicantFullAddress) {
		mto.ApplicantFullAddress = *_dto.ApplicantFullAddress
	}

	if _dto.ApplicantNote != nil {
		if !goutil.IsEmptyReal(_dto.ApplicantNote) {
			mto.ApplicantNote = _dto.ApplicantNote
		} else {
			mto.ApplicantNote = nil
		}
	}

	if !goutil.IsEmptyReal(_dto.RecipientName) {
		mto.RecipientName = *_dto.RecipientName
	}

	if _dto.RecipientFullAddress != nil {
		if !goutil.IsEmptyReal(_dto.RecipientFullAddress) {
			mto.RecipientFullAddress = *_dto.RecipientFullAddress
		} else {
			mto.RecipientFullAddress = ""
		}
	}

	if _dto.TesterNote != nil {
		if !goutil.IsEmptyReal(_dto.TesterNote) {
			mto.TesterNote = _dto.TesterNote
		} else {
			mto.TesterNote = nil
		}
	}

	if !goutil.IsEmptyReal(_dto.Status) {
		mto.Status = *_dto.Status
		switch mto.Status {
		case mtoEnum.MaterialTestOrderStatusAwaitingReview:
			mto.SubmittedAt = typeutil.Ptr(time.Now())
		case mtoEnum.MaterialTestOrderStatusDraft:
			mto.SubmittedAt = nil
		}
	}

	if _dto.OwnerUserIds != nil {
		if err := s.mtouRepository.DestroyOwnerByOrderId(mtoId); err != nil {
			return nil, err
		}

		if len(_dto.OwnerUserIds) > 0 && !goutil.IsEmptyReal(_dto.OwnerUserIds[0]) {
			for _, ownerUserId := range _dto.OwnerUserIds {
				mtous = append(mtous, &entity.MaterialTestOrderUser{
					OrderId: mtoId,
					UserId:  ownerUserId,
					Type:    materialTestOrderUserEnum.TypeOwner,
				})
			}
		}
	}

	if _dto.OrderedServices != nil {
		orderId := mtoId
		if err := s.mtosRepository.DestroyByOrderId(orderId); err != nil {
			return nil, err
		}
		mto.SubTotal = 0

		if len(_dto.OrderedServices) > 0 && !goutil.IsEmptyReal(_dto.OrderedServices[0]) {
			serviceIds := []string{}
			for _, orderedService := range _dto.OrderedServices {
				serviceIds = append(serviceIds, orderedService.ServiceId)
			}
			mtss, err := s.mtsRepository.FindByIds(serviceIds, nil)
			if err != nil {
				return nil, err
			}

			latestMtoss, err := s.mtosRepository.FindLatestThisYearPerServiceByServiceIds(serviceIds, nil)
			if err != nil {
				return nil, err
			}

			for _, orderedService := range _dto.OrderedServices {
				mts, _ := lo.Find(mtss, func(_mts *entity.MaterialTestService) bool {
					return _mts.Id == orderedService.ServiceId
				})

				latestMtos, latestMtosFound := lo.Find(latestMtoss, func(_mtos *entity.MaterialTestOrderService) bool {
					return _mtos.ServiceId == orderedService.ServiceId
				})

				mtosNextSeq := 1
				if latestMtosFound {
					latestMtosSampleNumber := latestMtos.SampleNumber

					lastThreeDigitsString := string(latestMtosSampleNumber[len(latestMtosSampleNumber)-3:])
					lastThreeDigits, _ := strconv.Atoi(lastThreeDigitsString)
					mtosNextSeq = lastThreeDigits + 1
				}

				sampleNumber := s.createSampleNumber(mto.Number, mts.ServiceCode, mtosNextSeq)
				for _, _mtos := range mtoss {
					if _mtos.SampleNumber == sampleNumber {
						mtosNextSeq++
						sampleNumber = s.createSampleNumber(mto.Number, mts.ServiceCode, mtosNextSeq)
					}
				}

				mtos := &entity.MaterialTestOrderService{
					Id:           s.idService.Generate(),
					OrderId:      orderId,
					ServiceId:    orderedService.ServiceId,
					SampleNumber: sampleNumber,
					SampleName:   orderedService.SampleName,
					Price:        mts.Price,
					Quantity:     orderedService.Quantity,
					LineTotal:    mts.Price * float64(orderedService.Quantity),
				}
				mtoss = append(mtoss, mtos)

				mtoses = append(mtoses, &entity.MaterialTestOrderServiceEvaluation{
					Id:             s.idService.Generate(),
					OrderServiceId: mtos.Id,
				})

				mto.SubTotal += mtos.LineTotal
			}
		}
	}

	if mto.SubTotal > 0 {
		mto.Total = mto.SubTotal + mto.Tax - mto.Discount
	} else {
		mto.Tax = 0
		mto.Discount = 0
		mto.Total = 0
	}

	if err := s.mtoRepository.Save(mto); err != nil {
		return nil, err
	}

	if err := s.mtouRepository.SaveMany(mtous); err != nil {
		return nil, err
	}

	if err := s.mtosRepository.SaveMany(mtoss); err != nil {
		return nil, err
	}

	if err := s.mtoseRepository.SaveMany(mtoses); err != nil {
		return nil, err
	}

	if mto, err = s.mtoRepository.First(q); err != nil {
		return nil, err
	}

	return mto, nil
}

func (saveMaterialTestOrderStep) createSampleNumber(orderNumber string, serviceCode string, sequence int) string {
	now := time.Now()
	sampleNumber := fmt.Sprintf(
		"%02d/%d.%s/%s/%03d",
		int(now.Month()), // MM
		now.Year(),       // YYYY
		orderNumber,      // ORDER_NUMBER
		serviceCode,      // SERVICE_CODE
		sequence,         // SEQ (3 digits)
	)
	return sampleNumber
}
