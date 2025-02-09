# sentinelAgent

**sentinelAgent** is a lightweight, fast, and efficient agent for monitoring Linux servers. It collects system metrics, monitors systemd units, gathers logs, and sends data via message brokers. The agent supports both **pull** and **push** models and is designed for easy deployment and high performance.

> **Note**: This project is still in development. Some features may be subject to changes, and certain functionality might not yet be fully implemented.

## Features

- **Metrics Collection**: Gathers system metrics such as CPU, RAM, Disk, and Network usage.
- **Systemd Monitoring**: Monitors systemd units, checks statuses, and tracks service restarts.
- **Log Collection**: Collects logs from systemd and other system logs.
- **Push Model**: Sends data to message brokers (e.g., RabbitMQ) in real-time.
- **Pull Model**: Exposes metrics through HTTP endpoints for pull-based collection (e.g., Prometheus).
- **High Performance**: Designed to be lightweight and fast, with minimal overhead, perfect for environments where speed is critical.
- **Easy Deployment**: Single binary for simple installation and usage.

## Project Structure

.
├── cmd
│   └── sentinel
│       └── main.go
├── config
│   └── config.yml
├── go.mod
├── go.sum
├── internal
│   ├── collector
│   │   ├── logs
│   │   │   ├── logging.go
│   │   │   └── systemd.go
│   │   └── metrics
│   │       ├── linux.go
│   │       └── metrics.go
│   ├── config
│   │   └── config.go
│   ├── core
│   │   ├── coordinator.go
│   │   └── pipeline.go
│   ├── health
│   │   └── checker.go
│   ├── storage
│   │   ├── postgres
│   │   │   └── postgres.go
│   │   └── redis
│   │       └── redis.go
│   ├── transport
│   │   ├── pull
│   │   │   └── http.go
│   │   └── push
│   │       └── rabbitmq.go
│   └── utils
│       └── utils.go
├── pkg
│   └── logger
│       ├── logger.go
│       └── sl
│           └── sl.go
└── README.md