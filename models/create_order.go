package models

type CreateOrder struct {
	AddressId  uint  `json:"addressId" validate:"required"`
	DiscountId *uint `json:"discountId"`
}
