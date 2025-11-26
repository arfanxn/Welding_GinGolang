package dto

type SaveMaterialTestService struct {
	Id          *string  `json:"id"`
	MachineId   *string  `json:"machine_id"`
	MethodId    *string  `json:"method_id"`
	ServiceType *string  `json:"service_type"`
	ServiceCode *string  `json:"service_code"`
	Unit        *string  `json:"unit"`
	Price       *float64 `json:"price"`
}

type DestroyMaterialTestService struct {
	Id string `json:"id"`
}
