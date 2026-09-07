package routes

import (
	"belajar_go/controllers"
	"net/http"
)

func SetupRoutes() http.Handler {
	mux := http.NewServeMux()

	// Root Endpoint
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		w.Write([]byte("API Belajar Go Berhasil Jalan!"))
	})

	// Auth & User Routes
	mux.HandleFunc("/register", controllers.Register)
	mux.HandleFunc("/login", controllers.Login)
	mux.HandleFunc("/users", controllers.GetUsers)

	// API Endpoints (Tambahkan /api/ agar cocok dengan request Thunder Client)
	mux.HandleFunc("/products", controllers.ProductHandler)
	mux.HandleFunc("/api/category", controllers.CategoryHandler)
	mux.HandleFunc("/category", controllers.CreateCategory)
	mux.HandleFunc("/rental", controllers.CreateRental)
	mux.HandleFunc("/item-instances", controllers.CreateItemInstance)
	mux.HandleFunc("/api/users", controllers.GetUsers)
	return mux
}
