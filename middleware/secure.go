package middleware

import "net/http"

func SecureHeaders(callback http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(
		func(response http.ResponseWriter, request *http.Request) {
			response.Header().Set("X-Content-Type-Options", "nosniff")
			response.Header().Set("X-Frame-Options", "DENY")
			response.Header().Set("X-XSS-Protection", "1; mode=block")
			response.Header().Set("Content-Security-Policy", "default-src 'self'")
			callback(response, request)
		},
	)
}
