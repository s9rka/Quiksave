package middleware

import (
	"context"
	"net/http"

	auth "notas/authentication"

	"github.com/gorilla/mux"
)

func CreateAuthMiddleware(secretKey []byte) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			userID, err := auth.ValidateAndExtractUserID(r, "access_token", secretKey, "access")
			if err != nil {
				http.Error(w, "Unauthorized: "+err.Error(), http.StatusUnauthorized)
				return
			}

			// Attach userID to the request context
			ctx := context.WithValue(r.Context(), "userID", userID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}