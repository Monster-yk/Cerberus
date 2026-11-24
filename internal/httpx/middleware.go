package httpx

import (
	"log"
	"net/http"
	"time"
)

type statusWriter struct {
	http.ResponseWriter
	status int
}

func (w *statusWriter) WriteHeader(status int) {
	w.status = status
	w.ResponseWriter.WriteHeader(status)
}

// RequestLogger creates a middleware that logs each request and its response
func RequestLogger(logger *log.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			sw := &statusWriter{
				ResponseWriter: w,
				status:         http.StatusOK,
			}

			start := time.Now()
			next.ServeHTTP(sw, r)
			duration := time.Since(start)

			logger.Printf(
				"method=%s path=%s status=%d duration=%s",
				r.Method,
				r.URL.Path,
				sw.status,
				duration,
			)

		})
	}
}
