# Start from the official Golang image for building
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Install git in a fixed version for go mod download
RUN apk add --no-cache git=2.40.1-r0

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -o langsmith-exporter main.go

# Use a minimal image for running
FROM alpine:3.18

# Create a non-root user and group
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

WORKDIR /home/appuser/

# Copy the built binary from builder and set ownership
COPY --from=builder --chown=appuser:appgroup /app/langsmith-exporter .

# Switch to the non-root user
USER appuser
# Expose port for the application
EXPOSE 8080

# Command to run
CMD ["./langsmith-exporter"]
