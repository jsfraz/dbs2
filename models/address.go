package models

type Address struct {
	ID       uint   `json:"id" validate:"required"`
	City     string `json:"city" validate:"required"`
	Street   string `json:"street" validate:"required"`
	PostCode int    `json:"postCode" validate:"required"`
	UserID   uint   `json:"userId" validate:"required"`
}

// Vytvoří novou adresu.
//
//	@param city
//	@param street
//	@param postCode
//	@param userId
//	@return *Address
func NewAddress(city string, street string, postCode int, userId uint) *Address {
	return &Address{
		City:     city,
		Street:   street,
		PostCode: postCode,
		UserID:   userId,
	}
}
