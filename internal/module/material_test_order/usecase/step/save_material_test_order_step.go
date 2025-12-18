package step

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/arfanxn/welding/internal/infrastructure/id"
	mtoRepository "github.com/arfanxn/welding/internal/module/material_test_order/domain/repository"
	"github.com/arfanxn/welding/internal/module/material_test_order/usecase/dto"
	mtosRepository "github.com/arfanxn/welding/internal/module/material_test_order_service/domain/repository"
	mtoseRepository "github.com/arfanxn/welding/internal/module/material_test_order_service_evaluation/domain/repository"
	materialTestOrderUserEnum "github.com/arfanxn/welding/internal/module/material_test_order_user/domain/enum"
	mtouRepository "github.com/arfanxn/welding/internal/module/material_test_order_user/domain/repository"
	mtsRepository "github.com/arfanxn/welding/internal/module/material_test_service/domain/repository"
	mediaEnum "github.com/arfanxn/welding/internal/module/media/domain/enum"
	mediaRepository "github.com/arfanxn/welding/internal/module/media/domain/repository"
	mediaDto "github.com/arfanxn/welding/internal/module/media/usecase/dto"
	mediaService "github.com/arfanxn/welding/internal/module/media/usecase/service"
	"github.com/arfanxn/welding/internal/module/shared/contextkey"
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	"github.com/arfanxn/welding/pkg/query"
	"github.com/davecgh/go-spew/spew"
	"github.com/gookit/goutil"
	"github.com/samber/lo"
	"go.uber.org/fx"
)

type SaveMaterialTestOrderStep interface {
	Handle(ctx context.Context, dto *dto.SaveMaterialTestOrder) (*entity.MaterialTestOrder, error)
}

type saveMaterialTestServiceStep struct {
	idService       id.IdService
	mtoRepository   mtoRepository.MaterialTestOrderRepository
	mtouRepository  mtouRepository.MaterialTestOrderUserRepository
	mtsRepository   mtsRepository.MaterialTestServiceRepository
	mtosRepository  mtosRepository.MaterialTestOrderServiceRepository
	mtoseRepository mtoseRepository.MaterialTestOrderServiceEvaluationRepository
	mediaRepository mediaRepository.MediaRepository
	mediaService    mediaService.MediaService
}

type NewSaveMaterialTestOrderStepParams struct {
	fx.In

	IdService                                    id.IdService
	MaterialTestOrderRepository                  mtoRepository.MaterialTestOrderRepository
	MaterialTestOrderUserRepository              mtouRepository.MaterialTestOrderUserRepository
	MaterialTestServiceRepository                mtsRepository.MaterialTestServiceRepository
	MaterialTestOrderServiceEvaluationRepository mtoseRepository.MaterialTestOrderServiceEvaluationRepository
	MaterialTestOrderServiceRepository           mtosRepository.MaterialTestOrderServiceRepository
	MediaRepository                              mediaRepository.MediaRepository
	MediaService                                 mediaService.MediaService
}

func NewSaveMaterialTestOrderStep(params NewSaveMaterialTestOrderStepParams) SaveMaterialTestOrderStep {
	return &saveMaterialTestServiceStep{
		idService:       params.IdService,
		mtoRepository:   params.MaterialTestOrderRepository,
		mtouRepository:  params.MaterialTestOrderUserRepository,
		mtsRepository:   params.MaterialTestServiceRepository,
		mtosRepository:  params.MaterialTestOrderServiceRepository,
		mtoseRepository: params.MaterialTestOrderServiceEvaluationRepository,
		mediaRepository: params.MediaRepository,
		mediaService:    params.MediaService,
	}
}

func (s *saveMaterialTestServiceStep) Handle(ctx context.Context, _dto *dto.SaveMaterialTestOrder) (
	*entity.MaterialTestOrder, error,
) {
	var (
		authUser                 = contextkey.GetUser(ctx)
		q                        = query.NewQuery()
		mto                      *entity.MaterialTestOrder
		mtoId                    string
		mtoNumber                string
		mtoCreatedAt             time.Time
		mtous                    = []*entity.MaterialTestOrderUser{}
		err                      error
		mtoss                    = []*entity.MaterialTestOrderService{}
		mtoses                   = []*entity.MaterialTestOrderServiceEvaluation{}
		createFromMultipartFiles = mediaDto.CreateFromMultipartFiles{}
	)

	if _dto.WorkCategoryId != nil {
		q = q.Include("work_category")
	}

	if _dto.OrderedServices != nil {
		q = q.Include("ordered_services.evaluation")
	}

	if _dto.Medias != nil {
		q = q.Include("medias")
	}

	// Retrieve existing material test order or create new one
	if !goutil.IsEmptyReal(_dto.Id) {
		// Update scenario: fetch existing mto
		mtoId = *_dto.Id
		q = q.FilterById(mtoId)
		mto, err = s.mtoRepository.First(q)
		if err != nil {
			return nil, err
		}
		mtoCreatedAt = mto.CreatedAt
	} else {
		// Create scenario: initialize new material test order with generated ID
		mtoId = s.idService.Generate()
		mtoNumber = string(mtoId[len(mtoId)-9:])
		mtoCreatedAt = time.Now()
		q = q.FilterById(mtoId)
		mto = &entity.MaterialTestOrder{Id: mtoId, Number: mtoNumber, CreatedAt: mtoCreatedAt}

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

	if _dto.TesterNote != nil {
		if !goutil.IsEmptyReal(_dto.TesterNote) {
			mto.TesterNote = _dto.TesterNote
		} else {
			mto.TesterNote = nil
		}
	}

	if !goutil.IsEmptyReal(_dto.Status) {
		mto.Status = *_dto.Status
	}

	if _dto.OwnerUserIds != nil {
		if err := s.mtouRepository.DestroyOwnerByOrderId(mtoId); err != nil {
			return nil, err
		}

		if !goutil.IsEmptyReal(_dto.OwnerUserIds[0]) {
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

		if !goutil.IsEmptyReal(_dto.OrderedServices[0]) {
			serviceIds := []string{}
			for _, orderedService := range _dto.OrderedServices {
				serviceIds = append(serviceIds, orderedService.ServiceId)
			}
			mtss, err := s.mtsRepository.FindByIds(serviceIds, nil)
			if err != nil {
				return nil, err
			}

			latestMtoss, err := s.mtosRepository.FindLatestPerServiceByServiceIds(serviceIds, nil)
			if err != nil {
				return nil, err
			}

			for _, orderedService := range _dto.OrderedServices {
				mts, _ := lo.Find(mtss, func(mt *entity.MaterialTestService) bool {
					return mt.Id == orderedService.ServiceId
				})

				latestMtos, latestMtosFound := lo.Find(latestMtoss, func(mt *entity.MaterialTestOrderService) bool {
					return mt.ServiceId == orderedService.ServiceId
				})

				nextLastThreeDigits := 1
				if latestMtosFound {
					latestMtosSampleNumber := latestMtos.SampleNumber

					lastThreeDigitsString := string(latestMtosSampleNumber[len(latestMtosSampleNumber)-3:])
					lastThreeDigits, _ := strconv.Atoi(lastThreeDigitsString)
					nextLastThreeDigits = lastThreeDigits + 1
				}

				sampledAt := mtoCreatedAt
				sampleNumber := fmt.Sprintf("%02d/%d.%.1f/%s/%03d", // eg: 01/2025.1/UTK/001
					sampledAt.Day(),
					sampledAt.Year(),
					float64(sampledAt.Month())/1,
					mts.ServiceCode,
					nextLastThreeDigits,
				)

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

	if _dto.Medias != nil {
		// TODO: not only destroy but also delete the files
		if err := s.mediaRepository.DestroyByModelTypeAndModelId(mediaEnum.ModelTypeMaterialTestOrder, mtoId); err != nil {
			return nil, err
		}

		if !goutil.IsEmptyReal(_dto.Medias[0]) {
			for _, m := range _dto.Medias {
				createFromMultipartFiles = append(createFromMultipartFiles, mediaDto.CreateFromMultipartFile{
					File: m.File,
					Name: *m.Name,
				})
			}
		}
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

	if _, err := s.mediaService.CreateFromMultipartFiles(ctx, createFromMultipartFiles); err != nil {
		return nil, err
	}

	if mto, err = s.mtoRepository.First(q); err != nil {
		return nil, err
	}

	spew.Dump(mto)

	return mto, nil
}
