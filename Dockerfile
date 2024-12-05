# Start with the official Go image
FROM golang:1.23

# Set the working directory
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the application code
COPY . .

# Build the Go application
RUN go build -o argus

# Set the default command
CMD ["./argus"]
