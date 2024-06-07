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
	// Create a custom metric
	myCounter := promauto.NewCounter(prometheus.CounterOpts{
		Name: "test_golang_metric",
		Help: "A custom metric to demonstrate pushing to Prometheus",
	})

	// Add value
	myCounter.Add(60)

	return promhttp.Handler().ServeHTTP
}
