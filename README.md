# Go REST API with Docker

A simple REST API built with Go, Fiber, and GORM for managing facts (questions and answers). Supports PostgreSQL and Oracle databases, containerized with Docker Compose, and includes hot reloading for development.

## Features

- RESTful API for CRUD operations on facts
- Support for PostgreSQL and Oracle databases
- Docker containerization with multi-stage builds
- Hot reloading with Air during development
- Graceful shutdown handling
- Clean architecture with handlers, models, and database layers
- Environment-based configuration
- Comprehensive error handling

## Prerequisites

- Docker and Docker Compose (for containerized development)
- Go 1.21+ (for local development)
- PostgreSQL or Oracle database

## Quick Start

### Option 1: Docker Development (Recommended)

1. Clone the repository:
   ```bash
   git clone https://github.com/vizchamz/go-rest-docker.git
   cd go-rest-docker
   ```

2. Copy the environment template:
   ```bash
   cp .env-sample .env
   ```

3. Update the `.env` file with your database credentials:
   ```env
   APPLICATION_DB=postgres
   POSTGRES_DB_HOST=localhost
   POSTGRES_DB_USER=your_db_user
   POSTGRES_DB_PASSWORD=your_db_password
   POSTGRES_DB_NAME=your_db_name
   ```

4. Run with Docker Compose:
   ```bash
   docker-compose up --build
   ```

### Option 2: Local Development

1. Clone and setup:
   ```bash
   git clone https://github.com/vizchamz/go-rest-docker.git
   cd go-rest-docker
   cp .env-sample .env
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Setup local PostgreSQL database and update `.env`

4. Build the application:
   ```bash
   go build -o main ./cmd
   ```

5. Run the application:
   ```bash
   ./main
   ```

   Or run with hot reload during development:
   ```bash
   air
   ```

The API will be available at `http://localhost:3000`.

## API Endpoints

### GET /
Returns a welcome message.

**Response:**
```json
"Hello, Visal Dharmasiri!"
```

### GET /fact
Lists all facts in the database.

**Response:**
```json
[
  {
    "ID": 1,
    "CreatedAt": "2026-01-06T06:27:03.025298+05:30",
    "UpdatedAt": "2026-01-06T06:27:03.025298+05:30",
    "DeletedAt": null,
    "question": "What is the capital of France?",
    "answer": "Paris"
  }
]
```

### POST /fact
Creates a new fact.

**Request Body:**
```json
{
  "question": "What is the capital of France?",
  "answer": "Paris"
}
```

**Response:**
```json
{
  "ID": 1,
  "CreatedAt": "2026-01-06T06:27:03.025298+05:30",
  "UpdatedAt": "2026-01-06T06:27:03.025298+05:30",
  "DeletedAt": null,
  "question": "What is the capital of France?",
  "answer": "Paris"
}
```

## Development

### Local Development Setup

1. **Install Go dependencies:**
   ```bash
   go mod tidy
   ```

2. **Setup database:**
   - For PostgreSQL: Create a database and user
   - Update `.env` with your database credentials

3. **Install Air for hot reloading:**
   ```bash
   go install github.com/cosmtrek/air@latest
   ```

4. **Run development server:**
   ```bash
   air
   ```

### Docker Development

The project includes Docker support for both development and production:

- **Development:** `docker-compose up --build`
- **Production:** Use the multi-stage Dockerfile for optimized builds

### Project Structure

```
.
├── .air.toml              # Air hot reload configuration
├── .dockerignore          # Docker ignore file
├── .env-sample            # Environment variables template
├── Dockerfile             # Multi-stage Docker build
├── docker-compose.yml     # Docker Compose configuration
├── cmd/
│   ├── main.go            # Application entry point with graceful shutdown
│   └── routes.go          # Route definitions
├── database/
│   └── database.go        # Database connection and GORM setup
├── handlers/
│   └── facts.go           # HTTP handlers for facts API
├── models/
│   └── models.go          # GORM data models
├── go.mod                 # Go module dependencies
├── go.sum                 # Go module checksums
└── README.md              # This file
```

## Environment Variables

### Required Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `APPLICATION_DB` | Database type (`postgres` or `oracle`) | `postgres` |

### PostgreSQL Variables

| Variable | Description | Example |
|----------|-------------|---------|
| `POSTGRES_DB_HOST` | Database host | `localhost` |
| `POSTGRES_DB_USER` | Database username | `myuser` |
| `POSTGRES_DB_PASSWORD` | Database password | `mypass` |
| `POSTGRES_DB_NAME` | Database name | `mydb` |

### Oracle Variables

| Variable | Description | Example |
|----------|-------------|---------|
| `ORACLE_DB_HOST` | Database host | `localhost` |
| `ORACLE_DB_PORT` | Database port | `1521` |
| `ORACLE_DB_SID` | Database SID | `ORCL` |
| `ORACLE_DB_USER` | Database username | `myuser` |
| `ORACLE_DB_PASSWORD` | Database password | `mypass` |

## Testing the API

### Create a Fact
```bash
curl -X POST http://localhost:3000/fact \
  -H "Content-Type: application/json" \
  -d '{"question": "What is 2+2?", "answer": "4"}'
```

### List All Facts
```bash
curl http://localhost:3000/fact
```

### Welcome Message
```bash
curl http://localhost:3000/
```

## Deployment

### Local Binary Build

For production deployment on a server without Docker:

1. Build the optimized binary:
   ```bash
   CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd
   ```

2. Run the application:
   ```bash
   ./main
   ```

### Docker Production Build

The Dockerfile includes multi-stage builds for optimized production images:

```bash
# Build production image
docker build -t go-rest-api .

# Run production container
docker run -p 3000:3000 --env-file .env go-rest-api
```

### Docker Compose Production

For production deployment with database:

```bash
docker-compose -f docker-compose.yml up -d
```
