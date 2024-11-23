package middleware

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
)

func genNonce() string {
	var b [20]byte
	if _, err := rand.Read(b[:]); err != nil {
		log.Fatal(err.Error())
	}
	return base64.StdEncoding.EncodeToString(b[:])
}

func SecureHeaders(handler http.Handler) http.HandlerFunc {
	return http.HandlerFunc(
		func(response http.ResponseWriter, request *http.Request) {

			nonce := genNonce()
			nonceString := "default-src 'self'; script-src 'self' 'nonce-%s'; frame-ancestors 'none'"
			nonceHeader := fmt.Sprintf(nonceString, nonce)

			response.Header().Set("X-Content-Type-Options", "nosniff")
			response.Header().Set("Content-Security-Policy", nonceHeader)
			response.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
			response.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
			handler.ServeHTTP(response, request)
		},
	)
}
