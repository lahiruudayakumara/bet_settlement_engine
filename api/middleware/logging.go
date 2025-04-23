package middleware

import (
	"log"
	"net/http"
	"time"
)

// LoggingMiddleware logs the details of incoming requests and their responses
func LoggingMiddleware(http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Log request details
		start := time.Now()
		log.Printf("Started %s %s", r.Method, r.URL.Path)

		// Use a custom ResponseWriter to capture the status code
		loggedResponseWriter := &LoggedResponseWriter{ResponseWriter: w}

		// Call the next handler
		//next.ServeHTTP(loggedResponseWriter, r)

		// Log response details
		duration := time.Since(start)
		log.Printf("Completed %s %s with status %d in %v", r.Method, r.URL.Path, loggedResponseWriter.statusCode, duration)
	})
}

// LoggedResponseWriter wraps the http.ResponseWriter to capture the status code
type LoggedResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

// Override WriteHeader to capture the status code
func (lrw *LoggedResponseWriter) WriteHeader(statusCode int) {
	lrw.statusCode = statusCode
	lrw.ResponseWriter.WriteHeader(statusCode)
}

// Override Write to ensure the status code is set before writing the body
func (lrw *LoggedResponseWriter) Write(p []byte) (n int, err error) {
	if lrw.statusCode == 0 {
		lrw.statusCode = http.StatusOK
	}
	return lrw.ResponseWriter.Write(p)
}
