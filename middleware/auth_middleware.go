package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/Siddheshk02/go-blog-platform/utils"
)

// JWTAuthMiddleware checks the JWT token validity
func JWTAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing Authorization Header", http.StatusUnauthorized)
			return
		}

		tokenStr := strings.Split(authHeader, "Bearer ")[1]
		claims, err := utils.ValidateJWT(tokenStr)
		if err != nil {
			http.Error(w, "Invalid JWT", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "userID", claims.UserID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
