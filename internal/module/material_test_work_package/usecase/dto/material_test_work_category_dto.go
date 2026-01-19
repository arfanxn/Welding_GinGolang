package dto

type SaveMaterialTestWorkPackage struct {
	Id          *string `json:"id"`
	Name        *string `json:"name"`
	Description *string `json:"description"`
}

type DestroyMaterialTestWorkPackage struct {
	Id string `json:"id"`
}
