package seeder

import (
	"github.com/arfanxn/welding/internal/infrastructure/id"
	mtWorkPackageRepository "github.com/arfanxn/welding/internal/module/material_test_work_package/domain/repository"
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	factoryGo "github.com/bluele/factory-go/factory"
	"go.uber.org/fx"
)

var _ Seeder = (*MaterialTestWorkPackageSeeder)(nil)

type MaterialTestWorkPackageSeeder struct {
	idService               id.IdService
	mtWorkPackageFactory    *factoryGo.Factory
	mtWorkPackageRepository mtWorkPackageRepository.MaterialTestWorkPackageRepository
}

type NewMaterialTestWorkPackageSeederParams struct {
	fx.In

	IdService                         id.IdService
	MaterialTestWorkPackageFactory    *factoryGo.Factory `name:"material_test_work_package_factory"`
	MaterialTestWorkPackageRepository mtWorkPackageRepository.MaterialTestWorkPackageRepository
}

func NewMaterialTestWorkPackageSeeder(
	params NewMaterialTestWorkPackageSeederParams,
) Seeder {
	return &MaterialTestWorkPackageSeeder{
		idService:               params.IdService,
		mtWorkPackageFactory:    params.MaterialTestWorkPackageFactory,
		mtWorkPackageRepository: params.MaterialTestWorkPackageRepository,
	}
}

func (s *MaterialTestWorkPackageSeeder) Seed() error {
	mtWorkPackageFactory := s.mtWorkPackageFactory

	mtWorkPackages := []*entity.MaterialTestWorkPackage{
		mtWorkPackageFactory.MustCreateWithOption(map[string]any{
			"Id":   "01J9X2A7MTWP4Q8Z6R3B5KCMNA",
			"Name": "Rehabilitasi Jalan Kabupaten Ruas Sukamaju–Sukajadi",
		}).(*entity.MaterialTestWorkPackage),
		mtWorkPackageFactory.MustCreateWithOption(map[string]any{
			"Id":   "01J9X2A8MTWP4Q8Z6R3B5KCMNB",
			"Name": "Pembangunan Gedung Kantor Kecamatan",
		}).(*entity.MaterialTestWorkPackage),
		mtWorkPackageFactory.MustCreateWithOption(map[string]any{
			"Id":   "01J9X2A9MTWP4Q8Z6R3B5KCMNC",
			"Name": "Pembangunan Jembatan Sungai Citarum",
		}).(*entity.MaterialTestWorkPackage),
		mtWorkPackageFactory.MustCreateWithOption(map[string]any{
			"Id":   "01J9X2B0MTWP4Q8Z6R3B5KCMND",
			"Name": "Peningkatan Jalan Lingkungan Kelurahan Mekarsari",
		}).(*entity.MaterialTestWorkPackage),
		mtWorkPackageFactory.MustCreateWithOption(map[string]any{
			"Id":   "01J9X2B1MTWP4Q8Z6R3B5KCMNE",
			"Name": "Pembangunan Gedung Puskesmas Pembantu",
		}).(*entity.MaterialTestWorkPackage),
		mtWorkPackageFactory.MustCreateWithOption(map[string]any{
			"Id":   "01J9X2B2MTWP4Q8Z6R3B5KCMNF",
			"Name": "Rehabilitasi Saluran Irigasi Desa Sindangjaya",
		}).(*entity.MaterialTestWorkPackage),
		mtWorkPackageFactory.MustCreateWithOption(map[string]any{
			"Id":   "01J9X2B3MTWP4Q8Z6R3B5KCMNG",
			"Name": "Pembangunan Rumah Susun Sederhana Sewa (Rusunawa)",
		}).(*entity.MaterialTestWorkPackage),
		mtWorkPackageFactory.MustCreateWithOption(map[string]any{
			"Id":   "01J9X2B4MTWP4Q8Z6R3B5KCMNH",
			"Name": "Pemeliharaan Berkala Jalan Provinsi",
		}).(*entity.MaterialTestWorkPackage),
		mtWorkPackageFactory.MustCreateWithOption(map[string]any{
			"Id":   "01J9X2B5MTWP4Q8Z6R3B5KCMNJ",
			"Name": "Pembangunan Embung Desa Karangnunggal",
		}).(*entity.MaterialTestWorkPackage),
		mtWorkPackageFactory.MustCreateWithOption(map[string]any{
			"Id":   "01J9X2B6MTWP4Q8Z6R3B5KCMNK",
			"Name": "Renovasi Gedung Sekolah Dasar Negeri 05",
		}).(*entity.MaterialTestWorkPackage),
	}

	if err := s.mtWorkPackageRepository.SaveMany(mtWorkPackages); err != nil {
		return err
	}

	return nil
}
