package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func NewResponseWriter(w http.ResponseWriter, status int) *responseWriter {
	return &responseWriter{
		ResponseWriter: w,
		statusCode:     status,
	}
}

func metricsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		startTime := time.Now()
		rw := NewResponseWriter(w, http.StatusOK)

		reqPath := r.URL.Path

		activeConnections.WithLabelValues(r.Method, reqPath).Inc()
		defer activeConnections.WithLabelValues(r.Method, reqPath).Dec()

		next.ServeHTTP(rw, r)

		status := strconv.Itoa(rw.statusCode)
		duration := time.Since(startTime).Milliseconds()

		httpRequestsTotal.WithLabelValues(r.Method, reqPath, status).Inc()
		requestDuration.WithLabelValues(r.Method, reqPath).Observe(float64(duration))

		fmt.Println("time taken", time.Since(startTime), "path", reqPath)

	})
}
