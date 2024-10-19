# Build stage
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Copy go mod and sum files
COPY go.mod ./
COPY go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code
COPY main.go ./
COPY handler.go ./
COPY query.go ./

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o webhook-server .

# Final stage
FROM alpine:latest

WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/webhook-server .

# Expose the port the app runs on
EXPOSE 8443

# Run the binary
CMD ["./webhook-server"]
