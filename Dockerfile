FROM golang:1.24.1-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o langsmith-exporter main.go

FROM alpine:3.18
RUN apk add --no-cache ca-certificates

RUN addgroup -S appgroup && adduser -S appuser -G appgroup

WORKDIR /home/appuser/

COPY --from=builder --chown=appuser:appgroup /app/langsmith-exporter .

USER appuser
EXPOSE 8080

CMD ["./langsmith-exporter"]
