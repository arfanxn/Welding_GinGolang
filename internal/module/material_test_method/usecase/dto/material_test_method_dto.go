package dto

type SaveMaterialTestMethod struct {
	Id          *string `json:"id"`
	Name        *string `json:"name"`
	Description *string `json:"description"`
}

type DestroyMaterialTestMethod struct {
	Id string `json:"id"`
}
