package di

import (
	"github.com/arfanxn/welding/internal/module/material_test_order/infrastructure/policy"
	repositoryImpl "github.com/arfanxn/welding/internal/module/material_test_order/infrastructure/repository"
	"github.com/arfanxn/welding/internal/module/material_test_order/presentation/http"
	"github.com/arfanxn/welding/internal/module/material_test_order/presentation/http/presenter"
	"github.com/arfanxn/welding/internal/module/material_test_order/usecase"
	"github.com/arfanxn/welding/internal/module/material_test_order/usecase/step"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"material_test_order",
	fx.Provide(
		repositoryImpl.NewGormMaterialTestOrderRepository,
		policy.NewMaterialTestOrderPolicy,
		step.NewSaveMaterialTestOrderStep,
		step.NewApproveMaterialTestOrderStep,
		step.NewRejectMaterialTestOrderStep,
		step.NewCancelMaterialTestOrderStep,
		step.NewSubmitPaymentMaterialTestOrderStep,
		step.NewApprovePaymentMaterialTestOrderStep,
		step.NewRejectPaymentMaterialTestOrderStep,
		step.NewTestMaterialTestOrderStep,
		step.NewRefundMaterialTestOrderStep,
		step.NewCompleteMaterialTestOrderStep,
		step.NewUpdateMaterialTestOrderServiceEvaluationStep,
		step.NewStoreMaterialTestOrderMediaStep,
		step.NewUpdateMaterialTestOrderMediaStep,
		step.NewDestroyMaterialTestOrderMediaStep,
		usecase.NewMaterialTestOrderUsecase,
		presenter.NewMaterialTestOrderPresenter,
		http.NewMaterialTestOrderHandler,
	),
)
