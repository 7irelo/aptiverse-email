# Aptiverse Email Service

A high-performance email microservice written in Go that processes email requests from RabbitMQ.

## Features

- RabbitMQ message consumption
- Concurrent email processing with configurable workers
- SMTP email delivery
- Docker and Docker Compose support
- Configurable retry logic
- Structured logging

## Quick Start

### Prerequisites

- Go 1.21+
- RabbitMQ
- SMTP credentials

### Environment Setup

1. Copy `.env.example` to `.env` and configure your settings:
```bash
cp .env.example .env