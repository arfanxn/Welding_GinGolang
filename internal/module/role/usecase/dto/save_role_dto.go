package dto

type SaveRole struct {
	Id            string   `json:"id"`
	Name          string   `json:"name"`
	PermissionIds []string `json:"permission_ids"` // permission id
}
