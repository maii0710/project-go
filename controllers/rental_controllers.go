package controllers

import (
	"belajar_go/models"
	"belajar_go/services"
	"encoding/json"
	"net/http"
)

func CreateRental(w http.ResponseWriter, r *http.Request) {
	
	if r.Method != http.MethodPost {
		http.Error(w, "Method tidak diizinkan", http.StatusMethodNotAllowed)
		return
	}


	var input models.CreateRentalInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Format JSON salah", http.StatusBadRequest)
		return
	}

	newRental, err := services.CreateRental(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newRental)
}
