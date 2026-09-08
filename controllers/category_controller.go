package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strings"

	"belajar_go/config"
	"belajar_go/models"
)

func CategoryHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		id := r.URL.Query().Get("id")
		if id != "" {
			GetCategoryByID(w, r, id)
			return
		}
		GetAllCategories(w, r)
	case http.MethodPost:
		CreateCategory(w, r)
	default:
		http.Error(w, `{"message":"Method tidak diizinkan"}`, http.StatusMethodNotAllowed)
	}
}

func CreateCategory(w http.ResponseWriter, r *http.Request) {
	var category models.Category
	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		http.Error(w, `{"message":"Format JSON tidak valid"}`, http.StatusBadRequest)
		return
	}
	if strings.TrimSpace(category.Name) == "" {
		http.Error(w, `{"message":"Nama kategori tidak boleh kosong"}`, http.StatusBadRequest)
		return
	}

	query := "INSERT INTO categories (name) VALUES (?)"
	result, err := config.DB.Exec(query, category.Name)
	if err != nil {
		http.Error(w, `{"message":"Gagal menyimpan data ke database: `+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}

	lastInsertID, _ := result.LastInsertId()
	category.ID = uint(lastInsertID)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(category)
}

func GetAllCategories(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query("SELECT id, name FROM categories")
	if err != nil {
		http.Error(w, `{"message":"Gagal mengambil data kategori: `+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	categories := make([]models.Category, 0)

	for rows.Next() {
		var cat models.Category
		if err := rows.Scan(&cat.ID, &cat.Name); err != nil {
			http.Error(w, `{"message":"Gagal membaca data"}`, http.StatusInternalServerError)
			return
		}
		categories = append(categories, cat)
	}
	if err := rows.Err(); err != nil {
		http.Error(w, `{"message":"Error saat iterasi data: `+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(categories)
}

func GetCategoryByID(w http.ResponseWriter, r *http.Request, id string) {
	var category models.Category

	query := "SELECT id, name FROM categories WHERE id = ?"
	err := config.DB.QueryRow(query, id).Scan(&category.ID, &category.Name)

	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, `{"message":"Kategori tidak ditemukan"}`, http.StatusNotFound)
			return
		}
		http.Error(w, `{"message":"Gagal mengambil data kategori: `+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(category)
}
