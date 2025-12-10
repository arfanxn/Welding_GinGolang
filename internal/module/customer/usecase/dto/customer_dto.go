package dto

type SaveCustomer struct {
	Id          *string `json:"id"`
	AddressId   *string `json:"address_id"`
	Name        *string `json:"name"`
	PhoneNumber *string `json:"phone_number"`
	Email       *string `json:"email"`
}

type DestroyCustomer struct {
	Id string `json:"id"`
}
