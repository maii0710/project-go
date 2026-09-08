package middlewares

import (
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

var SecretKey = []byte("kunci_rahasia_kamu") // Sesuaikan key JWT kamu

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, `{"message": "Akses ditolak, token tidak ditemukan"}`, http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return SecretKey, nil
		})

		if err != nil || !token.Valid {
			http.Error(w, `{"message": "Token tidak valid atau kadaluarsa"}`, http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	}
}