# langsmith-exporter

[![Build Status](https://github.com/ckrowiorsch/langsmith-exporter/actions/workflows/ci.yml/badge.svg)](https://github.com/ckrowiorsch/langsmith-exporter/actions/workflows/ci.yml)

A Prometheus exporter for the Langsmith tracing system.

## Features
- Exports metrics such as runs, failed runs, and total costs from a Langsmith project
- Prometheus-compatible /metrics endpoint
- Docker and Kubernetes ready

## Usage

### Requirements
- Go 1.22 or newer
- Access to the Langsmith API (API Key and Project ID)

### Build & Run

```bash
go build -o langsmith-exporter main.go
LANGSMITH_API_KEY=... LANGSMITH_PROJECT_ID=... ./langsmith-exporter
```

### Docker

```bash
docker build -t langsmith-exporter .
docker run -e LANGSMITH_API_KEY=... -e LANGSMITH_PROJECT_ID=... -p 8080:8080 langsmith-exporter
```

## Configuration
- `LANGSMITH_API_KEY`: API key for Langsmith
- `LANGSMITH_PROJECT_ID`: Project ID
- `EXPORTER_LISTEN_ADDR`: (optional) Address the exporter listens on (default: `:8080`)

## License
MIT License

## Example: Kubernetes Integration

To run the exporter in Kubernetes, you can use a Deployment and Service like this:

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
          value: "<your-api-key>"
        - name: LANGSMITH_PROJECT_ID
          value: "<your-project-id>"
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

The exporter will then be available at `http://<service-ip>:8080/metrics` and can be scraped by Prometheus.

## GitHub Actions: Docker Image Build & Push

A separate GitHub Action workflow automatically builds and pushes a Docker image to the GitHub Container Registry (`ghcr.io`) **when a pull request is merged into `main` and all tests are successful**. You can also trigger this workflow manually from the GitHub Actions tab.

- **Image Tag Format:** `yyMMdd_HHmm` (e.g., `250525_1530` for May 25, 2025, 15:30 UTC)
- **Registry:** `ghcr.io/<owner>/<repo>:<tag>`
- **Manual Trigger:** Go to the [Actions tab](/ckrowiorsch/langsmith-exporter/actions) and select the "Build and Push Docker Image" workflow, then click "Run workflow".

### Example: Using the Latest Image

Update your Kubernetes or Docker deployment to use the latest image tag:

```yaml
containers:
  - name: exporter
    image: ghcr.io/<owner>/<repo>:<tag>
    # ...
```

See `.github/workflows/docker-image.yml` for details.
