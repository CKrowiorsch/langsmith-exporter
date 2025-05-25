package exporter

import (
	"testing"
)

func TestSetMetrics(t *testing.T) {
	InitExporter()
	SetMetrics(10, 2, 5.5)
	if runsTotal == nil || failedRunsTotal == nil || totalCosts == nil {
		t.Fatal("Metriken wurden nicht initialisiert")
	}
	if runsTotal.Desc().String() == "" {
		t.Error("runsTotal Gauge nicht korrekt initialisiert")
	}
	if failedRunsTotal.Desc().String() == "" {
		t.Error("failedRunsTotal Gauge nicht korrekt initialisiert")
	}
	if totalCosts.Desc().String() == "" {
		t.Error("totalCosts Gauge nicht korrekt initialisiert")
	}
}
