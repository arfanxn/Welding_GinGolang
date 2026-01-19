package seeder

import (
	"github.com/arfanxn/welding/internal/infrastructure/id"
	addressRepository "github.com/arfanxn/welding/internal/module/address/domain/repository"
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	"github.com/bluele/factory-go/factory"
	"go.uber.org/fx"
)

var _ Seeder = (*AddressSeeder)(nil)

type AddressSeeder struct {
	idService         id.IdService
	addressFactory    *factory.Factory
	addressRepository addressRepository.AddressRepository
}

type NewAddressSeederParams struct {
	fx.In

	IdService         id.IdService
	AddressFactory    *factory.Factory `name:"address_factory"`
	AddressRepository addressRepository.AddressRepository
}

func NewAddressSeeder(params NewAddressSeederParams) Seeder {
	return &AddressSeeder{
		idService:         params.IdService,
		addressFactory:    params.AddressFactory,
		addressRepository: params.AddressRepository,
	}
}

func (s *AddressSeeder) Seed() (err error) {
	addressFactory := s.addressFactory

	addresses := []*entity.Address{
		// 1
		addressFactory.MustCreateWithOption(map[string]any{
			"Id": "01J8VADDR0K7WCF4T2R9MQ6HAE",
		}).(*entity.Address),
		// 2
		addressFactory.MustCreateWithOption(map[string]any{
			"Id": "01J8VADDR1X9PN8K4SH2E7F0QB",
		}).(*entity.Address),
		// 3
		addressFactory.MustCreateWithOption(map[string]any{
			"Id": "01J8VADDR2M4TDQ9G1W8K3R5YC",
		}).(*entity.Address),
		// 4
		addressFactory.MustCreateWithOption(map[string]any{
			"Id": "01J8VADDR3B7FQP2N0C6H9J4LM",
		}).(*entity.Address),
		// 5
		addressFactory.MustCreateWithOption(map[string]any{
			"Id": "01J8VADDR4Z2WKT6R5D1S8N3PV",
		}).(*entity.Address),
		// 6
		addressFactory.MustCreateWithOption(map[string]any{
			"Id": "01J8VADDR5H9MCG1Q8T4K2F7XS",
		}).(*entity.Address),
		// 7
		addressFactory.MustCreateWithOption(map[string]any{
			"Id": "01J8VADDR6T3RNB5M2P7Q0W9JD",
		}).(*entity.Address),
		// 8
		addressFactory.MustCreateWithOption(map[string]any{
			"Id": "01J8VADDR7Q8XHF0L9C3B6M5RW",
		}).(*entity.Address),
		// 9
		addressFactory.MustCreateWithOption(map[string]any{
			"Id": "01J8VADDR8D6PKM3T1W4S9V2JQ",
		}).(*entity.Address),
		// 10
		addressFactory.MustCreateWithOption(map[string]any{
			"Id": "01J8VADDR9W1SBL8H5M0Q2C7TG",
		}).(*entity.Address),
		// 11
		addressFactory.MustCreateWithOption(map[string]any{
			"Id": "01J8WADDR0P4NQK7B6X2H9T1MF",
		}).(*entity.Address),
		// 12
		addressFactory.MustCreateWithOption(map[string]any{
			"Id": "01J8WADDR1L9VFC3P0T7S5G2KB",
		}).(*entity.Address),
		// 13
		addressFactory.MustCreateWithOption(map[string]any{
			"Id": "01J8WADDR2C6HMB1W8K3N4D9QP",
		}).(*entity.Address),
		// 14
		addressFactory.MustCreateWithOption(map[string]any{
			"Id": "01J8WADDR3X0TRF6N4S9W1G2CQ",
		}).(*entity.Address),
		// 15
		addressFactory.MustCreateWithOption(map[string]any{
			"Id": "01J8WADDR4M7PQB2C1H5K8J9TR",
		}).(*entity.Address),
		// 16
		addressFactory.MustCreateWithOption(map[string]any{
			"Id": "01J8WADDR5B1GCW9T3M2F4Q8KS",
		}).(*entity.Address),
		// 17
		addressFactory.MustCreateWithOption(map[string]any{
			"Id": "01J8WADDR6Z8NSH0P4W7R3T2LD",
		}).(*entity.Address),
		// 18
		addressFactory.MustCreateWithOption(map[string]any{
			"Id": "01J8WADDR7H5DKQ2C9L1W8S4NP",
		}).(*entity.Address),
		// 19
		addressFactory.MustCreateWithOption(map[string]any{
			"Id": "01J8WADDR8K3WXM7R2H6C1P5VQ",
		}).(*entity.Address),
		// 20
		addressFactory.MustCreateWithOption(map[string]any{
			"Id": "01J8WADDR9T0QFG8B5N4M2S7JK",
		}).(*entity.Address),
		// 21
		addressFactory.MustCreateWithOption(map[string]any{
			"Id": "01J8XADDR0C7BHS4N2K8W5M1QF",
		}).(*entity.Address),
		// 22
		addressFactory.MustCreateWithOption(map[string]any{
			"Id": "01J8XADDR1R9QKP3D6F2T7S0LM",
		}).(*entity.Address),
		// 23
		addressFactory.MustCreateWithOption(map[string]any{
			"Id": "01J8XADDR2W3MTL5H0C9S7B8QJ",
		}).(*entity.Address),
		// 24
		addressFactory.MustCreateWithOption(map[string]any{
			"Id": "01J8XADDR3N1FVC8P4W2Q6K9SD",
		}).(*entity.Address),
		// 25
		addressFactory.MustCreateWithOption(map[string]any{
			"Id": "01J8XADDR4H8LSB0T6M3D7W1QP",
		}).(*entity.Address),
		// 26
		addressFactory.MustCreateWithOption(map[string]any{
			"Id": "01J8XADDR5Q6PGM2C9X1B4T8HS",
		}).(*entity.Address),
		// 27
		addressFactory.MustCreateWithOption(map[string]any{
			"Id": "01J8XADDR6D2XKQ7N5H8W0M3RL",
		}).(*entity.Address),
		// 28
		addressFactory.MustCreateWithOption(map[string]any{
			"Id": "01J8XADDR7M4WCB9S3P1T6F0HQ",
		}).(*entity.Address),
		// 29
		addressFactory.MustCreateWithOption(map[string]any{
			"Id": "01J8XADDR8F0TQD1K5N9W3S7LB",
		}).(*entity.Address),
		// 30
		addressFactory.MustCreateWithOption(map[string]any{
			"Id": "01J8XADDR9S5NWM8L2C7Q0H4TP",
		}).(*entity.Address),
		// 31
		addressFactory.MustCreateWithOption(map[string]any{
			"Id": "01J8YADDR0W9FBM1C6P3T8Q2KS",
		}).(*entity.Address),
		// 32
		addressFactory.MustCreateWithOption(map[string]any{
			"Id": "01J8YADDR1T3QXN9R4S0B6P7JM",
		}).(*entity.Address),
		// 33
		addressFactory.MustCreateWithOption(map[string]any{
			"Id": "01J8YADDR2H7MCQ5W1K8T4N9SB",
		}).(*entity.Address),
		// 34
		addressFactory.MustCreateWithOption(map[string]any{
			"Id": "01J8YADDR3K1PGT8B6R2W7C0QF",
		}).(*entity.Address),
		// 35
		addressFactory.MustCreateWithOption(map[string]any{
			"Id": "01J8YADDR4C8XNS6D3M1F9Q5WP",
		}).(*entity.Address),
		// 36
		addressFactory.MustCreateWithOption(map[string]any{
			"Id": "01J8YADDR5M2TLQ9P7C4H1S8KB",
		}).(*entity.Address),
		// 37
		addressFactory.MustCreateWithOption(map[string]any{
			"Id": "01J8YADDR6D4RPW1N8T0Q7M3SF",
		}).(*entity.Address),
		// 38
		addressFactory.MustCreateWithOption(map[string]any{
			"Id": "01J8YADDR7B6WKV2S9H3C0P5TL",
		}).(*entity.Address),
		// 39
		addressFactory.MustCreateWithOption(map[string]any{
			"Id": "01J8YADDR8Q0HMC7L5W2F9T3NS",
		}).(*entity.Address),
		// 40
		addressFactory.MustCreateWithOption(map[string]any{
			"Id": "01J8YADDR9S3TNB4P0C8W6M1FQ",
		}).(*entity.Address),
		// 41
		addressFactory.MustCreateWithOption(map[string]any{
			"Id": "01J8ZADDR0M8QRC5N7W1H3T9KB",
		}).(*entity.Address),
		// 42
		addressFactory.MustCreateWithOption(map[string]any{
			"Id": "01J8ZADDR1B2WTS9L4M7C0H8QF",
		}).(*entity.Address),
		// 43
		addressFactory.MustCreateWithOption(map[string]any{
			"Id": "01J8ZADDR2W6HMC3T1P8Q9S4LD",
		}).(*entity.Address),
		// 44
		addressFactory.MustCreateWithOption(map[string]any{
			"Id": "01J8ZADDR3C9KPW7F2N5B1M6TS",
		}).(*entity.Address),
		// 45
		addressFactory.MustCreateWithOption(map[string]any{
			"Id": "01J8ZADDR4T1QBF8H6R3W9C2KP",
		}).(*entity.Address),
		// 46
		addressFactory.MustCreateWithOption(map[string]any{
			"Id": "01J8ZADDR5H7WRM0S4P2T8Q1LC",
		}).(*entity.Address),
		// 47
		addressFactory.MustCreateWithOption(map[string]any{
			"Id": "01J8ZADDR6P3KSW5B9T1C7H0NF",
		}).(*entity.Address),
		// 48
		addressFactory.MustCreateWithOption(map[string]any{
			"Id": "01J8ZADDR7X5MTC1Q9W4S2H8PB",
		}).(*entity.Address),
		// 49
		addressFactory.MustCreateWithOption(map[string]any{
			"Id": "01J8ZADDR8L0QBN6S3T7M1W9HF",
		}).(*entity.Address),
		// 50
		addressFactory.MustCreateWithOption(map[string]any{
			"Id": "01J8ZADDR9F4TMH2C8W1Q5S7KP",
		}).(*entity.Address),
	}

	// Save all roles to the database
	err = s.addressRepository.SaveMany(addresses)
	if err != nil {
		return err
	}

	return nil
}
