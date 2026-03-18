# Prerequisite

- Go 1.25.5
- PostgreSQL
- RabbitMQ (for 3D optimization)
- AWS S3 (or S3-compatible storage)

# Quickstart
```bash
make install # or go mod download
make # or go run cmd/main.go
```

# API Endpoints

- **GraphQL API**: `/graphql`
- **GraphQL Playground**: `/`
- **Health Check**: `/health`
- **Optimization Webhook**: `/webhook/optimization` (POST)

# Features

## 3D Model Optimization

Asynchronous 3D model optimization using Argo Workflows with mesh reduction and DRACO compression.

- Upload GLB/GLTF models via pre-signed URLs
- Configure compression quality and quantization levels
- Automatic webhook notifications on completion
- Track optimization status in database

See [docs/3D_OPTIMIZATION.md](docs/3D_OPTIMIZATION.md) for detailed documentation.

# Environment Variables

The necessary environment variables can be seen in the .env.example file.

Required for 3D optimization:
- `RABBITMQ_URL` - RabbitMQ connection string
- `RABBITMQ_EXCHANGE` - Exchange for optimization jobs
- `RABBITMQ_ROUTING_KEY` - Routing key for job messages
- `OPTIMIZATION_WEBHOOK_URL` - Public webhook URL (optional, defaults to localhost)

# Development

## Generate GraphQL Code
```bash
make generate # or go generate ./...
```

## Format Code
```bash
make format
```

## Run Linter
```bash
make lint # or golangci-lint run ./...
```

## Run Test
```bash
make test # or go test ./...
```