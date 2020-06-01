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
	correctHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello from example application."))
	})
	wrongHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
	})

  r := prometheus.NewRegistry()
	r.MustRegister(RequestsTotal)
	http.Handle("/metrics", promhttp.HandlerFor(r, promhttp.HandlerOpts{}))
	http.Handle("/correct", promhttp.InstrumentHandlerCounter(RequestsTotal, correctHandler))
	http.Handle("/wrong", promhttp.InstrumentHandlerCounter(RequestsTotal, wrongHandler))

  log.Fatal(http.ListenAndServe(":2112", nil))
}
