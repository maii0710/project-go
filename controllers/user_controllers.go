package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"belajar_go/config"
	"belajar_go/models"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Ubah 'id' menjadi 'id_user' (sesuai nama kolom di MySQL)
	rows, err := config.DB.Query("SELECT id_user, username, email FROM users")
	if err != nil {
		log.Println("Error SQL Query:", err)
		http.Error(w, `{"message":"Gagal mengambil data user"}`, http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	users := make([]models.User, 0)

	for rows.Next() {
		var u models.User
		if err := rows.Scan(&u.ID, &u.Username, &u.Email); err != nil {
			log.Println("Error Scan Data:", err)
			http.Error(w, `{"message":"Gagal membaca data user"}`, http.StatusInternalServerError)
			return
		}
		users = append(users, u)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, `{"message":"Error saat iterasi data"}`, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(users)
}
