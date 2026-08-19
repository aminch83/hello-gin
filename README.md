# hello-gin

A hello world application with Go and Gin.

## Overview

`hello-gin` is a small example REST API built with [Gin](https://github.com/gin-gonic/gin), a
lightweight HTTP web framework for Go. It exposes a health check endpoint along with sample
`products` and `users` endpoints.

## Requirements

- Go 1.21.4 or later

## Getting Started

### Install dependencies

```sh
make deps
```

### Run the application (development mode)

```sh
make run
```

### Run the application (production mode)

```sh
make run-prod
```

The server listens on port `8080` by default.

### Run tests

```sh
make test
```

### Build only

```sh
make build
```

## API Endpoints

| Method | Path            | Description                  |
|--------|-----------------|-------------------------------|
| GET    | `/health`       | Health check                  |
| GET    | `/products`     | List all products             |
| GET    | `/products/:id` | Get a single product by ID    |

## Docker

Build and run the application using Docker:

```sh
docker build -t hello-gin .
docker run -p 8080:8080 hello-gin
```

## License

See the [LICENSE](./LICENSE) file for details.
