package models

type CreateRentalInput struct {
	IDUser          uint    `json:"id_user"`
	IDUnit          uint    `json:"id_unit"`
	DeliveryAddress string  `json:"delivery_address"`
	StartDate       string  `json:"start_date"`
	ReturnDate      string  `json:"return_date"`
	RentalFee       float64 `json:"rental_fee"`
}

type Rental struct {
	ID               uint    `json:"id"`
	IDUser           uint    `json:"id_user"`
	IDUnit           uint    `json:"id_unit"`
	DeliveryAddress  string  `json:"delivery_address"`
	StartDate        string  `json:"start_date"`
	ReturnDate       string  `json:"return_date"`
	ActualReturnDate *string `json:"actual_return_date"`
	RentalFee        float64 `json:"rental_fee"`
	LateFee          float64 `json:"late_fee"`
	RentalStatus     string  `json:"rental_status"`
}
