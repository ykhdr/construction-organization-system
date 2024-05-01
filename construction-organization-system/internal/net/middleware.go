package net

import (
	"construction-organization-system/internal/log"
	"net/http"
	"time"
)

type responseWriter struct {
	http.ResponseWriter
	status int
	length int
}

func (rw *responseWriter) WriteHeader(statusCode int) {
	rw.status = statusCode
	rw.ResponseWriter.WriteHeader(statusCode)
}

func (rw *responseWriter) Write(b []byte) (int, error) {
	if rw.status == 0 {
		rw.status = http.StatusOK
	}
	size, err := rw.ResponseWriter.Write(b)
	rw.length += size
	return size, err
}

func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		log.Logger.Infoln("Start processing request: ", r.Method, r.RequestURI)

		wrappedResponseWriter := &responseWriter{ResponseWriter: w}
		next.ServeHTTP(wrappedResponseWriter, r)

		log.Logger.Infoln("Finish processing request: ", r.Method, r.RequestURI, "Status: ", wrappedResponseWriter.status, "Length: ", wrappedResponseWriter.length, "Time: ", time.Since(start))
	})
}
