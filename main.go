package main

import (
	"fmt"
	"net/http"

	"belajar_go/config"
	"belajar_go/routes"
)

func main() {
	config.ConnectDB()

	router := routes.SetupRoutes()

routerWithCORS := enableCORS(router)

fmt.Println("Server berjalan di http://localhost:8080")

err := http.ListenAndServe(":8080", routerWithCORS)

if err != nil {
    panic(err)
}
}

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
