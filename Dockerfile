# Build API Gateway service
FROM golang:1.21-alpine AS builder

# Install git
RUN apk add --no-cache git

# Set working directory to api-gateway
WORKDIR /app/api-gateway

# Copy go mod files
COPY api-gateway/go.mod api-gateway/go.sum ./

# Download dependencies
RUN go mod download

# Copy API Gateway source code
COPY api-gateway/ .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Production stage
FROM alpine:latest

# Install ca-certificates for HTTPS
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy binary from builder
COPY --from=builder /app/api-gateway/main .

# Expose port
EXPOSE 8080

# Health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
  CMD wget --no-verbose --tries=1 --spider http://localhost:8080/health || exit 1

# Run the application
CMD ["./main"]
