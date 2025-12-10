package seeder

import (
	"github.com/arfanxn/welding/internal/infrastructure/id"
	mtWorkCategoryRepository "github.com/arfanxn/welding/internal/module/material_test_work_category/domain/repository"
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	factoryGo "github.com/bluele/factory-go/factory"
	"go.uber.org/fx"
)

var _ Seeder = (*MaterialTestWorkCategorySeeder)(nil)

type MaterialTestWorkCategorySeeder struct {
	idService                id.IdService
	mtWorkCategoryFactory    *factoryGo.Factory
	mtWorkCategoryRepository mtWorkCategoryRepository.MaterialTestWorkCategoryRepository
}

type NewMaterialTestWorkCategorySeederParams struct {
	fx.In

	IdService                          id.IdService
	MaterialTestWorkCategoryFactory    *factoryGo.Factory `name:"material_test_work_category_factory"`
	MaterialTestWorkCategoryRepository mtWorkCategoryRepository.MaterialTestWorkCategoryRepository
}

func NewMaterialTestWorkCategorySeeder(
	params NewMaterialTestWorkCategorySeederParams,
) Seeder {
	return &MaterialTestWorkCategorySeeder{
		idService:                params.IdService,
		mtWorkCategoryFactory:    params.MaterialTestWorkCategoryFactory,
		mtWorkCategoryRepository: params.MaterialTestWorkCategoryRepository,
	}
}

func (s *MaterialTestWorkCategorySeeder) Seed() error {
	mtWorkCategoryFactory := s.mtWorkCategoryFactory

	mtWorkCategories := []*entity.MaterialTestWorkCategory{
		mtWorkCategoryFactory.MustCreateWithOption(map[string]any{
			"Id":   "01ARZ3NDEKMTWC4RRFFQ69G5FA",
			"Name": "Industri",
		}).(*entity.MaterialTestWorkCategory),
		mtWorkCategoryFactory.MustCreateWithOption(map[string]any{
			"Id":   "01ARZ3NDEKMTWC7H2D8P9S4XKB",
			"Name": "IKM",
		}).(*entity.MaterialTestWorkCategory),
		mtWorkCategoryFactory.MustCreateWithOption(map[string]any{
			"Id":   "01ARZ3NDEKMTWC8A3C7ZB52R9C",
			"Name": "Pendidikan",
		}).(*entity.MaterialTestWorkCategory),
		mtWorkCategoryFactory.MustCreateWithOption(map[string]any{
			"Id":   "01ARZ3NDEKMTWC9R4F6XK8S5AD",
			"Name": "Other",
		}).(*entity.MaterialTestWorkCategory),
	}

	if err := s.mtWorkCategoryRepository.SaveMany(mtWorkCategories); err != nil {
		return err
	}

	return nil
}
