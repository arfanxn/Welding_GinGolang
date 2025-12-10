package dto

type SaveAddress struct {
	Id          *string `json:"id"`
	FullAddress *string `json:"full_address"`
}

type DestroyAddress struct {
	Id string `json:"id"`
}
