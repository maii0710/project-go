package middlewares

import (
	"belajar_go/utils" 
	"net/http"
	"strings"
)

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(`{"message": "Akses ditolak, token tidak ditemukan"}`))
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		_, err := utils.ValidateToken(tokenString)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(`{"message": "Token tidak valid atau kadaluarsa"}`))
			return
		}

		next.ServeHTTP(w, r)
	}
}