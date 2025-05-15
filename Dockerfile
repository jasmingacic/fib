FROM golang:1.23-alpine AS builder

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod ./
COPY go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o fibonacci-service

# Use a minimal alpine image for the final stage
FROM alpine:3.18

# Add ca certificates for HTTPS
RUN apk --no-cache add ca-certificates

# Set working directory
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/fibonacci-service .

# Expose the application port
EXPOSE 8080

# Run the application
ENTRYPOINT ["./fibonacci-service"]