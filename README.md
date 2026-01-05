# Go REST API with Docker

A simple REST API built with Go, Fiber, and GORM for managing facts (questions and answers). Supports PostgreSQL and Oracle databases, containerized with Docker Compose.

## Features

- RESTful API for CRUD operations on facts
- Support for PostgreSQL and Oracle databases
- Docker containerization
- Hot reloading with Air during development
- Clean architecture with handlers, models, and database layers

## Prerequisites

- Docker and Docker Compose
- Go 1.21+ (for local development)

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/vizchamz/go-rest-docker.git
   cd go-rest-docker
   ```

2. Create a `.env` file in the root directory with the following variables:
   ```env
   APPLICATION_DB=postgres
   POSTGRES_DB_USER=your_db_user
   POSTGRES_DB_PASSWORD=your_db_password
   POSTGRES_DB_NAME=your_db_name
   ```

3. Run the application with Docker Compose:
   ```bash
   docker-compose up --build
   ```

The API will be available at `http://localhost:3000`.

## API Endpoints

- `GET /` - Welcome message
- `GET /fact` - List all facts
- `POST /fact` - Create a new fact

### Create Fact Example

```bash
curl -X POST http://localhost:3000/fact \
  -H "Content-Type: application/json" \
  -d '{"question": "What is the capital of France?", "answer": "Paris"}'
```

## Development

For local development without Docker:

1. Install dependencies:
   ```bash
   go mod tidy
   ```

2. Set up your database (PostgreSQL or Oracle)

3. Run with Air for hot reloading:
   ```bash
   air
   ```

## Project Structure

```
.
├── cmd/
│   ├── main.go          # Application entry point
│   └── routes.go        # Route definitions
├── database/
│   └── database.go      # Database connection and configuration
├── handlers/
│   └── facts.go         # HTTP handlers for facts
├── models/
│   └── models.go        # Data models
├── Dockerfile           # Docker image definition
├── docker-compose.yml   # Docker Compose configuration
├── go.mod               # Go module file
└── README.md            # This file
```

## Environment Variables

- `APPLICATION_DB`: Database type (`postgres` or `oracle`)
- `POSTGRES_DB_USER`: PostgreSQL username
- `POSTGRES_DB_PASSWORD`: PostgreSQL password
- `POSTGRES_DB_NAME`: PostgreSQL database name
