package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	httpRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http.requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "path", "status"},
	)

	activeConnections = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "active_connections",
			Help: "Number of active connections",
		},
		[]string{"method", "path"},
	)

	requestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "HTTP request duration",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "path"},
	)
)

func init() {
	prometheus.MustRegister(httpRequestsTotal)
	prometheus.MustRegister(activeConnections)
	prometheus.MustRegister(requestDuration)
}

func instrumentHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hi there"))
}

func delayHandler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Millisecond * 500)
	w.Write([]byte("finally arrived"))
}

func main() {
	r := chi.NewRouter()
	r.Use(metricsMiddleware)

	r.Get("/", instrumentHandler)
	r.Get("/delay", delayHandler)
	r.Handle("/metrics", promhttp.Handler())


	log.Println("server started")
	log.Fatal(http.ListenAndServe(":8080", r))
}
