package models

type Product struct {
	ID          uint    `json:"id"`
	ProductName string  `json:"product_name"`
	PhotoURL    string  `json:"photo_url"`
	Deskripsi   string  `json:"deskripsi"`
	Spec        string  `json:"spec"`
	PricingPDay float64 `json:"pricing_p_day"`
	IDCategory  uint    `json:"id_category"`
}
