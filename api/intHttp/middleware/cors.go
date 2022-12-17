package middleware

import (
	"net/http"
)

func CORS(trustedOrigins []string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if allowedOrigin(trustedOrigins, r.Header.Get("Origin")) {
				w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
				w.Header().Set("Access-Control-Allow-Credentials", "true")
			}

			if r.Method == "OPTIONS" {
				if reqHeaders := r.Header.Get("Access-Control-Request-Headers"); reqHeaders != "" {
					w.Header().Set("Access-Control-Allow-Headers", reqHeaders)
				}
				w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE")
				w.Header().Set("Access-Control-Max-Age", "600")
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func allowedOrigin(trustedOrigin []string, origin string) bool {
	for _, trusted := range trustedOrigin {
		if trusted == "*" || trusted == origin {
			return true
		}
	}

	return false
}
