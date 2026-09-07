package models

type CreateItemInstanceInput struct {
	IDProducts uint   `json:"id_products"`
	AssetCode  string `json:"asset_code"`
	Status     string `json:"status"`
}

type ItemInstance struct {
	ID         uint   `json:"id"`
	IDProducts uint   `json:"id_products"`
	AssetCode  string `json:"asset_code"`
	Status     string `json:"status"`
}
