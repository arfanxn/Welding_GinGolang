package id

import "github.com/oklog/ulid/v2"

type IdService interface {
	Generate() string
}

type ulidIdService struct {
}

func NewULIDIdService() IdService {
	return &ulidIdService{}
}

func (s *ulidIdService) Generate() string {
	return ulid.Make().String()
}
