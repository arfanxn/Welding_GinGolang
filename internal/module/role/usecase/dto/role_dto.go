package dto

import "github.com/arfanxn/welding/internal/module/role/domain/enum"

type SaveRole struct {
	Id            string        `json:"id"`
	Name          enum.RoleName `json:"name"`
	PermissionIds []string      `json:"permission_ids"` // permission id
}

type SetDefaultRole struct {
	Id string `json:"id"`
}

type DestroyRole struct {
	Id string `json:"id"`
}
