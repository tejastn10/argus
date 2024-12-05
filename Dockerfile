# Stage 1: Build the application
FROM golang:1.23 as builder

# Set the working directory inside the builder stage
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the application source code
COPY . .

# Build the Go application statically
RUN CGO_ENABLED=0 GOOS=linux go build -o argus

# Stage 2: Create a lightweight runtime image
FROM alpine:latest

# Install certificates for HTTPS support
RUN apk add --no-cache ca-certificates

# Set the working directory inside the runtime image
WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/argus .

# Set the default command to run the binary
CMD ["./argus"]
