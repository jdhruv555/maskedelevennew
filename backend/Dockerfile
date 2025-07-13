# Multi-stage build for optimized production image
FROM golang:1.24.4-alpine AS builder

# Install build dependencies
RUN apk add --no-cache \
    git \
    ca-certificates \
    tzdata

# Set working directory
WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build with optimizations
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags="-w -s -X main.Version=$(git describe --tags --always --dirty)" \
    -a -installsuffix cgo \
    -o bin/server \
    cmd/server/main.go

# Security scanning stage
FROM aquasec/trivy:latest AS security-scanner
COPY --from=builder /app/bin/server /server
RUN trivy fs --exit-code 1 --severity HIGH,CRITICAL /server || true

# Production stage
FROM alpine:3.19

# Install runtime dependencies
RUN apk add --no-cache \
    ca-certificates \
    tzdata \
    curl \
    && rm -rf /var/cache/apk/*

# Create non-root user
RUN addgroup -g 1001 -S appgroup && \
    adduser -u 1001 -S appuser -G appgroup

# Set working directory
WORKDIR /app

# Copy binary from builder stage
COPY --from=builder /app/bin/server .

# Copy configuration files
COPY --from=builder /app/.env.example .env

# Create necessary directories
RUN mkdir -p /app/logs /app/tmp && \
    chown -R appuser:appgroup /app

# Switch to non-root user
USER appuser

# Expose port
EXPOSE 8080

# Health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD curl -f http://localhost:8080/health || exit 1

# Runtime configuration
ENV GIN_MODE=release
ENV CGO_ENABLED=0

# Security settings
ENV GODEBUG=netdns=go

# Performance settings
ENV GOMAXPROCS=1
ENV GOGC=100

# Start the application
CMD ["./server"]

# Alternative development stage
FROM golang:1.24.4-alpine AS development

# Install development dependencies
RUN apk add --no-cache \
    git \
    ca-certificates \
    tzdata \
    curl \
    vim \
    htop

# Set working directory
WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Expose port
EXPOSE 8080

# Development command
CMD ["go", "run", "cmd/server/main.go"]

# Alternative debug stage
FROM golang:1.24.4-alpine AS debug

# Install debug tools
RUN apk add --no-cache \
    git \
    ca-certificates \
    tzdata \
    curl \
    delve

# Set working directory
WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Expose ports
EXPOSE 8080 2345

# Debug command
CMD ["dlv", "debug", "cmd/server/main.go", "--listen=:2345", "--headless=true", "--api-version=2", "--accept-multiclient"]
