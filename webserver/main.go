package main

import (
	"log"
	"net/http"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	RequestsTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "requests_total",
	}, []string{"code", "method"})
)

func main() {
	wrongHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
	})

	teapotHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(418)
	})

	calmHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(420)
	})

  r := prometheus.NewRegistry()
	r.MustRegister(RequestsTotal)
	http.Handle("/metrics", promhttp.HandlerFor(r, promhttp.HandlerOpts{}))
	http.Handle("/teapot", promhttp.InstrumentHandlerCounter(RequestsTotal, teapotHandler))
	http.Handle("/calm", promhttp.InstrumentHandlerCounter(RequestsTotal, calmHandler))
	http.Handle("/wrong", promhttp.InstrumentHandlerCounter(RequestsTotal, wrongHandler))

  log.Fatal(http.ListenAndServe(":2112", nil))
}
