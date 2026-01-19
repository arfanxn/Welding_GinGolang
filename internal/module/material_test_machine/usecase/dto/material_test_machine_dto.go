package dto

type SaveMaterialTestMachine struct {
	Id          *string `json:"id"`
	Name        *string `json:"name"`
	Description *string `json:"description"`
}

type DestroyMaterialTestMachine struct {
	Id string `json:"id"`
}
