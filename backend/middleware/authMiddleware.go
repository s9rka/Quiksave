package middleware

import (
	"net/http"
	"strconv"
	"strings"

	auth "notas/authentication"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "Invalid Authorization header format", http.StatusUnauthorized)
			return
		}
		tokenStr := parts[1]

		claims, err := auth.ValidateJWT(tokenStr)
		if err != nil {
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		userIDStr := claims.UserID
		if userIDStr == "" {
			http.Error(w, "Invalid token payload", http.StatusUnauthorized)
			return
		}

		userID, err := strconv.Atoi(userIDStr)
		if err != nil {
			http.Error(w, "Invalid user ID format", http.StatusInternalServerError)
			return
		}

		ctx := auth.AddUserIDToContext(r.Context(), userID)

		// Pass the request to the next handler with the new context
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
