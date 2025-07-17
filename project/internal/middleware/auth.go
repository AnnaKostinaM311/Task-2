package middleware

import (
	"encoding/json"
	"net/http"
	"strings"

	"project/config"
	"project/pkg/logger"
)

func AuthMiddleware(cfg *config.Config, log logger.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Skip preflight requests
			if r.Method == http.MethodOptions {
				next.ServeHTTP(w, r)
				return
			}

			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				log.Error("Missing authorization header")
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(map[string]string{
					"error":   "Authorization header required",
					"details": "Format: Bearer <token>",
				})
				return
			}

			valid := false
			for _, token := range cfg.Server.AuthTokens {
				if strings.EqualFold(token, authHeader) {
					valid = true
					break
				}
			}

			if !valid {
				log.Error("Invalid auth token")
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(map[string]string{
					"error":   "Invalid token",
					"details": "Check your authorization token",
				})
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
