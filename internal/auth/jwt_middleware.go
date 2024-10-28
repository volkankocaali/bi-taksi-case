package auth

import (
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"strings"
)

func JWTAuthMiddleware(secret, expectedIssuer string, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing authorization token", http.StatusUnauthorized)
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok ||
			claims["iss"] != expectedIssuer ||
			claims["authenticated"] != true {
			http.Error(w, "Unauthorized access", http.StatusForbidden)
			return
		}

		r.Header.Set("userId", claims["userId"].(string))
		r.Header.Set("username", claims["username"].(string))

		next.ServeHTTP(w, r)
	}
}
