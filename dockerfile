# Build stage: use official Go image to compile the app
FROM golang:1.23.5-alpine AS builder

# Set working directory inside container
WORKDIR /app

# Copy go.mod and go.sum first (for caching dependencies)
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy all source files to /app
COPY . .

# Build the Go application (binary name: banking-app)
RUN go build -o banking-app main.go

# Final stage: small image to run the app
FROM alpine:latest

# Add ca-certificates for HTTPS if needed
RUN apk --no-cache add ca-certificates

# Copy the binary from builder stage
COPY --from=builder /app/banking-app /banking-app

# Expose port your app listens on (e.g., 8282)
EXPOSE 8282

# Command to run the binary
ENTRYPOINT ["/banking-app"]
CMD ["/banking-app"]
