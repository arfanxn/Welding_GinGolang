package service

import (
	"strconv"

	"github.com/arfanxn/welding/internal/infrastructure/id"
	"github.com/arfanxn/welding/pkg/numberutil"
)

type CodeService interface {
	Generate() string
}

type codeService struct {
	idService id.IdService
}

func NewCodeService(idService id.IdService) CodeService {
	return &codeService{
		idService: idService,
	}
}

func (s *codeService) Generate() string {
	return strconv.Itoa(numberutil.Random(100000, 999999))
}
