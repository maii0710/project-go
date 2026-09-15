package models

type Payment struct {
	IDPayment       int    `gorm:"primaryKey;column:id" json:"id_payment,omitempty"`
	IDRentals       int    `gorm:"column:id_rentals" json:"id_rentals"`
	PaymentDate     string `gorm:"column:payment_date" json:"payment_date"`
	Amount          int    `gorm:"column:amount" json:"amount"`
	PaymentMethod   string `gorm:"column:payment_method" json:"payment_method"`
	PaymentProofURL string `gorm:"column:payment_proof_url" json:"payment_proof_url"`
	PaymentStatus   string `gorm:"column:payment_status" json:"payment_status"`
}

func (Payment) TableName() string {
	return "payment"
}