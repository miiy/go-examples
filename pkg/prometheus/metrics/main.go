package main

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

func main() {
	make()
	new()
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8080", nil)
}
