package main

import (
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	requestsTotal = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
	)

	responseTime = prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Name:    "http_response_time_seconds",
			Help:    "Response time in seconds",
			Buckets: prometheus.DefBuckets,
		},
	)
)

func init() {
	prometheus.MustRegister(requestsTotal)
	prometheus.MustRegister(responseTime)
}

func handler(w http.ResponseWriter, r *http.Request) {
	start := time.Now() // Bắt đầu đo thời gian

	requestsTotal.Inc()

	w.Write([]byte("Hello, Prometheus!"))

	duration := time.Since(start).Seconds() // Tính thời gian xử lý
	responseTime.Observe(duration)          // Ghi nhận vào Histogram
}

func main() {
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/", handler)

	http.ListenAndServe(":8080", nil)
}
