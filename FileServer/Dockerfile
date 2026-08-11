# Build stage
FROM golang:1.23-alpine AS builder

# Set working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o fileserver

# Final stage
FROM alpine:latest

# Install necessary certificates
RUN apk --no-cache add ca-certificates

# Create directory for serving files
RUN mkdir -p /home/milind

WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/fileserver .

# Expose port 8080
EXPOSE 8080

# Run the application
CMD ["./fileserver"]
