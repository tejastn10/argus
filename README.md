
# Argus 🕵️‍♂️

Argus is a lightweight and efficient uptime monitoring service written in Go. It periodically checks the availability and response time of a given URL and logs the results. Designed with simplicity, reliability, and extensibility in mind, Argus is a great starting point for developers looking to monitor service health.

## Features

- ✅ Periodic URL monitoring
- ✅ Response time measurement
- ✅ Logs URL status (up or down) with error details
- ✅ Customizable monitoring intervals
- ✅ Lightweight and fast
- ✅ Easy to extend and deploy

## Getting Started

### Prerequisites

- [Go 1.23+](https://go.dev/doc/install) installed on your machine.
- [Git](https://git-scm.com/) for version control (optional).
- (Optional) Docker for containerized deployment.

---

### Installation

1. Clone this repository:

    ```bash
    git clone https://github.com/your-username/argus.git
    cd argus
    ```

2. Install dependencies:

    ```bash
    go mod tidy
    ```

3. Run the project:

    ```bash
    go run main.go
    ```

---

### Usage

1. Update the `main.go` file with the URL and monitoring interval:

    ```go
    url := "https://example.com"
    interval := 30 * time.Second
    ```

2. Build the project:

    ```bash
    go build -o argus
    ```

3. Execute the binary:

    ```bash
    ./argus
    ```

4. Argus will start monitoring the specified URL and log the status, response time, and errors (if any).

---

### Example Output

```md
Starting uptime monitoring for <https://example.com> every 30s
URL: <https://example.com> | Status: true | Response Time: 102ms
URL: <https://example.com> | Status: false | Error: dial tcp: lookup example.com: no such host

```

---

## Project Structure

```md
argus/
├── cmd/            # CLI entry point (future enhancement)
├── monitor/        # Core package with URL monitoring logic
├── internal/       # Non-exported modules for internal use
├── tests/          # Unit and integration tests (future)
├── go.mod          # Dependency management
├── main.go         # Main application entry point
└── README.md       # Project documentation
```

---

## Docker Support 🐳

Deploy Argus as a containerized service using Docker.

### Build Docker Image

```bash
docker build -t argus .
```

### Run Docker Container

```bash
docker run -d argus
```

### Example Dockerfile

```dockerfile
# Dockerfile
FROM golang:1.20 as builder
WORKDIR /app
COPY . .
RUN go build -o argus main.go

# Deploy stage
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/argus .
ENTRYPOINT ["./argus"]
```

---

## Contributing

Contributions are welcome! Please open an issue or submit a pull request.

### To-Do

- Add support for email notifications on downtime.
- Implement a REST API for managing monitored URLs.
- Store logs in a database for analysis.
- Build a dashboard to visualize uptime statistics.

---

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---

## Acknowledgments

- Inspired by the mythological **Argus**, the all-seeing guardian.
- Built with ❤️ and Go.
