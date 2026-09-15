package controllers

import (
	"belajar_go/config"
	"belajar_go/models"
	"encoding/json"
	"net/http"
)

func PaymentHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodPost:
		var payment models.Payment

		err := json.NewDecoder(r.Body).Decode(&payment)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{
				"message": "Format JSON tidak valid",
			})
			return
		}

		query := `INSERT INTO payment (id_rentals, payment_date, amount, payment_method, payment_proof_url, payment_status) VALUES (?, ?, ?, ?, ?, ?)`

		_, err = config.DB.Exec(query,
			payment.IDRentals,
			payment.PaymentDate,
			payment.Amount,
			payment.PaymentMethod,
			payment.PaymentProofURL,
			payment.PaymentStatus,
		)

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Pembayaran berhasil diproses dan disimpan ke database!",
			"data":    payment,
		})

	case http.MethodGet:
		query := "SELECT id, id_rentals, payment_date, amount, payment_method, payment_proof_url, payment_status FROM payment"
		rows, err := config.DB.Query(query)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{
				"message": "Gagal mengambil data dari database: " + err.Error(),
			})
			return
		}
		defer rows.Close()

		var payments []models.Payment = []models.Payment{}
		for rows.Next() {
			var p models.Payment
			err := rows.Scan(
				&p.IDPayment,
				&p.IDRentals,
				&p.PaymentDate,
				&p.Amount,
				&p.PaymentMethod,
				&p.PaymentProofURL,
				&p.PaymentStatus,
			)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(map[string]string{
					"message": "Gagal membaca baris data: " + err.Error(),
				})
				return
			}
			payments = append(payments, p)
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(payments)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Method tidak diizinkan",
		})
	}
}
