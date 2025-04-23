package middleware

import (
	"log"
	"net/http"
	"time"
)

func RateLimitMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		maxRequests := 100
		resetTime := time.Now().Add(time.Minute)

		requestCount := 0

		if requestCount > maxRequests {
			http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
			return
		}

		log.Printf("Request allowed - Count: %d, Reset Time: %s", requestCount, resetTime)
		next.ServeHTTP(w, r)
	})
}
