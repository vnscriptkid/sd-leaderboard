# Step 1: Build the Go application
FROM golang:1.22.4-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum
COPY go.mod go.sum ./

# Download necessary Go modules
RUN go mod download

# Copy the rest of the application
COPY . .

# Build the Go app
RUN go build -o main .

# Step 2: Create a small image to run the application
FROM alpine:latest

WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/main .

# Run the Go app
CMD ["./main"]
