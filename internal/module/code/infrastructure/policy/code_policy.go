package policy

import (
	codeRepository "github.com/arfanxn/welding/internal/module/code/domain/repository"
	roleRepository "github.com/arfanxn/welding/internal/module/role/domain/repository"
	"go.uber.org/fx"
)

type CodePolicy interface {
}

type codePolicy struct {
	codeRepository codeRepository.CodeRepository
	roleRepository roleRepository.RoleRepository
}

type NewCodePolicyParams struct {
	fx.In

	CodeRepository codeRepository.CodeRepository
	RoleRepository roleRepository.RoleRepository
}

func NewCodePolicy(params NewCodePolicyParams) CodePolicy {
	return &codePolicy{
		codeRepository: params.CodeRepository,
		roleRepository: params.RoleRepository,
	}
}
