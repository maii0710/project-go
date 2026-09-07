package controllers

import (
	"belajar_go/models"
	"belajar_go/services"
	"encoding/json"
	"net/http"
)

func CreateItemInstance(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method tidak diizinkan", http.StatusMethodNotAllowed)
		return
	}

	var input models.CreateItemInstanceInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Format JSON salah", http.StatusBadRequest)
		return
	}

	newItemInstance, err := services.CreateItemInstance(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newItemInstance)
}
