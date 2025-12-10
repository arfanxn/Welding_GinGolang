package seeder

import (
	"github.com/arfanxn/welding/internal/infrastructure/id"
	addressRepository "github.com/arfanxn/welding/internal/module/address/domain/repository"
	customerRepository "github.com/arfanxn/welding/internal/module/customer/domain/repository"
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	"github.com/bluele/factory-go/factory"
	"github.com/brianvoe/gofakeit/v7"
	"go.uber.org/fx"
)

var _ Seeder = (*CustomerSeeder)(nil)

type CustomerSeeder struct {
	idService          id.IdService
	customerFactory    *factory.Factory
	customerRepository customerRepository.CustomerRepository
	addressRepository  addressRepository.AddressRepository
}

type NewCustomerSeederParams struct {
	fx.In

	IdService          id.IdService
	CustomerFactory    *factory.Factory `name:"customer_factory"`
	CustomerRepository customerRepository.CustomerRepository
	AddressRepository  addressRepository.AddressRepository
}

func NewCustomerSeeder(params NewCustomerSeederParams) Seeder {
	return &CustomerSeeder{
		idService:          params.IdService,
		customerFactory:    params.CustomerFactory,
		customerRepository: params.CustomerRepository,
		addressRepository:  params.AddressRepository,
	}
}

func (s *CustomerSeeder) Seed() error {
	customerFactory := s.customerFactory

	addresses, err := s.addressRepository.Get(nil)
	if err != nil {
		return err
	}

	customers := []*entity.Customer{
		// 1
		customerFactory.MustCreateWithOption(map[string]any{
			"Id":        "01JACSTMR0K7WCF4T2R9MQ6HAE",
			"AddressId": addresses[gofakeit.Number(0, len(addresses)-1)].Id,
		}).(*entity.Customer),
		// 2
		customerFactory.MustCreateWithOption(map[string]any{
			"Id":        "01JACSTMR1X9PN8K4SH2E7F0QB",
			"AddressId": addresses[gofakeit.Number(0, len(addresses)-1)].Id,
		}).(*entity.Customer),
		// 3
		customerFactory.MustCreateWithOption(map[string]any{
			"Id":        "01JACSTMR2M4TDQ9G1W8K3R5YC",
			"AddressId": addresses[gofakeit.Number(0, len(addresses)-1)].Id,
		}).(*entity.Customer),
		// 4
		customerFactory.MustCreateWithOption(map[string]any{
			"Id":        "01JACSTMR3B7FQP2N0C6H9J4LM",
			"AddressId": addresses[gofakeit.Number(0, len(addresses)-1)].Id,
		}).(*entity.Customer),
		// 5
		customerFactory.MustCreateWithOption(map[string]any{
			"Id":        "01JACSTMR4Z2WKT6R5D1S8N3PV",
			"AddressId": addresses[gofakeit.Number(0, len(addresses)-1)].Id,
		}).(*entity.Customer),
		// 6
		customerFactory.MustCreateWithOption(map[string]any{
			"Id":        "01JACSTMR5H9MCG1Q8T4K2F7XS",
			"AddressId": addresses[gofakeit.Number(0, len(addresses)-1)].Id,
		}).(*entity.Customer),
		// 7
		customerFactory.MustCreateWithOption(map[string]any{
			"Id":        "01JACSTMR6T3RNB5M2P7Q0W9JD",
			"AddressId": addresses[gofakeit.Number(0, len(addresses)-1)].Id,
		}).(*entity.Customer),
		// 8
		customerFactory.MustCreateWithOption(map[string]any{
			"Id":        "01JACSTMR7Q8XHF0L9C3B6M5RW",
			"AddressId": addresses[gofakeit.Number(0, len(addresses)-1)].Id,
		}).(*entity.Customer),
		// 9
		customerFactory.MustCreateWithOption(map[string]any{
			"Id":        "01JACSTMR8D6PKM3T1W4S9V2JQ",
			"AddressId": addresses[gofakeit.Number(0, len(addresses)-1)].Id,
		}).(*entity.Customer),
		// 10
		customerFactory.MustCreateWithOption(map[string]any{
			"Id":        "01JACSTMR9W1SBL8H5M0Q2C7TG",
			"AddressId": addresses[gofakeit.Number(0, len(addresses)-1)].Id,
		}).(*entity.Customer),
		// 11
		customerFactory.MustCreateWithOption(map[string]any{
			"Id":        "01JBCSTMR0P4NQK7B6X2H9T1MF",
			"AddressId": addresses[gofakeit.Number(0, len(addresses)-1)].Id,
		}).(*entity.Customer),
		// 12
		customerFactory.MustCreateWithOption(map[string]any{
			"Id":        "01JBCSTMR1L9VFC3P0T7S5G2KB",
			"AddressId": addresses[gofakeit.Number(0, len(addresses)-1)].Id,
		}).(*entity.Customer),
		// 13
		customerFactory.MustCreateWithOption(map[string]any{
			"Id":        "01JBCSTMR2C6HMB1W8K3N4D9QP",
			"AddressId": addresses[gofakeit.Number(0, len(addresses)-1)].Id,
		}).(*entity.Customer),
		// 14
		customerFactory.MustCreateWithOption(map[string]any{
			"Id":        "01JBCSTMR3X0TRF6N4S9W1G2CQ",
			"AddressId": addresses[gofakeit.Number(0, len(addresses)-1)].Id,
		}).(*entity.Customer),
		// 15
		customerFactory.MustCreateWithOption(map[string]any{
			"Id":        "01JBCSTMR4M7PQB2C1H5K8J9TR",
			"AddressId": addresses[gofakeit.Number(0, len(addresses)-1)].Id,
		}).(*entity.Customer),
		// 16
		customerFactory.MustCreateWithOption(map[string]any{
			"Id":        "01JBCSTMR5B1GCW9T3M2F4Q8KS",
			"AddressId": addresses[gofakeit.Number(0, len(addresses)-1)].Id,
		}).(*entity.Customer),
		// 17
		customerFactory.MustCreateWithOption(map[string]any{
			"Id":        "01JBCSTMR6Z8NSH0P4W7R3T2LD",
			"AddressId": addresses[gofakeit.Number(0, len(addresses)-1)].Id,
		}).(*entity.Customer),
		// 18
		customerFactory.MustCreateWithOption(map[string]any{
			"Id":        "01JBCSTMR7H5DKQ2C9L1W8S4NP",
			"AddressId": addresses[gofakeit.Number(0, len(addresses)-1)].Id,
		}).(*entity.Customer),
		// 19
		customerFactory.MustCreateWithOption(map[string]any{
			"Id":        "01JBCSTMR8K3WXM7R2H6C1P5VQ",
			"AddressId": addresses[gofakeit.Number(0, len(addresses)-1)].Id,
		}).(*entity.Customer),
		// 20
		customerFactory.MustCreateWithOption(map[string]any{
			"Id":        "01JBCSTMR9T0QFG8B5N4M2S7JK",
			"AddressId": addresses[gofakeit.Number(0, len(addresses)-1)].Id,
		}).(*entity.Customer),
		// 21
		customerFactory.MustCreateWithOption(map[string]any{
			"Id":        "01JCCSTMR0C7BHS4N2K8W5M1QF",
			"AddressId": addresses[gofakeit.Number(0, len(addresses)-1)].Id,
		}).(*entity.Customer),
		// 22
		customerFactory.MustCreateWithOption(map[string]any{
			"Id":        "01JCCSTMR1R9QKP3D6F2T7S0LM",
			"AddressId": addresses[gofakeit.Number(0, len(addresses)-1)].Id,
		}).(*entity.Customer),
		// 23
		customerFactory.MustCreateWithOption(map[string]any{
			"Id":        "01JCCSTMR2W3MTL5H0C9S7B8QJ",
			"AddressId": addresses[gofakeit.Number(0, len(addresses)-1)].Id,
		}).(*entity.Customer),
		// 24
		customerFactory.MustCreateWithOption(map[string]any{
			"Id":        "01JCCSTMR3N1FVC8P4W2Q6K9SD",
			"AddressId": addresses[gofakeit.Number(0, len(addresses)-1)].Id,
		}).(*entity.Customer),
		// 25
		customerFactory.MustCreateWithOption(map[string]any{
			"Id":        "01JCCSTMR4H8LSB0T6M3D7W1QP",
			"AddressId": addresses[gofakeit.Number(0, len(addresses)-1)].Id,
		}).(*entity.Customer),
		// 26
		customerFactory.MustCreateWithOption(map[string]any{
			"Id":        "01JCCSTMR5Q6PGM2C9X1B4T8HS",
			"AddressId": addresses[gofakeit.Number(0, len(addresses)-1)].Id,
		}).(*entity.Customer),
		// 27
		customerFactory.MustCreateWithOption(map[string]any{
			"Id":        "01JCCSTMR6D2XKQ7N5H8W0M3RL",
			"AddressId": addresses[gofakeit.Number(0, len(addresses)-1)].Id,
		}).(*entity.Customer),
		// 28
		customerFactory.MustCreateWithOption(map[string]any{
			"Id":        "01JCCSTMR7M4WCB9S3P1T6F0HQ",
			"AddressId": addresses[gofakeit.Number(0, len(addresses)-1)].Id,
		}).(*entity.Customer),
		// 29
		customerFactory.MustCreateWithOption(map[string]any{
			"Id":        "01JCCSTMR8F0TQD1K5N9W3S7LB",
			"AddressId": addresses[gofakeit.Number(0, len(addresses)-1)].Id,
		}).(*entity.Customer),
		// 30
		customerFactory.MustCreateWithOption(map[string]any{
			"Id":        "01JCCSTMR9S5NWM8L2C7Q0H4TP",
			"AddressId": addresses[gofakeit.Number(0, len(addresses)-1)].Id,
		}).(*entity.Customer),
		// 31
		customerFactory.MustCreateWithOption(map[string]any{
			"Id":        "01JDCSTMR0W9FBM1C6P3T8Q2KS",
			"AddressId": addresses[gofakeit.Number(0, len(addresses)-1)].Id,
		}).(*entity.Customer),
		// 32
		customerFactory.MustCreateWithOption(map[string]any{
			"Id":        "01JDCSTMR1T3QXN9R4S0B6P7JM",
			"AddressId": addresses[gofakeit.Number(0, len(addresses)-1)].Id,
		}).(*entity.Customer),
		// 33
		customerFactory.MustCreateWithOption(map[string]any{
			"Id":        "01JDCSTMR2H7MCQ5W1K8T4N9SB",
			"AddressId": addresses[gofakeit.Number(0, len(addresses)-1)].Id,
		}).(*entity.Customer),
		// 34
		customerFactory.MustCreateWithOption(map[string]any{
			"Id":        "01JDCSTMR3K1PGT8B6R2W7C0QF",
			"AddressId": addresses[gofakeit.Number(0, len(addresses)-1)].Id,
		}).(*entity.Customer),
		// 35
		customerFactory.MustCreateWithOption(map[string]any{
			"Id":        "01JDCSTMR4C8XNS6D3M1F9Q5WP",
			"AddressId": addresses[gofakeit.Number(0, len(addresses)-1)].Id,
		}).(*entity.Customer),
		// 36
		customerFactory.MustCreateWithOption(map[string]any{
			"Id":        "01JDCSTMR5M2TLQ9P7C4H1S8KB",
			"AddressId": addresses[gofakeit.Number(0, len(addresses)-1)].Id,
		}).(*entity.Customer),
		// 37
		customerFactory.MustCreateWithOption(map[string]any{
			"Id":        "01JDCSTMR6D4RPW1N8T0Q7M3SF",
			"AddressId": addresses[gofakeit.Number(0, len(addresses)-1)].Id,
		}).(*entity.Customer),
		// 38
		customerFactory.MustCreateWithOption(map[string]any{
			"Id":        "01JDCSTMR7B6WKV2S9H3C0P5TL",
			"AddressId": addresses[gofakeit.Number(0, len(addresses)-1)].Id,
		}).(*entity.Customer),
		// 39
		customerFactory.MustCreateWithOption(map[string]any{
			"Id":        "01JDCSTMR8Q0HMC7L5W2F9T3NS",
			"AddressId": addresses[gofakeit.Number(0, len(addresses)-1)].Id,
		}).(*entity.Customer),
		// 40
		customerFactory.MustCreateWithOption(map[string]any{
			"Id":        "01JDCSTMR9S3TNB4P0C8W6M1FQ",
			"AddressId": addresses[gofakeit.Number(0, len(addresses)-1)].Id,
		}).(*entity.Customer),
		// 41
		customerFactory.MustCreateWithOption(map[string]any{
			"Id":        "01JECSTMR0M8QRC5N7W1H3T9KB",
			"AddressId": addresses[gofakeit.Number(0, len(addresses)-1)].Id,
		}).(*entity.Customer),
		// 42
		customerFactory.MustCreateWithOption(map[string]any{
			"Id":        "01JECSTMR1B2WTS9L4M7C0H8QF",
			"AddressId": addresses[gofakeit.Number(0, len(addresses)-1)].Id,
		}).(*entity.Customer),
		// 43
		customerFactory.MustCreateWithOption(map[string]any{
			"Id":        "01JECSTMR2W6HMC3T1P8Q9S4LD",
			"AddressId": addresses[gofakeit.Number(0, len(addresses)-1)].Id,
		}).(*entity.Customer),
		// 44
		customerFactory.MustCreateWithOption(map[string]any{
			"Id":        "01JECSTMR3C9KPW7F2N5B1M6TS",
			"AddressId": addresses[gofakeit.Number(0, len(addresses)-1)].Id,
		}).(*entity.Customer),
		// 45
		customerFactory.MustCreateWithOption(map[string]any{
			"Id":        "01JECSTMR4T1QBF8H6R3W9C2KP",
			"AddressId": addresses[gofakeit.Number(0, len(addresses)-1)].Id,
		}).(*entity.Customer),
		// 46
		customerFactory.MustCreateWithOption(map[string]any{
			"Id":        "01JECSTMR5H7WRM0S4P2T8Q1LC",
			"AddressId": addresses[gofakeit.Number(0, len(addresses)-1)].Id,
		}).(*entity.Customer),
		// 47
		customerFactory.MustCreateWithOption(map[string]any{
			"Id":        "01JECSTMR6P3KSW5B9T1C7H0NF",
			"AddressId": addresses[gofakeit.Number(0, len(addresses)-1)].Id,
		}).(*entity.Customer),
		// 48
		customerFactory.MustCreateWithOption(map[string]any{
			"Id":        "01JECSTMR7X5MTC1Q9W4S2H8PB",
			"AddressId": addresses[gofakeit.Number(0, len(addresses)-1)].Id,
		}).(*entity.Customer),
		// 49
		customerFactory.MustCreateWithOption(map[string]any{
			"Id":        "01JECSTMR8L0QBN6S3T7M1W9HF",
			"AddressId": addresses[gofakeit.Number(0, len(addresses)-1)].Id,
		}).(*entity.Customer),
		// 50
		customerFactory.MustCreateWithOption(map[string]any{
			"Id":        "01JECSTMR9F4TMH2C8W1Q5S7KP",
			"AddressId": addresses[gofakeit.Number(0, len(addresses)-1)].Id,
		}).(*entity.Customer),
	}

	// Save all roles to the database
	err = s.customerRepository.SaveMany(customers)
	if err != nil {
		return err
	}

	return nil
}
