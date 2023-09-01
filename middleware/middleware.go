package middleware

import (
	"fmt"
	"net/http"
	"time"
)

// LoggingMiddleware logs information about incoming requests.
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Pass the request to the next handler in the chain.
		next.ServeHTTP(w, r)

		// Calculate the elapsed time for the request.
		elapsed := time.Since(start)

		// Log the request details.
		fmt.Printf(
			"Request: %s %s %s %s\n",
			r.Method,
			r.RequestURI,
			r.RemoteAddr,
			elapsed,
		)
	})
}

// AuthenticationMiddleware is responsible for authenticating requests using JWT tokens.
func AuthenticationMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO: All the steps below were omited for simplicity
		// Get the Authorization header from the request
		// Check if the token is missing or doesn't start with "Bearer "
		// Extract the token
		// Validate the token
		// If the token is valid, proceed to the next handler
		next.ServeHTTP(w, r)
	})
}
