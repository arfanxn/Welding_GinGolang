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
		&entity.MaterialTestMethod{Id: "01KAXAS44M2M4M3Z4GJ6P0N2XH", Name: "ASTM E8", Description: typeutil.Ptr("Standard Test Methods for Tension Testing of Metallic Materials")},
		&entity.MaterialTestMethod{Id: "01KAXAS44N6Q6X6V7JQ3F9Y1DT", Name: "SNI 1974:2011", Description: typeutil.Ptr("Metode Uji Tarik Logam")},
		&entity.MaterialTestMethod{Id: "01KAXAS44PBQ8H9C8WZK6S7JQF", Name: "Internal SOP 2022", Description: typeutil.Ptr("Tensile Testing Procedure")},
	)

	// Material test methods, dummy but realistic data, for testing purposes
	mtms = append(mtms,
		&entity.MaterialTestMethod{Id: "01KAXAS43A1Q3E8VZD7J4X5Z9F", Name: "ASTM E18", Description: typeutil.Ptr("Rockwell Hardness Test")},
		&entity.MaterialTestMethod{Id: "01KAXAS43B7N6ZP8T4M29H1Q0G", Name: "ASTM E10", Description: typeutil.Ptr("Brinell Hardness Test")},
		&entity.MaterialTestMethod{Id: "01KAXAS43C4R8Q7H5F6VD8SY2T", Name: "ASTM E92", Description: typeutil.Ptr("Vickers Hardness Test")},
		&entity.MaterialTestMethod{Id: "01KAXAS43D9T1Y4NQ2XWQTSGKM", Name: "ASTM E9", Description: typeutil.Ptr("Compression Test")},
		&entity.MaterialTestMethod{Id: "01KAXAS43ECPZ1Q4GZCJ4VDTYB", Name: "ASTM E21", Description: typeutil.Ptr("Elevated Temperature Tensile Test")},
		&entity.MaterialTestMethod{Id: "01KAXAS43FJ8C6NSR4M3H8EPGJ", Name: "ASTM A370", Description: typeutil.Ptr("Mechanical Testing of Steel")},
		&entity.MaterialTestMethod{Id: "01KAXAS43GQK2N96V8Z2T1BM6F", Name: "ISO 6892-1", Description: typeutil.Ptr("Tensile Test at Room Temperature")},
		&entity.MaterialTestMethod{Id: "01KAXAS43HTX4P5BDQ8W3Y9XJD", Name: "ISO 6508", Description: typeutil.Ptr("Rockwell Hardness")},
		&entity.MaterialTestMethod{Id: "01KAXAS43J23NQ7R6HW2K8VT5P", Name: "ISO 6507", Description: typeutil.Ptr("Vickers Hardness")},
		&entity.MaterialTestMethod{Id: "01KAXAS43K68WZF21XJ9V2TQQW", Name: "ISO 6506", Description: typeutil.Ptr("Brinell Hardness")},

		&entity.MaterialTestMethod{Id: "01KAXAS43M2YH3C5W8XKHW4S3C", Name: "ASTM E23", Description: typeutil.Ptr("Charpy Impact Test")},
		&entity.MaterialTestMethod{Id: "01KAXAS43N7QK1FQPTA2QGHN2R", Name: "ASTM E112", Description: typeutil.Ptr("Grain Size Measurement")},
		&entity.MaterialTestMethod{Id: "01KAXAS43PCHWQ4D2RTN5MZHPP", Name: "ASTM E3", Description: typeutil.Ptr("Metallographic Sample Preparation")},
		&entity.MaterialTestMethod{Id: "01KAXAS43QH4TWD9A2EP9M8CGF", Name: "ASTM E407", Description: typeutil.Ptr("Microetching for Metallography")},
		&entity.MaterialTestMethod{Id: "01KAXAS43RKM1326DXC69AYQNS", Name: "ASTM G48", Description: typeutil.Ptr("Pitting Corrosion Test")},
		&entity.MaterialTestMethod{Id: "01KAXAS43TQ21SY8VGNYQ8KMWZ", Name: "ASTM B117", Description: typeutil.Ptr("Salt Spray (Fog) Test")},
		&entity.MaterialTestMethod{Id: "01KAXAS43W12KZ4Q6VQXNBH1SB", Name: "ISO 9227", Description: typeutil.Ptr("Salt Spray Test")},
		&entity.MaterialTestMethod{Id: "01KAXAS43X6PD2E4J7FQW2FSYF", Name: "ASTM G31", Description: typeutil.Ptr("Immersion Corrosion Test")},
		&entity.MaterialTestMethod{Id: "01KAXAS43YBWMH7XWVT6S2GMDH", Name: "ASTM G150", Description: typeutil.Ptr("Electrochemical Test")},

		&entity.MaterialTestMethod{Id: "01KAXAS440D3SYZEFX3TRW63KZ", Name: "ASTM E165", Description: typeutil.Ptr("Liquid Penetrant Test (PT)")},
		&entity.MaterialTestMethod{Id: "01KAXAS441GFM4TD2RZCM2H2D8", Name: "ASTM E1444", Description: typeutil.Ptr("Magnetic Particle Test (MT)")},
		&entity.MaterialTestMethod{Id: "01KAXAS442N7Q9MDQ3Y1XJ1B4K", Name: "ASTM E94", Description: typeutil.Ptr("Radiography Testing (RT)")},
		&entity.MaterialTestMethod{Id: "01KAXAS443RTP2V1NWX7PM5X8R", Name: "ASTM E125", Description: typeutil.Ptr("Magnetic Particle Indications")},
		&entity.MaterialTestMethod{Id: "01KAXAS444W9XMZR71C9S67QYF", Name: "ISO 9712", Description: typeutil.Ptr("NDT Personnel Certification")},

		&entity.MaterialTestMethod{Id: "01KAXAS445BFQ2WHVFM2Q8P5A2", Name: "ASTM D638", Description: typeutil.Ptr("Tensile Test for Plastics")},
		&entity.MaterialTestMethod{Id: "01KAXAS446EJ1KQP9MHZ1WQC29", Name: "ASTM D256", Description: typeutil.Ptr("Izod Impact")},
		&entity.MaterialTestMethod{Id: "01KAXAS447JH8ZRQ2NSDW8FM4C", Name: "ASTM D790", Description: typeutil.Ptr("Flexural Properties of Plastics")},
		&entity.MaterialTestMethod{Id: "01KAXAS448NTCKX5Q3PFBZQ1KG", Name: "ISO 527", Description: typeutil.Ptr("Tensile Properties of Plastics")},

		&entity.MaterialTestMethod{Id: "01KAXAS449RZCY1S66JW4A51QF", Name: "ASTM C39", Description: typeutil.Ptr("Compressive Strength of Concrete")},
		&entity.MaterialTestMethod{Id: "01KAXAS44ATW2KQF0MTQ7JZPQH", Name: "ASTM C109", Description: typeutil.Ptr("Compressive Strength of Mortar")},
		&entity.MaterialTestMethod{Id: "01KAXAS44C35MFZTDXZN2S8ZD4", Name: "ASTM C143", Description: typeutil.Ptr("Slump Test")},
		&entity.MaterialTestMethod{Id: "01KAXAS44D7QJ2SQHQRMZWS4V9", Name: "SNI 1972", Description: typeutil.Ptr("Kuat Tekan Beton")},
		&entity.MaterialTestMethod{Id: "01KAXAS44ECP8NY7G4XHFH8ARP", Name: "SNI 4431", Description: typeutil.Ptr("Beton Ready Mix")},

		&entity.MaterialTestMethod{Id: "01KAXAS44FDX9MG2NTQ4B6B8C1", Name: "ASTM E415", Description: typeutil.Ptr("Chemical Analysis of Carbon Steel")},
		&entity.MaterialTestMethod{Id: "01KAXAS44GJAW3X1QDTCV7WZ33", Name: "ASTM E1086", Description: typeutil.Ptr("Optical Emission Spectrometry")},
		&entity.MaterialTestMethod{Id: "01KAXAS44HHQPCQ6T7X6MCG4YF", Name: "ASTM D129", Description: typeutil.Ptr("Sulfur in Petroleum")},
		&entity.MaterialTestMethod{Id: "01KAXAS44J1XHDH78QY4EFB79J", Name: "ASTM D1744", Description: typeutil.Ptr("Water Content by Karl Fischer")},
		&entity.MaterialTestMethod{Id: "01KAXAS44K5Z6KDZ9TNV7H22WF", Name: "ISO 17025", Description: typeutil.Ptr("General Laboratory Requirements")},
	)

	err := s.mtmRepository.SaveMany(mtms)
	if err != nil {
		return err
	}

	return nil
}
