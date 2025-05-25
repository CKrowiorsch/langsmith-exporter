package exporter

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	runsTotal = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "langsmith_runs_total",
		Help: "Gesamtanzahl der Runs im Langsmith-Projekt.",
	})
	failedRunsTotal = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "langsmith_failed_runs_total",
		Help: "Anzahl der fehlgeschlagenen Runs im Langsmith-Projekt.",
	})
	totalCosts = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "langsmith_total_costs",
		Help: "Gesamtkosten des Langsmith-Projekts.",
	})
)

func InitExporter() {
	prometheus.MustRegister(runsTotal)
	prometheus.MustRegister(failedRunsTotal)
	prometheus.MustRegister(totalCosts)
}

func SetMetrics(runs, failedRuns int, costs float64) {
	runsTotal.Set(float64(runs))
	failedRunsTotal.Set(float64(failedRuns))
	totalCosts.Set(costs)
}

func StartExporter(addr string) {
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(addr, nil)
}
