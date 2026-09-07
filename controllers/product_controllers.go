package controllers

import (
	"belajar_go/models"
	"belajar_go/services"
	"encoding/json"
	"net/http"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {

    products, err := services.GetProducts() 
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)

    json.NewEncoder(w).Encode(products)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method tidak diizinkan", http.StatusMethodNotAllowed)
		return
	}

	var input models.Product
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Format JSON salah", http.StatusBadRequest)
		return
	}

	newProduct, err := services.CreateProduct(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newProduct)
}

func ProductHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		GetProducts(w, r)
	} else if r.Method == http.MethodPost {
		CreateProduct(w, r)
	} else {
		http.Error(w, "Method tidak diizinkan", http.StatusMethodNotAllowed)
	}
}
