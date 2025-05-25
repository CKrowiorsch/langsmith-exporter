package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "langsmith-exporter",
		Short: "Prometheus Exporter für Langsmith Tracing-Projekte",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Langsmith Exporter läuft...")
		},
	}
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
