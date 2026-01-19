package seeder

import (
	"github.com/arfanxn/welding/internal/infrastructure/id"
	mtmRepository "github.com/arfanxn/welding/internal/module/material_test_machine/domain/repository"
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
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
		&entity.MaterialTestMachine{Id: "01KB051BW3GDGTC2NYS9QB0B4P", Name: "Mesin Timbangan", Description: nil},
		&entity.MaterialTestMachine{Id: "01KB051BW331Z7TEEQB292EM2X", Name: "Mesin UTM", Description: nil},
		&entity.MaterialTestMachine{Id: "01KB051BW3DS6F1J0QYKNFQBH0", Name: "Mesin Impact", Description: nil},
		&entity.MaterialTestMachine{Id: "01KB051BW3CNQ3ZGBHPW9EMV7N", Name: "Komputer", Description: nil},
		&entity.MaterialTestMachine{Id: "01M2Z0VGW7Q9PA4M3X6ZKJ9F9R", Name: "Mesin LAS SMAW/GMAW", Description: nil},
		&entity.MaterialTestMachine{Id: "01M2Z0VGW81Q5R6Z8DYX7T4H2M", Name: "Mesin LAS TIG", Description: nil},
		&entity.MaterialTestMachine{Id: "01M2Z0VGW8S6QWTH9K5C2MF0D8", Name: "Mesin CNC Lathe", Description: nil},
		&entity.MaterialTestMachine{Id: "01M2Z0VGW90FT8SCD2F4Z3V8QD", Name: "Mesin CNC Millin Table 80 cm", Description: nil},
		&entity.MaterialTestMachine{Id: "01M2Z0VGW9ZJ1GFDC89KD7R0FS", Name: "Mesin CNC Millin Table 100 cm", Description: nil},
		&entity.MaterialTestMachine{Id: "01M2Z0VGWA4C6H9ZB35P7WF6QX", Name: "Mesin Wire Cut", Description: nil},
		&entity.MaterialTestMachine{Id: "01M2Z0VGWAGA3V9NFXZ4H78QHP", Name: "Mesin Sprectometer Ferro", Description: nil},
		&entity.MaterialTestMachine{Id: "01M2Z0VGWAQPPX4ZHD7B9T2S4J", Name: "Mesin Sprectometer Non Ferro", Description: nil},
		&entity.MaterialTestMachine{Id: "01M2Z0VGWB0W8DNH7R24ZK1T06", Name: "Mesin Hardness", Description: nil},
		&entity.MaterialTestMachine{Id: "01M2Z0VGWB8E4CKST2Q9JZYGFG", Name: "Mesin Vertical Machine Center", Description: nil},
		&entity.MaterialTestMachine{Id: "01M2Z0VGWBG4T8VYQ2H1X6D0AS", Name: "Mesin General Purpose Tunning", Description: nil},
		&entity.MaterialTestMachine{Id: "01M2Z0VGWBPXJ6K9Z4TD3CFW2J", Name: "Mesin Electric Discharge Machi", Description: nil},
		&entity.MaterialTestMachine{Id: "01M2Z0VGWBZV5K6JXH03C8FT0G", Name: "Mesin Heat Treatment (HTP) Fur", Description: nil},
		&entity.MaterialTestMachine{Id: "01M2Z0VGWC7X12FH9N6J5BG4DQ", Name: "Mesin Heat Treatment (HTP) Mul", Description: nil},
		&entity.MaterialTestMachine{Id: "01M2Z0VGWCGTRJ3ZFW1BT87K0P", Name: "Mesin Universal Cutter and Too", Description: nil},
		&entity.MaterialTestMachine{Id: "01M2Z0VGWCPZ7VB5DNF6X92JMG", Name: "Mesin Struktur Micro", Description: nil},
		&entity.MaterialTestMachine{Id: "01M2Z0VGWCXFPG3D7TY4KSJ5PN", Name: "Mesin STT (Salt Spray Test)", Description: nil},
		&entity.MaterialTestMachine{Id: "01M2Z0VGWD3WV9S2Q7ZJ4KBD8H", Name: "Mesin CNC Surface Grinding", Description: nil},
		&entity.MaterialTestMachine{Id: "01M2Z0VGWDC6MN47ZF1X0G8RFZ", Name: "Mesin Brenc Drilling", Description: nil},
		&entity.MaterialTestMachine{Id: "01M2Z0VGWDLXW8HD1G7N2S0QAY", Name: "Mesin Drilling dan Milling (Re", Description: nil},
		&entity.MaterialTestMachine{Id: "01M2Z0VGWDW2KDH6Z9QX3JP4BV", Name: "Mesin Kompressor", Description: nil},
		&entity.MaterialTestMachine{Id: "01M2Z0VGWE2MK1T3ZVH8C7QJDN", Name: "Mesin Grinder", Description: nil},
		&entity.MaterialTestMachine{Id: "01M2Z0VGWE9S6W4T1D2FYQZKMB", Name: "Manual Milling", Description: nil},
		&entity.MaterialTestMachine{Id: "01M2Z0VGWEHQFD0X8V5S9KTJ6P", Name: "Manual Bubut", Description: nil},
		&entity.MaterialTestMachine{Id: "01M2Z0VGWEWQXG1Z5C9H27JQGR", Name: "UTM", Description: nil},
		&entity.MaterialTestMachine{Id: "01M2Z0VGWF4HTN8Z12X9MC6RTN", Name: "Las", Description: nil},
		&entity.MaterialTestMachine{Id: "01M2Z0VGWFC8MH5ZGN1T9BX40S", Name: "Gerinda", Description: nil},
	)

	err := s.mtmRepository.SaveMany(mtms)
	if err != nil {
		return err
	}

	return nil
}
