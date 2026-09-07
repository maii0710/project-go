package services

import (
	"belajar_go/config"
	"belajar_go/models"
	"errors"
)

func CreateProduct(product models.Product) (models.Product, error) {
	if product.ProductName == "" {
		return models.Product{}, errors.New("nama produk tidak boleh kosong")
	}

	query := `INSERT INTO products (product_name, photo_url, deskripsi, spec, pricing_p_day, id_category) 
	          VALUES (?, ?, ?, ?, ?, ?)`

	result, err := config.DB.Exec(query,
		product.ProductName,
		product.PhotoURL,
		product.Deskripsi,
		product.Spec,
		product.PricingPDay,
		product.IDCategory,
	)
	if err != nil {
		return models.Product{}, err
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return models.Product{}, err
	}

	product.ID = uint(lastID)
	return product, nil
}

func GetProducts() ([]models.Product, error) {
	
	products := []models.Product{
		{
			ID:          1,
			ProductName: "Produk A", 
			PhotoURL:    "https://example.com/photo.jpg",
			Deskripsi:   "Deskripsi produk A",
			Spec:        "Spesifikasi A",
			PricingPDay: 50000,
			IDCategory:  1,
		},
		{
			ID:          2,
			ProductName: "Produk B", 
			PhotoURL:    "https://example.com/photo2.jpg",
			Deskripsi:   "Deskripsi produk B",
			Spec:        "Spesifikasi B",
			PricingPDay: 75000,
			IDCategory:  1,
		},
	}

	return products, nil
}