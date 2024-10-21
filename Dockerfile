# Start with the official Golang image
FROM golang:1.23-alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY . .

# Download all dependencies. Dependencies will be cached if they haven't changed
RUN go mod download

# Copy the rest of the application code
COPY .env /app/

# Build the Go app
RUN go build -o main .

# Start a new stage from scratch
FROM alpine:latest  

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the pre-built binary from the builder stage
COPY --from=builder /app/main .

# Expose port (e.g., 8080)
EXPOSE 8080

# Command to run the Go app
CMD ["./main"]
