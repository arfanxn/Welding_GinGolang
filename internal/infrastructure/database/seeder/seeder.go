package seeder

type Seeder interface {
	Seed() error
}
