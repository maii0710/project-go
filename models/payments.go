package models

type Payment struct {
	IDPayment       int    `json:"id_payment,omitempty"`
	IDRentals       int    `json:"id_rentals"`
	PaymentDate     string `json:"payment_date"`
	Amount          int    `json:"amount"`
	PaymentMethod   string `json:"payment_method"`
	PaymentProofURL string `json:"payment_proof_url"`
	PaymentStatus   string `json:"payment_status"`
}