package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	fmt.Println("Starting server on :5000")
	http.ListenAndServe(":5000", router())
}

func router() chi.Router {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	r.Get("/metrics", testMetricCounter())

	return r
}

func testMetricCounter() func(w http.ResponseWriter, r *http.Request) {
	// Create a custom metric for the test-metrics endpoint
	myCounter := promauto.NewCounter(prometheus.CounterOpts{
		Name: "test_custom_metric",
		Help: "A custom metric to demonstrate pushing to Prometheus",
	})

	// Increment the custom metric on each request
	myCounter.Inc()

	return promhttp.Handler().ServeHTTP
}
