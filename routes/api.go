package routes

import (
	"belajar_go/controllers"
	"net/http"
)

func InitAPIRoutes() {
	// Endpoint Auth
	http.HandleFunc("/api/register", controllers.Register)
	http.HandleFunc("/api/login", controllers.Login)

	// Endpoint Users
	http.HandleFunc("/api/users", controllers.GetUsers)

	// Endpoint Products
	http.HandleFunc("/api/products", controllers.GetProducts)
}