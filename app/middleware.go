package main

import (
	"net/http"
	"time"
)

// LogResponseWriter allows us to capture the response status code.
// .
type LogResponseWriter struct {
	http.ResponseWriter
	status int
}

func (w *LogResponseWriter) WriteHeader(status int) {
	w.status = status
	w.ResponseWriter.WriteHeader(status)
}

// Logger is a middleware which logs the http request and response status.
// .
func (s *Server) Logger(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		ww := &LogResponseWriter{
			ResponseWriter: w,
			status:         http.StatusOK,
		}

		// Defer the logging call.
		defer func(start time.Time) {

			s.infoLog.Printf(
				"%s %d %s %s %s",
				"[request]",
				ww.status,
				r.Method,
				r.URL.RequestURI(),
				time.Since(start),
			)

		}(time.Now())

		// Call the next handler
		next.ServeHTTP(ww, r)
	}
}
