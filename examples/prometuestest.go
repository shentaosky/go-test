package main

import (
    "fmt"
    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promauto"
    "github.com/prometheus/client_golang/prometheus/promhttp"
    "net/http"
)

func main() {
    // watch prefix of package
    Tmconntimeout := promauto.NewCounter(prometheus.CounterOpts{Name: "number_tm_connection_timeout", Help: "collect connect to tess master timeout number"})
    Tmconntimeout.Inc()
    http.Handle("/metrics", promhttp.Handler())
        fmt.Println("Starting server 127.0.0.1:9090/metrics")
        if err := http.ListenAndServe(":9090", nil); err != nil && err != http.ErrServerClosed {
            fmt.Println("Failed to server 9090: %v", err)
            return
        }
}