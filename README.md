# langsmith-exporter

[![Build Status](https://github.com/ckrowiorsch/langsmith-exporter/actions/workflows/ci.yml/badge.svg)](https://github.com/ckrowiorsch/langsmith-exporter/actions/workflows/ci.yml)

Ein Prometheus Exporter für das Langsmith Tracing-System.

## Features
- Exportiert Metriken wie Runs, fehlgeschlagene Runs und Gesamtkosten aus einem Langsmith-Projekt
- Prometheus-kompatibles /metrics-Endpoint
- Docker- und Kubernetes-fähig

## Nutzung

### Voraussetzungen
- Go 1.22 oder neuer
- Zugriff auf die Langsmith API (API Key und Project ID)

### Build & Start

```bash
go build -o langsmith-exporter main.go
LANGSMITH_API_KEY=... LANGSMITH_PROJECT_ID=... ./langsmith-exporter
```

### Docker

```bash
docker build -t langsmith-exporter .
docker run -e LANGSMITH_API_KEY=... -e LANGSMITH_PROJECT_ID=... -p 8080:8080 langsmith-exporter
```

## Konfiguration
- `LANGSMITH_API_KEY`: API Key für Langsmith
- `LANGSMITH_PROJECT_ID`: Projekt-ID
- `EXPORTER_LISTEN_ADDR`: (optional) Adresse, auf der der Exporter lauscht (Standard: `:8080`)

## Lizenz
MIT License

## Beispiel: Kubernetes Integration

Um den Exporter in Kubernetes zu betreiben, kann ein Deployment und ein Service wie folgt aussehen:

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: langsmith-exporter
spec:
  replicas: 1
  selector:
    matchLabels:
      app: langsmith-exporter
  template:
    metadata:
      labels:
        app: langsmith-exporter
    spec:
      containers:
      - name: exporter
        image: ghcr.io/ckrowiorsch/langsmith-exporter:latest
        env:
        - name: LANGSMITH_API_KEY
          value: "<dein-api-key>"
        - name: LANGSMITH_PROJECT_ID
          value: "<deine-projekt-id>"
        ports:
        - containerPort: 8080
---
apiVersion: v1
kind: Service
metadata:
  name: langsmith-exporter
spec:
  selector:
    app: langsmith-exporter
  ports:
  - protocol: TCP
    port: 8080
    targetPort: 8080
```

Der Exporter ist dann unter `http://<service-ip>:8080/metrics` erreichbar und kann von Prometheus gescraped werden.
