package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"
	"github.com/yourusername/langsmith-exporter/exporter"
	"github.com/yourusername/langsmith-exporter/langsmith"
)

func main() {
	var apiKey, projectID, listenAddr string

	var rootCmd = &cobra.Command{
		Use:   "langsmith-exporter",
		Short: "Prometheus Exporter für Langsmith Tracing-Projekte",
		Run: func(cmd *cobra.Command, args []string) {
			apiKey = os.Getenv("LANGSMITH_API_KEY")
			projectID = os.Getenv("LANGSMITH_PROJECT_ID")
			listenAddr = os.Getenv("EXPORTER_LISTEN_ADDR")
			if listenAddr == "" {
				listenAddr = ":8080"
			}
			if apiKey == "" || projectID == "" {
				log.Fatal("LANGSMITH_API_KEY und LANGSMITH_PROJECT_ID müssen als Umgebungsvariablen gesetzt sein.")
			}

			exporter.InitExporter()
			client := langsmith.NewClient(apiKey, projectID)

			go func() {
				for {
					runs, err := client.GetRuns()
					if err != nil {
						log.Println("Fehler beim Abrufen der Runs:", err)
						runs = 0
					}
					failed, err := client.GetFailedRuns()
					if err != nil {
						log.Println("Fehler beim Abrufen der fehlgeschlagenen Runs:", err)
						failed = 0
					}
					costs, err := client.GetTotalCosts()
					if err != nil {
						log.Println("Fehler beim Abrufen der Kosten:", err)
						costs = 0.0
					}
					exporter.SetMetrics(runs, failed, costs)
					time.Sleep(60 * time.Second)
				}
			}()

			log.Printf("Exporter läuft auf %s/metrics", listenAddr)
			exporter.StartExporter(listenAddr)
		},
	}
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
