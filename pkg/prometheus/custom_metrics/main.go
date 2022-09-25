package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"math"
	"math/rand"
	"net/http"
	"time"
)

var (
	TestCounter = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "test_counter",
			Help: "counter test",
		},
	)

	TestGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "test_gauge",
		Help: "gauge test",
	})

	TestHistogram = prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Name:    "test_histogram",
			Help:    "histogram test",
			Buckets: prometheus.LinearBuckets(20, 5, 5),
		},
	)

	TestSummary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "test_summary",
		Help: "summary test",
		Objectives: map[float64]float64{
			0.5:  0.05,
			0.9:  0.01,
			0.99: 0.001,
		},
	})
)

func main() {

	// 注册指标
	prometheus.MustRegister(TestCounter)
	prometheus.MustRegister(TestGauge)
	prometheus.MustRegister(TestHistogram)
	prometheus.MustRegister(TestSummary)

	// counter
	go func() {
		for {
			//TestCounter.Add(1)
			TestCounter.Inc()
			time.Sleep(time.Second * 2)
		}
	}()

	// gauge
	go func() {
		i := 0
		for {
			if i%2 == 0 {
				TestGauge.Add(1)
			} else {
				TestGauge.Sub(1)
			}
			time.Sleep(time.Second * 2)
			i += 1
		}
	}()

	// histogram
	go func() {
		i := 0.0
		for {
			TestHistogram.Observe(30 + math.Floor(float64(rand.Intn(120))*math.Sin(i*0.1))/10)
			time.Sleep(time.Second * 2)
			i += 1
		}
	}()

	// summary
	go func() {
		i := 0.0
		for {
			TestSummary.Observe(30 + math.Floor(float64(rand.Intn(120))*math.Sin(i*0.1))/10)
			time.Sleep(time.Second * 2)
			i += 1
		}
	}()

	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/counter", sayHello)
	http.ListenAndServe(":8080", nil)
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	TestCounter.Inc()
	fmt.Fprintf(w, "Hello World!")
}
