package middleware

import (
	"net/http"
	"os"
)

// CORSMiddleware sets the appropriate CORS headers
func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		allowedOrigins := map[string]bool{
			os.Getenv("ADMIN_ORIGIN"):  true,
			os.Getenv("ADMIN_ORIGIN2"): true,
			os.Getenv("CLIENT_ORIGIN"): true,
			// Add more origins if you want
		}

		if _, ok := allowedOrigins[r.Header.Get("Origin")]; ok {
			w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PATCH, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
		}

		// If it's a preflight request, respond with 200
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
