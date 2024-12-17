# Stage 1: Build the Go binary
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum first to cache module downloads
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the application
RUN go build -o /out/receipt-processor ./cmd/server/main.go

# Stage 2: Create a minimal runtime image
FROM alpine:3.17
WORKDIR /app

# Copy binary from builder stage
COPY --from=builder /out/receipt-processor /usr/local/bin/receipt-processor

# Expose the application port
EXPOSE 8080

# Start the application
CMD ["receipt-processor"]