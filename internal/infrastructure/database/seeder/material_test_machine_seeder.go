package seeder

import (
	"github.com/arfanxn/welding/internal/infrastructure/id"
	mtmRepository "github.com/arfanxn/welding/internal/module/material_test_machine/domain/repository"
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	"github.com/arfanxn/welding/pkg/typeutil"
	"go.uber.org/fx"
)

var _ Seeder = (*MaterialTestMachineSeeder)(nil)

type MaterialTestMachineSeeder struct {
	idService     id.IdService
	mtmRepository mtmRepository.MaterialTestMachineRepository
}

type NewMaterialTestMachineSeederParams struct {
	fx.In

	IdService     id.IdService
	MtmRepository mtmRepository.MaterialTestMachineRepository
}

func NewMaterialTestMachineSeeder(
	params NewMaterialTestMachineSeederParams,
) Seeder {
	return &MaterialTestMachineSeeder{
		idService:     params.IdService,
		mtmRepository: params.MtmRepository,
	}
}

func (s *MaterialTestMachineSeeder) Seed() error {
	var mtms []*entity.MaterialTestMachine

	// Welding company's available material test machines
	mtms = append(mtms,
		&entity.MaterialTestMachine{Id: "01KB051BW3GDGTC2NYS9QB0B4P", Name: "Universal Testing Machine", Description: typeutil.Ptr("Tests tensile and compressive strength")},
		&entity.MaterialTestMachine{Id: "01KB051BW331Z7TEEQB292EM2X", Name: "Oven", Description: typeutil.Ptr("Heats samples under controlled temperature")},
		&entity.MaterialTestMachine{Id: "01KB051BW3DS6F1J0QYKNFQBH0", Name: "pH Meter", Description: typeutil.Ptr("Measures acidity of materials")},
		&entity.MaterialTestMachine{Id: "01KB051BW3CNQ3ZGBHPW9EMV7N", Name: "CNC Lathe Machine", Description: typeutil.Ptr("Precision cutting and shaping tools")},
	)

	// Material test machines, dummy but realistic data, for testing purposes
	mtms = append(mtms,
		&entity.MaterialTestMachine{Id: "01M2Z0VGW7Q9PA4M3X6ZKJ9F9R", Name: "Hardness Tester", Description: typeutil.Ptr("Measures material hardness level")},
		&entity.MaterialTestMachine{Id: "01M2Z0VGW81Q5R6Z8DYX7T4H2M", Name: "Charpy Impact Tester", Description: typeutil.Ptr("Measures material impact toughness")},
		&entity.MaterialTestMachine{Id: "01M2Z0VGW8S6QWTH9K5C2MF0D8", Name: "Izod Impact Tester", Description: typeutil.Ptr("Tests impact resistance strength")},
		&entity.MaterialTestMachine{Id: "01M2Z0VGW90FT8SCD2F4Z3V8QD", Name: "Fatigue Testing Machine", Description: typeutil.Ptr("Tests cyclic load endurance")},
		&entity.MaterialTestMachine{Id: "01M2Z0VGW9ZJ1GFDC89KD7R0FS", Name: "Creep Testing Machine", Description: typeutil.Ptr("Measures long-term deformation behavior")},
		&entity.MaterialTestMachine{Id: "01M2Z0VGWA4C6H9ZB35P7WF6QX", Name: "Dynamic Mechanical Analyzer", Description: typeutil.Ptr("Analyzes viscoelastic properties")},
		&entity.MaterialTestMachine{Id: "01M2Z0VGWAGA3V9NFXZ4H78QHP", Name: "Differential Scanning Calorimeter", Description: typeutil.Ptr("Measures thermal transition behavior")},
		&entity.MaterialTestMachine{Id: "01M2Z0VGWAQPPX4ZHD7B9T2S4J", Name: "Thermogravimetric Analyzer", Description: typeutil.Ptr("Measures thermal weight changes")},
		&entity.MaterialTestMachine{Id: "01M2Z0VGWB0W8DNH7R24ZK1T06", Name: "XRF Spectrometer", Description: typeutil.Ptr("Analyzes elemental material composition")},
		&entity.MaterialTestMachine{Id: "01M2Z0VGWB8E4CKST2Q9JZYGFG", Name: "XRD Machine", Description: typeutil.Ptr("Identifies crystal material structures")},
		&entity.MaterialTestMachine{Id: "01M2Z0VGWBG4T8VYQ2H1X6D0AS", Name: "Metallurgical Microscope", Description: typeutil.Ptr("Observes metal microstructures")},
		&entity.MaterialTestMachine{Id: "01M2Z0VGWBPXJ6K9Z4TD3CFW2J", Name: "Scanning Electron Microscope", Description: typeutil.Ptr("Provides high-magnification imaging")},
		&entity.MaterialTestMachine{Id: "01M2Z0VGWBZV5K6JXH03C8FT0G", Name: "Microtome Machine", Description: typeutil.Ptr("Slices thin material samples")},
		&entity.MaterialTestMachine{Id: "01M2Z0VGWC7X12FH9N6J5BG4DQ", Name: "Ultrasonic Flaw Detector", Description: typeutil.Ptr("Detects internal material defects")},
		&entity.MaterialTestMachine{Id: "01M2Z0VGWCGTRJ3ZFW1BT87K0P", Name: "Magnetic Particle Tester", Description: typeutil.Ptr("Reveals magnetic surface cracks")},
		&entity.MaterialTestMachine{Id: "01M2Z0VGWCPZ7VB5DNF6X92JMG", Name: "Dye Penetrant Tester", Description: typeutil.Ptr("Detects surface-level cracks")},
		&entity.MaterialTestMachine{Id: "01M2Z0VGWCXFPG3D7TY4KSJ5PN", Name: "Roughness Tester", Description: typeutil.Ptr("Measures surface roughness values")},
		&entity.MaterialTestMachine{Id: "01M2Z0VGWD3WV9S2Q7ZJ4KBD8H", Name: "Profilometer", Description: typeutil.Ptr("Measures surface profile geometries")},
		&entity.MaterialTestMachine{Id: "01M2Z0VGWDC6MN47ZF1X0G8RFZ", Name: "Torque Testing Machine", Description: typeutil.Ptr("Measures rotational force strength")},
		&entity.MaterialTestMachine{Id: "01M2Z0VGWDLXW8HD1G7N2S0QAY", Name: "Compression Testing Machine", Description: typeutil.Ptr("Tests compressive load strength")},
		&entity.MaterialTestMachine{Id: "01M2Z0VGWDW2KDH6Z9QX3JP4BV", Name: "Wear Tester", Description: typeutil.Ptr("Measures friction and wear")},
		&entity.MaterialTestMachine{Id: "01M2Z0VGWE2MK1T3ZVH8C7QJDN", Name: "Salt Spray Chamber", Description: typeutil.Ptr("Tests corrosion resistance rates")},
		&entity.MaterialTestMachine{Id: "01M2Z0VGWE9S6W4T1D2FYQZKMB", Name: "Environmental Chamber", Description: typeutil.Ptr("Simulates environmental testing conditions")},
		&entity.MaterialTestMachine{Id: "01M2Z0VGWEHQFD0X8V5S9KTJ6P", Name: "UV Weathering Tester", Description: typeutil.Ptr("Simulates sunlight material degradation")},
		&entity.MaterialTestMachine{Id: "01M2Z0VGWEWQXG1Z5C9H27JQGR", Name: "Heat Treatment Furnace", Description: typeutil.Ptr("Heats metals for treatment")},
		&entity.MaterialTestMachine{Id: "01M2Z0VGWF4HTN8Z12X9MC6RTN", Name: "Muffle Furnace", Description: typeutil.Ptr("Burns samples at high temperatures")},
		&entity.MaterialTestMachine{Id: "01M2Z0VGWFC8MH5ZGN1T9BX40S", Name: "Kiln Dryer", Description: typeutil.Ptr("Dries materials using heat")},
		&entity.MaterialTestMachine{Id: "01M2Z0VGWFNQCY4JX0DG8S2QKW", Name: "Hydraulic Press", Description: typeutil.Ptr("Applies strong compressive force")},
		&entity.MaterialTestMachine{Id: "01M2Z0VGWFW2V9GJH4KTQ53D2E", Name: "Laser Cutter Machine", Description: typeutil.Ptr("Cuts materials with laser")},
		&entity.MaterialTestMachine{Id: "01M2Z0VGWG2S7VS5HZF9KQJ8CG", Name: "Waterjet Cutting Machine", Description: typeutil.Ptr("Cuts materials using waterjet")},
		&entity.MaterialTestMachine{Id: "01M2Z0VGWGAZK1N5VY4QSX8TMB", Name: "Spectrophotometer", Description: typeutil.Ptr("Measures light absorption levels")},
		&entity.MaterialTestMachine{Id: "01M2Z0VGWGPF4DBC6NQZ8KT1R0", Name: "Rheometer", Description: typeutil.Ptr("Measures material flow properties")},
		&entity.MaterialTestMachine{Id: "01M2Z0VGWGXKPR7HFD1Z93V4W6", Name: "Viscometer", Description: typeutil.Ptr("Measures fluid viscosity levels")},
		&entity.MaterialTestMachine{Id: "01M2Z0VGWH3E8XYC5TNS9F71MH", Name: "Moisture Analyzer", Description: typeutil.Ptr("Measures sample moisture content")},
		&entity.MaterialTestMachine{Id: "01M2Z0VGWHCQ0NT5ZJ82K9SW3G", Name: "Gloss Meter", Description: typeutil.Ptr("Measures surface gloss levels")},
		&entity.MaterialTestMachine{Id: "01M2Z0VGWHM7G2F8V0ZCXYT95H", Name: "Color Meter", Description: typeutil.Ptr("Measures surface color values")},
		&entity.MaterialTestMachine{Id: "01M2Z0VGWHW3Y4NSQ5JX1P7M8T", Name: "Tensile Grip Fixture", Description: typeutil.Ptr("Holds samples during tension")},
		&entity.MaterialTestMachine{Id: "01M2Z0VGWJ3Y9FHQ8BW6K21RTP", Name: "Bending Test Fixture", Description: typeutil.Ptr("Supports bending test samples")},
		&entity.MaterialTestMachine{Id: "01M2Z0VGWJC47X2PTHZ9K58SBF", Name: "Impact Testing Chamber", Description: typeutil.Ptr("Conditions samples before impact")},
		&entity.MaterialTestMachine{Id: "01M2Z0VGWJLH3Q5W8FZ7D1PG9E", Name: "Rotary Abrasion Tester", Description: typeutil.Ptr("Tests rotary wear resistance")},
		&entity.MaterialTestMachine{Id: "01M2Z0VGWJVX9RN4Z2TY7K1PGX", Name: "Shore Hardness Tester", Description: typeutil.Ptr("Measures rubber hardness levels")},
	)

	err := s.mtmRepository.SaveMany(mtms)
	if err != nil {
		return err
	}

	return nil
}
