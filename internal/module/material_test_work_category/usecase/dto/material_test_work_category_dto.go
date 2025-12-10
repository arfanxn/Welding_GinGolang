package dto

type SaveMaterialTestWorkCategory struct {
	Id          *string `json:"id"`
	Name        *string `json:"name"`
	Description *string `json:"description"`
}

type DestroyMaterialTestWorkCategory struct {
	Id string `json:"id"`
}
