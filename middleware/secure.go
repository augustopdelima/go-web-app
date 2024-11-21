package middleware

import "net/http"

func SecureHeaders(handler http.Handler) http.HandlerFunc {
	return http.HandlerFunc(
		func(response http.ResponseWriter, request *http.Request) {
			response.Header().Set("X-Content-Type-Options", "nosniff")
			response.Header().Set("Content-Security-Policy", "default-src 'self'; script-src 'self'; frame-ancestors 'none'")
			response.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
			response.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
			handler.ServeHTTP(response, request)
		},
	)
}
