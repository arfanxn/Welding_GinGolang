package seeder

import (
	"github.com/arfanxn/welding/internal/infrastructure/id"
	mtmRepository "github.com/arfanxn/welding/internal/module/material_test_method/domain/repository"
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	"github.com/arfanxn/welding/pkg/typeutil"
	"go.uber.org/fx"
)

var _ Seeder = (*MaterialTestMethodSeeder)(nil)

type MaterialTestMethodSeeder struct {
	idService     id.IdService
	mtmRepository mtmRepository.MaterialTestMethodRepository
}

type NewMaterialTestMethodSeederParams struct {
	fx.In

	IdService     id.IdService
	MtmRepository mtmRepository.MaterialTestMethodRepository
}

func NewMaterialTestMethodSeeder(
	params NewMaterialTestMethodSeederParams,
) Seeder {
	return &MaterialTestMethodSeeder{
		idService:     params.IdService,
		mtmRepository: params.MtmRepository,
	}
}

func (s *MaterialTestMethodSeeder) Seed() error {
	var mtms []*entity.MaterialTestMethod

	// Welding company's available material test methods
	mtms = append(mtms,
		&entity.MaterialTestMethod{Id: "01KAXAS44M2M4M3Z4GJ6P0N2XH", Name: "JIS Z 2241:2011", Description: typeutil.Ptr("Logam dan produk logam")},
		&entity.MaterialTestMethod{Id: "01KAXAS44N6Q6X6V7JQ3F9Y1DT", Name: "SNI 8389:2017", Description: typeutil.Ptr("Logam dan produk logam; Baja Profil H Canai Panas (Bj PHC); Baja Profil Siku Sama Kaki Proses Canai Panas (Bj P Siku Sama Kaki); Baja Tulangan Beton")},
		&entity.MaterialTestMethod{Id: "01KAXAS44PBQ8H9C8WZK6S7JQF", Name: "JIS Z 2248:2006", Description: typeutil.Ptr("Logam dan produk logam")},
		&entity.MaterialTestMethod{Id: "01KAXAS43A1Q3E8VZD7J4X5Z9F", Name: "SNI 0410:2017", Description: typeutil.Ptr("Logam dan produk logam; Baja Profil H Canai Panas (Bj PHC); Baja Profil Siku Sama Kaki Proses Canai Panas (Bj P Siku Sama Kaki); Baja Tulangan Beton")},
		&entity.MaterialTestMethod{Id: "01KAXAS43B7N6ZP8T4M29H1Q0G", Name: "JIS Z 2242:2018", Description: typeutil.Ptr("Logam dan produk logam")},
		&entity.MaterialTestMethod{Id: "01KAXAS43C4R8Q7H5F6VD8SY2T", Name: "SNI 07-8732-2002", Description: typeutil.Ptr("Logam dan produk logam")},
		&entity.MaterialTestMethod{Id: "01KAXAS43D9T1Y4NQ2XWQTSGKM", Name: "JIS Z 2245:2016", Description: typeutil.Ptr("Logam dan produk logam")},
		&entity.MaterialTestMethod{Id: "01KAXAS43ECPZ1Q4GZCJ4VDTYB", Name: "SNI 8388:201", Description: typeutil.Ptr("Logam dan produk logam")},
		&entity.MaterialTestMethod{Id: "01KAXAS43FJ8C6NSR4M3H8EPGJ", Name: "JIS Z 2243:2008", Description: typeutil.Ptr("Logam dan produk logam")},
		&entity.MaterialTestMethod{Id: "01KAXAS43GQK2N96V8Z2T1BM6F", Name: "SNI 8387:2017", Description: typeutil.Ptr("Logam dan produk logam")},
		&entity.MaterialTestMethod{Id: "01KAXAS43HTX4P5BDQ8W3Y9XJD", Name: "SNI 2610:2011 Butir 9.1", Description: typeutil.Ptr("Baja Profil H Canai Panas (Bj PHC)")},
		&entity.MaterialTestMethod{Id: "01KAXAS43J23NQ7R6HW2K8VT5P", Name: "SNI 2610:2011 Butir 9.3.2", Description: typeutil.Ptr("Baja Profil H Canai Panas (Bj PHC)")},
		&entity.MaterialTestMethod{Id: "01KAXAS43K68WZF21XJ9V2TQQW", Name: "IK/UPTD-LAB/U-18 (kuantitatif)", Description: typeutil.Ptr("Baja Profil Siku Sama Kaki Proses Canai Panas (Bj P Siku Sama Kaki); Baja Tulangan Beton")},
		&entity.MaterialTestMethod{Id: "01KAXAS43M2YH3C5W8XKHW4S3C", Name: "SNI 2052:2017 Butir 8.2", Description: typeutil.Ptr("Baja Tulangan Beton")},
		&entity.MaterialTestMethod{Id: "01KAXAS43N7QK1FQPTA2QGHN2R", Name: "SNI 1974:2011", Description: typeutil.Ptr("Beton Silinder dan Beton Kubus")},
		&entity.MaterialTestMethod{Id: "01KAXAS43PCHWQ4D2RTN5MZHPP", Name: "IK/UPTD-LAB/5.4-25 (mesin tank tekan)", Description: typeutil.Ptr("Beton Silinder dan Beton Kubus")},
		&entity.MaterialTestMethod{Id: "01KAXAS43QH4TWD9A2EP9M8CGF", Name: "SNI 03-0691-1996 Butir 7.3", Description: typeutil.Ptr("Bata Beton (Paving Block)")},
		&entity.MaterialTestMethod{Id: "01KAXAS43RKM1326DXC69AYQNS", Name: "ASTM E 1251-11", Description: typeutil.Ptr("Logam Non Ferro Basis Alumunium")},
		&entity.MaterialTestMethod{Id: "01KAXAS43TQ21SY8VGNYQ8KMWZ", Name: "IK-UPTD-LAB/U-9 (Optical Emission Spectrometer", Description: typeutil.Ptr("Logam Non Ferro Basis Alumunium")},
		&entity.MaterialTestMethod{Id: "01KAXAS43W12KZ4Q6VQXNBH1SB", Name: "IK/UPTD-LAB/5.4-3 (Optical Emission Spectrometer)", Description: typeutil.Ptr("Logam Non Ferro Basis Alumunium")},
		&entity.MaterialTestMethod{Id: "01KAXAS43X6PD2E4J7FQW2FSYF", Name: "IK-UPTD-LAB/U-18 (Atomic Emission Spectrometer", Description: typeutil.Ptr("Baja Karbon dan Paduan Rendah")},
		&entity.MaterialTestMethod{Id: "01KAXAS43YBWMH7XWVT6S2GMDH", Name: "ASTM E 1086-14", Description: typeutil.Ptr("Baja Tahan Karat")},
		&entity.MaterialTestMethod{Id: "01KAXAS440D3SYZEFX3TRW63KZ", Name: "IK-UPTD-LAB/U-16", Description: typeutil.Ptr("Baja Tahan Karat")},
	)

	err := s.mtmRepository.SaveMany(mtms)
	if err != nil {
		return err
	}

	return nil
}
