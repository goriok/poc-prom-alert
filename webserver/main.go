package main

import (
        "net/http"
        "github.com/prometheus/client_golang/prometheus/promhttp"
        "fmt"
)

func main() {
        fmt.Println("running at port: 2112")
        http.Handle("/metrics", promhttp.Handler())
        http.HandleFunc("/logs", handler)
        http.ListenAndServe(":2112", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Println(w, "log handler has been called")
}
