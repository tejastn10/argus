# Argus 🕵️‍♂️

[![Docker Pulls](https://img.shields.io/docker/pulls/tejastn10/argus)](https://hub.docker.com/r/tejastn10/argus)

Argus is a lightweight and efficient uptime monitoring service written in Go. It periodically checks the availability and response time of a given URL and logs the results. Designed with simplicity, reliability, and extensibility in mind, Argus is a great starting point for developers looking to monitor service health.

## Features

- [x] Periodic URL monitoring
- [x] Response time measurement
- [x] Logs URL status (up or down) with error details
- [x] Customizable monitoring intervals
- [x] Lightweight and fast
- [x] Easy to extend and deploy

## Getting Started

### Prerequisites

- [Go 1.23+](https://go.dev/doc/install) installed on your machine.
- [Git](https://git-scm.com/) for version control (optional).
- (Optional) Docker for containerized deployment.

---

### Installation

1. Clone this repository:

    ```bash
    git clone https://github.com/tejastn10/argus.git
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

1. **Build the project:**

    ```bash
    go build -o argus
    ```

2. **Run the service with URL and monitoring interval flags:**

    You can pass the URL and monitoring interval as command-line flags:

    ```bash
    ./argus -url=https://example.com -interval=30s
    ```

3. Argus will start monitoring the specified URL and log the status, response time, and errors (if any).

4. **Example Docker Usage:**

    If you're running Argus in Docker, you can use the pre-configured `docker-compose.yml` file, which automatically passes the required flags to the container.

    ```bash
    docker-compose up -d
    ```

---

### Example Output

```md
INFO    : Starting uptime monitoring for https://example.com every 30 seconds
SUCCESS : URL: https://example.com | Response Time: 1.113759875s | Status: 200
SUCCESS : URL: https://example.com | Response Time: 299.630625ms | Status: 200
```

---

## Project Structure

```md
argus/
├── logs/           # Core package with console and file logging logic 
├── monitor/        # Core package with URL monitoring logic
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

### Example Docker Compose File

We have included an example [docker-compose.yml](./docker-compose.yml) file that can be used to run Argus easily in a Docker container.

---

## Docker Registry

You can pull the Argus Docker image directly from the Docker Hub registry.

Docker Hub Link: [tejastn10/argus](https://hub.docker.com/r/tejastn10/argus)

Contributing

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
