package controllers

import (
	"belajar_go/models"
	"encoding/json"
	"net/http"
)

func PaymentHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodPost:
		var payment models.Payment

		// Decode JSON request body dari Thunder Client
		err := json.NewDecoder(r.Body).Decode(&payment)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{
				"message": "Format JSON tidak valid",
			})
			return
		}

		// TODO: Masukkan logika query INSERT ke database di sini

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Pembayaran berhasil diproses!",
			"data":    payment,
		})

	case http.MethodGet:
		// Logika untuk mengambil data payment (jika diperlukan)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Endpoint GET payments",
		})

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Method tidak diizinkan",
		})
	}
}